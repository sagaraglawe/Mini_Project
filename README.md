# Mini Project

* It is oriented towards storing the log file and retrieving the information according to search parameters provided

* MySQL database, GIN framework and GORM used for database Query

* Generalised struct is used to store the information in the database. Declare is storing the entire JSON string

```
type Product struct {
	Username        string `json:"username"`
	UserID          int    `json:"user_id"`
	Price           int    `json:"price"`
	PhoneNo         string `json:"phone_no"`
	OrderPlaced     string `json:"order_placed"`
	Password        string `json:"password"`
	Declare         []byte
}

```

* for uploading the multiple files and store them into the database use the following route

```
http://localhost:8080/multiupload

```
* for showing the data to the user use following route with search parameter name="****"

```
http://localhost:8080/user/show
```

* for showing the data to the admin use the following route with search parameter name="****"

```
http://localhost:8080/admin/show
```
