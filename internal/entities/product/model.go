package product

type ProductDatabaseStruct struct {
	Product_id  int64   `json:"product_id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	In_stock    bool    `json:"in_stock"`
	Category_id int64   `json:"category_id"`
}
