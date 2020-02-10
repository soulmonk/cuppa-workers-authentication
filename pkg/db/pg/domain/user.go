package domain

import (
	"time"
)

type User struct {
	ID        uint64    `bson:"id" json:"id"`
	Name      string    `bson:"name" json:"name"`
	Email     string    `bson:"email" json:"email"`
	Password  string    `bson:"password,skip" json:"password,skip"`
	Enabled   bool      `bson:"enable" json:"enable"`
	CreatedAt time.Time `bson:"createdAt,omitempty" json:"createdAt,omitempty" db:"created_at"`
	UpdatedAt time.Time `bson:"updatedAt,omitempty" json:"updatedAt,omitempty" db:"updated_at"`
}

type Users struct {
	List []User `bson:"list" json:"list"`
}
