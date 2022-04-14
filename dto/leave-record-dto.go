package dto

// Model response when user is creating a leave request
type CreateLeaveRecordModel struct {
	Request_on  string `json:"request_on"`
	From_date   string `json:"from_date"`
	To_date     string `json:"to_date"`
	Return_date string `json:"return_date"`
	Amount      string `json:"amount"`
	Reason      string `json:"reason"`
	Mobile      string `json:"mobile"`
	Address     string `json:"address"`
	Status_id   string `json:"status_id"`
	Leave_id    string `json:"leave_id"`
	User_id     int    `json:"user_id"`
}
