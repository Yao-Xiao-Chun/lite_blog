package bins

import (
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	dos "mywork/pkg/dto"
)

/**
执行定时任务的脚本
*/
type Send struct {
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
func ClearLogs() {

	//循环遍历当前下的数据
	logs.Info(readAll(""))
}

func readAll(path string) []dos.FileInfo {

	var allFile []dos.FileInfo

	if path == "" {

		path = "logs"
	}

	info, _ := ioutil.ReadDir(path)

	for _, x := range info {

		realPath := dos.FileInfo{FileName: x.Name(), FileSize: x.Size(), FilePath: path + "/" + x.Name()}

		if x.IsDir() {

			allFile = append(allFile, readAll(path+"/"+x.Name())...)
		} else {
			allFile = append(allFile, realPath)
		}
	}

	return allFile
}
