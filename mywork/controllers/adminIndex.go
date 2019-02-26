package controllers

import (
	"mywork/models"
	"time"
	"github.com/astaxie/beego/logs"
	"github.com/tealeg/xlsx"
	"path"
	"os"
	"runtime"
	"io/ioutil"
	"strings"
)

type AdminIndexController struct {

	AdminBaseController
}

type LogInfo struct{

	models.LiteLog

	LogsType string
}

/**
	遍历目录中的文件
 */
type FileInfo struct {
	FileName string //名称
	FileSize int64	//大小
	FilePath string //路径
	FileSizeStr string //转换
}


//excel 文件注解结构体

type SheetName struct {
	Sheet string //xls中的分段
	HeadData map[int]map[int]string
	ContentData map[int]map[int]string
}




/**
	注解路由 后台首页
 */


// @router /admin [get] 后台首页
func (this *AdminIndexController)AdminIndex(){
	this.SetLogs("我被请求到了，看看日志是否存在呢！")
	this.TplName = "admin/index.html"
}

// @router /main [get] 后台 main
func (this *AdminIndexController)AdminMain(){

	this.getTimeLines()

	this.countMessage()

	this.getSystem()

	this.TplName = "admin/main.html"
}


func (this *AdminIndexController) getSystem(){

	var sys string

	sys = runtime.GOOS+" 架构:"+runtime.GOARCH

	cpu := runtime.GOMAXPROCS(0)
	this.Data["Sys"] = sys
	this.Data["Cpu"] = cpu
}


// @router /admin/login [get] 后台 登录页面
func (this *AdminIndexController)AdminLogin(){

	this.TplName = "admin/login.html"
}

// @router /admin/addtime [get] 后台 获取时间先路由
func (this *AdminIndexController)AdminGettime(){

	this.Data["key"] = this.GetUUID()

	this.TplName = "admin/news/addtime.html"
}




/**
	查询 后台的时间线
 */

 func (this *AdminIndexController)getTimeLines(){

	 line, err := models.GetAdminTimeLine() //多维结构体

	 arr := make(map[int]map[string]interface{},10)

	 if err != nil {

	 	arr[0]["code"] = "没有数据"

	 	arr[0]["content"] = "没有数据"

	 }else{

		 //变量赋值
		 for key,val := range line{

			 data := make(map[string]interface{},4) // 每次使用都要初始化一次

			 data["code"] = val.Title

			 data["content"] = val.Content

			 data["id"] = val.ID

			 data["token"] = val.Token

			 arr[key] = data

		 }
	 }

	 num,_:= models.GetHomeCountTimeLine()

	 this.Data["count"] = num

	 this.Data["line"] = arr

 }


 /**
 	基本资料
  */
// @router /admin/user/message [get] 后台 获取用户资料
func (this *AdminIndexController) GetUserMessage(){

	//获取当前登录用户的id

	id := this.User.ID

	user,err := models.FindUser(int(id));

	if err != nil{

	}else{

		this.Data["userinfo"] = map[string]interface{}{
			"title":user.Nikename,
			"status":user.Status,
			"is_admin":user.Is_admin,
			"head_img":user.Head_img,
			"account":user.Account,
			"email":user.Email,
			"password":user.Password,
			"id":user.ID,
		}

	}

	this.TplName = "admin/user/edit.html"
}



/**
	统计今天新增的留言数量
	后台
 */
 func (this *AdminIndexController) countMessage(){

 	//获取今天的时间
 	timeStr := time.Now().Format("2006-01-02 00:00:00")

 	count,_:= models.GetWhereReviewCount(timeStr)

 	this.Data["reviewCount"] = count

 }


 /**
 	后台关于
  */
// @router /admin/baseseting [get] 后台 获取用户资料
func (this *AdminIndexController) SetAbort(){


	this.Data["abort"],_ = models.GetAbort()

	this.TplName = "admin/abort/index.html"
}

/**
	后台关于 数据处理
 */
// @router /admin/baseseting [post] 后台 获取用户资料
func (this *AdminIndexController) AbortFormData() {

	data := this.GetString("content") //获取数据



	if data == ""{

		this.Data["json"] = map[string]interface{}{
			"code":1003,
			"errmsg":"数据丢失",
		}
	}else{
		models.UpdateBase(data)

		this.Data["json"] = map[string]interface{}{
			"code":0,
			"errmsg":"更新成功",
		}
	}

	this.ServeJSON()
}


