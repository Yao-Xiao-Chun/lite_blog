package common

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/toolbox"
	"mywork/bins"
	"mywork/internal/app/admin"
	"mywork/models"
	"strings"
	"time"
)

/**
定时任务控制器
*/
type CronTabController struct {
	admin.AdminBaseController
}

/**
获取新增参数结构体
*/
type TaskData struct {
	ID           int    `form:"-"`
	TaskName     string `form:"task_name"`
	TaskDescript string `form:"task_descript"`
	MonthStart   string `form:"month_start"`
	MonthEnd     string `form:"month_end"`
	WeekStart    string `form:"week_start"`
	WeekEnd      string `form:"week_end"`
	Time         string `form:"time"`
	Task         string `form:"task"`
}

// @router /admin/crontab/index [get] 定时任务列表
func (c *CronTabController) Task() {

	c.Data["num"], _ = models.CountPage()
	c.TplName = "admin/crontab/index.html"
}

// @router /admin/crontab/add [get] 创建页面
func (c *CronTabController) TaskAdd() {

	c.TplName = "admin/crontab/add.html"
}

// @router /admin/crontab/add [post] 创建
func (c *CronTabController) Create() {

	//获取传输的参数
	var data TaskData

	if err := c.ParseForm(&data); err != nil {
		c.Data["json"] = map[string]interface{}{
			"code": "1003",
			"msg":  err,
		}
	} else {

		var taskData models.LiteCrontab

		taskData.TaskName = data.TaskName
		taskData.Descript = data.TaskDescript
		taskData.Frequency = c.toTaskStr(data)
		taskData.Status = 0
		taskData.CreateName = c.User.Nikename
		taskData.TaskId = data.Task

		models.TaskAdd(taskData)

		c.Data["json"] = map[string]interface{}{
			"code": "0",
			"msg":  "创建计划任务成功",
		}
	}

	c.ServeJSON()
}

/**
页面获取
*/
// @router /admin/crontab/page/?:key [get]
func (c *CronTabController) GetTaskPage() {

	var page int

	var size string

	c.Ctx.Input.Bind(&page, "page")

	c.Ctx.Input.Bind(&size, "size")

	result, err := models.FindTask(page, 10)

	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"code": "1003",
			"msg":  err,
		}
	} else {

		c.Data["json"] = map[string]interface{}{
			"code": "0",
			"msg":  "获取成功",
			"data": result,
		}
	}

	c.ServeJSON()
}

/**
删除
*/
// @router /admin/crontab/delete/?:key [get]
func (c *CronTabController) CrontabDelete() {

	var taskId int

	c.Ctx.Input.Bind(&taskId, "id")

	//检查次id是否在执行状态
	resStatus, _ := models.FindInfoTask(taskId)

	if resStatus.Status == 0 {
		//可以删除
		models.DeleteTask(taskId)
		c.Data["json"] = map[string]interface{}{
			"code": "0",
			"msg":  "删除成功",
		}

	} else {
		//启用状态，停止后可以删除
		c.Data["json"] = map[string]interface{}{
			"code": "1002",
			"msg":  "请停止该定时任务后，进行删除",
		}
	}
	c.ServeJSON()
}

/**
启用任务
*/
// @router /admin/crontab/startTask/?:key [get] 开通任务计划
func (c *CronTabController) StartTask() {

	var taskId int

	c.Ctx.Input.Bind(&taskId, "task_id")
	//查询要启用的任务
	result, _ := models.FindInfoTask(taskId)

	c.setTask(result.Frequency, result.TaskName, result.TaskId) //设定任务

	//更新后台数据库
	result.Status = 1
	models.UpdateTaskStatus(result)
	//写入日志
	c.ReadLog("用户："+c.User.Nikename+"开启定时任务:"+result.TaskName+"。", 3)
	c.Data["json"] = map[string]interface{}{
		"code": "0",
		"msg":  "创建定时任务成功",
	}

	c.ServeJSON()
}

/**
停用任务
*/
// @router /admin/crontab/stopTask/?:key [get] 开通任务计划
func (c *CronTabController) StopTask() {

	var taskId int

	c.Ctx.Input.Bind(&taskId, "task_id")

	//查询要启用的任务
	result, _ := models.FindInfoTask(taskId)

	DeleteTask(result.TaskName)

	//更新后台数据库
	result.Status = 0

	models.UpdateTaskStatus(result)

	c.ReadLog("用户："+c.User.Nikename+"停用定时任务:"+result.TaskName+"。", 3)

	c.Data["json"] = map[string]interface{}{
		"code": "0",
		"msg":  "停止定时任务：" + result.TaskName,
	}

	c.ServeJSON()
}

/**
停止全部任务
*/
// @router /admin/crontab/allstop [get] 停止所有任务
func (c *CronTabController) StopAllTask() {

	flag := c.GetRunTask()

	if flag {
		toolbox.StopTask() //停止所有任务执行
		c.Data["json"] = map[string]string{
			"code": "0",
			"msg":  "所有定时任务停止成功",
		}
	} else {
		c.Data["json"] = map[string]string{
			"code": "1003",
			"msg":  "停止失败",
		}
	}

	c.ServeJSON()
}

