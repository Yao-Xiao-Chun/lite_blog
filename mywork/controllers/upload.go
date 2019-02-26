package controllers

import (
	"github.com/astaxie/beego/logs"
	"path"
	"mime/multipart"
	"os"
	"io/ioutil"
	"sort"
	"mywork/models"
	"strings"
)

/**
	文件上传控制器
 */
type UploadController struct {
	AdminBaseController
}


type UploadNames struct {
	Name string //文件名称
	Ids  string //模块分支
	File multipart.File //文件内容
	Suffix string //后缀名
	Guid string //标示id

}

/**
	接收上传文件
 */
// @router /admin/upload/success [post] 合并上传文件
func (this *UploadController) SetUpload(){
	//获取合并的文件夹
	guid := this.GetString("guid")

	name := this.GetString("name")
	//获取文件后缀
	tags := this.GetStrings("tag")

	tag := strings.Join(tags,",")

	suffix := path.Ext(name)//文件后缀

	if guid == "" || name == ""{

		this.Data["json"] = map[string]string{
			"code":"1002",
			"msg":"未获取到正常的文件合并路径",
		}

	}else{

		//合并文件路径
		if ok := this.doneMergeFile(guid,suffix);ok{

			//写入上传的小说
			models.CreateFiction(models.LiteFiction{Name:name,Tags:tag,Status:1,FileName:"download/"+guid+suffix,Users:this.User.Nikename})

			this.Data["json"] = map[string]string{
				"code":"0",
				"msg":"合并完成",
			}
		}else{
			this.Data["json"] = map[string]string{
				"code":"1002",
				"msg":"合并文件错误，不存在此文件",
			}
		}
	}

	this.ServeJSON()
}

// @router /admin/upload/file [post] 大文件上传 demo试验
func (this *UploadController) UploadFileDone(){
	//接收获取的参数
	fileName := this.GetString("name") //上传名称

	index := this.GetString("chunk") //当前分块的类目

	guid := this.GetString("guid") //文件上传标示

	//获取上传的数据
	f,h,err := this.GetFile("file")

	fileSuffix := path.Ext(h.Filename) //获取文件后缀名称 .xlsx

	defer f.Close()

	if err != nil{

		logs.Error(err)
	}

	if index == ""{

		index = "0"
	}

	data := UploadNames{fileName,index,f,fileSuffix,guid}

	this.createFile(data)

	this.Data["json"] = map[string]interface{}{
		"code":0,
		"msg":"上传成功",
	}

	this.ServeJSON()

}

// @router /admin/upload/index [get] 文件上传首页
func (this *UploadController) UploadIndex(){

	list,_ := models.FindTagTypeTwo()

	this.Data["Tag"] = list

	this.TplName = "admin/upload/index.html"
}
/**
	创建临时目录
 */
func (this *UploadController) createFile(data UploadNames){

	if ok,_:= PathExists("download/"+data.Guid);ok{

	}else{

		os.Mkdir("download/"+data.Guid,0777)
	}

	//保存文件
	this.SaveToFile("file","download/"+data.Guid+"/"+data.Ids+data.Suffix)
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
func (this *UploadController) doneMergeFile(guid,suffix string)bool{
	//检查此文件是否存在
	if ok,_:=PathExists("download/"+guid);ok{
			//遍历当前下的文件
			var data []string
			data = make([]string,0)
			fileInfo,_ := ioutil.ReadDir("download/"+guid)
			//文件内容排序
			for _,val := range fileInfo{
				data = append(data,val.Name())
			}
			//排序
			sort.Strings(data)

			this.mergeFile("download/"+guid+suffix,data,"download/"+guid)
		return true
	}else{

		return false;
	}
}
/**
	合并文件
	@param string 合并后的文件名称，[]string 需要合并的文件名 string 合并的路径
	@return
 */
func (this *UploadController) mergeFile(names string,data []string,path string){

	//创建文件
	f,_ := os.OpenFile(names,os.O_CREATE|os.O_APPEND|os.O_RDWR,0777)

	for _,val := range data{
		//打开当前文件下内容
		contents,_:= ioutil.ReadFile(path+"/"+val)

		f.Write(contents)//写入文件

		//删除当前的文件
		os.Remove(path+"/"+val)
	}

	os.Remove(path)

	defer f.Close()
}