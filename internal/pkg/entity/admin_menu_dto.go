package entity

type Cat struct {
	ID         int
	MenuParent int
	MenuName   string
	Level      int
	MenuStatus int
	Sort       int
	Children   []*Cat ``
}
