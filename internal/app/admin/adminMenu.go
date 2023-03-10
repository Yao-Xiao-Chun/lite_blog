package admin

import (
	"mywork/models"
	"strconv"

	"encoding/json"
	"fmt"
	"strings"
)

type AdminMenuController struct {
	AdminBaseController
}

type Cat struct {
	ID          int
	Menu_parent int
	Menu_name   string
	Level       int
	Menu_status int
	sort        int
	Children    []*Cat ``
}

/**
列表页
*/
// @router /admin/menu [get] 后台 tag列表
func (this *AdminMenuController) GetList() {

	data := this.GetMenuList(true)

	this.Data["list"] = data //总条数*/

	this.TplName = "admin/menu/list.html"
}

/**
新增菜单页面
*/
// @router /admin/menu/add [get] 后台 tag列表
func (this *AdminMenuController) GetAdd() {

	var id string

	this.Ctx.Input.Bind(&id, "id") //是否有选中的值

	menuData := this.GetMenuList(false) //获取列表

	this.Data["menu"] = menuData

	ids, _ := strconv.Atoi(id)

	this.Data["ids"] = ids

	this.TplName = "admin/menu/add.html"
}

/**
新增页面数据处理
*/
// @router /admin/menu/add [post] 后台
func (this *AdminMenuController) MenuAddForm() {

	menuName := this.GetString("menu_name")

	menuLevel := this.GetString("menu_level")

	menuStatus := this.GetString("menu_status")

	menuSort := this.GetString("menu_sort")

	if menuLevel == "" || menuName == "" || menuSort == "" || menuStatus == "" {

		this.Data["json"] = map[string]interface{}{
			"code": "1003",
			"msg":  "参数不完整",
		}
	} else {

		var menu models.LiteAdminMenu

		menu.Menu_name = menuName //名称

		menu.Menu_parent, _ = strconv.Atoi(menuLevel) //fid

		menu.Menu_status, _ = strconv.Atoi(menuStatus) //启用状态

		menu.Menu_roule, _ = strconv.Atoi(menuSort) //排序

		key, num := this.getFidAndLevel(menuLevel) // 获取创建的等级

		menu.Menu_key = key //排序等级

		menu.Menu_level = num //等级

		models.SetMenu(menu)

		this.Data["json"] = map[string]interface{}{
			"code": "0",
			"msg":  "创建成功",
		}

	}

	this.ServeJSON()
}

/**
删除菜单
*/
// @router /admin/menu/del/?:key [get] 后台 菜单删除
func (this *AdminMenuController) DeleteMenu() {
	var id string

	this.Ctx.Input.Bind(&id, "id")

	if id == "" {
		this.Data["json"] = map[string]interface{}{
			"code":   "1006",
			"errmsg": "删除失败！",
		}
	} else {
		ids, _ := strconv.Atoi(id)

		models.DeleteMenu(uint(ids))

		this.Data["json"] = map[string]interface{}{
			"code":   "0",
			"errmsg": "删除成功！",
		}
	}

	this.ServeJSON()
}

/**
编辑菜单列表
*/
// @router /admin/menu/edit/?:key [get] 后台 菜单删除
func (this *AdminMenuController) EditMenu() {

	var id string

	this.Ctx.Input.Bind(&id, "id")

	menuData := this.GetMenuList(false) //获取列表

	this.Data["menu"] = menuData

	ids, _ := strconv.Atoi(id)

	list, _ := models.GetMenuAndFindInfo(uint(ids))

	this.Data["list"] = list

	this.Data["ids"] = list.Menu_parent

	this.TplName = "admin/menu/edit.html"
}

/**
编辑菜单数据提交 post
*/
// @router /admin/menu/edit [post] 后台 菜单更新
func (this *AdminMenuController) EditPost() {

	menuName := this.GetString("menu_name")

	menuLevel := this.GetString("menu_level")

	menuStatus := this.GetString("menu_status")

	menuSort := this.GetString("menu_sort")

	menuId := this.GetString("menu_id")

	if menuLevel == "" || menuName == "" || menuSort == "" || menuStatus == "" || menuId == "" {

		this.Data["json"] = map[string]interface{}{
			"code": "1003",
			"msg":  "参数不完整",
		}
	} else {

		var menu models.LiteAdminMenu

		menu.Menu_name = menuName //名称

		menu.Menu_parent, _ = strconv.Atoi(menuLevel) //fid

		menu.Menu_status, _ = strconv.Atoi(menuStatus) //启用状态

		menu.Menu_roule, _ = strconv.Atoi(menuSort) //排序

		key, num := this.getFidAndLevel(menuLevel) // 获取创建的等级

		ids, _ := strconv.Atoi(menuId)

		menu.ID = uint(ids) //更新id

		menu.Menu_key = key //排序等级

		menu.Menu_level = num //等级

		models.EditMenu(menu)

		this.Data["json"] = map[string]interface{}{
			"code": "0",
			"msg":  "更新成功",
		}

	}

	this.ServeJSON()
}

func tree(list []*Cat) string {

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
func buildData(list []*Cat) map[int]map[int]*Cat {

	var data map[int]map[int]*Cat = make(map[int]map[int]*Cat)

	for _, v := range list {

		id := v.ID

		fid := v.Menu_parent

		if _, ok := data[fid]; !ok {

			data[fid] = make(map[int]*Cat)
		}

		data[fid][id] = v
	}

	return data
}

/**
结构树
*/
func makeTreeCore(index int, data map[int]map[int]*Cat) []*Cat {

	tmp := make([]*Cat, 0)

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

func (this *AdminMenuController) getFidAndLevel(fid string) (str string, num int) {

	if fid == "0" {

		return "0-", 0

	} else {

		id, _ := strconv.Atoi(fid)

		menkey, _ := models.GetMenuAndRoule(id)

		num := strings.Count(menkey.Menu_key, "-")

		return menkey.Menu_key + fid + "-", num //0-1-2-3-
	}
}

/**
递归实现无线菜单分类 核心
*/
func superCategory(allCate []Cat, pid int) []Cat {
	var arr []Cat
	for _, v := range allCate {
		if pid == v.Menu_parent {
			arr = append(arr, v)
			sonCate := superCategory(allCate, v.ID)
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
func (this *AdminMenuController) GetMenuList(flag bool) (data map[int]map[string]interface{}) {

	var menu []models.LiteAdminMenu
	//是获取所有的还是获取选择的
	if flag {
		menu, _ = models.GetAll()
	} else {
		menu, _ = models.GetMenuInfo() //根据条件获取
	}

	var list []Cat

	list = make([]Cat, len(menu))

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

	menus := superCategory(list, 0)

	var menuData map[int]map[string]interface{}

	menuData = make(map[int]map[string]interface{})

	var menuArr map[string]interface{}

	for key, val := range menus {

		menuArr = make(map[string]interface{})

		var str string

		if val.Level > 0 {

			str = setSpace(val.Level)
		}

		menuArr["name"] = str + val.Menu_name

		menuArr["id"] = val.ID

		menuArr["pid"] = val.Menu_parent

		menuArr["status"] = val.Menu_status
		menuArr["sort"] = val.sort

		menuData[key] = menuArr

	}

	return menuData
}
