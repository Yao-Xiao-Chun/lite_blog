package admin

import (
	"encoding/json"
	"fmt"
	_ "github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	"lite_blog/internal/app/common/dto"
	"lite_blog/internal/pkg/entity"
	"lite_blog/syserror"
	"strconv"
	"strings"
	_ "strings"
)

type AdminUserController struct {
	AdminBaseController
}

//处理提交过来的json格式

// @router /admin/dologin [post] 后台 登录页面
func (c *AdminUserController) DoLogin() {

	account := c.CheckMustKey("account", "账户不能为空")

	account = c.CheckEmail(account)

	pwd := c.CheckMustKey("password", "密码不能为空")

	pwd = c.SetMd5Pwd(pwd)

	user, err := dto.QueryAccountAndPwd(account, pwd) //获取查询结果
	fmt.Println(user.Account)
	if err != nil {

		c.ReadLog("账户:"+account+",在Ip地址为："+c.GetIP()+",进行尝试登陆，登陆失败", 1)

		c.Abort500(syserror.New("登录失败", err))

	} else {
		//更新最后一次登录的ip地址
		ip := c.Ctx.Request.RemoteAddr

		ip = ip[0:strings.LastIndex(ip, ":")]

		dto.UpdateIP(ip, int(user.ID))
	}

	//设置session
	c.SetSession(SESSION_ADMIN_KEY, user)
	//写入日志
	c.ReadLog("账户:"+account+",在Ip地址为："+c.GetIP()+",进行登陆，登陆成功", 1)

	TaskPushData("账户:" + account + "登陆上线了") //提示登录的用户
	//跳转页面进入后台
	c.Ctx.Redirect(302, "/admin") //登录成功重定向这个方法

}

// @router /admin/loginout [get] 后台 登录页面
func (c *AdminUserController) LoginOut() {

	//写入日志
	c.ReadLog("账户:"+c.User.Account+",在Ip地址为："+c.GetIP()+",进行登出，登出成功", 1)

	c.DelSession(SESSION_ADMIN_KEY)

	TaskPushData("账户:" + c.User.Account + "退出了登录") //提示登录的用户
	//跳转页面进入后台
	c.Ctx.Redirect(302, "/admin") //登录成功重定向这个方法

}

// @router /admin/user [get] 后台 用户列表
func (c *AdminUserController) UserList() {

	num, _ := dto.GetUserNum()

	c.Data["num"] = num

	c.TplName = "admin/user/list.html"
}

/**
列表页分页代码
*/

// @router /admin/user/page/?:page [get] 后台 分页 用户列表
func (c *AdminUserController) UserPage() {

	var page string

	c.Ctx.Input.Bind(&page, "page") //绑定分页

	pages, _ := strconv.Atoi(page)

	if pages == 0 {

		pages = 1
	}

	users, err := dto.QueryUserList(pages, 10)

	var arrList map[int]map[string]interface{}

	var arr map[string]interface{}

	if err != nil {

		c.Data["json"] = map[string]interface{}{
			"code": "1001",
			"msg":  "非法操作",
		}
	} else {

		arrList = make(map[int]map[string]interface{}, 10)

		for key, val := range users {

			arr = make(map[string]interface{})

			arr["created_at"] = val.CreatedAt.Format("2006-01-02") //创建时间
			arr["updated_at"] = val.UpdatedAt.Format("2006-01-02") //更新时间
			arr["uid"] = val.ID
			arr["nickname"] = val.Nikename           //昵称
			arr["head_img"] = val.Head_img           //头像
			arr["last_login_ip"] = val.Last_login_ip //最后一次登录ip
			arr["status"] = val.Status
			arr["is_admin"] = val.Is_admin
			arr["account"] = val.Account
			arr["email"] = val.Email
			arr["id"] = val.ID

			arrList[key] = arr
		}

		c.Data["json"] = map[string]interface{}{
			"code": "0",
			"data": arrList,
			"msg":  "请求成功",
		}
	}

	c.ServeJSON()

}

/**
处理创建用户方法
*/

