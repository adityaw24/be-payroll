package model

import (
	"time"

	"github.com/google/uuid"
)

type PayrollRecord struct {
	Payroll_id     uuid.UUID `json:"payroll_id"`
	Payment_period string    `json:"payment_period"`
	Payment_date   time.Time `json:"payment_date"`
	Basic_salary   int       `json:"basic_salary"`
	Bpjs           int       `json:"bpjs"`
	Tax            int       `json:"tax"`
	Total_salary   int       `json:"total_salary"`
	Status_id      uuid.UUID `json:"status_id"`
	User_id        uuid.UUID `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Is_delete      bool      `json:"is_delete"`
}

type PayrollRecordDetailModel struct {
	Payroll_id     uuid.UUID `json:"payroll_id"`
	Name           string    `json:"name"`
	Payment_period string    `json:"payment_period"`
	Payment_date   time.Time `json:"payment_date"`
	Basic_salary   int       `json:"basic_salary"`
	Bpjs           int       `json:"bpjs"`
	Tax            int       `json:"tax"`
	Total_salary   int       `json:"total_salary"`
	Status_name    string    `json:"status_name"`
}

type PayrollRecordListModel struct {
	Payroll_id     uuid.UUID `json:"payroll_id"`
	Name           string    `json:"name"`
	Payment_period string    `json:"payment_period"`
	Payment_date   string    `json:"payment_date"`
	Status_name    string    `json:"status"`
}

type CreatePayrollRecordModel struct {
	Payment_period string    `json:"payment_period"`
	Payment_date   string    `json:"payment_date"`
	Basic_salary   int       `json:"basic_salary"`
	Bpjs           int       `json:"bpjs"`
	Tax            int       `json:"tax"`
	Total_salary   int       `json:"total_salary"`
	Status_id      uuid.UUID `json:"status_id"`
	User_id        uuid.UUID `json:"user_id"`
}

type UpdatePayrollRecordModel struct {
	Payment_period string    `json:"payment_period"`
	Payment_date   string    `json:"payment_date"`
	Basic_salary   int       `json:"basic_salary"`
	Bpjs           int       `json:"bpjs"`
	Tax            int       `json:"tax"`
	Total_salary   int       `json:"total_salary"`
	Status_id      uuid.UUID `json:"status_id"`
	User_id        uuid.UUID `json:"user_id"`
}
