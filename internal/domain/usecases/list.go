package domain

type ListProducts struct {
	ID         int64    `json:"id"`
	CategoryID int64    `json:"category_id"`
	Name       string   `json:"name"`
	Desc       string   `json:"description"`
	Price      float64  `json:"price"`
	Qty        int      `json:"qty"`
	IsActive   bool     `json:"is_active"`
	Category   Category `json:"category"`
}

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
type SearchItemList struct {
	Name     string `query:"name"`
	Category string `query:"category"`
	Limit    int    `query:"limit"`
	Offset   int    `query:"offset"`
}
