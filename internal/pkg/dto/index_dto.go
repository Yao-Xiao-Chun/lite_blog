package dto

type FileInfo struct {
	FileName    string //名称
	FileSize    int64  //大小
	FilePath    string //路径
	FileSizeStr string //转换
}

//excel 文件注解结构体

type SheetName struct {
	Sheet       string //xls中的分段
	HeadData    map[int]map[int]string
	ContentData map[int]map[int]string
}
