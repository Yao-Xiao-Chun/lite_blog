package dto

/**
遍历目录中的文件
*/
type FileInfo struct {
	FileName    string //名称
	FileSize    int64  //大小
	FilePath    string //路径
	FileSizeStr string //转换
}
