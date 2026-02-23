package products

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var products = []Product{
	{ID: 1, Name: "Gary Slippers", Price: 29.99},
	{ID: 2, Name: "Fish Slippers", Price: 19.99},
	{ID: 3, Name: "Avocado Slippers", Price: 29.99},
	{ID: 4, Name: "Croissant Slippers", Price: 19.99},
	{ID: 5, Name: "Cat Slippers", Price: 19.99},
}

func GetProducts() []Product {
	return products
}