// @router /admin/user/create [post] 后台 创建用户
func (c *AdminUserController) CreateUser() {

	data := c.CheckMustKey("data", "接受参数错误") //获取参数 json格式

	user := entity.Users{}

	err := json.Unmarshal([]byte(data), &user) //值进行绑定

	if err != nil {

		c.Data["json"] = map[string]interface{}{
			"code":   "1004",
			"errmsg": "json数据请求错误！",
		}

	} else {

		//数据验证
		acc := c.CheckEmail(user.Account)

		if acc == "" {
			c.Data["json"] = map[string]interface{}{
				"code":   "1005",
				"errmsg": "登录账户不合法！",
			}
		} else {

			users, err := dto.GetIsAccount(acc) //检测当前用户是否存在

			if !gorm.IsRecordNotFoundError(err) {

				c.Data["json"] = map[string]interface{}{
					"code":   "1006",
					"errmsg": "此账户已存在，请勿重复添加！",
				}

			} else {

				users.Account = user.Account

				users.Status, _ = strconv.Atoi(user.Status) //转换为整型

				users.Uid = int(c.User.ID) //获取登录的用

				users.Is_admin, _ = strconv.Atoi(user.IsAdmin) //是否是管理员 允许后台登录

				users.Email = user.Email

				users.Password = c.SetMd5Pwd(user.Password)

				users.Head_img = user.TitleImg //头像地址

				users.Nikename = user.Title //昵称

				dto.CreateUser(&users) //创建

				//写入日志
				c.ReadLog("账号:"+c.User.Nikename+" 操作：创建用户:'"+user.Account+"',状态：成功", 2)
				c.Data["json"] = map[string]interface{}{
					"code":   "0",
					"errmsg": "创建账户成功",
				}

			}

		}

	}

	c.ServeJSON()
}

// @router /admin/user/add [get] 后台首页
func (c *ArticleController) GetUser() {

	c.TplName = "admin/user/add.html"
}

/**
后台删除用户
*/
// @router /admin/user/del/?:id [get] 删除用户
func (c *AdminUserController) DelUser() {

	var id int

	c.Ctx.Input.Bind(&id, "id")

	if id == 0 {

		c.Abort("500")
	}

	_, err := dto.DelUser(id)

	if err == nil {

		//写入日志
		c.ReadLog("账号:"+c.User.Nikename+" 操作：删除用户id:'"+string(id)+"',状态：成功", 2)

		c.Data["json"] = map[string]interface{}{
			"code":   "0",
			"errmsg": "删除成功",
		}
	} else {

		//写入日志
		c.ReadLog("账号:"+c.User.Nikename+" 操作：删除用户id:'"+string(id)+"',状态：失败", 2)

		c.Data["json"] = map[string]interface{}{
			"code":   "1006",
			"errmsg": "删除失败！",
		}
	}

	c.ServeJSON()

}

/**
后台用户编辑页面
*/
// @router /admin/user/edit/?:id [get] 编辑用户
func (c *AdminUserController) EditUser() {

	var id int

	c.Ctx.Input.Bind(&id, "id") //绑定获取的数值

	user, err := dto.FindUser(id)

	if err != nil {

	} else {

		c.Data["userinfo"] = map[string]interface{}{
			"title":    user.Nikename,
			"status":   user.Status,
			"is_admin": user.Is_admin,
			"head_img": user.Head_img,
			"account":  user.Account,
			"email":    user.Email,
			"password": user.Password,
			"id":       user.ID,
		}

	}

	c.TplName = "admin/user/edit.html"
}

/**
后台编辑页面提交功能
*/
// @router /admin/user/update [post] 编辑用户提交
func (c *AdminUserController) EditUserData() {

	data := c.CheckMustKey("data", "接受参数错误") //获取参数 json格式

	user := entity.Users{}

	err := json.Unmarshal([]byte(data), &user) //值进行绑定
	//修改后台登录
	if user.IsAdmin == "" {

		user.IsAdmin = "0"
	}
	if err != nil {

		c.Data["json"] = map[string]interface{}{
			"code":   "1004",
			"errmsg": "json数据请求错误！",
		}

	} else {

		//数据验证
		acc := c.CheckEmail(user.Account)

		if acc == "" {
			c.Data["json"] = map[string]interface{}{
				"code":   "1005",
				"errmsg": "登录账户不合法！",
			}
		} else {

			users, _ := dto.GetIsAccount(acc) //检测当前用户是否存在

			//users.Account = user.Account //不准修改账号

			users.Status, _ = strconv.Atoi(user.Status) //转换为整型

			users.Is_admin, _ = strconv.Atoi(user.IsAdmin) //是否是管理员 允许后台登录

			users.Email = user.Email

			if user.Password != "" {

				users.Password = c.SetMd5Pwd(user.Password)
			}

			users.Head_img = user.TitleImg //头像地址

			users.Nikename = user.Title //昵称

			ids := users.ID

			dto.EditUser(ids, users)

			//写入日志
			c.ReadLog("账号:"+c.User.Nikename+" 操作：编辑用户:'"+user.Account+"',状态：成功", 2)

			c.Data["json"] = map[string]interface{}{
				"code":   "0",
				"errmsg": "修改用户成功",
			}

		}

	}

	c.ServeJSON()
}
