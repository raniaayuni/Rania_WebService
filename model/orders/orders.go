package order

type Order struct {
	ID           int    `json:"id"`
	User_id      int    `json:"user_id"`
	Product_id   int    `json:"product_id"`
	Order_status string `json:"order_status"`
}
