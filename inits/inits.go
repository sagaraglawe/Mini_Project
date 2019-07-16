package inits

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

//this is making the pointer for the database connection which is Global
	var Db *gorm.DB



//init function always runs before start of the program
func init(){

//error to handle the error while the connection is in Process
	var err error
	Db,err =gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/books?charset=utf8&parseTime=True")
	if err!=nil{
		log.Panic(err)
	}
	log.Println("you are good")
}
