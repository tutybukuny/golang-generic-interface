package mysqlstorage

import (
	"log"

	"github.com/tutybukuny/golang-generic-interface/entity"
)

type InsertStorage[E entity.IEntity] struct {
	*BaseStorage
}

func NewInsertStorage[E entity.IEntity](baseStorage *BaseStorage) *InsertStorage[E] {
	return &InsertStorage[E]{baseStorage}
}

func (s *InsertStorage[E]) Insert(obj E) {
	log.Printf("%s insert new object type %s: %v", s.GetDB(), obj.GetType(), obj)
}
