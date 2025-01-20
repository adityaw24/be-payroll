package model

import (
	"time"

	"github.com/google/uuid"
)

type Status struct {
	Status_id uuid.UUID `json:"status_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Is_delete bool      `json:"is_delete"`
}
