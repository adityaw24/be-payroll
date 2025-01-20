package model

import (
	"time"

	"github.com/google/uuid"
)

type Position struct {
	Position_id uuid.UUID `json:"position_id"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Is_delete   bool      `json:"is_delete"`
}
