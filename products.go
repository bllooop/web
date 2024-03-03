package web

type ProductList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ShopList struct {
	Id     int
	ShopId int
	ListId int
}

type ProductItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Expiration  string `json:"expiration"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}
