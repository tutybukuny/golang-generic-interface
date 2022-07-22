package mysqlstorage

import (
	"log"

	"github.com/tutybukuny/golang-generic-interface/entity"
)

type DeleteStorage[E entity.IEntity, K any] struct {
	*BaseStorage
}

func NewDeleteStorage[E entity.IEntity, K any](baseStorage *BaseStorage) *DeleteStorage[E, K] {
	return &DeleteStorage[E, K]{baseStorage}
}

func (s *DeleteStorage[E, K]) Delete(obj E, id K) {
	log.Printf("%s delete object type %s with key %d: %v", s.GetDB(), obj.GetType(), id, obj)
}
