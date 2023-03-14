package home

import (
	"math/rand"
	"mywork/internal/app/common"
	"mywork/internal/pkg"
	"mywork/internal/pkg/dto"
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
	Children    []*dto.Cat ``
}

/**
  这个函数主要是为了用户扩展用的，这个函数会在下面定义的这些 Method 方法之前执行，用户可以重写这个函数实现类似用户验证之类
*/

func (c *HomeBaseController) Prepare() {

	c.Data["Path"] = c.Ctx.Request.RequestURI //获取当前中的url

	menuData := c.GetMenuList(false) //获取列表

	c.Data["Menus"] = menuData

}

func (c *HomeBaseController) GetMenuList(flag bool) (data map[int]map[string]interface{}) {

	var menu []models.LiteAdminMenu
	//是获取所有的还是获取选择的
	if flag {
		menu, _ = models.GetAll()
	} else {
		menu, _ = models.GetMenuInfo() //根据条件获取
	}

	var list []dto.Cat

	list = make([]dto.Cat, len(menu))

	/*var listA []*Cat

	listA = make([]*Cat,len(menu))*/

	for i, v := range menu {

		list[i].ID = int(v.ID)
		list[i].MenuName = v.Menu_name
		list[i].MenuParent = v.Menu_parent
		list[i].Level = v.Menu_level
		list[i].Sort = v.Menu_roule
		list[i].MenuStatus = v.Menu_status
		//listA[i] = &list[i]
	}

	menus := pkg.SuperCategory(list, 0)

	var menuData map[int]map[string]interface{}

	menuData = make(map[int]map[string]interface{})

	var menuArr map[string]interface{}

	for key, val := range menus {

		menuArr = make(map[string]interface{})

		var str string

		/*if val.Level > 0{

			str = setSpace(val.Level)
		}*/

		menuArr["name"] = str + val.MenuName

		menuArr["id"] = val.ID

		menuArr["pid"] = val.MenuParent

		menuArr["status"] = val.MenuStatus
		menuArr["sort"] = val.Sort

		menuData[key] = menuArr

	}

	return menuData
}

/**
	生成随机字符串
	@param int 长度
    @return string
*/
func (c *HomeBaseController) GetRandomString(lens int) string {

	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	bytes := []byte(str)

	result := []byte{}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < lens; i++ {

		result = append(result, bytes[r.Intn(len(bytes))])
	}

	return string(result)
}
