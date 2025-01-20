package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/dafiqarba/be-payroll/model"
	"github.com/dafiqarba/be-payroll/utils"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type PayrollRecordRepo interface {
	//Read
	GetPayrollRecordList(ctx context.Context) ([]model.PayrollRecordListModel, error)
	GetPayrollRecordDetail(ctx context.Context, id uuid.UUID) (model.PayrollRecordDetailModel, error)
	//Create
	// CreatePayrollRecord(ctx context.Context, tx *sqlx.Tx, p model.PayrollRecord) (model.PayrollRecord, error)
	CreatePayrollRecord(ctx context.Context, tx *sqlx.Tx, p model.PayrollRecord) (uuid.UUID, error)
	//Update
	UpdatePayrollRecord(ctx context.Context, tx *sqlx.Tx, id uuid.UUID, p model.PayrollRecord) (uuid.UUID, error)
	// UpdatePayrollRecord(ctx context.Context, id int, p model.PayrollRecord) (model.PayrollRecord, error)
}

type payrollRecordRepo struct {
	connection *sqlx.DB
}

func NewPayrollRecordRepo(db *sqlx.DB) PayrollRecordRepo {
	return &payrollRecordRepo{
		connection: db,
	}
}

func (db *payrollRecordRepo) GetPayrollRecordList(ctx context.Context) ([]model.PayrollRecordListModel, error) {
	payrollRecordList := make([]model.PayrollRecordListModel, 0)

	// err := db.connection.QueryRow("SELECT * FROM payroll_records WHERE employee_id = ? AND year = ?", id, year).Scan(&payrollRecordList)

	// query := fmt.Sprintf(`
	// 	SELECT
	// 		p.payroll_id, u.name, p.payment_period, p.payment_date, s.status_name
	// 	FROM
	// 		payroll_records p
	// 			INNER JOIN status s ON p.status_id = s.status_id
	// 			INNER JOIN users u ON p.user_id = u.user_id
	// 	WHERE
	// 		p.user_id = %v AND p.payment_date = '%v';`, id, year)

	query := `
		SELECT
			p.payroll_id, u.name, p.payment_period, p.payment_date, s.status_name
		FROM
			payroll_records p
				INNER JOIN status s ON p.status_id = s.status_id
				INNER JOIN users u ON p.user_id = u.user_id
		;`

	rows, err := db.connection.QueryxContext(ctx, query)

	if err != nil {
		utils.LogError("Repo", "GetPayrollRecordList", err)
		return payrollRecordList, err
	}

	defer rows.Close()

	for rows.Next() {
		var payrollRecord model.PayrollRecordListModel

		err = rows.Scan(
			&payrollRecord.Payroll_id,
			&payrollRecord.Name,
			&payrollRecord.Payment_period,
			&payrollRecord.Payment_date,
			&payrollRecord.Status_name,
		)

		if err != nil {
			utils.LogError("Repo", "GetPayrollRecordList scan data", err)
			return payrollRecordList, err
		}
		payrollRecordList = append(payrollRecordList, payrollRecord)
	}

	if len(payrollRecordList) == 0 {
		err := errors.New("sql: no data found")
		utils.LogError("Repo", "GetPayrollRecordList", err)
		return payrollRecordList, err
	}

	utils.CloseDB(rows)

	return payrollRecordList, err
}

func (db *payrollRecordRepo) GetPayrollRecordDetail(ctx context.Context, id uuid.UUID) (model.PayrollRecordDetailModel, error) {
	var (
		payrollRecord model.PayrollRecordDetailModel
	)

	// err := db.connection.QueryRow("SELECT * FROM payroll_records WHERE employee_id = ? AND year = ?", id, year).Scan(&payrollRecord)
	query := fmt.Sprintf(`
		SELECT
			p.payroll_id, u.name, p.payment_period, p.payment_date, p.basic_salary, p.bpjs, p.tax, p.total_salary, s.status_name
		FROM
			payroll_records p
				INNER JOIN status s ON p.status_id = s.status_id
				INNER JOIN users u ON p.user_id = u.user_id
		WHERE
			p.payroll_id = %v;`, id)

	err := db.connection.QueryRowxContext(ctx, query).Scan(
		&payrollRecord.Payroll_id,
		&payrollRecord.Name,
		&payrollRecord.Payment_period,
		&payrollRecord.Payment_date,
		&payrollRecord.Basic_salary,
		&payrollRecord.Bpjs,
		&payrollRecord.Tax,
		&payrollRecord.Total_salary,
		&payrollRecord.Status_name,
	)

	if err != nil {
		utils.LogError("Repo", "func GetPayrollRecordDetail", err)
		return payrollRecord, err
	}

	return payrollRecord, err
}

