package invoicerepository

import (
	"github.com/tutybukuny/golang-generic-interface/entity"
	"github.com/tutybukuny/golang-generic-interface/repository"
)

type IInvoiceRepo interface {
	repository.IInsertRepo[*entity.Invoice]
	repository.IUpdateRepo[*entity.Invoice, string]
	repository.IDeleteRepo[*entity.Invoice, string]

	FindByProductID(productID int64) []*entity.Invoice
}
