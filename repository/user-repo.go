package repository

import (
	"context"

	"github.com/dafiqarba/be-payroll/model"
	"github.com/dafiqarba/be-payroll/utils"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// TODO: search how to return a populated struct from function, if struct is defined

type UserRepo interface {
	//Create
	CreateUser(ctx context.Context, tx *sqlx.Tx, u model.User) (string, error)
	//Read
	FindByEmail(ctx context.Context, email string) (model.UserResponse, error)
	GetUserList(ctx context.Context) ([]model.User, error)
	GetUserDetail(ctx context.Context, id uuid.UUID) (model.UserDetailModel, error)
}

type userConnection struct {
	connection *sqlx.DB
}

func NewUserRepo(dbConn *sqlx.DB) UserRepo {
	return &userConnection{
		connection: dbConn,
	}
}

func (db *userConnection) GetUserList(ctx context.Context) ([]model.User, error) {
	//Variable to store collection of users
	users := make([]model.User, 0)

	//Execute SQL Query
	query := `SELECT * FROM users`
	rows, err := db.connection.QueryxContext(ctx, query)

	//Error Handling
	if err != nil {
		utils.LogError("Repo", "func GetUserList", err)
		return users, err
	}

	//Close the Execution of SQL Query
	defer rows.Close()

	//Iterate over all available rows and strore the data
	for rows.Next() {
		var user model.User
		// scan and assign into destination variable
		err = rows.Scan(
			&user.User_id,
			&user.Email,
			&user.Password,
			&user.Name,
			&user.Position_id,
			&user.Nik,
			&user.Role_id,
		)

		if err != nil {
			utils.LogError("Repo", "GetUserList scan data", err)
			return users, err
		}
		// append to users slice
		users = append(users, user)
	}

	utils.CloseDB(rows)
	// returns populated data
	return users, err
}

func (db *userConnection) GetUserDetail(ctx context.Context, id uuid.UUID) (model.UserDetailModel, error) {

	var (
		userDetail model.UserDetailModel
	)

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
	err := db.connection.QueryRowxContext(
		ctx,
		query,
		id,
	).Scan(
		&userDetail.User_id,
		&userDetail.Username,
		&userDetail.Name,
		&userDetail.Position_id,
		&userDetail.Nik,
		&userDetail.Role_id,
		&userDetail.Role_name,
		&userDetail.Position_name,
	)

	//Err Handling
	if err != nil {
		utils.LogError("Repo", "func GetUserDetail", err)
		return userDetail, err
	}

	// returns login data
	return userDetail, err
}

func (db *userConnection) CreateUser(ctx context.Context, tx *sqlx.Tx, u model.User) (string, error) {
	//Variable that holds registered user email
	var (
		createdUser string
	)

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
	err := db.connection.QueryRowxContext(
		ctx,
		query,
		u.Username,
		u.Name,
		u.Password,
		u.Email,
		u.Nik,
		u.Role_id,
		u.Position_id,
	).Scan(
		&createdUser,
	)

	//Err Handling
	if err != nil {
		utils.LogError("Repo", "func CreateUser", err)
		return createdUser, err
	}

	//Returns registered user email and nil error
	return createdUser, err
}

func (db *userConnection) FindByEmail(ctx context.Context, email string) (model.UserResponse, error) {
	// Var to be populated with user data
	var (
		userData model.UserResponse
	)

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
	err := db.connection.QueryRowxContext(
		ctx,
		query,
		email,
	).Scan(
		&userData.User_id,
		&userData.Username,
		&userData.Email,
		&userData.Password,
		&userData.Role_id,
	)

	//Err Handling
	if err != nil {
		utils.LogError("Repo", "func FindByEmail", err)
		return userData, err
	}

	// returns login data
	return userData, err
}
