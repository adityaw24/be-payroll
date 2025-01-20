package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/dafiqarba/be-payroll/model"
	"github.com/dafiqarba/be-payroll/utils"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type LeaveRecordRepo interface {
	//Read
	GetLeaveRecordDetail(ctx context.Context, req_id uuid.UUID, id uuid.UUID) (model.LeaveRecord, error)
	GetLeaveRecordList(ctx context.Context, id uuid.UUID, year string) ([]model.LeaveRecordListModel, error)
	//Create
	CreateLeaveRecord(ctx context.Context, tx *sqlx.Tx, d model.LeaveRecord) (uuid.UUID, error)
	//Update
	//Delete
}

type leaveRecordConnection struct {
	connection *sqlx.DB
}

func NewLeaveRecordRepo(dbConn *sqlx.DB) LeaveRecordRepo {
	return &leaveRecordConnection{
		connection: dbConn,
	}
}

func (db *leaveRecordConnection) GetLeaveRecordDetail(ctx context.Context, req_id uuid.UUID, id uuid.UUID) (model.LeaveRecord, error) {
	//Variable to store leave record detail
	var (
		leaveRecordDetail model.LeaveRecord
	)

	//Query
	query := `
		SELECT 
			* 
		FROM 
			leave_records 
		WHERE 
			request_id=$1 AND user_id=$2;
	`
	//Execute SQL Query
	err := db.connection.QueryRowxContext(
		ctx,
		query,
		req_id,
		id,
	).Scan(
		&leaveRecordDetail.Request_id,
		&leaveRecordDetail.Request_on,
		&leaveRecordDetail.From_date,
		&leaveRecordDetail.To_date,
		&leaveRecordDetail.Return_date,
		&leaveRecordDetail.Amount,
		&leaveRecordDetail.Reason,
		&leaveRecordDetail.Mobile,
		&leaveRecordDetail.Address,
		&leaveRecordDetail.Status_id,
		&leaveRecordDetail.Leave_id,
		&leaveRecordDetail.User_id,
	)

	//Err Handling
	if err != nil {
		utils.LogError("Repo", "func GetLeaveRecordDetail", err)
		return leaveRecordDetail, err
	}

	// returns populated data
	return leaveRecordDetail, err
}

func (db *leaveRecordConnection) GetLeaveRecordList(ctx context.Context, id uuid.UUID, year string) ([]model.LeaveRecordListModel, error) {
	// ORDER BY request_on DESC
	leaveRecordList := make([]model.LeaveRecordListModel, 0)

	//Execute SQL Query
	query := fmt.Sprintf(`
		SELECT 
			l.request_id, l.request_on, t.leave_name, l.reason, s.status_name, l.user_id 
		FROM 
			leave_records as l 
				INNER JOIN status as s 
					ON s.status_id = l.status_id 
				INNER JOIN leave_types as t 
					ON t.leave_id = l.leave_id 
		WHERE 
			l.user_id = %v ORDER BY l.request_on %v;`, id, year)

	rows, err := db.connection.QueryxContext(ctx, query)
	if err != nil {
		utils.LogError("Repo", "func GetLeaveRecordList", err)
		return leaveRecordList, err
	}

	defer rows.Close()

	for rows.Next() {
		var leaveRecord model.LeaveRecordListModel
		// scan and assign into leaveRecord
		err = rows.Scan(
			&leaveRecord.Request_id,
			&leaveRecord.Request_on,
			&leaveRecord.Leave_name,
			&leaveRecord.Reason,
			&leaveRecord.Status_name,
			&leaveRecord.User_id,
		)

		if err != nil {
			utils.LogError("Repo", "GetLeaveRecordList scan data", err)
			return leaveRecordList, err
		}
		// append to leaveRecordList slices
		leaveRecordList = append(leaveRecordList, leaveRecord)
	}
	// Check for empty result
	if len(leaveRecordList) == 0 {
		err := errors.New("sql: no results")
		utils.LogError("Repo", "GetLeaveRecordList", err)
		return leaveRecordList, err
	}

	utils.CloseDB(rows)
	// return leaveRecordlist populated with results
	return leaveRecordList, err
}

func (db *leaveRecordConnection) CreateLeaveRecord(ctx context.Context, tx *sqlx.Tx, d model.LeaveRecord) (uuid.UUID, error) {
	var (
		req_id uuid.UUID
	)

	// Insert SQL Query
	query := `
		INSERT INTO
			leave_records 
				(request_on, from_date, to_date, return_date, amount, reason, mobile, address, status_id, leave_id, user_id)
		VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING request_id	
			;
		`
	err := tx.QueryRowxContext(
		ctx,
		query,
		d.Request_on,
		d.From_date,
		d.To_date,
		d.Return_date,
		d.Amount,
		d.Reason,
		d.Mobile,
		d.Address,
		d.Status_id,
		d.Leave_id,
		d.User_id,
	).Scan(
		&req_id,
	)

	if err != nil {
		utils.LogError("Repo", "func CreateLeaveRecord", err)
		return req_id, err
	}

	return req_id, err
}
