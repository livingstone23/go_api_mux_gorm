package dto

type CategoryDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProductDto struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
	CategoryID  uint   `json:"category_id"`
}
