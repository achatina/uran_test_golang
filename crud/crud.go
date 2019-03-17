package crud

import (
	"uran_test/mysql"
)

const productType = "product_type"
const category = "category"
const product = "product"

func GetProducts(q string) ([]Product, error) {
	db := mysql.Get()

	var p []ProductDb
	var products []Product

	if err := db.Table(product).Find(&p, "name LIKE ?", "%" + q + "%").Error; err != nil {
		return nil, err
	}

	for _, prod := range p {
		c, err := getCategoryById(prod.CategoryId)
		if err != nil {
			return nil, err
		}
		pType, err := getProductTypeById(prod.ProductTypeId)
		if err != nil {
			return nil, err
		}
		products = append(products, Product{
			Id:          prod.Id,
			ProductType: *pType,
			Category:    *c,
			Name:        prod.Name,
			Description: prod.Description,
			Image:       prod.Image,
		})

	}

	return products, nil
}

func getProductDbById(id int) (*ProductDb, error) {
	db := mysql.Get()
	var p ProductDb

	if err := db.Table(product).First(&p, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &p, nil
}

func getCategoryById(id int) (*Category, error) {
	db := mysql.Get()
	var c Category

	if err := db.Table(category).Where("id = ?", id).Find(&c).Error; err != nil {
		return nil, err
	}

	return &c, nil
}

func getProductTypeById(id int) (*ProductType, error) {
	db := mysql.Get()
	var pType ProductType

	if err := db.Table(productType).First(&pType, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &pType, nil
}

func GetProductById(id int) (*Product, error) {
	p, err := getProductDbById(id)
	if err != nil {
		return nil, err
	}

	c, err := getCategoryById(p.CategoryId)
	if err != nil {
		return nil, err
	}

	pType, err := getProductTypeById(p.ProductTypeId)
	if err != nil {
		return nil, err
	}

	return &Product{
		Id:          p.Id,
		ProductType: *pType,
		Category:    *c,
		Name:        p.Name,
		Description: p.Description,
		Image:       p.Image,
	}, nil
}

func EditProductImage(id int, url string) (*Product, error) {

	p, err := getProductDbById(id)

	if err != nil {
		return nil, err
	}

	p.Image = url

	return updateProduct(*p)
}

func EditProduct(newP Product) (*Product, error) {

	p, err := getProductDbById(newP.Id)

	if err != nil {
		return nil, err
	}

	p.Name = newP.Name
	p.ProductTypeId = newP.ProductType.Id
	p.CategoryId = newP.Category.Id
	p.Description = newP.Description

	return updateProduct(*p)
}

func updateProduct(p ProductDb) (*Product, error) {
	db := mysql.Get()
	if err := db.Table(product).Save(&p).Error; err != nil {
		return nil, err
	}

	return GetProductById(p.Id)
}

func AddProduct(p Product) (*Product, error) {
	db := mysql.Get()

	prodType := p.ProductType
	c := p.Category

	if err := db.Table(productType).Create(&prodType).Error; err != nil {
		return nil, err
	}

	if err := db.Table(category).Create(&c).Error; err != nil {
		return nil, err
	}

	prod := ProductDb{
		ProductTypeId: prodType.Id,
		CategoryId:    c.Id,
		Name:          p.Name,
		Description:   p.Description,
		Image:         p.Image,
	}

	if err := db.Table(product).Create(&prod).Error; err != nil {
		return nil, err
	}

	return &Product{
		Id:          prod.Id,
		ProductType: prodType,
		Category:    c,
		Name:        prod.Name,
		Description: prod.Description,
		Image:       prod.Image,
	}, nil
}
