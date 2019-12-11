package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	// "net"
	"net/http"
	// "net/http/httputil"
	// "net/url"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// gorm 字段映射 首字母需要大写
type Product struct {
	// gorm.Model
	Id int
	Pn string
	Mfs string
	SupplierPn string
	SupplierId int
}
func (Product) TableName() string {
	return "pm_product"
}


func main() {
	// fmt.Println("dbconnect：%s", err)
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!!!")
	})

	router.GET("/param/:param", func(c *gin.Context) {
		param := c.Param("param")
		c.String(http.StatusOK, "Param： %s", param)
	})

	router.GET("/search", func(c *gin.Context) {
		db, err := gorm.Open("postgres", "host=192.168.3.120 port=5432 user=pm dbname=pm password=aaaaaaaa sslmode=disable")
		if err != nil {
			panic("failed to connect database")
		}
		part := c.Query("part")
		var products []Product
		db.Where("pn = ? ", part).Find(&products)
		if len(products) > 0 {
			fmt.Printf("%+v\n", products[0])
		}
		c.JSON(http.StatusOK, gin.H{
			"search": part,
			"data": products,
			"status": 200,
		})
	})

	router.POST("/post", func(c *gin.Context) {
		message := c.PostForm("message")
		c.JSON(http.StatusOK, gin.H{
			"message": message,
			"status": 200,
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}