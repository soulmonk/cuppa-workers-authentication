package domain

import (
	"time"
)

type UserVerification struct {
	ID     uint64
	UserID uint64
	Code   string

	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty" db:"updated_at"`
}
