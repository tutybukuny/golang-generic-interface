package mysqlstorage

import (
	"log"

	"github.com/tutybukuny/golang-generic-interface/entity"
)

type UpdateStorage[E entity.IEntity, K any] struct {
	*BaseStorage
}

func NewUpdateStorage[E entity.IEntity, K any](baseStorage *BaseStorage) *UpdateStorage[E, K] {
	return &UpdateStorage[E, K]{baseStorage}
}

func (s *UpdateStorage[E, K]) Update(obj E, id K) {
	log.Printf("%s update object type %s with key %v: %v", s.GetDB(), obj.GetType(), id, obj)
}
