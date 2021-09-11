package products

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	ImageUrl    string  `json:"imageUrl"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}
