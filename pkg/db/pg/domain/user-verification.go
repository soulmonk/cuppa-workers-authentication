package domain

import (
	"time"
)

type UserVerification struct {
	ID     uint64 `bson:"id" json:"id"`
	UserID uint64 `bson:"user_id" json:"user_id" db:"user_id"`
	Code   string `bson:"code" json:"code"`

	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty" db:"updated_at"`
}
