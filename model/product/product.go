package product

type Product struct {
	ID         int `json:"id"`
	Name    string `json:"name"`
	Description    string `json:"description"`
	Price float32 `json:"price"`
	Created_at   string `json:"created_at"`
}
