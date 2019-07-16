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


	//always load the entire HTML files or say frontend files before its use
	r.LoadHTMLFiles("index.html")

	r.GET("/admin",handlers.Admindata)
	r.GET("/user",handlers.Userdata)
	r.POST("/admin/useid",handlers.Showdata)
	r.POST("admin/show",handlers.AdminShow)
	r.POST("user/show",handlers.UserShow)

	r.POST("/store",handlers.StoreData)

	//this to handle the html files
	r.GET("/uploadfile",handlers.UploadFile)
	r.POST("/uploadfile",handlers.TakeFile)


	err:=r.Run(":8080")
	if err!=nil{
		log.Fatal(err)
	}

	defer inits.Db.Close()

}