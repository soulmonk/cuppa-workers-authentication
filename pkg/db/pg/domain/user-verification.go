package domain

import (
	"time"
)

type UserVerification struct {
	ID     uint64 `bson:"id" json:"id"`
	UserID uint64 `bson:"userId" json:"userId" db:"user_id"`
	Code   string `bson:"code" json:"code"`

	CreatedAt time.Time `bson:"createdAt,omitempty" json:"createdAt,omitempty" db:"created_at"`
	UpdatedAt time.Time `bson:"updatedAt,omitempty" json:"updatedAt,omitempty" db:"updated_at"`
}
