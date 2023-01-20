package config

import(
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

func Connect(db_user string, db_password string, db_host string, db_port string, db_database string){
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", db_user, db_password, db_host, db_port, db_database)
	d, err := gorm.Open("mysql", DBURL)
	if err != nil{
		panic(err)
	} else {
		fmt.Println("Connected To DB")
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}