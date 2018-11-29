package controllers

import (
	"mywork/models"
	"mywork/syserror"
	_"github.com/astaxie/beego/logs"
	_"strings"
	"strconv"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"strings"

)

type AdminUserController struct {

	AdminBaseController
}


//处理提交过来的json格式

type Users struct{

	Title string `json:"title"`
	Account string `json:"account"`
	Password string `json:"password"`
	Email string	`json:"email"`
	File string	`json:"file"`
	Status string `json:"status"`
	Is_admin string	`json:"is_admin"`
	Create_at string `json:"create_at"`
	Title_img string `json:"title_img"`
}


// @router /admin/dologin [post] 后台 登录页面
func (this *AdminUserController)DoLogin(){

		account := this.CheckMustKey("account","账户不能为空")

		account = this.CheckEmail(account)

		pwd := this.CheckMustKey("password","密码不能为空")

		user,err := models.QueryAccountAndPwd(account,pwd)//获取查询结果

		if err != nil{

			this.ReadLog("账户:"+account+",在Ip地址为："+this.GetIP()+",进行尝试登陆，登陆失败",1)

			this.Abort500(syserror.New("登录失败",err))

		}else{
			//更新最后一次登录的ip地址
			ip:=this.Ctx.Request.RemoteAddr

			ip=ip[0:strings.LastIndex(ip, ":")]


			models.UpdateIP(ip,int(user.ID))
		}

		//设置session
		this.SetSession(SESSION_ADMIN_KEY,user)
		//写入日志
		this.ReadLog("账户:"+account+",在Ip地址为："+this.GetIP()+",进行登陆，登陆成功",1)

		//跳转页面进入后台
		this.Ctx.Redirect(302,"/admin") //登录成功重定向这个方法

}

// @router /admin/loginout [get] 后台 登录页面
func (this *AdminUserController)LoginOut() {

	//写入日志
	this.ReadLog("账户:"+this.User.Account+",在Ip地址为："+this.GetIP()+",进行登出，登出成功",1)

	this.DelSession(SESSION_ADMIN_KEY)

	//跳转页面进入后台
	this.Ctx.Redirect(302,"/admin") //登录成功重定向这个方法

}

// @router /admin/user [get] 后台 用户列表
func (this *AdminUserController) UserList(){

	num,_ := models.GetUserNum()

	this.Data["num"] = num

	this.TplName = "admin/user/list.html"
}

/**
	列表页分页代码
 */

// @router /admin/user/page/?:page [get] 后台 分页 用户列表
func (this *AdminUserController) UserPage(){

	var page string

	this.Ctx.Input.Bind(&page,"page") //绑定分页

	pages,_ := strconv.Atoi(page)

	if pages== 0{

		pages = 1
	}

	users,err :=models.QueryUserList(pages,10)

	var arrList map[int]map[string]interface{}

	var arr map[string]interface{}

	if err != nil{

		this.Data["json"] = map[string]interface{}{
			"code":"1001",
			"msg":"非法操作",
		}
	}else{

		arrList = make(map[int]map[string]interface{},10)

		for key,val := range users{

			arr = make(map[string]interface{})

			arr["created_at"] = val.CreatedAt.Format("2006-01-02") //创建时间
			arr["updated_at"] = val.UpdatedAt.Format("2006-01-02") //更新时间
			arr["uid"] = val.ID
			arr["nickname"] = val.Nikename//昵称
			arr["head_img"] = val.Head_img//头像
			arr["last_login_ip"] = val.Last_login_ip//最后一次登录ip
			arr["status"] = val.Status
			arr["is_admin"] = val.Is_admin
			arr["account"] = val.Account
			arr["email"] = val.Email
			arr["id"] = val.ID

			arrList[key] = arr
		}

		this.Data["json"] = map[string]interface{}{
			"code":"0",
			"data":arrList,
			"msg":"请求成功",
		}
	}

	this.ServeJSON()



}


/**
	处理创建用户方法
 */

