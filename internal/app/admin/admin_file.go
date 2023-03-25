package admin

import (
	"lite_blog/internal/app/common/dto"
	models "lite_blog/internal/app/common/dto"
	"lite_blog/internal/pkg/entity"
)

/**
后台上传小说功能，自己的使用
*/

type FileController struct {
	UploadController
}

func (c *FileController) FileIndex() {

	c.Data["num"], _ = dto.CountFictionNum()

	c.TplName = "admin/fiction/index.html"
}

/**
禁止下载
*/
func (c *FileController) SetFictionStatus() {

	var id int

	c.Ctx.Input.Bind(&id, "id")

	models.UpdateFictionStatus(id)

	c.Data["json"] = map[string]string{
		"code": "0",
		"msg":  "禁止下载成功",
	}

	c.ServeJSON()
}

/**
分页数据
*/
func (c *FileController) FilePage() {

	var page, size int

	c.Ctx.Input.Bind(&page, "page")

	c.Ctx.Input.Bind(&size, "size")

	data, err := models.FindFictionData(page, 10)

	if err != nil {

		c.Data["json"] = map[string]string{
			"code": "1002",
			"msg":  "消息错误",
		}
	} else {
		var res []entity.FictionList

		res = make([]entity.FictionList, 0)

		for _, val := range data {

			ids, _ := models.FictionOperation(int(val.ID))

			dataRes := entity.FictionList{LiteFiction: val, Times: val.CreatedAt.Format("2006-01-02 13:04:05"), DownloadNum: ids.DownloadNum}

			if val.Tags != "" {

				res := models.FictionAndTag(val.Tags)

				dataRes.Tags = res.Fname

			}

			res = append(res, dataRes)
		}

		c.Data["json"] = map[string]interface{}{
			"code": "0",
			"data": res,
			"msg":  "查询成功",
		}

	}

	c.ServeJSON()
}

/**
日志小说首页
*/
func (c *FileController) FictionLog() {

	c.Data["num"], _ = models.CountFictionLog()

	c.TplName = "admin/log/fictionlog.html"
}

/**
小说日志查询
*/
func (c *FileController) FictionLogPage() {

	var page, size int

	c.Ctx.Input.Bind(&page, "page")

	c.Ctx.Input.Bind(&size, "size")

	data, err := models.GetFictionLogList(page, 10)

	if err != nil {

		c.Data["json"] = map[string]string{
			"code": "1002",
			"msg":  "消息错误",
		}
	} else {

		c.Data["json"] = map[string]interface{}{
			"code": "0",
			"data": data,
			"msg":  "查询成功",
		}

	}

	c.ServeJSON()
}

/**
加入黑名单
*/
func (c *FileController) FictionBanned() {

	var ip string

	c.Ctx.Input.Bind(&ip, "ip")

	if ip == "" {

		c.Data["json"] = map[string]string{
			"code": "1002",
			"msg":  "获取失败，error",
		}
	} else {

		//查询是否存在

		num, _ := models.QueryBanned(ip)

		if num > 0 {

			c.Data["json"] = map[string]string{
				"code": "1002",
				"msg":  "已存在黑名单中，请勿重复添加",
			}
		} else {

			models.AddBanned(ip)

			c.Data["json"] = map[string]string{
				"code": "0",
				"msg":  "添加成功",
			}
		}

	}

	c.ServeJSON()
}
