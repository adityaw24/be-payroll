package dto

//Login data model
type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//Register data model
type RegisterUser struct {
	Name        string `json:"name"`
	Username	string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Nik         string `json:"nik"`
	Role_id     int    `json:"role_id"`
	Position_id int    `json:"position_id"`
}
