package admin

import (
	"mywork/internal/app/common/dto"
	"mywork/models"
	"strconv"
)

/**
友情链接控制器
*/
type AdminLinkController struct {
	AdminBaseController
}

// @router /admin/link/index [get] 友情首页
func (c *AdminLinkController) Index() {

	c.Data["num"], _ = dto.GetLinkTotal()

	c.TplName = "admin/link/index.html"
}

// @router /admin/link/add [get] 友情新增
func (c *AdminLinkController) LinkAdd() {

	c.TplName = "admin/link/add.html"
}

// @router /admin/link/add [post] 友情新增
func (c *AdminLinkController) LinkAddForm() {
	linkName := c.GetString("link_name")

	linkUrl := c.GetString("link_url")

	linkStatus := c.GetString("link_status")

	linkSort := c.GetString("link_sort")

	if linkName == "" || linkUrl == "" || linkSort == "" || linkStatus == "" {

		c.Data["json"] = map[string]interface{}{
			"code": "1003",
			"msg":  "参数不完整",
		}
	} else {

		var link models.LiteLink

		link.Name = linkName
		link.Url = linkUrl
		link.Sort, _ = strconv.Atoi(linkSort)
		link.Status, _ = strconv.Atoi(linkStatus)

		err := dto.AddLinkData(link)

		if err == nil {

			c.ReadLog("账号:"+c.User.Nikename+" 操作：添加友链:'"+linkName+"',状态：成功", 2) //写入操作日志
		}

		c.Data["json"] = map[string]interface{}{
			"code": "0",
			"msg":  "创建成功",
		}

	}

	c.ServeJSON()
}

/**
获取前台展示数据
*/
// @router /admin/link/page/?:key [get]
func (c *AdminLinkController) GetLinkIndex() {
	var page int

	c.Ctx.Input.Bind(&page, "page")

	var res []models.LiteLink

	res = make([]models.LiteLink, 0)

	if page == 0 {

		res, _ = dto.SelectLink(1)
	} else {

		res, _ = dto.SelectLink(page)
	}

	c.Data["json"] = map[string]interface{}{
		"code": "0",
		"data": res,
	}

	c.ServeJSON()
}

// @router /admin/link/edit/?:key [get] 编辑页面
func (c *AdminLinkController) LinkInfo() {

	var id int

	c.Ctx.Input.Bind(&id, "id")

	c.Data["link"], _ = dto.FindLinkInfo(id)

	c.TplName = "admin/link/edit.html"
}

// @router /admin/link/edit [post] 编辑保存
func (c *AdminLinkController) LinkFormData() {

	linkName := c.GetString("link_name")

	linkUrl := c.GetString("link_url")

	linkStatus := c.GetString("link_status")

	linkSort := c.GetString("link_sort")

	id := c.GetString("id")

	if linkName == "" || linkUrl == "" || linkSort == "" || linkStatus == "" || id == "" {

		c.Data["json"] = map[string]interface{}{
			"code": "1003",
			"msg":  "参数不完整",
		}
	} else {

		ids, _ := strconv.Atoi(id)

		var link models.LiteLink

		link.Name = linkName
		link.Url = linkUrl
		link.Sort, _ = strconv.Atoi(linkSort)
		link.Status, _ = strconv.Atoi(linkStatus)
		link.ID = uint(ids)

		err := dto.SaveLink(link)

		if err == nil {

			c.ReadLog("账号:"+c.User.Nikename+" 操作：修改友链:'"+linkName+"',状态：成功", 2) //写入操作日志
		}

		c.Data["json"] = map[string]interface{}{
			"code": "0",
			"msg":  "修改成功",
		}

	}

	c.ServeJSON()
}

/**
删除
*/
// @router /admin/link/delete/?:key [get] 编辑保存
func (c *AdminLinkController) LinkDel() {

	var id int

	c.Ctx.Input.Bind(&id, "id")

	if id == 0 {

		c.Data["json"] = map[string]interface{}{
			"code": "1003",
			"msg":  "参数不完整",
		}
	} else {

		err := dto.DeleteLink(id)

		if err == nil {

			c.ReadLog("账号:"+c.User.Nikename+" 操作：删除友链:'"+strconv.Itoa(id)+"',状态：成功", 2) //写入操作日志
		}

		c.Data["json"] = map[string]interface{}{
			"code": "0",
			"msg":  "删除成功",
		}
	}

	c.ServeJSON()
}
