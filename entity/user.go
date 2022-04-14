package entity

//User represents users table in the database
type User struct {
	User_id     int    `json:"user_id"`
	Username    string `json:"username"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Nik         string `json:"nik"`
	Role_id     int    `json:"role_id"`
	Position_id int    `json:"position"`
}

type UserDetailModel struct {
	User_id       int    `json:"user_id"`
	Username      string `json:"username"`
	Name          string `json:"name"`
	Position_id   int    `json:"position_id"`
	Nik           string `json:"nik"`
	Role_id       int    `json:"role_id"`
	Role_name     string `json:"role"`
	Position_name string `json:"position"`
}

type UserLogin struct {
	User_id  int    `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Role_id  int    `json:"role_id"`
	Token    string `json:"token"`
}
