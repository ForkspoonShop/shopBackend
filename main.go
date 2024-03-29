package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

type Product struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	URL         string `db:"url"`
	Cost        string `db:"cost"`
	Category    string `db:"category"`
	Description string `db:"description"`
}

const baseLogin = "root:qwe123@/shop?charset=utf8&"

func main() {
	db, err := gorm.Open("mysql", baseLogin)
	if err != nil {
		log.Println("Connection Failed to Open", err)
	}
	defer db.Close()
	db.AutoMigrate(Product{})
	router := gin.Default()
	router.GET("/allproducts", handlerProductsAll)
	router.GET("/products/:category", handlerProductsGroup)
	router.GET("/product/:id", handlerProductId)
	router.POST("/addproduct", handlerProductAdd)
	router.POST("/updateproduct/:id", handlerProductUpdate)
	router.DELETE("/delproduct/:id", handlerProductDel)
	router.POST("/upload", handlerImageUpload)
	router.OPTIONS("/addproduct", handlerOption)
	router.OPTIONS("/delproduct/:id", handlerOption)
	router.OPTIONS("/updateproduct/:id", handlerOption)
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
	db, err := gorm.Open("mysql", baseLogin)
	if err != nil {
		c.Header("Access-Control-Allow-Origin", "*")
		c.String(500, "failed to open database")
		log.Println("Connection Failed to Open", err)
	}
	defer db.Close()
	var products []Product
	db.Find(&products)
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, products)
}

func handlerProductsGroup(c *gin.Context) {
	db, err := gorm.Open("mysql", baseLogin)
	if err != nil {
		c.Header("Access-Control-Allow-Origin", "*")
		c.String(500, "failed to open database")
		log.Println("Connection Failed to Open", err)
	}
	defer db.Close()
	var products []Product
	db.Where("category = ?", c.Param("category")).Find(&products)
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, products)
}

func handlerProductId(c *gin.Context) {
	db, err := gorm.Open("mysql", baseLogin)
	if err != nil {
		c.Header("Access-Control-Allow-Origin", "*")
		c.String(500, "failed to open database")
		log.Println("Connection Failed to Open", err)
	}
	defer db.Close()
	var products []Product
	db.Where("id = ?", c.Param("id")).First(&products)
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, products)
}

func handlerProductAdd(c *gin.Context) {
	db, err := gorm.Open("mysql", baseLogin)
	if err != nil {
		c.Header("Access-Control-Allow-Origin", "*")
		c.String(500, "failed to open database")
		log.Println("Connection Failed to Open", err)
	}
	defer db.Close()
	var products Product
	_ = c.ShouldBindJSON(&products)
	db.Create(&products)
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, "accept")
}

func handlerProductUpdate(c *gin.Context) {
	db, err := gorm.Open("mysql", baseLogin)
	if err != nil {
		c.Header("Access-Control-Allow-Origin", "*")
		c.String(500, "failed to open database")
		log.Println("Connection Failed to Open", err)
	}
	defer db.Close()
	var products Product
	_ = c.ShouldBindJSON(&products)
	db.Model(&products).Where("id = ?", c.Param("id")).Update(&products)
	c.Header("Access-Control-Allow-Origin", "*")
	c.String(200, "accept")
}

func handlerProductDel(c *gin.Context) {
	db, err := gorm.Open("mysql", baseLogin)
	if err != nil {
		c.Header("Access-Control-Allow-Origin", "*")
		c.String(500, "failed to open database")
		log.Println("Connection Failed to Open", err)
	}
	defer db.Close()
	db.Where("id = ?", c.Param("id")).Delete(Product{})
	c.Header("Access-Control-Allow-Origin", "*")
	c.String(200, "accept")
}

func handlerImageUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.Header("Access-Control-Allow-Origin", "*")
		c.String(500, err.Error()+"qwe")
		fmt.Println(err)
	}
	path := "./img/" + file.Filename
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		c.Header("Access-Control-Allow-Origin", "*")
		c.String(500, err.Error())
		fmt.Println(err)
	}
	time.Sleep(time.Second)
	c.Header("Access-Control-Allow-Origin", "*")
	c.String(200, file.Filename)
}
