package repository

import (
	"context"

	"github.com/dafiqarba/be-payroll/model"
	"github.com/dafiqarba/be-payroll/utils"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type RoleRepo interface {
	//Create
	CreateRole(ctx context.Context, tx *sqlx.Tx, u model.Role) (uuid.UUID, error)
	//Read
	GetRoleList(ctx context.Context) ([]model.Role, error)
	GetRoleDetail(ctx context.Context, id uuid.UUID) (model.Role, error)
	//Update
	UpdateRole(ctx context.Context, tx *sqlx.Tx, u model.Role) (uuid.UUID, error)
	//Delete
	DeleteRole(ctx context.Context, tx *sqlx.Tx, id uuid.UUID) (uuid.UUID, error)
}

type roleRepository struct {
	db *sqlx.DB
}

func NewRoleRepo(dbConn *sqlx.DB) RoleRepo {
	return &roleRepository{
		db: dbConn,
	}
}

func (r *roleRepository) GetRoleList(ctx context.Context) ([]model.Role, error) {
	//Variable to store collection of roles
	roles := make([]model.Role, 0)

	//Execute SQL Query
	query := `
		SELECT 
			r.role_id,
			r.name,
			r.created_at, 
			r.updated_at, 
			r.is_delete 
		FROM 
			roles r 
		ORDER BY name ASC;
		`
	rows, err := r.db.QueryxContext(ctx, query)

	//Error Handling
	if err != nil {
		utils.LogError("Repo", "func GetRoleList", err)
		return roles, err
	}

	//Close the Execution of SQL Query
	defer rows.Close()

	//Iterate over all available rows and strore the data
	for rows.Next() {
		var role model.Role
		// scan and assign into destination variable
		err = rows.Scan(
			&role.Role_id,
			&role.Name,
			&role.CreatedAt,
			&role.UpdatedAt,
			&role.Is_delete,
		)

		if err != nil {
			utils.LogError("Repo", "GetRoleList scan data", err)
			return roles, err
		}
		// append to roles slice
		roles = append(roles, role)
	}

	utils.CloseDB(rows)
	// returns populated data
	return roles, err
}

func (r *roleRepository) GetRoleDetail(ctx context.Context, id uuid.UUID) (model.Role, error) {

	var (
		role model.Role
	)

	//SQL Query
	query := `
		SELECT 
			r.role_id,
			r.name,
			r.created_at, 
			r.updated_at, 
			r.is_delete 
		FROM 
			roles r
		WHERE r.role_id=$1;
		`

	//Execute SQL Query
	err := r.db.QueryRowxContext(
		ctx,
		query,
		id,
	).Scan(
		&role.Role_id,
		&role.Name,
		&role.CreatedAt,
		&role.UpdatedAt,
		&role.Is_delete,
	)

	//Err Handling
	if err != nil {
		utils.LogError("Repo", "func GetRoleDetail", err)
		return role, err
	}

	// returns login data
	return role, err
}

func (r *roleRepository) CreateRole(ctx context.Context, tx *sqlx.Tx, u model.Role) (uuid.UUID, error) {
	//Variable that holds registered user email
	var (
		role_id uuid.UUID
	)

	//Query
	query := `
		INSERT INTO 
			roles (name) 
		VALUES
			($1)
		RETURNING role_id
			;
	`
	//Execute query and Populating createdUser variable
	err := tx.QueryRowxContext(
		ctx,
		query,
		u.Name,
	).Scan(
		&role_id,
	)

	//Err Handling
	if err != nil {
		utils.LogError("Repo", "func CreateRole", err)
		return role_id, err
	}

	//Returns registered user email and nil error
	return role_id, err
}

func (r *roleRepository) UpdateRole(ctx context.Context, tx *sqlx.Tx, u model.Role) (uuid.UUID, error) {
	var (
		role_id uuid.UUID
	)

	//Query
	query := `
		UPDATE 
			roles
		SET
			name = $1,
			updated_at = now()
		VALUES
			($2)
		WHERE
			role_id = $1
		RETURNING role_id
			;
	`

	err := tx.QueryRowxContext(
		ctx,
		query,
		u.Role_id,
		u.Name,
	).Scan(
		&role_id,
	)

	//Err Handling
	if err != nil {
		utils.LogError("Repo", "func UpdateRole", err)
		return role_id, err
	}

	return role_id, err
}

func (r *roleRepository) DeleteRole(ctx context.Context, tx *sqlx.Tx, id uuid.UUID) (uuid.UUID, error) {
	var (
		role_id uuid.UUID
	)

	//Query
	query := `
		UPDATE 
			roles
		SET
			updated_at = now(),
			is_delete = true
		WHERE 
			role_id = $1
		RETURNING role_id
			;
	`

	err := tx.QueryRowxContext(
		ctx,
		query,
		id,
	).Scan(
		&role_id,
	)

	//Err Handling
	if err != nil {
		utils.LogError("Repo", "func DeleteRole", err)
		return role_id, err
	}

	return role_id, err
}
