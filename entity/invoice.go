package entity

type Invoice struct {
	ID        string
	ProductID int64
	Quantity  int64
	Price     int64
}

func (Invoice) GetType() string {
	return "Invoice"
}