/**
设置定时任务
* * * * * *
*/
//前6个字段分别表示：
//       秒钟：0-59
//       分钟：0-59
//       小时：1-23
//       日期：1-31
//       月份：1-12
//       星期：0-6（0 表示周日）
//还可以用一些特殊符号：
//       *： 表示任何时刻
//       ,：　表示分割，如第三段里：2,4，表示 2 点和 4 点执行
//　　    －：表示一个段，如第三端里： 1-5，就表示 1 到 5 点
//       /n : 表示每个n的单位执行一次，如第三段里，* /1, 就表示每隔 1 个小时执行一次命令。也可以写成1-23/1.
/////////////////////////////////////////////////////////
//  0/30 * * * * *                        每 30 秒 执行
//  0 43 21 * * *                         21:43 执行
//  0 15 05 * * * 　　                     05:15 执行
//  0 0 17 * * *                          17:00 执行
//  0 0 17 * * 1                          每周一的 17:00 执行
//  0 0,10 17 * * 0,2,3                   每周日,周二,周三的 17:00和 17:10 执行
//  0 0-10 17 1 * *                       毎月1日从 17:00 到 7:10 毎隔 1 分钟 执行
//  0 0 0 1,15 * 1                        毎月1日和 15 日和 一日的 0:00 执行
//  0 42 4 1 * * 　 　                     毎月1日的 4:42 分 执行
//  0 0 21 * * 1-6　　                     周一到周六 21:00 执行
//  0 0,10,20,30,40,50 * * * *　           每隔 10 分 执行
//  0 */10 * * * * 　　　　　　              每隔 10 分 执行
//  0 * 1 * * *　　　　　　　　               从 1:0 到 1:59 每隔 1 分钟 执行
//  0 0 1 * * *　　　　　　　　               1:00 执行
//  0 0 */1 * * *　　　　　　　               毎时 0 分 每隔 1 小时 执行
//  0 0 * * * *　　　　　　　　               毎时 0 分 每隔 1 小时 执行
//  0 2 8-20/3 * * *　　　　　　             8:02,11:02,14:02,17:02,20:02 执行
//  0 30 5 1,15 * *　　　　　　              1 日 和 15 日的 5:30 执

func (c *CronTabController) setTask(taskTime, name, taskId string) {

	logs.Info(taskTime, taskId)

	if taskTime == "" {

		taskTime = "0/10 * * * * *"
	}

	if name == "" {

		name = "myTask"
	}

	tk := toolbox.NewTask(name, taskTime, func() error {

		switch taskId {
		case "0":
			return bins.StartYx()

		case "1":

			return bins.StartYeWu()
		case "2":

			return admin.TaskPushData("")
		default:

			return ToStr()
		}
	})

	toolbox.AddTask(name, tk)

	tk.SetNext(time.Now()) //动态创建多个任务执行的时候，需要执行此次方法，否则会出现其他任务

	logs.Info("创建任务成功....")

	toolbox.StartTask() //开始执行当前任务

	logs.Info("开始执行任务")
}

/**
输出数据测试demo
*/
func ToStr() error {

	logs.Info("我在发送营销短信....")

	return nil
}

/**
删除定时任务
*/
func DeleteTask(name string) {

	if name == "" {

		name = "myTask"
	}

	//toolbox.StopTask()

	toolbox.DeleteTask(name)

	logs.Info("删除任务成功")

}

/**
定时执行系统
@param struct 上传参数结构体
@return string 执行任务格式
*/
func (c *CronTabController) toTaskStr(data TaskData) string {

	var task, s, m, d, date, month, week string

	s, m, d, date, month, week = "*", "*", "*", "*", "*", "*"

	//判断月份是否为空
	if data.MonthStart != "" {
		//设置月份开始时间
		infoStart := strings.Split(data.MonthStart, "-")

		if data.MonthEnd != "" {

			infoEnd := strings.Split(data.MonthStart, "-")

			//月份
			if infoStart[0] == infoEnd[0] {

				month = infoStart[0]

			} else {

				month = infoStart[0] + "-" + infoEnd[0] //设定开始月份和结束月份 例如 11-12
			}
			//天
			if infoStart[1] == infoEnd[1] {
				date = infoStart[1]
			} else {
				date = infoStart[1] + "-" + infoEnd[1] //设定开始天和结束天 例如 1-31
			}
		} else {

			month = infoStart[0] //月份

			date = infoStart[1] //天
		}
	}

	//判断星期
	if data.WeekStart != "" {

		if data.WeekEnd != "" {
			week = data.WeekStart + "-" + data.WeekEnd
		} else {
			week = data.WeekStart
		}

	}

	//判断时间
	if data.Time != "" {
		timeList := strings.Split(data.Time, ":")
		s = timeList[2] //秒
		m = timeList[1] //分
		d = timeList[0] //时

		if m == "00" && d == "00" && s != "00" {

			s = "0/" + s //拼接定时任务 0/10 0/20 , 10s执行一次，20s执行一次

			m = "*"

			d = "*"
		}
	}

	task = s + " " + m + " " + d + " " + date + " " + month + " " + week //   0/1 * * * * * 任务列表

	return task
}

/**
获取系统正在执行的定时任务
*/
func (c *CronTabController) GetRunTask() bool {

	list, err := models.RunTask()

	if err != nil {
		return false
	}

	for _, val := range list {

		val.Status = 0

		models.UpdateTaskStatus(val)

		//写入日志
		c.ReadLog("用户："+c.User.Nikename+"停止定时任务:"+val.TaskName+"。", 3)
	}

	return true
}
