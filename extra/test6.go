package extra


import(
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"os"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"

)


//this is for creating the db connection Global
var db *gorm.DB

//init function always starts in the start of the fuction run
func init(){

	//define error and database connection
	var err error
	db,err =gorm.Open("mysql", "root:Sagaraglawe@26@tcp(127.0.0.1:3306)/books?charset=utf8&parseTime=True")
	if err!=nil{
		log.Println("you are not good")
	}
	log.Println("you are good")


//do not use it here because if so then when the init finishes it will call the defer function
//	defer db.Close()



}

//this struct define what we are having in the database
type product struct {
	//gorm.Model
	Username        string `json:"username"`
	OrderValid      string `json:"order_valid"`
	UserID          int    `json:"user_id"`
	Price           int    `json:"price"`
	PhoneNo         string `json:"phone_no"`
	OrderPlaced     string `json:"order_placed"`
	Password        string `json:"password"`
	OrderNumber     int    `json:"order_number"`
	ProductWeight   int    `json:"product_weight"`
	NumInstallments int    `json:"num_installments"`
}


func main(){


//opening the Json file
	jsonFile, err := os.Open("JsonFile/sample.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened sample.json")
	defer jsonFile.Close()



//filling table into the database
	var prod []product

	file, _ := ioutil.ReadFile("JsonFile/sample.json")

	json.Unmarshal([]byte(file),&prod)

	//this is for importing the data in the database
	//for i:=0;i<len(prod);i++{
	//	db.Create(prod[i])
	//}


	//fmt.Println(len(prod))

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