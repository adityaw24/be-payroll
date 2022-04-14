package entity

import "time"

//Represents leave_records table on the database
type LeaveRecord struct {
	Request_id  int       `json:"request_id"`
	Request_on  time.Time `json:"request_on"`
	From_date   time.Time `json:"from_date"`
	To_date     time.Time `json:"to_date"`
	Return_date time.Time `json:"return_date"`
	Amount      int       `json:"amount"`
	Reason      string    `json:"reason"`
	Mobile      string    `json:"mobile"`
	Address     string    `json:"address"`
	Status_id   int       `json:"status_id"`
	Leave_id    int       `json:"leave_id"`
	User_id     int       `json:"user_id"`
}

//View model for leave_records and status table
type LeaveRecordListModel struct {
	Request_id  int    `json:"request_id"`
	Request_on  string `json:"request_on"`
	Leave_name  string `json:"leave_type"`
	Reason      string `json:"reason"`
	Status_name string `json:"status"`
	User_id     int    `json:"user_id"`
}
