package model

import (
	"time"

	"github.com/google/uuid"
)

type LeaveBalance struct {
	Leave_id     int       `json:"leave_id"`
	Leave_year   string    `json:"leave_year"`
	Cuti_tahunan int       `json:"cuti_tahunan"`
	Cuti_diambil int       `json:"cuti_diambil"`
	Cuti_balance int       `json:"cuti_balance"`
	Cuti_izin    int       `json:"cuti_izin"`
	Cuti_sakit   int       `json:"cuti_sakit"`
	User_id      uuid.UUID `json:"user_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Is_delete    bool      `json:"is_delete"`
}

type UpdateLeaveBalanceModel struct {
	User_id      uuid.UUID `json:"-"`
	Year         string    `json:"year"`
	Amounts      int       `json:"amounts"`
	Leave_id     int       `json:"leave_id"`
	Cuti_diambil int       `json:"-"`
}
