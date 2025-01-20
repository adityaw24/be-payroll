package repository

import (
	"context"
	"fmt"

	"github.com/dafiqarba/be-payroll/model"
	"github.com/dafiqarba/be-payroll/utils"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// TODO: search how to return a populated struct from function, if struct is defined

type LeaveBalanceRepo interface {
	//Read
	GetLeaveBalance(ctx context.Context, id uuid.UUID, year string) (model.LeaveBalance, error)
	//Create
	// InsertUser(user model.User) (model.User, error)
	CreateLeaveBalance(ctx context.Context, tx *sqlx.Tx, leaveBalance model.LeaveBalance) (model.LeaveBalance, error)
	//Update
	UpdateLeaveBalance(ctx context.Context, tx *sqlx.Tx, updatedData model.UpdateLeaveBalanceModel, leave_type string) (int, error)
	//Delete
}

type leaveBalanceConnection struct {
	connection *sqlx.DB
}

func NewLeaveBalanceRepo(dbConn *sqlx.DB) LeaveBalanceRepo {
	return &leaveBalanceConnection{
		connection: dbConn,
	}
}

func (db *leaveBalanceConnection) GetLeaveBalance(ctx context.Context, id uuid.UUID, year string) (model.LeaveBalance, error) {
	//Variable to store leave balance detail
	var (
		leaveBalance model.LeaveBalance
	)

	//Execute SQL Query
	query := `SELECT * FROM leave_balance WHERE user_id=$1 AND leave_year=$2`
	err := db.connection.QueryRowxContext(ctx, query, id, year).Scan(
		&leaveBalance.Leave_id,
		&leaveBalance.Leave_year,
		&leaveBalance.Cuti_tahunan,
		&leaveBalance.Cuti_diambil,
		&leaveBalance.Cuti_balance,
		&leaveBalance.Cuti_izin,
		&leaveBalance.Cuti_sakit,
		&leaveBalance.User_id,
	)

	if err != nil {
		utils.LogError("Repo", "func GetLeaveBalance", err)
		return leaveBalance, err
	}

	// returns populated data
	return leaveBalance, err
}

func (db *leaveBalanceConnection) CreateLeaveBalance(ctx context.Context, tx *sqlx.Tx, leaveBalance model.LeaveBalance) (model.LeaveBalance, error) {
	var (
		leave model.LeaveBalance
	)

	query := `
			INSERT INTO
				leave_balance (leave_year, cuti_tahunan, cuti_diambil, cuti_balance, cuti_izin, cuti_sakit, user_id)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
			RETURNING
				leave_id,
				leave_year,
				cuti_tahunan,
				cuti_diambil,
				cuti_balance,
				cuti_izin,
				cuti_sakit,
				user_id;
		`

	err := tx.QueryRowxContext(ctx, query,
		leaveBalance.Leave_id,
		leaveBalance.Leave_year,
		leaveBalance.Cuti_tahunan,
		leaveBalance.Cuti_diambil,
		leaveBalance.Cuti_balance,
		leaveBalance.Cuti_izin,
		leaveBalance.Cuti_sakit,
		leaveBalance.User_id,
	).Scan(
		&leave.Leave_id,
		&leave.Leave_year,
		&leave.Cuti_tahunan,
		&leave.Cuti_diambil,
		&leave.Cuti_balance,
		&leave.Cuti_izin,
		&leave.Cuti_sakit,
		&leave.User_id,
	)

	if err != nil {
		utils.LogError("Repo", "func CreateLeaveBalance", err)
		return leave, err
	}

	return leave, err
}

func (db *leaveBalanceConnection) UpdateLeaveBalance(ctx context.Context, tx *sqlx.Tx, updatedData model.UpdateLeaveBalanceModel, leave_type string) (int, error) {
	var (
		updatedColumn int
	)

	query := fmt.Sprintf(`
		UPDATE 
			leave_balance 
		SET 
			%v=%v, cuti_diambil=%v
		WHERE 
			user_id=%v
		RETURNING leave_id;
		`, leave_type, updatedData.Amounts, updatedData.Cuti_diambil, updatedData.User_id)

	err := tx.QueryRowxContext(ctx, query).Scan(
		&updatedColumn,
	)

	if err != nil {
		utils.LogError("Repo", "func UpdateLeaveBalance", err)
		return updatedColumn, err
	}

	return updatedColumn, err
}
