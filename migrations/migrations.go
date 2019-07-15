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
