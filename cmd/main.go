package main

import (
	"log"

	"github.com/tutybukuny/golang-generic-interface/entity"
	"github.com/tutybukuny/golang-generic-interface/repository/invoice"
	"github.com/tutybukuny/golang-generic-interface/repository/product"
	"github.com/tutybukuny/golang-generic-interface/storage/mysql/invoice"
	"github.com/tutybukuny/golang-generic-interface/storage/mysql/product"
)

func main() {
	var productRepo productrepository.IProductRepo
	var invoiceRepo invoicerepository.IInvoiceRepo

	dbConnection := "Mysql"
	productRepo = productstore.New(dbConnection)
	invoiceRepo = invoicestorage.New(dbConnection)

	//region product repo
	product := &entity.Product{
		ID:    1,
		Name:  "Nintendo Switch",
		Price: 100,
	}
	productRepo.Insert(product)
	productRepo.Update(product, 1)
	productRepo.Delete(product, 1)
	products := productRepo.GetAll()
	products[0].Price = 1000
	log.Printf("product price %d", products[0].Price)
	//endregion

	//region invoice repo
	invoice := &entity.Invoice{
		ID:        "I001",
		ProductID: 1,
		Quantity:  10,
		Price:     100,
	}
	invoiceRepo.Insert(invoice)
	invoiceRepo.Update(invoice, "I001")
	invoiceRepo.Delete(invoice, "I001")
	invoices := invoiceRepo.FindByProductID(1)
	log.Printf("list invoices: %v", invoices)
	//endregion
}
