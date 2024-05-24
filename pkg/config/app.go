package config

import(
	// "github.com/jinzhu/gorm"
	// _ "gorm.io/driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db * gorm.DB
)

func Connect(){
	// import "gorm.io/driver/mysql"
	// refer: https://gorm.io/docs/connecting_to_the_database.html#MySQL
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	// d, err := gorm.Open("mysql", "user:pass@/simplerest?charset=utf8mb4&parseTime=True&loc=Local")
	// import "gorm.io/driver/mysql"
	// refer: https://gorm.io/docs/connecting_to_the_database.html#MySQL
	dsn := "user:pass@tcp(localhost:9090)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	

	if err!= nil{
		panic(err)
	}
	db=d
}

func GetDB() *gorm.DB{
	return db;
}