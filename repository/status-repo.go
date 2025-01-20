package repository

import (
	"context"

	"github.com/dafiqarba/be-payroll/model"
	"github.com/dafiqarba/be-payroll/utils"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type StatusRepo interface {
	//Create
	CreateStatus(ctx context.Context, tx *sqlx.Tx, u model.Status) (uuid.UUID, error)
	//Read
	GetStatusList(ctx context.Context) ([]model.Status, error)
	GetStatusDetail(ctx context.Context, id uuid.UUID) (model.Status, error)
	//Update
	UpdateStatus(ctx context.Context, tx *sqlx.Tx, u model.Status) (uuid.UUID, error)
	//Delete
	DeleteStatus(ctx context.Context, tx *sqlx.Tx, id uuid.UUID) (uuid.UUID, error)
}

type statusRepository struct {
	db *sqlx.DB
}

func NewStatusRepo(dbConn *sqlx.DB) StatusRepo {
	return &statusRepository{
		db: dbConn,
	}
}

func (r *statusRepository) GetStatusList(ctx context.Context) ([]model.Status, error) {
	list := make([]model.Status, 0)

	//Execute SQL Query
	query := `
		SELECT 
			s.status_id,
			s.name,
			s.created_at, 
			s.updated_at, 
			s.is_delete 
		FROM 
			status s 
		ORDER BY name ASC;
		`
	rows, err := r.db.QueryxContext(ctx, query)

	//Error Handling
	if err != nil {
		utils.LogError("Repo", "func GetStatusList", err)
		return list, err
	}
	//Close the Execution of SQL Query
	defer rows.Close()

	//Iterate over all available rows and strore the data
	for rows.Next() {
		var status model.Status
		// scan and assign into destination variable
		err = rows.Scan(
			&status.Status_id,
			&status.Name,
			&status.CreatedAt,
			&status.UpdatedAt,
			&status.Is_delete,
		)

		if err != nil {
			utils.LogError("Repo", "GetStatusList scan data", err)
			return list, err
		}
		// append to list slice
		list = append(list, status)
	}

	utils.CloseDB(rows)
	// returns populated data
	return list, err
}

func (r *statusRepository) GetStatusDetail(ctx context.Context, id uuid.UUID) (model.Status, error) {

	var (
		status model.Status
	)

	//SQL Query
	query := `
		SELECT 
			s.status_id,
			s.name,
			s.created_at, 
			s.updated_at, 
			s.is_delete 
		FROM 
			status s
		WHERE s.status_id,=$1;
		`

	//Execute SQL Query
	err := r.db.QueryRowxContext(
		ctx,
		query,
		id,
	).Scan(
		&status.Status_id,
		&status.Name,
		&status.CreatedAt,
		&status.UpdatedAt,
		&status.Is_delete,
	)

	//Err Handling
	if err != nil {
		utils.LogError("Repo", "func GetStatusDetail", err)
		return status, err
	}

	// returns login data
	return status, err
}

func (r *statusRepository) CreateStatus(ctx context.Context, tx *sqlx.Tx, u model.Status) (uuid.UUID, error) {
	//Variable that holds registered user email
	var (
		status_id uuid.UUID
	)

	//Query
	query := `
		INSERT INTO 
			status (name) 
		VALUES
			($1)
		RETURNING status_id
			;
	`
	//Execute query and Populating createdUser variable
	err := tx.QueryRowxContext(
		ctx,
		query,
		u.Name,
	).Scan(
		&status_id,
	)

	//Err Handling
	if err != nil {
		utils.LogError("Repo", "func CreateStatus", err)
		return status_id, err
	}

	return status_id, err
}

func (r *statusRepository) UpdateStatus(ctx context.Context, tx *sqlx.Tx, u model.Status) (uuid.UUID, error) {
	var (
		status_id uuid.UUID
	)

	//Query
	query := `
		UPDATE 
			status
		SET
			name = $1,
			updated_at = now()
		VALUES
			($2)
		WHERE
			status_id = $1
		RETURNING status_id
		;
	`

	err := tx.QueryRowxContext(
		ctx,
		query,
		u.Status_id,
		u.Name,
	).Scan(
		&status_id,
	)

	//Err Handling
	if err != nil {
		utils.LogError("Repo", "func UpdateStatus", err)
		return status_id, err
	}

	return status_id, err
}

func (r *statusRepository) DeleteStatus(ctx context.Context, tx *sqlx.Tx, id uuid.UUID) (uuid.UUID, error) {
	var (
		status_id uuid.UUID
	)

	//Query
	query := `
		UPDATE 
			status
		SET
			updated_at = now(),
			is_delete = true
		WHERE 
			status_id = $1
		RETURNING status_id
		;
	`

	err := tx.QueryRowxContext(
		ctx,
		query,
		id,
	).Scan(
		&status_id,
	)

	//Err Handling
	if err != nil {
		utils.LogError("Repo", "func DeleteStatus", err)
		return status_id, err
	}

	return status_id, err
}
