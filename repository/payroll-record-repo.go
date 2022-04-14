package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/dafiqarba/be-payroll/entity"
)

type PayrollRecordRepo interface {
	//Read
	GetPayrollRecordList() ([]entity.PayrollRecordListModel, error)
	GetPayrollRecordDetail(id int) (entity.PayrollRecordDetailModel, error)
	//Create
	// CreatePayrollRecord(p entity.PayrollRecord) (entity.PayrollRecord, error)
	CreatePayrollRecord(p entity.PayrollRecord) (int, error)
	//Update
	UpdatePayrollRecord(id int, p entity.PayrollRecord) (int, error)
	// UpdatePayrollRecord(id int, p entity.PayrollRecord) (entity.PayrollRecord, error)
}

type payrollRecordRepo struct {
	connection *sql.DB
}

func NewPayrollRecordRepo(db *sql.DB) PayrollRecordRepo {
	return &payrollRecordRepo{
		connection: db,
	}
}

func (db *payrollRecordRepo) GetPayrollRecordList() ([]entity.PayrollRecordListModel, error) {
	var payrollRecordList []entity.PayrollRecordListModel

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

	rows, err := db.connection.Query(query)

	if err != nil {
		log.Printf("Cannot execute query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var payrollRecord entity.PayrollRecordListModel
		err = rows.Scan(&payrollRecord.Payroll_id, &payrollRecord.Name, &payrollRecord.Payment_period, &payrollRecord.Payment_date, &payrollRecord.Status_name)
		if err != nil {
			log.Printf("Cannot retrieve data: %v", err)
		}
		payrollRecordList = append(payrollRecordList, payrollRecord)
	}

	if len(payrollRecordList) == 0 {
		log.Println("| " + errors.New("sql: no data found").Error())
		err := errors.New("sql: no data found")
		return payrollRecordList, err
	}
	return payrollRecordList, err
}

func (db *payrollRecordRepo) GetPayrollRecordDetail(id int) (entity.PayrollRecordDetailModel, error) {
	var payrollRecord entity.PayrollRecordDetailModel

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

	rows := db.connection.QueryRow(query)

	err := rows.Scan(&payrollRecord.Payroll_id, &payrollRecord.Name, &payrollRecord.Payment_period, &payrollRecord.Payment_date, &payrollRecord.Basic_salary, &payrollRecord.Bpjs, &payrollRecord.Tax, &payrollRecord.Total_salary, &payrollRecord.Status_name)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("| " + err.Error())
			return payrollRecord, err
		} else {
			log.Println("| " + err.Error())
			return payrollRecord, err
		}
	}

	return payrollRecord, err
}

// func (db *payrollRecordRepo) CreatePayrollRecord(p entity.PayrollRecord) (entity.PayrollRecord, error) {
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

func (db *payrollRecordRepo) CreatePayrollRecord(p entity.PayrollRecord) (int, error) {
	var user_id int

	query := `
		INSERT INTO payroll_records(
			user_id, payment_period, payment_date, basic_salary, bpjs, tax, total_salary, status_id
		) VALUES(
			$1, $2, $3, $4, $5, $6, $7, $8
		) RETURNING payroll_id;`

	err := db.connection.QueryRow(query, p.User_id, p.Payment_period, p.Payment_date, p.Basic_salary, p.Bpjs, p.Tax, p.Total_salary, p.Status_id).Scan(&user_id)

	if err != nil {
		log.Println("| " + err.Error())
		return 0, err
	}

	return user_id, err
}

// func (db *payrollRecordRepo) UpdatePayrollRecord(id int, p entity.PayrollRecord) (entity.PayrollRecord, error) {
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

func (db *payrollRecordRepo) UpdatePayrollRecord(id int, p entity.PayrollRecord) (int, error) {
	stmt, err := db.connection.Prepare(`
		UPDATE payroll_records SET
			user_id = $1, payment_period = $2, payment_date = $3, basic_salary = $4, bpjs = $5, tax = $6, total_salary = $7, status_id = $8
		WHERE
			payroll_id = $9;`)

	stmt.Exec(p.User_id, p.Payment_period, p.Payment_date, p.Basic_salary, p.Bpjs, p.Tax, p.Total_salary, p.Status_id, p.Payroll_id)

	if err != nil {
		log.Println("| " + err.Error())
		return 0, err
	}

	return p.Payroll_id, err
}
