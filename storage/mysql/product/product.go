package productstore

import (
	"github.com/tutybukuny/golang-generic-interface/entity"
	"github.com/tutybukuny/golang-generic-interface/storage/mysql"
)

type ProductStore struct {
	*mysqlstorage.BaseStorage
	*mysqlstorage.InsertStorage[*entity.Product]
	*mysqlstorage.UpdateStorage[*entity.Product, int64]
	*mysqlstorage.DeleteStorage[*entity.Product, int64]
}

func New(dbConnection string) *ProductStore {
	baseRepo := mysqlstorage.NewBaseStorage(dbConnection)
	insertRepo := mysqlstorage.NewInsertStorage[*entity.Product](baseRepo)
	updateRepo := mysqlstorage.NewUpdateStorage[*entity.Product, int64](baseRepo)
	deleteRepo := mysqlstorage.NewDeleteStorage[*entity.Product, int64](baseRepo)
	return &ProductStore{baseRepo, insertRepo, updateRepo, deleteRepo}
}
