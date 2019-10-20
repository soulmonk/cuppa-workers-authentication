package domain

import (
	"time"
)

type User struct {
	ID        uint64
	Name      string
	Email     string
	Password  string
	Salt      string
	Enabled   string
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty" db:"updated_at"`
}

type Users struct {
	List []User
}
