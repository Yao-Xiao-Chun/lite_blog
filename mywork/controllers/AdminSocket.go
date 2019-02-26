package controllers

import (
	"log"
	"github.com/gorilla/websocket"
	"fmt"

	"github.com/astaxie/beego/logs"
	"encoding/json"
	"net/http"
)

/**
	websocket 推送相关代码
 */
type PushSocketController struct {

	AdminBaseController
}

type PushData struct {
	Message string `json:"message"`
	Count int	`json:"count"`
	Token string `json:"token"`
	Account string `json:"account"`
	Code int 	`json:"code"`
}

/**
	设置全局调用
 */
var (
	clients   = make(map[*websocket.Conn]bool)

	broadcast = make(chan PushData)
)

// 初始化调用
func init() {

	go handleMessages()
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024, //读取存储空间带线
	WriteBufferSize: 1024, //设置写入大小
	CheckOrigin: func(r *http.Request) bool {
		return true
	},//允许跨域
}

/**
	请求的地址 前台访问确定链接
 */
func (this *PushSocketController) Get() {

	ws, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil) //beego 框架中的接收 完成http的应答

	if err != nil {
		log.Fatal(err)
	}
	 //defer ws.Close()

	clients[ws] = true

	//不断的广播发送到页面上
	for {

		_, p, err := ws.ReadMessage() //获取前端提交的值

		datas := string(p[:]) //返回的结果值 json格式

		res,errs := toArray(datas)

		if res.Code == 200{ //此状态代表着登录
			if !checkOauth(res){
				//写入日志
				this.ReadLog("用户："+res.Account+"连接socket错误任务:",4)
				delete(clients, ws)
				goto ERR
				break
			}
		}

		//断开连接，那么需要移除当前的推送
		if err != nil || errs != nil{

			fmt.Print(err)
			delete(clients, ws)
			goto ERR
			break
		}

	}

	ERR:
		//关闭连接
		ws.Close()
}

/**
	定时检查数据库是否存在推送
 */
func TaskPushData(str string)  error{

	if str == ""{
		str = "您有新的消息提醒"
	}

	//模拟查询到了数据
	msg := PushData{Message: str,Account:"949656336@qq.com",Token:"abcdefghjklmn",Count:2,Code:0}


	broadcast <- msg

	return nil
}


/**
	socket 页面发送到数据
 */
func handleMessages() {
	for {
		msg := <-broadcast

		//fmt.Println("clients len ", len(clients))

		/*if len(clients) < 1{
			break //没有连接了，退出
		}*/
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				logs.Info("client.WriteJSON error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}




/**
	转换json数据
 */
func toArray(str string)(list PushData,errs error){

	var data PushData

	err := json.Unmarshal([]byte(str),&data)

	if err != nil{

		return data,err
	}

	return data,nil

}


/**
	查询此用户的的连接权限
 */
func checkOauth(user PushData)bool{

	if user.Account == "949656336@qq.com" && user.Token == "asdfghjkl"{
		return  true
	}else{
		return false
	}

}