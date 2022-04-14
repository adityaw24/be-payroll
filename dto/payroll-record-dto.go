package dto

type CreatePayrollRecordModel struct {
	Payment_period string `json:"payment_period"`
	Payment_date   string `json:"payment_date"`
	Basic_salary   int    `json:"basic_salary"`
	Bpjs           int    `json:"bpjs"`
	Tax            int    `json:"tax"`
	Total_salary   int    `json:"total_salary"`
	Status_id      int    `json:"status_id"`
	User_id        int    `json:"user_id"`
}

type UpdatePayrollRecordModel struct {
	Payment_period string `json:"payment_period"`
	Payment_date   string `json:"payment_date"`
	Basic_salary   int    `json:"basic_salary"`
	Bpjs           int    `json:"bpjs"`
	Tax            int    `json:"tax"`
	Total_salary   int    `json:"total_salary"`
	Status_id      int    `json:"status_id"`
	User_id        int    `json:"user_id"`
}
