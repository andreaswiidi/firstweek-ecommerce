package request

type CreateProductRequest struct {
	UPI   string  `validate:"required" json:"UPI"`
	Title string  `validate:"required" json:"title"`
	Price float64 `validate:"required" json:"price"`
	// BrandId uuid.UUID `validate:"required" json:"brand_id"`
}

type UpdateProductRequest struct {
	UPI   string  `validate:"required" json:"UPI"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
	// BrandId uuid.UUID `json:"brand_id"`
}
