package invoicestorage

import (
	"log"

	"github.com/tutybukuny/golang-generic-interface/entity"
	"github.com/tutybukuny/golang-generic-interface/storage/mysql"
)

type InvoiceStorage struct {
	*mysqlstorage.BaseStorage
	*mysqlstorage.InsertStorage[*entity.Invoice]
	*mysqlstorage.UpdateStorage[*entity.Invoice, string]
	*mysqlstorage.DeleteStorage[*entity.Invoice, string]
}

func New(dbConnection string) *InvoiceStorage {
	baseRepo := mysqlstorage.NewBaseStorage(dbConnection)
	insertRepo := mysqlstorage.NewInsertStorage[*entity.Invoice](baseRepo)
	updateRepo := mysqlstorage.NewUpdateStorage[*entity.Invoice, string](baseRepo)
	deleteRepo := mysqlstorage.NewDeleteStorage[*entity.Invoice, string](baseRepo)
	return &InvoiceStorage{baseRepo, insertRepo, updateRepo, deleteRepo}
}

func (s *InvoiceStorage) FindByProductID(productID int64) []*entity.Invoice {
	log.Printf("%s get all invoice with product id %d", s.GetDB(), productID)
	//fake data
	return []*entity.Invoice{
		{
			ID:        "I001",
			ProductID: 1,
			Quantity:  10,
			Price:     100,
		},
		{
			ID:        "I002",
			ProductID: 2,
			Quantity:  8,
			Price:     300,
		},
	}
}
