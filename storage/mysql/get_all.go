package mysqlstorage

import "github.com/tutybukuny/golang-generic-interface/entity"

type GetAllStorage[E entity.IEntity] struct {
	*BaseStorage
}

func NewGetAllStorage[E entity.IEntity](baseRepo *BaseStorage) *GetAllStorage[E] {
	return &GetAllStorage[E]{baseRepo}
}

func (s *GetAllStorage[E]) GetAll() []*E {
	return []*E{
		new(E),
	}
}
