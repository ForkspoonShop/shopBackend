package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type Product struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	URL         string `db:"url"`
	Cost        string `db:"cost"`
	Category    string `db:"category"`
	Description string `db:"description"`
}

func main() {
	db, err := gorm.Open("mysql", "root:qwe123@/shop?charset=utf8&")
	if err != nil {
		log.Println("Connection Failed to Open", err)
	}
	defer db.Close()
	db.AutoMigrate(Product{})
	router := gin.Default()
	router.GET("/allproducts", handlerProductsAll)
	router.GET("/products/:category", handlerProductsGroup)
	router.GET("/product/:id", handlerProductId)
	router.OPTIONS("/addproduct", handlerOption)
	router.POST("/addproduct", handlerProductAdd)
	router.OPTIONS("/delproduct/:id", handlerOption)
	router.DELETE("/delproduct/:id", handlerProductDel)
	router.POST("/upload", handlerProductUpload)
	router.OPTIONS("/upload", handlerOption)
	_ = router.Run()
}

func handlerOption(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.JSON(200, "accept")
}

func handlerProductsAll(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:qwe123@/shop?charset=utf8&")
	if err != nil {
		log.Println("Connection Failed to Open", err)
	}
	defer db.Close()
	log.Println("Connection Established")
	var products []Product
	db.Find(&products)
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, products)
}

func handlerProductsGroup(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:qwe123@/shop?charset=utf8&")
	if err != nil {
		log.Println("Connection Failed to Open", err)
	}
	defer db.Close()
	log.Println("Connection Established")
	var products []Product
	db.Where("category = ?", c.Param("category")).Find(&products)
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, products)
}

func handlerProductId(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:qwe123@/shop?charset=utf8&")
	if err != nil {
		log.Println("Connection Failed to Open", err)
	}
	defer db.Close()
	log.Println("Connection Established")
	var products []Product
	db.Where("id = ?", c.Param("id")).First(&products)
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, products)
}

func handlerProductAdd(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:qwe123@/shop?charset=utf8&")
	if err != nil {
		log.Println("Connection Failed to Open", err)
	}
	defer db.Close()
	log.Println("Connection Established2")
	var products Product
	_ = c.ShouldBindJSON(&products)
	log.Println(products)
	db.Create(&products)
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, "accept")
}

func handlerProductDel(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:qwe123@/shop?charset=utf8&")
	if err != nil {
		log.Println("Connection Failed to Open", err)
	}
	defer db.Close()
	log.Println(c.Param("id"))
	db.Where("id = ?", c.Param("id")).Delete(Product{})
	c.Header("Access-Control-Allow-Origin", "*")
	c.String(200, "accept")
}

func handlerProductUpload(c *gin.Context) {
	log.Println(c.FormFile("file"))
	c.Header("Access-Control-Allow-Origin", "*")
	c.String(200, "accept")
}