// func (db *payrollRecordRepo) CreatePayrollRecord(ctx context.Context, p model.PayrollRecord, tx *sqlx.Tx, tx *sqlx.Tx) (model.PayrollRecord, error) {
// 	stmt, err := db.connection.Prepare(`
// 		INSERT INTO payroll_records(
// 			user_id, payment_period, payment_date, basic_salary, bpjs, tax, total_salary, status_id
// 		) VALUES(
// 			$1, $2, $3, $4, $5, $6, $7, $8
// 		) RETURNING payroll_id;`)

// 	stmt.Exec(p.User_id, p.Payment_period, p.Payment_date, p.Basic_salary, p.Bpjs, p.Tax, p.Total_salary, p.Status_id)

// 	if err != nil {
// 		log.Println("| " + err.Error())
// 		return p, err
// 	}

// 	return p, err
// }

func (db *payrollRecordRepo) CreatePayrollRecord(ctx context.Context, tx *sqlx.Tx, p model.PayrollRecord) (uuid.UUID, error) {
	var (
		user_id uuid.UUID
	)

	query := `
		INSERT INTO payroll_records(
			user_id, payment_period, payment_date, basic_salary, bpjs, tax, total_salary, status_id
		) VALUES(
			$1, $2, $3, $4, $5, $6, $7, $8
		) RETURNING payroll_id;`

	err := tx.QueryRowxContext(
		ctx,
		query,
		p.User_id,
		p.Payment_period,
		p.Payment_date,
		p.Basic_salary,
		p.Bpjs,
		p.Tax,
		p.Total_salary,
		p.Status_id,
	).Scan(
		&user_id,
	)

	if err != nil {
		utils.LogError("Repo", "func CreatePayrollRecord", err)
		return user_id, err
	}

	return user_id, err
}

// func (db *payrollRecordRepo) UpdatePayrollRecord(ctx context.Context, id int, p model.PayrollRecord) (model.PayrollRecord, error) {
// 	stmt, err := db.connection.Prepare(`
// 		UPDATE payroll_records SET
// 			user_id = $1, payment_period = $2, payment_date = $3, basic_salary = $4, bpjs = $5, tax = $6, total_salary = $7, status_id = $8
// 		WHERE
// 			payroll_id = $9;`)

// 	stmt.Exec(p.User_id, p.Payment_period, p.Payment_date, p.Basic_salary, p.Bpjs, p.Tax, p.Total_salary, p.Status_id, id)

// 	if err != nil {
// 		log.Println("| " + err.Error())
// 		return p, err
// 	}

// 	return p, err
// }

func (db *payrollRecordRepo) UpdatePayrollRecord(ctx context.Context, tx *sqlx.Tx, id uuid.UUID, p model.PayrollRecord) (uuid.UUID, error) {
	query := `
		UPDATE payroll_records SET
			user_id = $1, payment_period = $2, payment_date = $3, basic_salary = $4, bpjs = $5, tax = $6, total_salary = $7, status_id = $8
		WHERE
			payroll_id = $9;`

	err := tx.QueryRowxContext(
		ctx,
		query,
		p.User_id,
		p.Payment_period,
		p.Payment_date,
		p.Basic_salary,
		p.Bpjs,
		p.Tax,
		p.Total_salary,
		p.Status_id,
		p.Payroll_id,
	).Scan(
		&p.Payroll_id,
	)

	if err != nil {
		utils.LogError("Repo", "func UpdateLeaveBalance", err)
		return p.Payroll_id, err
	}

	return p.Payroll_id, err
}
