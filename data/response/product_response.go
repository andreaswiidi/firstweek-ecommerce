package response

type ProductResponse struct {
	UPI   string  `json:"UPI"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
}
