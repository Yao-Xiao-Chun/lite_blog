package entity

// Users  试题
type Users struct {
	Title    string `json:"title"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Email    string `json:"email"`
	File     string `json:"file"`
	Status   string `json:"status"`
	IsAdmin  string `json:"is_admin"`
	CreateAt string `json:"create_at"`
	TitleImg string `json:"title_img"`
}
