package bins

import (
	"github.com/astaxie/beego/logs"
	"io/ioutil"
)

/**
	执行定时任务的脚本
 */
type Send struct {


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

//test
func StartYx() error {

	logs.Info("发送营销短信...")

	return nil
}

func StartYeWu() error {

	logs.Info("发送行业短信...")

	return nil
}



/**
	定时清理日志功能
 */
func ClearLogs(){

	//循环遍历当前下的数据
	logs.Info(readAll(""))
}

func  readAll(path string) []FileInfo {

	var all_file []FileInfo

	if path == ""{

		path = "logs"
	}

	finfo, _ := ioutil.ReadDir(path)

	for _ ,x := range finfo {

		realPath := FileInfo{x.Name(),x.Size(),path + "/" + x.Name(),""}

		if x.IsDir() {

			all_file = append(all_file,readAll(path + "/" + x.Name())...)
		}else {
			all_file = append(all_file,realPath)
		}
	}

	return all_file
}