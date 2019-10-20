package domain

import (
	"time"
)

type User struct {
	ID        uint64    `bson:"id" json:"id"`
	Name      string    `bson:"name" json:"name"`
	Email     string    `bson:"email" json:"email"`
	Password  string    `bson:"password,skip" json:"password,skip"`
	Salt      string    `bson:"salt,delete" json:"salt,delete"` // todo remove field
	Enabled   bool      `bson:"enable" json:"enable"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty" db:"updated_at"`
}

type Users struct {
	List []User `bson:"list" json:"list"`
}
