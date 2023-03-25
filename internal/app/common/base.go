package common

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/satori/go.uuid"
	"github.com/tabalt/ipquery"
	"github.com/tealeg/xlsx"
	"io"
	"lite_blog/models"
	"strconv"
	"strings"
	"time"
)

/**
基础控制器
*/
type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {

	c.Data["Path"] = c.Ctx.Request.RequestURI //获取当前中的url

}

func (c *BaseController) Abort500(err error) {

	c.Data["error"] = err

	c.Abort("500")
}

/**
生成随机TOKen
*/
func (c *BaseController) SetToken() {

}

/**
生成加密密码
*/
func (c *BaseController) SetMd5Pwd(pwd string) string {

	//假设用户名abc，密码123456
	h := md5.New()

	_, err := io.WriteString(h, pwd)
	if err != nil {
		return ""
	}

	//pwmd5等于e10adc3949ba59abbe56e057f20f883e
	//93bcab4ab719fde430e5ad90656a240e
	pwmd5 := fmt.Sprintf("%x", h.Sum(nil))

	return pwmd5
}

/**
生成 uuid
*/
func (c *BaseController) GetUUID() interface{} {

	uuidStr := uuid.NewV4()

	return uuidStr
}

/**
生成格式化时间
*/

func (c *BaseController) Date(times string) string {

	datetime := time.Now().Format(times)

	return datetime
}

/**
设置全局的图片上传类型
*/
func (c *BaseController) GetUploadTypeImage() (arr map[string]string) {

	return map[string]string{
		"0": "jpg",
		"1": "gif",
		"2": "jpeg",
		"3": "png",
	}
}

func JsonFormat(retcode int, retmsg string, retdata interface{}, stime time.Time) (json map[string]interface{}) {
	cost := time.Now().Sub(stime).Seconds()
	if retcode == 1 {
		json = map[string]interface{}{
			"code": retcode,
			"data": retdata,
			"desc": retmsg,
			"cost": cost,
		}
	} else {
		json = map[string]interface{}{
			"code": retcode,
			"desc": retmsg,
			"cost": cost,
		}
	}

	return json
}

/**
获取ip省份
*/
func (c *BaseController) GetAddress(ip string) string {

	if ip == "" {

		return "未知省份"
	}

	df := "conf/testdata/test_10000.txt"

	err := ipquery.Load(df)
	if err != nil {

		fmt.Println(err)
	}

	dt, err := ipquery.Find(ip)

	if err != nil {

		return "未知地址"

	} else {

		ips := strings.Split(string(dt), "	")

		ips = append(ips[:3], ips[4:]...) //移除第三个切片

		ipStr := strings.Join(ips, "-")

		return ipStr
	}
}

/**
设置日志
*/
func (c *BaseController) SetLogs(str interface{}) {

	beego.Debug(str)

	beego.SetLogger("file", `{"filename":"logs/test.log"}`)
	beego.BeeLogger.DelLogger("console")
	beego.SetLevel(beego.LevelInformational)
	beego.SetLogFuncCall(true)
}

/**
  	获取用户提交的ip
    @param  会尝试获取nginx代理的真实ip 或自带ip
	@return ip
*/
func (c *BaseController) GetIP() string {

	var ip string

	ips := c.Ctx.Request.Header.Get("X-Real-IP") //nginx设置反向代理以后获取的值

	if ips == "" {

		ip = c.Ctx.Request.RemoteAddr

		ip = ip[0:strings.LastIndex(ip, ":")]

	} else {

		ip = ips
	}

	return ip
}

/**
比较两个时间
@param newTime 对比时间
@param oldTime 老时间
@return 路径文件
*/
func (c *BaseController) TimeRangeComparison(oldTime time.Time) string {

	newTime := time.Now()

	num := c.timeSub(newTime, oldTime)

	if num > 1 {

		return oldTime.Format("2006-01-02 15:04:05")

	} else {
		//判断写入的时间
		strMon := newTime.Sub(oldTime)

		mon := int(strMon.Minutes()) / 60

		if mon > 0 {

			return strconv.Itoa(mon) + "小时前"

		} else {

			if int(strMon.Minutes()) > 0 {

				return strconv.Itoa(int(strMon.Minutes())) + "分钟前"

			} else {

				return strconv.Itoa(int(strMon.Seconds())) + "秒前"
			}

		}
	}

}

func (c *BaseController) timeSub(t1, t2 time.Time) int {

	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)

	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)

	return int(t1.Sub(t2).Hours() / 24)
}

/**
导出excel
*/
func (c *BaseController) SetToExcel(name, path string, data []models.LiteLog) string {

	headArr := []string{"编号", "时间", "内容"}

	//创建新的问题
	file := xlsx.NewFile()

	//设置sheet
	sheet, _ := file.AddSheet("日志表格")

	//新增一行表头
	row := sheet.AddRow()

	//设置每行的高度
	row.SetHeightCM(1)

	for _, val := range headArr {

		cell := row.AddCell()

		cell.Value = val
	}

	//处理内容
	for _, v := range data {

		rows := sheet.AddRow() //新增一行

		cells := rows.AddCell()

		cells.Value = strconv.Itoa(int(v.ID))

		cells = rows.AddCell()

		cells.Value = v.CreatedAt.Format("2006-01-02 15:04:05")

		cells = rows.AddCell()

		cells.Value = v.Content
	}

	err := file.Save(path + name + ".xlsx")

	if err != nil {

		panic(err)
	} else {

		return path + name + ".xlsx"
	}

}

/**
转换单位
@param
*/
func (c *BaseController) GetToUnit(size int) string {

	if size < 1024 {
		return strconv.Itoa(size) + "b"
	} else if size >= 1024 && size < (1024*1024) {
		return strconv.Itoa(size/1024) + "kb"
	} else if size >= (1024*1024) && size < (1024*1024*1024) {

		return strconv.Itoa(size/(1024*1024)) + "M"
	} else if size >= (1024 * 1024 * 1024) {

		return strconv.Itoa(size/(1024*1024*1024)) + "G"
	}

	return ""
}
