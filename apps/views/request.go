package views

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Variety     string  `json:"variety" binding:"required"`
	Rating      float32 `json:"rating" binding:"required"`
	Stock       int     `json:"stock" binding:"required"`
}

type UpdateProductRequest struct {
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Price       *float64 `json:"price,omitempty"`
	Variety     *string  `json:"variety,omitempty"`
	Rating      *float32 `json:"rating,omitempty"`
	Stock       *int     `json:"stock,omitempty"`
}
