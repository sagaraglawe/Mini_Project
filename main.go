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
	r.LoadHTMLFiles("temp.html")

//Show the data based on name query
	r.POST("admin/show",handlers.AdminShow)
	r.POST("user/show",handlers.UserShow)

//store the struct in the database using the link from the Already store file and pass the variable path as the path of the file to be store
	r.POST("/store",handlers.StoreData)

//this to handle the html files
	r.GET("/uploadfile",handlers.UploadFile)
	r.POST("/uploadfile",handlers.TakeFile)

//Uploading multiple files
	r.GET("/multiupload",handlers.MultiUpload)
	r.POST("/multiupload",handlers.StoreMultiUpload)

//it is for running the server over the address htttp://localhost:8080
	err:=r.Run(":8080")
	if err!=nil{
		log.Fatal(err)
	}

	err=inits.Db.Close()
	if err!=nil{
		log.Fatal(err)
	}



}