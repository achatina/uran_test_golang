package crud

import (
	"fmt"
	"uran_test/consts"
)

type ProductType struct {
	Id   int    `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

type Category struct {
	Id   int    `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

type DbProduct struct {
	Id            int    `gorm:"column:id"`
	ProductTypeId int    `gorm:"column:product_type_id"`
	CategoryId    int    `gorm:"column:category_id"`
	Name          string `gorm:"column:name"`
	Description   string `gorm:"column:description"`
	Image         string `gorm:"column:image"`
}

type Product struct {
	Id          int         `json:"id"`
	ProductType ProductType `json:"product_type"`
	Category    Category    `json:"category"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Image       string      `json:"image"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

func (p Product) HandleProductChanges() *Error {
	if len(p.Name) == 0 {
		return &Error{Code: consts.CodeEmptyProductName, Message: consts.MessageEmptyProductName}
	}
	if p.Category == (Category{}) {
		return &Error{Code: consts.CodeEmptyProductCategory, Message: consts.MessageEmptyProductCategory}
	}
	if p.ProductType == (ProductType{}) {
		return &Error{Code: consts.CodeEmptyProductType, Message: consts.MessageEmptyProductType}
	}

	return nil
}

func CreateError(code int, message string) Error {
	return Error{Code: code, Message: message}
}
