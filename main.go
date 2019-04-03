package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io"
	"os"
	"strconv"
	"uran_test/consts"
	"uran_test/crud"
	"uran_test/mysql"
)

func main() {

	mysql.Open()

	defer mysql.Close()

	router := gin.Default()
	router.Use(static.Serve("/images", static.LocalFile("./images", true)))
	api := router.Group("/api")
	products := api.Group("/products")
	{
		products.GET("/", searchProductsList)
	}
	product := api.Group("/product")
	{
		product.GET("/:productId", getProduct)
		product.POST("/", postProduct)
		product.PATCH("/", updateProduct)
		product.POST("/:productId/image", updateProductImage)
	}

	router.Run()

}

func respondWithError(c *gin.Context, err crud.Error) {
	c.JSON(consts.CodeError, gin.H{
		"error": err,
	})
	c.Abort()
}

func response(body interface{}) gin.H {
	return gin.H{
		"body": body,
	}
}

func searchProductsList(c *gin.Context) {
	q := c.Query("q")
	p, err := crud.GetProducts(q)

	if err != nil {
		respondWithError(c, crud.CreateError(consts.CodeDbError, consts.MessageErrorGettingProducts))
		return
	}

	c.JSON(consts.CodeSuccess, response(p))
}

func getProduct(c *gin.Context) {
	idString := c.Param("productId")
	id, err := strconv.Atoi(idString)
	if err != nil {
		respondWithError(c, crud.CreateError(consts.CodeWrongProductIdParameter, consts.MessageInvalidProductId))
	}
	p, err := crud.GetProductById(id)

	if err != nil {
		respondWithError(c, crud.CreateError(consts.CodeDbError, consts.MessageErrorGettingProduct))
		return
	}

	c.JSON(consts.CodeSuccess, response(p))
}

func postProduct(c *gin.Context) {
	var product crud.Product
	c.Bind(&product)

	p, err := crud.AddProduct(product)
	if err != nil {
		respondWithError(c, crud.CreateError(consts.CodeDbError, consts.MessageErrorSavingProduct))
		return
	}

	c.JSON(consts.CodeSuccess, response(p))
}

func updateProductImage(c *gin.Context) {
	idS := c.Param("productId")
	id, err := strconv.Atoi(idS)
	if err != nil {
		respondWithError(c, crud.CreateError(consts.CodeWrongProductIdParameter, consts.MessageInvalidProductId))
	}

	img, _, err := c.Request.FormFile("image")

	out, err := os.Create("./images/" + "product_" + idS + ".png")
	defer out.Close()

	if err != nil {
		respondWithError(c, crud.CreateError(consts.CodeCantSaveImage, consts.MessageCantSaveImage))
		return
	}

	_, err = io.Copy(out, img)
	if err != nil {
		respondWithError(c, crud.CreateError(consts.CodeCantSaveImage, consts.MessageCantSaveImage))
		return
	}

	p, err := crud.EditProductImage(id, consts.URL+"images/"+"product_"+idS+".png")

	if err != nil {
		respondWithError(c, crud.CreateError(consts.CodeDbError, consts.MessageErrorGettingProduct))
		return
	}

	c.JSON(consts.CodeSuccess, response(p))

}

func updateProduct(c *gin.Context) {
	var product crud.Product
	c.Bind(&product)

	if err := product.HandleProductChanges(); err != nil {
		respondWithError(c, *err)
		return
	}

	p, err := crud.EditProduct(product)

	if err != nil {
		respondWithError(c, crud.CreateError(consts.CodeDbError, consts.MessageErrorUpdatingProduct))
		return
	}

	c.JSON(consts.CodeSuccess, response(p))
}
