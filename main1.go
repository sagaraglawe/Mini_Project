package main


import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"github.com/jinzhu/gorm"
	"net/http"

	//"net/http"
	"log"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


//this is making the pointer for the database connection which is Global
var db *gorm.DB



//init function always runs before start of the program
func init(){

	//error to handle the error while the connection is in Process
	var err error
	db,err =gorm.Open("mysql", "root:Sagaraglawe@26@tcp(127.0.0.1:3306)/books?charset=utf8&parseTime=True")
	if err!=nil{
		log.Panic(err)
	}
	log.Println("you are good")
}



//the structure used for the migration and creation of table in the database


//Kept only those fields which needs to be searchable

type Product struct {
	Username        string `json:"username"`
	UserID          int    `json:"user_id"`
	Price           int    `json:"price"`
	PhoneNo         string `json:"phone_no"`
	OrderPlaced     string `json:"order_placed"`
	Password        string `json:"password"`
	Declare         []byte
}

type Tproduct struct {
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

type Obfuscate struct{
	PhoneNo string
	Password string

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
	var prod []Product

	//reading the Json file into the file
	file, _ := ioutil.ReadFile("JsonFile/sample.json")

	//converting the Json file into the slice of bytes
	err:=json.Unmarshal([]byte(file),&prod)

	//converting the entire fields to set into the Declare column
	var temp []Tproduct
	err=json.Unmarshal([]byte(file),&temp)



	//some modifications

	var pp []map[string]interface{}

	err=json.Unmarshal([]byte(file),&pp)

	//if error happens then call panic
	if err!=nil{
		fmt.Println(err)
	}

	//for k,v:=range pp{
	//	fmt.Println(k,v)
	//}

	//migration to create the table in the Database
	//db.AutoMigrate(&Product{})

	//filling the database table
	for i:=0;i<len(pp);i++{
		//this is getting the entire field and setting that to the declare field
		//, _ := json.Marshal(temp[i])

		byte,_:=json.Marshal(pp[i])

		//setting the prod[i] declare field
		prod[i].Declare=byte
		//now sending the entry to the database
		//db.Create(prod[i])
	}

	//getting the data

	//user:=Product{}
	//db.Where("username=?","Rishabh").First(&user)
	//fmt.Println(string(user.Declare))

	//creating the route general
	r:=gin.Default()
	r.GET("/admin",admindata)
	r.GET("/user",userdata)
	r.POST("/admin/useid",showdata)
	r.POST("admin/show",adminShow)
	r.POST("user/show",userShow)
	r.Run(":8080")


	defer db.Close()

}



func admindata(c *gin.Context){

	var ten[] Product
	db.Limit(4).Find(&ten)
	for i:=0;i<len(ten);i++{
		fmt.Println(ten[i])
	}

}


func userdata(c *gin.Context){
	var ten[] Product
	db.Limit(4).Select([]string{"user_id","price","order_placed"}).Find(&ten)
	for i:=0;i<len(ten);i++{
		fmt.Println(ten[i])
	}

}


func showdata(c *gin.Context){
	fmt.Println("you are in the showdata")
	fmt.Println(c.Query("name"))
}


func adminShow(c *gin.Context){
	user:=c.Query("name");
	tt:=[]Product{}
	db.Where("username=?",user).Find(&tt)
	zz:=[] json.RawMessage{}
	for i:=0;i<len(tt);i++{
		//zz.append()
		zz=append(zz,tt[i].Declare)
	}
	c.JSON(http.StatusOK,zz)
	return

}


func userShow(c *gin.Context){
	user:=c.Query("name");
	tt:=[]Product{}
	db.Where("username=?",user).Find(&tt)

	var zz [] json.RawMessage

	for i:=0;i<len(tt);i++{
	var pp map[string]interface{}
		json.Unmarshal(tt[i].Declare,&pp)

		for k,_:=range pp{
			if k=="username"{
				pp[k]="********"
			}
		}

		tpt,_:=json.Marshal(pp)

		zz=append(zz,tpt)
	}

	c.JSON(http.StatusOK,zz)
	return
}



//Q.What mean by Json.rawmessagae?
//Q.Difference between json.Marshal and marshalJSON
