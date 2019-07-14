package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sagaraglawe/miniProject/inits"
	"github.com/sagaraglawe/miniProject/migrations"
	"net/http"
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
	for i:=0;i<len(tt);i++{
		//zz.append()
		zz=append(zz,tt[i].Declare)
	}
	c.JSON(http.StatusOK,zz)
	return

}


func UserShow(c *gin.Context){
	user:=c.Query("name");
	tt:=[]migrations.Product{}
	inits.Db.Where("username=?",user).Find(&tt)

	//fmt.Printf("%T",tt[0].Declare)

	var zz [] migrations.Tproduct

	for i:=0;i<len(tt);i++{
		pp:=migrations.Tproduct{}
		json.Unmarshal(tt[i].Declare,&pp)
		pp.PhoneNo=pp.PhoneNo[:2] + "******" + pp.PhoneNo[8:]
		pp.Password="********"
		zz=append(zz,pp)
	}

	c.JSON(http.StatusOK,zz)
	return
}
