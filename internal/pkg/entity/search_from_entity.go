package entity

type SearchFromEntity struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Keywords    string `json:"keywords"`
	Description string `json:"descript"`
}
