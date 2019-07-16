package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sagaraglawe/miniProject/inits"
	"github.com/sagaraglawe/miniProject/migrations"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"sync"
)

func AdminShow(c *gin.Context){
//getting the parameter value user with the key name in the post request
	user:=c.Query("name");
	ShowData:=[]migrations.Product{}
//Querying
	inits.Db.Where("username=?",user).Find(&ShowData)
	JsonMessage:=[] json.RawMessage{}
	for i:=0;i<len(ShowData);i++{
			JsonMessage=append(JsonMessage,ShowData[i].Declare)
	}
	c.JSON(http.StatusOK,JsonMessage)
	return
}

func UserShow(c *gin.Context) {
	user := c.Query("name");
	ShowData := []migrations.Product{}
//Query
	inits.Db.Where("username=?", user).Find(&ShowData)

	var JsonMessage [] json.RawMessage
	for i := 0; i < len(ShowData); i++ {
			var pp map[string]interface{}
			err:=json.Unmarshal(ShowData[i].Declare,&pp)
			if err!=nil{
				log.Panic(err)
			}
//securing the data which is not to be display to the user
			for k, _ := range pp {
				if k == "phone_no" {
					pp[k] = "********"
				}
				if k == "password" {
					pp[k] = "**********"
				}
			}

			tpt, _ := json.Marshal(pp)

			JsonMessage = append(JsonMessage, tpt)
	}
	c.JSON(http.StatusOK, JsonMessage)
	return
}

func StoreData(c *gin.Context){
	path:=c.Query("path")
	var prod []migrations.Product

//reading the Json file into the file
	file, _ := ioutil.ReadFile(path)

//converting the Json file into the slice of bytes
	err:=json.Unmarshal([]byte(file),&prod)

//handling unstructured data

	var pp []map[string]interface{}
	err=json.Unmarshal([]byte(file),&pp)

//if error happens then call panic
	if err!=nil{
		fmt.Println(err)
	}

//migration to create the table in the Database
	inits.Db.AutoMigrate(&migrations.Product{})

//filling the database table
	for i:=0;i<len(pp);i++{
//this is getting the entire field and setting that to the declare field

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
	var prod []migrations.Product

//reading the Json file into the file
	file1, _:= ioutil.ReadFile(path)

//converting the Json file into the slice of bytes
	err=json.Unmarshal([]byte(file1),&prod)

//Handling Unstructured data
	var pp []map[string]interface{}

	err=json.Unmarshal([]byte(file1),&pp)

//if error happens then call panic
	if err!=nil{
		fmt.Println(err)
	}

//migration to create the table in the Database
	inits.Db.AutoMigrate(&migrations.Product{})

//filling the database table
	for i:=0;i<len(pp);i++{
//this is getting the entire field and setting that to the declare field

	byte,_:=json.Marshal(pp[i])

//setting the prod[i] declare field
	prod[i].Declare=byte
//now sending the entry to the database
	inits.Db.Create(prod[i])
	}
}

//this is for getting the HTML view
func MultiUpload(c *gin.Context){
	c.HTML(200, "temp.html", nil)
}



func StoreMultiUpload(c *gin.Context){
		form,_:=c.MultipartForm()
		files:=form.File["multiplefiles"]
		inits.Db.AutoMigrate(&migrations.Product{})
		var wg sync.WaitGroup
		for _,file:=range files{
			err:=c.SaveUploadedFile(file,"JsonFile/"+file.Filename)
			if err!=nil{
				log.Fatal(err)
			}
			wg.Add(1)
			go CreateDatabase(file)
			defer wg.Done()
		}

		wg.Wait()

}


func CreateDatabase(file2 *multipart.FileHeader){

	path:="JsonFile/"+file2.Filename

	var prod []migrations.Product

//reading the Json file into the file
	file1, _:= ioutil.ReadFile(path)

//converting the Json file into the slice of bytes
	err:=json.Unmarshal([]byte(file1),&prod)

	var pp []map[string]interface{}

	err=json.Unmarshal([]byte(file1),&pp)

//if error happens then call panic
	if err!=nil{
		fmt.Println(err)
	}

//filling the database table
	for i:=0;i<len(pp);i++{
//this is getting the entire field and setting that to the declare field
	byte,_:=json.Marshal(pp[i])

//setting the prod[i] declare field
	prod[i].Declare=byte
//now sending the entry to the database
	inits.Db.Create(prod[i])
	}
}