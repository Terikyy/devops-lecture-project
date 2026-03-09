package products

// FindProductByID returns the product with the given ID, or nil if not found.
func FindProductByID(id int) *Product {
	for _, product := range products {
		if product.ID == id {
			return &product
		}
	}
	return nil
}
