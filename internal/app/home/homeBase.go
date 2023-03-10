package controllers

import (
	"math/rand"
	"mywork/internal/app/admin"
	"mywork/internal/app/common"
	"mywork/models"
	"time"
)

/**
	前台基础控制器
	@param
    @Auther 姚春
	@Date 2018-08-25
*/
type HomeBaseController struct {
	common.BaseController
}

type cat struct {
	ID          int
	Menu_parent int
	Menu_name   string
	Level       int
	Menu_status int
	sort        int
	Children    []*admin.Cat ``
}

/**
  这个函数主要是为了用户扩展用的，这个函数会在下面定义的这些 Method 方法之前执行，用户可以重写这个函数实现类似用户验证之类
*/

func (this *HomeBaseController) Prepare() {

	this.Data["Path"] = this.Ctx.Request.RequestURI //获取当前中的url

	menuData := this.GetMenuList(false) //获取列表

	this.Data["Menus"] = menuData

}

func (this *HomeBaseController) GetMenuList(flag bool) (data map[int]map[string]interface{}) {

	var menu []models.LiteAdminMenu
	//是获取所有的还是获取选择的
	if flag {
		menu, _ = models.GetAll()
	} else {
		menu, _ = models.GetMenuInfo() //根据条件获取
	}

	var list []admin.Cat

	list = make([]admin.Cat, len(menu))

	/*var listA []*Cat

	listA = make([]*Cat,len(menu))*/

	for i, v := range menu {

		list[i].ID = int(v.ID)
		list[i].Menu_name = v.Menu_name
		list[i].Menu_parent = v.Menu_parent
		list[i].Level = v.Menu_level
		list[i].sort = v.Menu_roule
		list[i].Menu_status = v.Menu_status
		//listA[i] = &list[i]
	}

	menus := admin.superCategory(list, 0)

	var menuData map[int]map[string]interface{}

	menuData = make(map[int]map[string]interface{})

	var menuArr map[string]interface{}

	for key, val := range menus {

		menuArr = make(map[string]interface{})

		var str string

		/*if val.Level > 0{

			str = setSpace(val.Level)
		}*/

		menuArr["name"] = str + val.Menu_name

		menuArr["id"] = val.ID

		menuArr["pid"] = val.Menu_parent

		menuArr["status"] = val.Menu_status
		menuArr["sort"] = val.sort

		menuData[key] = menuArr

	}

	return menuData
}

/**
	生成随机字符串
	@param int 长度
    @return string
*/
func (this *HomeBaseController) GetRandomString(lens int) string {

	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	bytes := []byte(str)

	result := []byte{}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < lens; i++ {

		result = append(result, bytes[r.Intn(len(bytes))])
	}

	return string(result)
}
