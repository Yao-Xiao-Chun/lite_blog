package controllers

type AdminTaskController struct {
	AdminBaseController
}


/**
	列表页
 */
func (this *AdminTaskController) TaskIndex()  {

	this.TplName = "admin/task/index.html"
}

/**
	新增设置
 */
