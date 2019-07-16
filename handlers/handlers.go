package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sagaraglawe/miniProject/inits"
	"github.com/sagaraglawe/miniProject/migrations"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func Admindata(c *gin.Context){

	var ten[] migrations.Product
	inits.Db.Limit(4).Find(&ten)
	for i:=0;i<len(ten);i++{
		fmt.Println(ten[i])
	}

}


func Userdata(c *gin.Context){
	var ten[] migrations.Product
	inits.Db.Limit(4).Select([]string{"user_id","price","order_placed"}).Find(&ten)
	for i:=0;i<len(ten);i++{
		fmt.Println(ten[i])
	}

}


func Showdata(c *gin.Context){
	fmt.Println("you are in the showdata")
	fmt.Println(c.Query("name"))
}


func AdminShow(c *gin.Context){
	user:=c.Query("name");
	tt:=[]migrations.Product{}
	inits.Db.Where("username=?",user).Find(&tt)
	zz:=[] json.RawMessage{}
	var wg sync.WaitGroup
	for i:=0;i<len(tt);i++{
		//zz.append()
		//zz=append(zz,tt[i].Declare)
		wg.Add(1)
		go func (message []byte){
			zz=append(zz,message)
			wg.Done()
		}(tt[i].Declare)
	}
	wg.Wait()
	c.JSON(http.StatusOK,zz)
	return

}

func UserShow(c *gin.Context) {
	user := c.Query("name");
	tt := []migrations.Product{}
	inits.Db.Where("username=?", user).Find(&tt)

	var zz [] json.RawMessage
	var wg sync.WaitGroup
	for i := 0; i < len(tt); i++ {

		wg.Add(1)
		go func(message []byte){
			var pp map[string]interface{}
			err:=json.Unmarshal(message,&pp)
			if err!=nil{
				log.Panic(err)
			}

			for k, _ := range pp {
				if k == "phone_no" {
					pp[k] = "********"
				}
				if k == "password" {
					pp[k] = "**********"
				}
			}

			tpt, _ := json.Marshal(pp)

			zz = append(zz, tpt)
			wg.Done()
		}(tt[i].Declare)


	}
	wg.Wait()
	c.JSON(http.StatusOK, zz)
	return
}



//
//func UserShow(c *gin.Context){
//	user:=c.Query("name");
//	tt:=[]migrations.Product{}
//	inits.Db.Where("username=?",user).Find(&tt)
//
//	//fmt.Printf("%T",tt[0].Declare)
//
//	var zz [] migrations.Tproduct
//
//	for i:=0;i<len(tt);i++{
//		pp:=migrations.Tproduct{}
//		json.Unmarshal(tt[i].Declare,&pp)
//		pp.PhoneNo=pp.PhoneNo[:2] + "******" + pp.PhoneNo[8:]
//		pp.Password="********"
//		zz=append(zz,pp)
//	}
//
//	c.JSON(http.StatusOK,zz)
//	return
//}


func StoreData(c *gin.Context){
	path:=c.Query("path")

	//fmt.Println(path)

	var prod []migrations.Product

	//reading the Json file into the file
	file, _ := ioutil.ReadFile(path)

	//converting the Json file into the slice of bytes
	err:=json.Unmarshal([]byte(file),&prod)

	//converting the entire fields to set into the Declare column
	//var temp []Tproduct
	//err=json.Unmarshal([]byte(file),&temp)



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
	inits.Db.AutoMigrate(&migrations.Product{})

	//filling the database table
	for i:=0;i<len(pp);i++{
		//this is getting the entire field and setting that to the declare field
		//, _ := json.Marshal(temp[i])

		byte,_:=json.Marshal(pp[i])

		//setting the prod[i] declare field
		prod[i].Declare=byte
		//now sending the entry to the database
		inits.Db.Create(prod[i])
	}
}



//Handling the HTML file
func UploadFile(c *gin.Context) {
	//it is for calling the html file to be loaded
	c.HTML(200, "index.html", nil)
}

func TakeFile(c *gin.Context){
//this is for receiving the uploaded content in the name file
//the "myFile" comes from the index.html where we used this attribute to represent the name of the file
	file,err:=c.FormFile("myFile")

//if error happens while uploading then panic
	if err!=nil{
		fmt.Println(err)
		log.Panic(err)
	}
//it is to store the file thus get into the destination directory required and set by us
	err=c.SaveUploadedFile(file,"JsonFile/"+file.Filename)

	if err!=nil{
		log.Fatal(err)
	}

	path:="JsonFile/"+file.Filename

	//fmt.Println(path)

	var prod []migrations.Product

	//reading the Json file into the file
	file1, _:= ioutil.ReadFile(path)

	//converting the Json file into the slice of bytes
	err=json.Unmarshal([]byte(file1),&prod)

	//converting the entire fields to set into the Declare column
	//var temp []Tproduct
	//err=json.Unmarshal([]byte(file),&temp)



	//some modifications

	var pp []map[string]interface{}

	err=json.Unmarshal([]byte(file1),&pp)

	//if error happens then call panic
	if err!=nil{
		fmt.Println(err)
	}

	//for k,v:=range pp{
	//	fmt.Println(k,v)
	//}

	//migration to create the table in the Database
	inits.Db.AutoMigrate(&migrations.Product{})

	//filling the database table
	for i:=0;i<len(pp);i++{
		//this is getting the entire field and setting that to the declare field
		//, _ := json.Marshal(temp[i])

		byte,_:=json.Marshal(pp[i])

		//setting the prod[i] declare field
		prod[i].Declare=byte
		//now sending the entry to the database
		inits.Db.Create(prod[i])
	}
}