package admin

import (
	"github.com/jinzhu/gorm"
	"mywork/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego/logs"
)

type AdminTagController struct {
	AdminBaseController
}

/**
后台tag列表
*/
// @router /admin/tag [get] 后台 tag列表
func (c *AdminTagController) TagList() {

	count, _ := models.QueryCountTag()

	c.Data["num"] = count

	c.TplName = "admin/tag/list.html"
}

/**
获取分页数据
*/
// @router /admin/tag/list/?:key [get] 后台 tag列表
func (c *AdminTagController) GetTagList() {

	var page string

	tagSize := 10

	c.Ctx.Input.Bind(&page, "page")

	if page == "" {

		page = "1"
	}

	list, _, err := models.QueryTagList(page, tagSize)

	if !gorm.IsRecordNotFoundError(err) {

		var arr map[int]map[string]interface{}

		var data map[string]interface{}

		arr = make(map[int]map[string]interface{})

		for index, val := range list {

			data = make(map[string]interface{})

			data["tid"] = val.ID

			data["tag_name"] = val.Tag_name

			data["tag_status"] = val.Is_status

			arr[index] = data
		}

		c.Data["json"] = map[string]interface{}{
			"code": "0",
			"data": arr,
			"msg":  "请求成功",
		}

	} else {

		c.Data["json"] = map[string]interface{}{
			"code":   "1003",
			"errmsg": "获取分页失败",
		}

	}

	c.ServeJSON()
}

/**
新增tag
*/
// @router /admin/tag/add [get] 后台 tag列表
func (c *AdminTagController) TagAdd() {

	c.TplName = "admin/tag/add.html"
}

/**
删除 tag
*/
// @router /admin/tag/del/?:key [get] 后台 tag列表
func (c *AdminTagController) DelTag() {

	var id string

	c.Ctx.Input.Bind(&id, "id")

	ids, _ := strconv.Atoi(id)

	if ids == 0 {

		c.Data["json"] = map[string]interface{}{
			"code":   "1006",
			"errmsg": "删除失败！",
		}
	} else {
		logs.Info(ids)
		_, err := models.TagDel(ids)

		if err == nil {

			c.Data["json"] = map[string]interface{}{
				"code":   "0",
				"errmsg": "删除成功",
			}
		}

	}

	c.ServeJSON()

}

/**
处理新增的标签
*/

// @router /admin/tag/add [post] 后台 tag列表
func (c *AdminTagController) TagAddPost() {

	tag := c.GetString("tag_name")

	tagStatus := c.GetString("status")

	if len(tag) < 0 {

		c.Data["json"] = map[string]interface{}{
			"code":   "1003",
			"errmsg": "参数不完整，请输入签名",
		}

	} else {
		var arr []string

		arr = c.stringFromData(tag)

		var tag models.LiteTag

		for _, val := range arr {

			tag.Tag_name = strings.Trim(val, " ")

			tag.Is_status, _ = strconv.Atoi(tagStatus)

			if val != "" {

				models.CreateTag(tag)
			}

		}

		c.Data["json"] = map[string]interface{}{
			"code":   "0",
			"errmsg": "创建成功",
		}

	}

	c.ServeJSON()

}

/**
获取当前用户详情
*/

// @router /admin/tag/edit/?:key [get] 后台 tag列表
func (c *AdminTagController) GatTagInfo() {

	var id string

	c.Ctx.Input.Bind(&id, "id")

	ids, _ := strconv.Atoi(id)

	if ids == 0 {

		c.Abort("500")

	}

	tag, err := models.QueryTagFirst(ids)

	if !gorm.IsRecordNotFoundError(err) {

		c.Data["tag"] = map[string]interface{}{
			"tag_name": tag.Tag_name,
			"status":   tag.Is_status,
			"tid":      tag.ID,
		}
	}

	c.TplName = "admin/tag/edit.html"
}

/**
更新标签信息
*/
// @router /admin/tag/update [post] 后台 tag列表
func (c *AdminTagController) SetTagInfo() {

	tags := c.GetString("tag_name")

	id := c.GetString("tid")

	tagStatus := c.GetString("status")

	if len(tags) < 0 || id == "" {

		c.Data["json"] = map[string]interface{}{
			"code":   "1003",
			"errmsg": "参数不完整，请输入签名",
		}

	} else {

		var tag models.LiteTag

		ids, _ := strconv.Atoi(id)

		tag.ID = uint(ids)

		tag.Tag_name = tags

		tag.Is_status, _ = strconv.Atoi(tagStatus)

		models.UpdateTagFirst(tag)

		c.Data["json"] = map[string]interface{}{
			"code":   "0",
			"errmsg": "更新成功",
		}

	}

	c.ServeJSON()
}

/**
批量处理字符串
*/

func (c *AdminTagController) stringFromData(data string) []string {

	var arr []string

	arr = make([]string, 0)

	arr = strings.Split(data, "\n")

	return arr

}