/**
	后台公告
 */
// @router /admin/baseplacard [get] 后台 获取用户资料
 func (this *AdminIndexController) SetPlacard(){

 	this.Data["abort"],_ = models.GetPlacard()

 	this.TplName = "admin/placard/index.html"
 }

/**
   后台公告 数据处理
*/
// @router /admin/baseplacard [post] 后台 获取用户资料
func (this *AdminIndexController) PlacardFormData()  {

	data := this.GetString("content") //获取数据



	if data == ""{

		this.Data["json"] = map[string]interface{}{
			"code":1003,
			"errmsg":"数据丢失",
		}
	}else{
		models.UpdatePlacard(data)

		this.Data["json"] = map[string]interface{}{
			"code":0,
			"errmsg":"更新成功",
		}
	}

	this.ServeJSON()
}



/**
	获取日志记录
 */
// @router /admin/log/index [get]
func (this *AdminIndexController) GetLogList(){

	this.Data["num"],_ = models.CountLog()

	this.TplName = "admin/log/index.html"
}

/**
	日志数据获取
 */

// @router /admin/log/page/?:key [get]
func (this *AdminIndexController) GetLogPage(){

	var page int

	this.Ctx.Input.Bind(&page,"page")

	var res []models.LiteLog

	res = make([]models.LiteLog,0)

	if page == 0{

		res,_ = models.SelectLog(1)
	}else{

		res,_ = models.SelectLog(page)
	}

	var data []LogInfo

	data = make([]LogInfo,0)

	for _,v := range res{

		arr := LogInfo{v,this.LogTypeStr(v.Level)}

		data =  append(data,arr)
	}

	this.Data["json"] = map[string]interface{}{
		"code":"0",
		"data":data,
	}

	this.ServeJSON()
}

/**
 	生成日志
  */
// @router /download/log [get] 导出日志
func (this *AdminIndexController) DownLog(){

	//获取数据

	list,_ := models.FindLogAll()

	file := this.SetToExcel("登录日志表"+time.Now().Format("2006-01-02"),"download/",list)

	//写入日志
	this.ReadLog("账号:"+this.User.Nikename+" 操作：日志:'导出日志',状态：成功",2)

	this.Ctx.Output.Download(file)
}



/**
	上传excel
 */
// @router /upload/excel [get] 上传excel
func (this *AdminIndexController) ToExcel(){

	this.TplName = "admin/excel/index.html"

}

// @router /upload/done [post] 上传excel
func (this *AdminIndexController) ReadExcel(){
	//获取上传的数据
	f,h,err := this.GetFile("file")

	defer f.Close()

	if err != nil{

		logs.Error(err)
	}

	fileSuffix := path.Ext(h.Filename) //获取文件后缀名称 .xlsx


	if !(fileSuffix != ".xlsx" || fileSuffix != ".xls"){

		logs.Warning("后缀名错误")
	}

	Name := this.GetRandomString(20)+fileSuffix
	//保存文件
	this.SaveToFile("file", "download/" + Name) // 保存位置在 static/upload, 没有文件夹要先创建

	excelFileName := "download/"+Name

	xlFile, err := xlsx.OpenFile(excelFileName)

	if err != nil {
		logs.Warning(err)
	}

	var sheetData  []SheetName

	sheetData = make([]SheetName,0)

	var headData map[int]map[int]string

	var contentData map[int]map[int]string

	headData = make(map[int]map[int]string)//存放表头信息

	contentData = make(map[int]map[int]string)//存放内容信息

	for _, sheet := range xlFile.Sheets {

		//fmt.Printf("Sheet Name: %s\n", sheet.Name)
		for index, row := range sheet.Rows {

			hData := make(map[int]string)

			cData := make(map[int]string)

			for k, cell := range row.Cells {

				text := cell.String()

				if index == 0{

					hData[k] = text
				}else{
					cData[k] = text
				}

				//fmt.Printf("%s\n", text)
			}

			if index == 0{

				headData[index] = hData

			}else{

				contentData[index] = cData
			}
		}

		data := SheetName{sheet.Name,headData,contentData} //把值放入结构体中

		sheetData = append(sheetData,data)

	}

	//logs.Info(sheetData)

	//删除使用后的文件
	defer func() {
		os.Remove(excelFileName)
	}()

	this.Data["json"] = map[string]interface{}{
		"code":0,
		"data":sheetData,
	}
	this.ServeJSON()
}

