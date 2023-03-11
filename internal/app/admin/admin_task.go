package admin

type AdminTaskController struct {
	AdminBaseController
}

/**
列表页
*/
func (c *AdminTaskController) TaskIndex() {

	c.TplName = "admin/task/index.html"
}

/**
新增设置
*/
