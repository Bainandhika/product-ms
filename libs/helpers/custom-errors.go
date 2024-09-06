package helpers

type ProductNotFound struct {
	Message string
}

func (p *ProductNotFound) Error() string {
	return "product not found"
}