// @router /admin/clear [get] 清楚日志界面
func (this *AdminIndexController) SetClear(){

	this.TplName = "admin/log/file.html"
}

// @router /admin/clear/log/?:key [get] 缓存日志界面
func (this *AdminIndexController) GetFileName(){

	var dirName string

	this.Ctx.Input.Bind(&dirName,"dir_name")

	this.Data["json"] = map[string]interface{}{
		"code":0,
		"data":this.ReadFile(dirName),
	}
	this.ServeJSON()
}

// @router /admin/clear/download/?:key [get] 缓存文件
func (this *AdminIndexController) DownCacheLog(){

	var path,names string

	this.Ctx.Input.Bind(&path,"names")

	pathInfo := strings.Split(path,"/")

	names = pathInfo[len(pathInfo)-1]//获取文件名称

	list := this.ReadFile(pathInfo[0])//获取路径

	var flag bool

	flag = false

	for _,val := range list{

		if val.FileName == names{

			flag = true
			break

		}
	}

	if flag{

		this.Data["json"] = map[string]interface{}{
			"code":"0",
			"msg":path,
		}

	}else{
		this.Data["json"] = map[string]interface{}{
			"code":"1003",
			"msg":"没有找到该文件，请稍后再试",
		}
	}

	this.ServeJSON()
}

// @router /admin/clear/download/logs/?:key [get] 下载文件
func (this *AdminIndexController) DownFile(){

	var str string

	this.Ctx.Input.Bind(&str,"file")

	//判断是否在可供下载的目录

	if str == ""{
		this.Abort("404")
	}

	dir := strings.Split(str,"/")

	//设置默认可以访问的文件夹
	arr := map[string]string{
		"download":"download",
		"logs":"logs",
	}

	if _,ok := arr[dir[0]];ok{
		this.Ctx.Output.Download(str)
	}else{
		this.Abort("404")
	}
}

// @router /admin/clear/delete/?:key [get] 删除文件
func (this *AdminIndexController) DeleteFile(){

	var path,names string

	this.Ctx.Input.Bind(&path,"names")

	pathInfo := strings.Split(path,"/")

	names = pathInfo[len(pathInfo)-1]//获取文件名称

	//我要获取前面的所有路径，包括其他
	pathNames := pathInfo[1:]

	lastPath := strings.Join(pathNames,"/") //拼接路径，防止多个问题

	list := this.ReadFile(pathInfo[0])

	var flag bool

	flag = false

	for _,val := range list{

		if val.FileName == names{

			flag = true
			break

		}
	}

	if flag {

		err := os.Remove(pathInfo[0]+"/"+lastPath)

		if err != nil{
			this.Data["json"] = map[string]interface{}{
				"code":"1004",
				"msg":"删除失败",
			}
		}else{

			this.Data["json"] = map[string]interface{}{
				"code":"0",
				"msg":"删除成功",
			}
		}


	}else{
		this.Data["json"] = map[string]interface{}{
			"code":"1003",
			"msg":"没有找到该文件，请稍后再试",
		}
	}

	this.ServeJSON()
}



/**
	遍历日志目录下的文件
	@param key string 查询的文件 default
 */
func (this *AdminIndexController) ReadFile(key string) (list []FileInfo){

	//设置默认可以访问的文件夹
	arr := map[string]string{
		"download":"download",
		"logs":"logs",
	}

	if _,ok := arr[key];ok{

		list = this.readAll(key)

	}

	return list
}

/**
	获取遍历文件
	@param string 目录
	@return [] 文件详情
 */
func (this *AdminIndexController) readAll(path string) []FileInfo {

	var all_file []FileInfo

	finfo, _ := ioutil.ReadDir(path)

	for _ ,x := range finfo {

		realPath := FileInfo{x.Name(),x.Size(),path + "/" + x.Name(),this.GetToUnit(int(x.Size()))}

		if x.IsDir() {

			all_file = append(all_file,this.readAll(path + "/" + x.Name())...)
		}else {
			all_file = append(all_file,realPath)
		}
	}

	return all_file
}


