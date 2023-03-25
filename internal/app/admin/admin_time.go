package admin

import (
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	"mywork/internal/app/common/dto"
)

/**
时间线处理控制器；
*/
type AdminTimeController struct {
	AdminBaseController
}

// @router /admin/timepost/:key [post] Time新增处理方法
func (c *AdminTimeController) AddTime() {
	//获取key
	token := c.Ctx.Input.Param(":key")

	//查询token是否存在
	line, err := dto.QueryToken(token)

	//判断是否查找到记录
	if !gorm.IsRecordNotFoundError(err) {

		c.Data["json"] = map[string]interface{}{
			"code": 2,
			"msg":  "token创建失败",
		}
	} else {

		//获取标题
		title := c.CheckMustKey("title", "没有获取到正确的标题")

		content := c.GetString("content")

		status := c.CheckMustKey("status", "请选择启用状态")

		logs.Info(status)

		line.Title = title

		line.Content = content

		line.Uid = int(c.User.ID) //用户id

		line.Token = token

		line.Status = status

		dto.CreateTimeLine(line)

		c.Data["json"] = map[string]interface{}{
			"code": 0,
			"msg":  "创建成功",
		}

	}

	c.ServeJSON()
}

/**
删除 时间线操作
*/

// @router /admin/deltime/:key [get] Time新增处理方法
func (c *AdminTimeController) DelTime() {

	tid := c.GetString("id")

	token := c.GetString("token")

	if tid == "" || token == "" {

		c.Data["json"] = map[string]interface{}{
			"code": 1002,
			"msg":  "数据丢失",
		}

	} else {

		_, err := dto.SetDelTimes(tid, token)

		if err == nil {

			c.Data["json"] = map[string]interface{}{
				"code": 0,
				"msg":  "删除成功",
			}
		} else {

			c.Data["json"] = map[string]interface{}{
				"code": 1002,

				"msg": "删除失败",
			}

		}

	}

	c.ServeJSON()
}

/**
修改时间线页面
*/
// @router /admin/timeinfo/?:id [get]
func (c *AdminTimeController) GetTimeInfo() {

	var tid int
	c.Ctx.Input.Bind(&tid, "id")

	var token string

	c.Ctx.Input.Bind(&token, "token")

	line, _ := dto.GetTileLineFind(tid, token)

	data := map[string]interface{}{

		"title":   line.Title,
		"content": line.Content,
		"id":      line.ID,
		"token":   line.Token,
		"status":  line.Status,
	}
	c.Data["lines"] = data

	c.TplName = "admin/news/time.html"

}

/**
修改提交表单
*/

//@router /admin/formtime/:key [post]
func (c *AdminTimeController) SetTimeInfo() {
	//获取key
	token := c.Ctx.Input.Param(":key")

	//查询token是否存在
	line, err := dto.QueryToken(token)

	//判断是否查找到记录
	if !gorm.IsRecordNotFoundError(err) {
		//获取标题
		title := c.CheckMustKey("title", "没有获取到正确的标题")

		content := c.GetString("content")

		id := c.GetString("id")

		status := c.CheckMustKey("status", "请选择启用状态")

		line.Status = status

		line.Content = content

		line.Title = title

		dto.SetTimeInfo(id, line)

		c.Data["json"] = map[string]interface{}{
			"code": 0,
			"msg":  "修改成功",
		}

	} else {
		c.Data["json"] = map[string]interface{}{
			"code": 2,
			"msg":  "数据不存在，禁止操作",
		}
	}

	c.ServeJSON()
}
