package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sagaraglawe/miniProject/handlers"
	"github.com/sagaraglawe/miniProject/inits"
	"github.com/sagaraglawe/miniProject/migrations"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"log"
)

func main() {


//opening the Json file
//	jsonFile, err := os.Open("JsonFile/sample.json")
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println("Successfully Opened sample.json")
//	defer jsonFile.Close()



//it is for using the data of the sample.json into that of the prod of Product type
var prod []migrations.Product

//reading the Json file into the file
file, _ := ioutil.ReadFile("JsonFile/sample.json")

//converting the Json file into the slice of bytes
err:=json.Unmarshal([]byte(file),&prod)

//converting the entire fields to set into the Declare column
var temp []migrations.Tproduct
err=json.Unmarshal([]byte(file),&temp)

//if error happens then call panic
if err!=nil{
log.Panic(err)
}


//migration to create the table in the Database
//db.AutoMigrate(&Product{})

//filling the database table
for i:=0;i<len(temp);i++{
//this is getting the entire field and setting that to the declare field
byte2, _ := json.Marshal(temp[i])
//setting the prod[i] declare field
prod[i].Declare=byte2
//now sending the entry to the database
//db.Create(prod[i])
}

//getting the data

user:=migrations.Product{}
inits.Db.Where("username=?","Rishabh").First(&user)
fmt.Println(string(user.Declare))

//creating the route general
r:=gin.Default()
r.GET("/admin",handlers.Admindata)
r.GET("/user",handlers.Userdata)
r.POST("/admin/useid",handlers.Showdata)
r.POST("admin/show",handlers.AdminShow)
r.POST("user/show",handlers.UserShow)
r.Run(":8080")


defer inits.Db.Close()

}