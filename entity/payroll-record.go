package entity

type PayrollRecord struct {
	Payroll_id     int    `json:"payroll_id"`
	Payment_period string `json:"payment_period"`
	Payment_date   string `json:"payment_date"`
	Basic_salary   int    `json:"basic_salary"`
	Bpjs           int    `json:"bpjs"`
	Tax            int    `json:"tax"`
	Total_salary   int    `json:"total_salary"`
	Status_id      int    `json:"status_id"`
	User_id        int    `json:"user_id"`
}

type PayrollRecordDetailModel struct {
	Payroll_id     int    `json:"payroll_id"`
	Name           string `json:"name"`
	Payment_period string `json:"payment_period"`
	Payment_date   string `json:"payment_date"`
	Basic_salary   int    `json:"basic_salary"`
	Bpjs           int    `json:"bpjs"`
	Tax            int    `json:"tax"`
	Total_salary   int    `json:"total_salary"`
	Status_name    string `json:"status_name"`
}

type PayrollRecordListModel struct {
	Payroll_id     int    `json:"payroll_id"`
	Name           string `json:"name"`
	Payment_period string `json:"payment_period"`
	Payment_date   string `json:"payment_date"`
	Status_name    string `json:"status"`
}
