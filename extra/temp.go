package extra

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"


)

var db *gorm.DB

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

func main() {


	var err error
//creating the database connection
	db,err = gorm.Open("mysql", "root:Sagaraglawe@26@tcp(127.0.0.1:3306)/books?charset=utf8&parseTime=True")
	if err!=nil{
		log.Println("you are not good")
	}
	log.Println("you are good")



//opening the Json file
	jsonFile, err := os.Open("JsonFile/sample.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened sample.json")


//filling table into the database
	var prod []product

	file, _ := ioutil.ReadFile("JsonFile/sample.json")
	json.Unmarshal([]byte(file),&prod)

	fmt.Println(len(prod))
	//for i:=0;i<len(prod);i++{
	//	db.Create(prod[i])
	//}

	tempp()

	defer jsonFile.Close()
	defer db.Close()
}

func tempp(){
	var ten product
	db.First(&ten)
	fmt.Println(ten)
}