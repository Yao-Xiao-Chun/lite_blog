package admin

import (
	"mywork/internal/pkg/dto"
	"mywork/models"
	"strconv"

	"encoding/json"
	"fmt"
	"strings"
)

type AdminMenuController struct {
	AdminBaseController
}

/**
列表页
*/
// @router /admin/menu [get] 后台 tag列表
func (c *AdminMenuController) GetList() {

	data := c.GetMenuList(true)

	c.Data["list"] = data //总条数*/

	c.TplName = "admin/menu/list.html"
}

/**
新增菜单页面
*/
// @router /admin/menu/add [get] 后台 tag列表
func (c *AdminMenuController) GetAdd() {

	var id string

	c.Ctx.Input.Bind(&id, "id") //是否有选中的值

	menuData := c.GetMenuList(false) //获取列表

	c.Data["menu"] = menuData

	ids, _ := strconv.Atoi(id)

	c.Data["ids"] = ids

	c.TplName = "admin/menu/add.html"
}

/**
新增页面数据处理
*/
// @router /admin/menu/add [post] 后台
func (c *AdminMenuController) MenuAddForm() {

	menuName := c.GetString("menu_name")

	menuLevel := c.GetString("menu_level")

	menuStatus := c.GetString("menu_status")

	menuSort := c.GetString("menu_sort")

	if menuLevel == "" || menuName == "" || menuSort == "" || menuStatus == "" {

		c.Data["json"] = map[string]interface{}{
			"code": "1003",
			"msg":  "参数不完整",
		}
	} else {

		var menu models.LiteAdminMenu

		menu.Menu_name = menuName //名称

		menu.Menu_parent, _ = strconv.Atoi(menuLevel) //fid

		menu.Menu_status, _ = strconv.Atoi(menuStatus) //启用状态

		menu.Menu_roule, _ = strconv.Atoi(menuSort) //排序

		key, num := c.getFidAndLevel(menuLevel) // 获取创建的等级

		menu.Menu_key = key //排序等级

		menu.Menu_level = num //等级

		models.SetMenu(menu)

		c.Data["json"] = map[string]interface{}{
			"code": "0",
			"msg":  "创建成功",
		}

	}

	c.ServeJSON()
}

/**
删除菜单
*/
// @router /admin/menu/del/?:key [get] 后台 菜单删除
func (c *AdminMenuController) DeleteMenu() {
	var id string

	c.Ctx.Input.Bind(&id, "id")

	if id == "" {
		c.Data["json"] = map[string]interface{}{
			"code":   "1006",
			"errmsg": "删除失败！",
		}
	} else {
		ids, _ := strconv.Atoi(id)

		models.DeleteMenu(uint(ids))

		c.Data["json"] = map[string]interface{}{
			"code":   "0",
			"errmsg": "删除成功！",
		}
	}

	c.ServeJSON()
}

/**
编辑菜单列表
*/
// @router /admin/menu/edit/?:key [get] 后台 菜单删除
func (c *AdminMenuController) EditMenu() {

	var id string

	c.Ctx.Input.Bind(&id, "id")

	menuData := c.GetMenuList(false) //获取列表

	c.Data["menu"] = menuData

	ids, _ := strconv.Atoi(id)

	list, _ := models.GetMenuAndFindInfo(uint(ids))

	c.Data["list"] = list

	c.Data["ids"] = list.Menu_parent

	c.TplName = "admin/menu/edit.html"
}

/**
编辑菜单数据提交 post
*/
// @router /admin/menu/edit [post] 后台 菜单更新
func (c *AdminMenuController) EditPost() {

	menuName := c.GetString("menu_name")

	menuLevel := c.GetString("menu_level")

	menuStatus := c.GetString("menu_status")

	menuSort := c.GetString("menu_sort")

	menuId := c.GetString("menu_id")

	if menuLevel == "" || menuName == "" || menuSort == "" || menuStatus == "" || menuId == "" {

		c.Data["json"] = map[string]interface{}{
			"code": "1003",
			"msg":  "参数不完整",
		}
	} else {

		var menu models.LiteAdminMenu

		menu.Menu_name = menuName //名称

		menu.Menu_parent, _ = strconv.Atoi(menuLevel) //fid

		menu.Menu_status, _ = strconv.Atoi(menuStatus) //启用状态

		menu.Menu_roule, _ = strconv.Atoi(menuSort) //排序

		key, num := c.getFidAndLevel(menuLevel) // 获取创建的等级

		ids, _ := strconv.Atoi(menuId)

		menu.ID = uint(ids) //更新id

		menu.Menu_key = key //排序等级

		menu.Menu_level = num //等级

		models.EditMenu(menu)

		c.Data["json"] = map[string]interface{}{
			"code": "0",
			"msg":  "更新成功",
		}

	}

	c.ServeJSON()
}

