package main

import (
	"fmt"
	"time"
	"strconv"
	"strings"
	// "net"
	"net/http"
	// "net/http/httputil"
	// "net/url"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// "github.com/spf13/viper"
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
// GO没有类，只有结构体和结构方法
func (Product) TableName() string {
	return "product"
}


// gorm 字段映射 首字母需要大写
type BloodPressure struct {
	// gorm.Model
	Id int
	Dp int
	Sp int
	HeartRate int
	CreateTime time.Time
}
// GO没有类，只有结构体和结构方法
func (BloodPressure) TableName() string {
	return "blood_pressure"
}

type Database struct {
	host string
	username string
	dbname string
	password string
	port int
}

var database Database
var db *gorm.DB
func init() {
	// viper.SetConfigName("config")
	// viper.SetConfigType("yaml")
	// viper.AddConfigPath("./conf/")

	// err := viper.ReadInConfig()
	// if err != nil {
	// 	panic("failed to read conf")
	// }
	// // 初始化数据库配置
	// host := viper.GetString("database.host")
	// port := viper.GetString("database.port")
	// username := viper.GetString("database.username")
	// dbname := viper.GetString("database.dbname")
	// password := viper.GetString("database.password")
	// fmt.Println("config init")
	// portInt, err := strconv.Atoi(port);
	database = Database{
		host: "172.18.153.61",
		port: 54321,
		username: "postgres",
		dbname: "postgres",
		password: "postgres",
		// host: host,
		// port: portInt,
		// username: username,
		// dbname: dbname,
		// password: password,
	}
	fmt.Println(database)

	dialect := GetGormDialect()
	dbInstance, err := gorm.Open("postgres", dialect)
	if err != nil {
		panic("failed to connect database")
	}
	db = dbInstance
}

func GetGormDialect() string {
	s := []string{"host=", database.host, 
		" port=", strconv.Itoa(database.port),
		" user=", database.username,
		" dbname=", database.dbname,
		" password=", database.password,
		" sslmode=disable"}
	return strings.Join(s, "")
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

	router.GET("/blood-pressure", func(c *gin.Context) {

		var records []BloodPressure
		db.Find(&records)
		if len(records) > 0 {
			fmt.Printf("%+v\n", records[0])
		}
		// { 'date': '2019-11-16', '收缩压SP': 127, '舒张压DP': 84, '心率': 69 },
		var list map[string]interface{} = make(map[string]interface{})
		for i := 0; i < len(records); i++ {
			list[strconv.Itoa(records[i].Id)] = records[i]
		}
		c.JSON(http.StatusOK, gin.H{
			"data": list,
			"status": 200,
		})
	})
	router.POST("/blood-pressure", func(c *gin.Context) {
		var json BloodPressure
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// dp := strconv.Atoi(c.PostForm("dp"))
		bloodPressure := BloodPressure{
			Dp: json.Dp,
			Sp: json.Sp,
			HeartRate: json.HeartRate,
			CreateTime: time.Now(),
		}
		db.Create(&bloodPressure)
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}