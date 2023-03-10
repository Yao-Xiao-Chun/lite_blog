package admin

import (
	"mywork/internal/app/common"
	"mywork/models"
)

/**
后台上传小说功能，自己的使用
*/

type AdminFileController struct {
	common.UploadController
}

type FictionList struct {
	models.LiteFiction
	TagsName    string
	Times       string
	DownloadNum int
}

func (this *AdminFileController) FileIndex() {

	this.Data["num"], _ = models.CountFictionNum()

	this.TplName = "admin/fiction/index.html"
}

/**
禁止下载
*/
func (this *AdminFileController) SetFictionStatus() {

	var id int

	this.Ctx.Input.Bind(&id, "id")

	models.UpdateFictionStatus(id)

	this.Data["json"] = map[string]string{
		"code": "0",
		"msg":  "禁止下载成功",
	}

	this.ServeJSON()
}

/**
分页数据
*/
func (this *AdminFileController) FilePage() {

	var page, size int

	this.Ctx.Input.Bind(&page, "page")

	this.Ctx.Input.Bind(&size, "size")

	data, err := models.FindFictionData(page, 10)

	if err != nil {

		this.Data["json"] = map[string]string{
			"code": "1002",
			"msg":  "消息错误",
		}
	} else {
		var res []FictionList

		res = make([]FictionList, 0)

		for _, val := range data {

			ids, _ := models.FictionOperation(int(val.ID))

			dataRes := FictionList{val, "", val.CreatedAt.Format("2006-01-02 13:04:05"), ids.DownloadNum}

			if val.Tags != "" {

				res := models.FictionAndTag(val.Tags)

				dataRes.Tags = res.Fname

			}

			res = append(res, dataRes)
		}

		this.Data["json"] = map[string]interface{}{
			"code": "0",
			"data": res,
			"msg":  "查询成功",
		}

	}

	this.ServeJSON()
}

/**
日志小说首页
*/
func (this *AdminFileController) FictionLog() {

	this.Data["num"], _ = models.CountFictionLog()

	this.TplName = "admin/log/fictionlog.html"
}

/**
小说日志查询
*/
func (this *AdminFileController) FictionLogPage() {

	var page, size int

	this.Ctx.Input.Bind(&page, "page")

	this.Ctx.Input.Bind(&size, "size")

	data, err := models.GetFictionLogList(page, 10)

	if err != nil {

		this.Data["json"] = map[string]string{
			"code": "1002",
			"msg":  "消息错误",
		}
	} else {

		this.Data["json"] = map[string]interface{}{
			"code": "0",
			"data": data,
			"msg":  "查询成功",
		}

	}

	this.ServeJSON()
}

/**
加入黑名单
*/
func (this *AdminFileController) FictionBanned() {

	var ip string

	this.Ctx.Input.Bind(&ip, "ip")

	if ip == "" {

		this.Data["json"] = map[string]string{
			"code": "1002",
			"msg":  "获取失败，error",
		}
	} else {

		//查询是否存在

		num, _ := models.QueryBanned(ip)

		if num > 0 {

			this.Data["json"] = map[string]string{
				"code": "1002",
				"msg":  "已存在黑名单中，请勿重复添加",
			}
		} else {

			models.AddBanned(ip)

			this.Data["json"] = map[string]string{
				"code": "0",
				"msg":  "添加成功",
			}
		}

	}

	this.ServeJSON()
}