func tree(list []*dto.Cat) string {

	data := buildData(list)

	result := makeTreeCore(0, data)

	body, err := json.Marshal(result)

	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}

/**
结构生成可以维护的数据
*/
func buildData(list []*dto.Cat) map[int]map[int]*dto.Cat {

	var data map[int]map[int]*dto.Cat = make(map[int]map[int]*dto.Cat)

	for _, v := range list {

		id := v.ID

		fid := v.MenuParent

		if _, ok := data[fid]; !ok {

			data[fid] = make(map[int]*dto.Cat)
		}

		data[fid][id] = v
	}

	return data
}

/**
结构树
*/
func makeTreeCore(index int, data map[int]map[int]*dto.Cat) []*dto.Cat {

	tmp := make([]*dto.Cat, 0)

	for id, item := range data[index] {

		if data[id] != nil {

			item.Children = makeTreeCore(id, data)
		}

		tmp = append(tmp, item)
	}

	return tmp
}

/**
获取当前中父类等级
*/

func (c *AdminMenuController) getFidAndLevel(fid string) (str string, num int) {

	if fid == "0" {

		return "0-", 0

	} else {

		id, _ := strconv.Atoi(fid)

		menkey, _ := models.GetMenuAndRoule(id)

		num := strings.Count(menkey.Menu_key, "-")

		return menkey.Menu_key + fid + "-", num //0-1-2-3-
	}
}

// SuperCategory /**
func SuperCategory(allCate []dto.Cat, pid int) []dto.Cat {
	var arr []dto.Cat
	for _, v := range allCate {
		if pid == v.MenuParent {
			arr = append(arr, v)
			sonCate := SuperCategory(allCate, v.ID)
			arr = append(arr, sonCate...)
		}
	}
	return arr
}

/**
遍历符号
*/
func setSpace(num int) string {

	str := ""

	for i := 0; i < num; i++ {

		str += "&nbsp;&nbsp;"
	}

	return str + "|-&nbsp;&nbsp;"
}

/**
菜单列表
@param bool 条件筛选
*/
func (c *AdminMenuController) GetMenuList(flag bool) (data map[int]map[string]interface{}) {

	var menu []models.LiteAdminMenu
	//是获取所有的还是获取选择的
	if flag {
		menu, _ = models.GetAll()
	} else {
		menu, _ = models.GetMenuInfo() //根据条件获取
	}

	var list []dto.Cat

	list = make([]dto.Cat, len(menu))

	/*var listA []*dto.Cat

	listA = make([]*dto.Cat,len(menu))*/

	for i, v := range menu {

		list[i].ID = int(v.ID)
		list[i].MenuName = v.Menu_name
		list[i].MenuParent = v.Menu_parent
		list[i].Level = v.Menu_level
		list[i].Sort = v.Menu_roule
		list[i].MenuStatus = v.Menu_status
		//listA[i] = &list[i]
	}

	menus := SuperCategory(list, 0)

	var menuData map[int]map[string]interface{}

	menuData = make(map[int]map[string]interface{})

	var menuArr map[string]interface{}

	for key, val := range menus {

		menuArr = make(map[string]interface{})

		var str string

		if val.Level > 0 {

			str = setSpace(val.Level)
		}

		menuArr["name"] = str + val.MenuName

		menuArr["id"] = val.ID

		menuArr["pid"] = val.MenuParent

		menuArr["status"] = val.MenuStatus
		menuArr["sort"] = val.Sort

		menuData[key] = menuArr

	}

	return menuData
}
