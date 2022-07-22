package entity

type Product struct {
	ID    int64
	Name  string
	Price int64
}

func (Product) GetType() string {
	return "Product"
}
