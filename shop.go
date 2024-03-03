package web

type Shop struct {
	Id       int    `json:"-" db:"id"`
	ShopName string `json:"shopname" binding:"required"`
	Password string `json:"password" binding:"required"`
}
