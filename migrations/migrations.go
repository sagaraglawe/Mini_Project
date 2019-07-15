package migrations


type Product struct {
	Username        string `json:"username"`
	UserID          int    `json:"user_id"`
	Price           int    `json:"price"`
	PhoneNo         string `json:"phone_no"`
	OrderPlaced     string `json:"order_placed"`
	Password        string `json:"password"`
	Declare         []byte
}

//type Tproduct struct {
//	Username        string `json:"username"`
//	OrderValid      string `json:"order_valid"`
//	UserID          int    `json:"user_id"`
//	Price           int    `json:"price"`
//	PhoneNo         string `json:"phone_no"`
//	OrderPlaced     string `json:"order_placed"`
//	Password        string `json:"password"`
//	OrderNumber     int    `json:"order_number"`
//	ProductWeight   int    `json:"product_weight"`
//	NumInstallments int    `json:"num_installments"`
//}
