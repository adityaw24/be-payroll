package repository

import (
	"database/sql"
	"log"

	"github.com/dafiqarba/be-payroll/entity"
	_ "github.com/lib/pq"
)

// TODO: search how to return a populated struct from function, if struct is defined

type UserRepo interface {
	//Create
	CreateUser(user entity.User) (string, error)
	//Read
	FindByEmail(email string) (entity.UserLogin, error)
	GetUserList() ([]entity.User, error)
	GetUserDetail(id int) (entity.UserDetailModel, error)
}

type userConnection struct {
	connection *sql.DB
}

func NewUserRepo(dbConn *sql.DB) UserRepo {
	return &userConnection{
		connection: dbConn,
	}
}

func (db *userConnection) GetUserList() ([]entity.User, error) {
	//Variable to store collection of users
	var users []entity.User
	//Execute SQL Query
	rows, err := db.connection.Query(`SELECT * FROM users`)
	//Error Handling
	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}
	log.Println("|  Requst received")
	//Close the Execution of SQL Query
	defer rows.Close()
	//Iterate over all available rows and strore the data
	for rows.Next() {
		var user entity.User
		// scan and assign into destination variable
		err = rows.Scan(&user.User_id, &user.Email, &user.Password, &user.Name, &user.Position_id, &user.Nik, &user.Role_id)
		if err != nil {
			log.Fatalf("tidak bisa mengambil data. %v", err)
		}
		// append to users slice
		users = append(users, user)
	}
	// returns populated data
	return users, err
}

func (db *userConnection) GetUserDetail(id int) (entity.UserDetailModel, error) {

	var userDetail entity.UserDetailModel
	//SQL Query
	query := `
		SELECT 
			u.user_id, u.username, u.name, u.position_id, u.nik, u.role_id, r.role_name, p.position_name 
		FROM users AS u 
			INNER JOIN roles AS r 
				ON r.role_id = u.role_id 
			INNER JOIN positions AS p 
				ON p.position_id = u.position_id 
		WHERE user_id=$1`
	//Execute SQL Query
	row := db.connection.QueryRow(query,id)
	err := row.Scan(&userDetail.User_id, &userDetail.Username, &userDetail.Name, &userDetail.Position_id, &userDetail.Nik, &userDetail.Role_id, &userDetail.Role_name, &userDetail.Position_name)
	
	//Err Handling
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("| "+err.Error())
			return userDetail, err
		} else {
			log.Println("| "+err.Error())
			return userDetail, err
		}
	}
	
	// returns login data
	return userDetail, err
}


func (db *userConnection) CreateUser (u entity.User) (string, error) {
	//Variable that holds registered user email
	var createdUser string
	//Query
	query := `
		INSERT INTO 
			users (username, name, password, email, nik, role_id, position_id) 
		VALUES
			($1, $2, $3, $4, $5, $6, $7)
		RETURNING email
			;
	`
	//Execute query and Populating createdUser variable
	err := db.connection.QueryRow(query, u.Username, u.Name, u.Password, u.Email, u.Nik, u.Role_id, u.Position_id).Scan(&createdUser)
	
	if err != nil {
		log.Println("| " + err.Error())
		return "", err
	}
	//Returns registered user email and nil error
	return createdUser, err
}

func (db *userConnection) FindByEmail (emailToCheck string) (entity.UserLogin, error) {
	// Var to be populated with user data
	var userData entity.UserLogin
	//Query
	query := `
		SELECT 
			u.user_id, u.username, u.email, u.password, u.role_id
		FROM
			users AS u
		WHERE
			email = $1;
	`
	//Execute
	row := db.connection.QueryRow(query, emailToCheck)
	err := row.Scan(&userData.User_id, &userData.Username, &userData.Email, &userData.Password, &userData.Role_id)
	//Err Handling
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("| "+err.Error())
			return userData, err
		} else {
			log.Println("| "+err.Error())
			return userData, err
		}
	}
	
	// returns login data
	return userData, err
}