// @router /admin/user/create [post] 后台 创建用户
func (this *AdminUserController) CreateUser(){

	data := this.CheckMustKey("data","接受参数错误") //获取参数 json格式

	user := Users{}

	err := json.Unmarshal([]byte(data),&user) //值进行绑定

	if err != nil{

		this.Data["json"] = map[string]interface{}{
			"code":"1004",
			"errmsg":"json数据请求错误！",
		}

	}else{

		//数据验证
		acc := this.CheckEmail(user.Account)

		if acc == ""{
			this.Data["json"] = map[string]interface{}{
				"code":"1005",
				"errmsg":"登录账户不合法！",
			}
		}else{

			users,err := models.GetIsAccount(acc) //检测当前用户是否存在

			if !gorm.IsRecordNotFoundError(err){

				this.Data["json"] = map[string]interface{}{
					"code":"1006",
					"errmsg":"此账户已存在，请勿重复添加！",
				}

			}else{

				users.Account = user.Account

				users.Status,_ = strconv.Atoi(user.Status)//转换为整型

				users.Uid = int(this.User.ID) //获取登录的用

				users.Is_admin,_ = strconv.Atoi(user.Is_admin)//是否是管理员 允许后台登录

				users.Email = user.Email

				users.Password = user.Password

				users.Head_img = user.Title_img //头像地址

				users.Nikename = user.Title //昵称

				models.CreateUser(&users)//创建


				this.Data["json"] = map[string]interface{}{
					"code":"0",
					"errmsg":"创建账户成功",
				}

			}

		}

	}

	this.ServeJSON()
}



// @router /admin/user/add [get] 后台首页
func (this *AdminArticleController) GetUser(){

	this.TplName = "admin/user/add.html"
}

/**
	后台删除用户
 */
// @router /admin/user/del/?:id [get] 删除用户
 func (this *AdminUserController) DelUser(){

 	var id int

 	this.Ctx.Input.Bind(&id,"id")

 	if id == 0 {

 		this.Abort("500")
	}

	_,err := models.DelUser(id)

	if err == nil{

		this.Data["json"] = map[string]interface{}{
			"code":"0",
			"errmsg":"删除成功",
		}
	}else{
		this.Data["json"] = map[string]interface{}{
			"code":"1006",
			"errmsg":"删除失败！",
		}
	}

	this.ServeJSON()

 }


 /**
 	后台用户编辑页面
  */
// @router /admin/user/edit/?:id [get] 编辑用户
  func (this *AdminUserController) EditUser(){

  	var id int

  	this.Ctx.Input.Bind(&id,"id") //绑定获取的数值

  	user,err := models.FindUser(id);

  	if err != nil{

	}else{

		this.Data["userinfo"] = map[string]interface{}{
			"title":user.Nikename,
			"status":user.Status,
			"is_admin":user.Is_admin,
			"head_img":user.Head_img,
			"account":user.Account,
			"email":user.Email,
			"password":user.Password,
			"id":user.ID,
		}

	}

	this.TplName = "admin/user/edit.html"
  }



  /**
  	后台编辑页面提交功能
   */
// @router /admin/user/update [post] 编辑用户提交
   func (this *AdminUserController) EditUserData(){

	   data := this.CheckMustKey("data","接受参数错误") //获取参数 json格式

	   user := Users{}

	   err := json.Unmarshal([]byte(data),&user) //值进行绑定
		//修改后台登录
	   if user.Is_admin == ""{

	   		user.Is_admin = "0";
	   }
	   if err != nil{

		   this.Data["json"] = map[string]interface{}{
			   "code":"1004",
			   "errmsg":"json数据请求错误！",
		   }

	   }else{

		   //数据验证
		   acc := this.CheckEmail(user.Account)

		   if acc == ""{
			   this.Data["json"] = map[string]interface{}{
				   "code":"1005",
				   "errmsg":"登录账户不合法！",
			   }
		   }else{

			   users,_ := models.GetIsAccount(acc) //检测当前用户是否存在

				   //users.Account = user.Account //不准修改账号

				   users.Status,_ = strconv.Atoi(user.Status)//转换为整型

				   users.Is_admin,_ = strconv.Atoi(user.Is_admin)//是否是管理员 允许后台登录

				   users.Email = user.Email

				   users.Password = user.Password

				   users.Head_img = user.Title_img //头像地址

				   users.Nikename = user.Title //昵称

				   ids := users.ID

				   models.EditUser(ids,users)


				   this.Data["json"] = map[string]interface{}{
					   "code":"0",
					   "errmsg":"修改用户成功",
				   }

		   }

	   }


   	this.ServeJSON()
   }