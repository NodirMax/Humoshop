package product

type ProductDatabaseStruct struct {
	Product_id int64   `json:"product_id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	In_stock   bool    `json:"in_stock"`
}

type GetProductStruct struct {
	Id int64 `json:"id"`
}

type CreateProductStruct struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	In_stock bool    `json:"in_stock"`
}

type UpdateProductStruct struct {
	Product_id int64   `json:"product_id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	In_stock   bool    `json:"in_stock"`
}

type DeleteProductStruct struct {
	Product_id int64 `json:"product_id"`
}
