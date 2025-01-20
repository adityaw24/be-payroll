package repository

import (
	"context"

	"github.com/dafiqarba/be-payroll/model"
	"github.com/dafiqarba/be-payroll/utils"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PositionRepo interface {
	//Create
	CreatePosition(ctx context.Context, tx *sqlx.Tx, u model.Position) (uuid.UUID, error)
	//Read
	GetPositionList(ctx context.Context) ([]model.Position, error)
	GetPositionDetail(ctx context.Context, id uuid.UUID) (model.Position, error)
	//Update
	UpdatePosition(ctx context.Context, tx *sqlx.Tx, u model.Position) (uuid.UUID, error)
	//Delete
	DeletePosition(ctx context.Context, tx *sqlx.Tx, id uuid.UUID) (uuid.UUID, error)
}

type positionRepository struct {
	db *sqlx.DB
}

func NewPositionRepo(dbConn *sqlx.DB) PositionRepo {
	return &positionRepository{
		db: dbConn,
	}
}

func (r *positionRepository) GetPositionList(ctx context.Context) ([]model.Position, error) {
	//Variable to store collection of roles
	positions := make([]model.Position, 0)

	//Execute SQL Query
	query := `
		SELECT 
			p.position_id,
			p.name,
			p.created_at, 
			p.updated_at, 
			p.is_delete 
		FROM 
			positions p 
		ORDER BY name ASC;
		`
	rows, err := r.db.QueryxContext(ctx, query)

	//Error Handling
	if err != nil {
		utils.LogError("Repo", "func GetPositionList", err)
		return positions, err
	}

	//Close the Execution of SQL Query
	defer rows.Close()

	//Iterate over all available rows and strore the data
	for rows.Next() {
		var position model.Position
		// scan and assign into destination variable
		err = rows.Scan(
			&position.Position_id,
			&position.Name,
			&position.CreatedAt,
			&position.UpdatedAt,
			&position.Is_delete,
		)

		if err != nil {
			utils.LogError("Repo", "GetPositionList scan data", err)
			return positions, err
		}
		// append to positions slice
		positions = append(positions, position)
	}

	utils.CloseDB(rows)
	// returns populated data
	return positions, err
}

func (r *positionRepository) GetPositionDetail(ctx context.Context, id uuid.UUID) (model.Position, error) {

	var (
		position model.Position
	)

	//SQL Query
	query := `
		SELECT 
			p.position_id,
			p.name,
			p.created_at, 
			p.updated_at, 
			p.is_delete 
		FROM 
			positions p
		WHERE r.position_id=$1;
		`

	//Execute SQL Query
	err := r.db.QueryRowxContext(
		ctx,
		query,
		id,
	).Scan(
		&position.Position_id,
		&position.Name,
		&position.CreatedAt,
		&position.UpdatedAt,
		&position.Is_delete,
	)

	//Err Handling
	if err != nil {
		utils.LogError("Repo", "func GetPositionDetail", err)
		return position, err
	}

	// returns login data
	return position, err
}

func (r *positionRepository) CreatePosition(ctx context.Context, tx *sqlx.Tx, u model.Position) (uuid.UUID, error) {
	//Variable that holds registered user email
	var (
		position_id uuid.UUID
	)

	//Query
	query := `
		INSERT INTO 
			positions (name) 
		VALUES
			($1)
		RETURNING position_id
			;
	`
	//Execute query and Populating createdUser variable
	err := tx.QueryRowxContext(
		ctx,
		query,
		u.Name,
	).Scan(
		&position_id,
	)

	//Err Handling
	if err != nil {
		utils.LogError("Repo", "func CreatePosition", err)
		return position_id, err
	}

	return position_id, err
}

func (r *positionRepository) UpdatePosition(ctx context.Context, tx *sqlx.Tx, u model.Position) (uuid.UUID, error) {
	var (
		position_id uuid.UUID
	)

	//Query
	query := `
		UPDATE 
			positions
		SET
			name = $1,
			updated_at = now()
		VALUES
			($2)
		WHERE
			position_id = $1
		RETURNING position_id
			;
	`

	err := tx.QueryRowxContext(
		ctx,
		query,
		u.Position_id,
		u.Name,
	).Scan(
		&position_id,
	)

	//Err Handling
	if err != nil {
		utils.LogError("Repo", "func UpdatePosition", err)
		return position_id, err
	}

	return position_id, err
}

func (r *positionRepository) DeletePosition(ctx context.Context, tx *sqlx.Tx, id uuid.UUID) (uuid.UUID, error) {
	var (
		position_id uuid.UUID
	)

	//Query
	query := `
		UPDATE 
			positions
		SET
			updated_at = now(),
			is_delete = true
		WHERE 
			position_id = $1
		RETURNING position_id
		;
	`

	err := tx.QueryRowxContext(
		ctx,
		query,
		id,
	).Scan(
		&position_id,
	)

	//Err Handling
	if err != nil {
		utils.LogError("Repo", "func DeletePosition", err)
		return position_id, err
	}

	return position_id, err
}
