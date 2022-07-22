package repository

import "github.com/tutybukuny/golang-generic-interface/entity"

type IInsertRepo[E entity.IEntity] interface {
	Insert(obj E)
}

type IUpdateRepo[E entity.IEntity, K any] interface {
	Update(obj E, id K)
}

type IDeleteRepo[E entity.IEntity, K any] interface {
	Delete(obj E, id K)
}

type IGetAllRepo[E entity.IEntity] interface {
	GetAll() []*E
}
