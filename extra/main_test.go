package extra

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"log"
	//"os"
)



//this is making the pointer for the database connection which is Global
var db *gorm.DB




//the structure used for the migration and creation of table in the database


//Kept only those fields which needs to be searchable

type Product struct {
	Username        string `json:"username"`
	UserID          int    `json:"user_id"`
	Price           int    `json:"price"`
	PhoneNo         string `json:"phone_no"`
	OrderPlaced     string `json:"order_placed"`
	Password        string `json:"password"`
	Declare         json.RawMessage
}


//init function always runs before start of the program
func init(){

	//error to handle the error while the connection is in Process
	var err error
	db,err =gorm.Open("mysql", "root:Sagaraglawe@26@tcp(127.0.0.1:3306)/books?charset=utf8&parseTime=True")
	if err!=nil{
		log.Println("you are not good")
		log.Panic(err)
	}
	log.Println("you are good")
}





func main() {


//opening the Json file
	//	jsonFile, err := os.Open("JsonFile/sample.json")
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Println("Successfully Opened sample.json")
	//	defer jsonFile.Close()



//it is for using the data of the sample.json into that of the prod of Product type
	var prod []product

//reading the Json file into the file
	file, _ := ioutil.ReadFile("JsonFile/sample.json")

//converting the Json file into the slice of bytes
	err:=json.Unmarshal([]byte(file),&prod)

//if error happens then call panic
	if err!=nil{
		log.Panic(err)
	}


//migration to create the tabel in the Database
	db.AutoMigrate(&Product{})

//filling the database table
	for i:=0;i<len(prod);i++{
		db.Create(prod[i])
		byte2, _ := json.Marshal(prod[i])
		db.Model(&prod[i]).UpdateColumn("Declare", byte2)
	}


//creating the route general
	r:=gin.Default()
	r.GET("/admin",admindata)
	r.GET("/user",userdata)
	r.Run(":8080")


	defer db.Close()

}



func admindata(c *gin.Context){

var ten[] product
db.Limit(4).Find(&ten)
for i:=0;i<len(ten);i++{
fmt.Println(ten[i])
}

}


func userdata(c *gin.Context){
	var ten[] product
	db.Limit(4).Select([]string{"user_id","price","order_placed"}).Find(&ten)
	for i:=0;i<len(ten);i++{
		fmt.Println(ten[i])
	}

}