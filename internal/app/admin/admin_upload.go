package admin

import (
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"lite_blog/internal/app/common/dto"
	"lite_blog/models"
	"mime/multipart"
	"os"
	"path"
	"sort"
	"strings"
)

/**
文件上传控制器
*/
type UploadController struct {
	AdminBaseController
}

type UploadNames struct {
	Name   string         //文件名称
	Ids    string         //模块分支
	File   multipart.File //文件内容
	Suffix string         //后缀名
	Guid   string         //标示id

}

/**
接收上传文件
*/
// @router /admin/upload/success [post] 合并上传文件
func (c *UploadController) SetUpload() {
	//获取合并的文件夹
	guid := c.GetString("guid")

	name := c.GetString("name")
	//获取文件后缀
	tags := c.GetStrings("tag")

	tag := strings.Join(tags, ",")

	suffix := path.Ext(name) //文件后缀

	if guid == "" || name == "" {

		c.Data["json"] = map[string]string{
			"code": "1002",
			"msg":  "未获取到正常的文件合并路径",
		}

	} else {

		//合并文件路径
		if ok := c.doneMergeFile(guid, suffix); ok {

			//写入上传的小说
			dto.CreateFiction(models.LiteFiction{Name: name, Tags: tag, Status: 1, FileName: "download/" + guid + suffix, Users: c.User.Nikename})

			c.Data["json"] = map[string]string{
				"code": "0",
				"msg":  "合并完成",
			}
		} else {
			c.Data["json"] = map[string]string{
				"code": "1002",
				"msg":  "合并文件错误，不存在此文件",
			}
		}
	}

	c.ServeJSON()
}

// @router /admin/upload/file [post] 大文件上传 demo试验
func (c *UploadController) UploadFileDone() {
	//接收获取的参数
	fileName := c.GetString("name") //上传名称

	index := c.GetString("chunk") //当前分块的类目

	guid := c.GetString("guid") //文件上传标示

	//获取上传的数据
	f, h, err := c.GetFile("file")

	fileSuffix := path.Ext(h.Filename) //获取文件后缀名称 .xlsx

	defer f.Close()

	if err != nil {

		logs.Error(err)
	}

	if index == "" {

		index = "0"
	}

	data := UploadNames{fileName, index, f, fileSuffix, guid}

	c.createFile(data)

	c.Data["json"] = map[string]interface{}{
		"code": 0,
		"msg":  "上传成功",
	}

	c.ServeJSON()

}

// @router /admin/upload/index [get] 文件上传首页
func (c *UploadController) UploadIndex() {

	list, _ := dto.FindTagTypeTwo()

	c.Data["Tag"] = list

	c.TplName = "admin/upload/index.html"
}

/**
创建临时目录
*/
func (c *UploadController) createFile(data UploadNames) {

	if ok, _ := PathExists("download/" + data.Guid); ok {

	} else {

		os.Mkdir("download/"+data.Guid, 0777)
	}

	//保存文件
	c.SaveToFile("file", "download/"+data.Guid+"/"+data.Ids+data.Suffix)
}

/**
判断文件夹是否存在
*/
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

/**
合并文件
*/
func (c *UploadController) doneMergeFile(guid, suffix string) bool {
	//检查此文件是否存在
	if ok, _ := PathExists("download/" + guid); ok {
		//遍历当前下的文件
		var data []string
		data = make([]string, 0)
		fileInfo, _ := ioutil.ReadDir("download/" + guid)
		//文件内容排序
		for _, val := range fileInfo {
			data = append(data, val.Name())
		}
		//排序
		sort.Strings(data)

		c.mergeFile("download/"+guid+suffix, data, "download/"+guid)
		return true
	} else {

		return false
	}
}

/**
合并文件
@param string 合并后的文件名称，[]string 需要合并的文件名 string 合并的路径
@return
*/
func (c *UploadController) mergeFile(names string, data []string, path string) {

	//创建文件
	f, _ := os.OpenFile(names, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)

	for _, val := range data {
		//打开当前文件下内容
		contents, _ := ioutil.ReadFile(path + "/" + val)

		f.Write(contents) //写入文件

		//删除当前的文件
		os.Remove(path + "/" + val)
	}

	os.Remove(path)

	defer f.Close()
}
