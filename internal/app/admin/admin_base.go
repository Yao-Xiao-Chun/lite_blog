package admin

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"log"
	"math/rand"
	"mime/multipart"
	"mywork/internal/app/common"
	"mywork/models"
	"regexp"
	"strings"
	"time"
)

/**
后台基础控制器
*/
type AdminBaseController struct {
	common.BaseController
	IsLogin bool
	User    models.LiteOauthUser
}

//设置session Key

const SESSION_ADMIN_KEY = "SESSION_ADMIN_KEY"

/**
后台是否用户登录
*/
func (c *AdminBaseController) Prepare() {

	c.EnableXSRF = false

	user, ok := c.GetSession(SESSION_ADMIN_KEY).(models.LiteOauthUser)

	c.IsLogin = false

	if ok {

		c.User = user

		c.IsLogin = true

		c.Data["AdminUsers"] = c.User

	} else {

		//如果没有登录 就重定向到登录页面 为了防止无限重复定向
		if c.Ctx.Request.RequestURI != "/admin/login" && c.Ctx.Request.RequestURI != "/admin/dologin" {

			c.Ctx.Redirect(302, "/admin/login")
		}

		/*logs.Info(c.Ctx.Request.RequestURI)*/
	}

	c.Data["IsLogin"] = c.IsLogin //是否登录

}

/**
是否是需要的参数
*/
func (c *AdminBaseController) CheckMustKey(key string, msg string) string {

	key = strings.Trim(key, " ")

	result := c.GetString(key)

	if len(result) == 0 {

		c.Abort500(errors.New(msg))
	}

	return result
}

/*
	判断是否是账户邮箱
*/
func (c *AdminBaseController) CheckEmail(email string) string {

	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, email); !m {

		return ""

	} else {

		return email

	}

}

/**
判断是否是合法的时间
*/
func (c *AdminBaseController) CheckTime(date string) {

}

/**
后台基础图片上传
*/

func (c *AdminBaseController) Upload(key string) (img map[string]interface{}) {

	var imgName string
	f, h, err := c.GetFile(key)

	//判断是否是文件上传
	imgStr := "image/jpeg,image/png,image/gif,image/jpg,image/png"

	imgType := h.Header["Content-Type"][0]

	//判断该类型是否在请求的类型中出现
	if strings.Index(imgStr, imgType) < 0 {

		return map[string]interface{}{
			"code": "1003",
			"msg":  "请求类型错误",
		}
	}

	//获取文件后缀
	imgSuffixArr := strings.Split(h.Filename, ".")

	count := len(imgSuffixArr)

	imgSuffix := imgSuffixArr[count-1] //获取最后一个后缀

	flag := false

	imgSuffix = strings.ToLower(imgSuffix)

	for _, val := range c.GetUploadTypeImage() {

		if val == imgSuffix {
			flag = true
			break
		}
	}

	//检测是否存在
	if !flag {

		return map[string]interface{}{
			"code": "1001",
			"msg":  "文件上传错误，获取后缀名错误",
		}
	}

	imgName = c.GetRandomString(16) + "." + imgSuffix

	if err != nil {

		log.Fatal("getfile err ", err)

	}

	defer func(f multipart.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	err = c.SaveToFile(key, "static/upload/"+imgName)
	if err != nil {
		return nil
	} // 保存位置在 static/upload, 没有文件夹要先创建

	return map[string]interface{}{
		"code":  "0",
		"msg":   "保存成功",
		"path":  "static/upload/" + imgName,
		"name":  imgName,
		"names": "/static/upload/" + imgName,
	}
}

// 生成32位MD5
func (c *AdminBaseController) MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

/**
	生成随机字符串
	@param int 长度
    @return string
*/
func (c *AdminBaseController) GetRandomString(lens int) string {

	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	bytes := []byte(str)

	result := []byte{}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < lens; i++ {

		result = append(result, bytes[r.Intn(len(bytes))])
	}

	return string(result)
}

/**
admin日志写入
*/
func (c *AdminBaseController) ReadLog(content string, level int) {

	models.InsertLog(content, level)
}

func (c *AdminBaseController) LogTypeStr(key int) string {

	arr := map[int]string{
		1: "登录日志",
		2: "操作日志",
		3: "定时任务日志",
	}

	for index, val := range arr {

		if index == key {

			return val
		}
	}

	return "不识别状态"
}
