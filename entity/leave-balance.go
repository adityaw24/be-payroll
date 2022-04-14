package entity

type LeaveBalance struct {
	Balance_id   string `json:"balance_id"`
	Balance_year string `json:"balance_year"`
	Cuti_tahunan int    `json:"cuti_tahunan"`
	Cuti_diambil int    `json:"cuti_diambil"`
	Cuti_balance int    `json:"cuti_balance"`
	Cuti_izin    int    `json:"cuti_izin"`
	Cuti_sakit   int    `json:"cuti_sakit"`
	User_id      string `json:"user_id"`
}

type UpdateLeaveBalanceModel struct {
	User_id      int    `json:"-"`
	Year         string `json:"year"`
	Amounts      int    `json:"amounts"`
	Leave_id     int    `json:"leave_id"`
	Cuti_diambil int    `json:"-"`
}
