package model

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	Role_id   uuid.UUID `json:"role_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Is_delete bool      `json:"is_delete"`
}
