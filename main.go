package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sagaraglawe/miniProject/handlers"
	"github.com/sagaraglawe/miniProject/inits"
	"log"
)

func main() {

//creating the route general
	r:=gin.Default()
	r.GET("/admin",handlers.Admindata)
	r.GET("/user",handlers.Userdata)
	r.POST("/admin/useid",handlers.Showdata)
	r.POST("admin/show",handlers.AdminShow)
	r.POST("user/show",handlers.UserShow)
	r.POST("/store",handlers.StoreData)
	err:=r.Run(":8080")
	if err!=nil{
		log.Fatal(err)
	}

	defer inits.Db.Close()

}