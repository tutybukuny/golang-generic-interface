package productrepository

import (
	"github.com/tutybukuny/golang-generic-interface/entity"
	"github.com/tutybukuny/golang-generic-interface/repository"
)

type IProductRepo interface {
	repository.IInsertRepo[*entity.Product]
	repository.IUpdateRepo[*entity.Product, int64]
	repository.IDeleteRepo[*entity.Product, int64]
}
