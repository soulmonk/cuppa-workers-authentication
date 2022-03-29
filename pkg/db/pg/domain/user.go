package domain

import (
	"database/sql/driver"
	"github.com/pborman/uuid"
	"time"
)

// NullUUID can be used with the standard sql package to represent a
// UUID value that can be NULL in the database
type NullUUID struct {
	UUID  uuid.UUID
	Valid bool
}

var Nil = uuid.UUID{}

// Value implements the driver.Valuer interface.
func (u NullUUID) Value() (driver.Value, error) {
	if !u.Valid {
		return nil, nil
	}
	// Delegate to UUID Value function
	return u.UUID.Value()
}

// Scan implements the sql.Scanner interface.
func (u *NullUUID) Scan(src any) error {
	if src == nil {
		u.UUID, u.Valid = Nil, false
		return nil
	}

	// Delegate to UUID Scan function
	u.Valid = true
	return u.UUID.Scan(src)
}

type User struct {
	ID           uint64    `bson:"id" json:"id"`
	Name         string    `bson:"name" json:"name"`
	Email        string    `bson:"email" json:"email"`
	Password     string    `bson:"password,skip" json:"password,skip"`
	Enabled      bool      `bson:"enable" json:"enable"`
	CreatedAt    time.Time `bson:"createdAt,omitempty" json:"createdAt,omitempty" db:"created_at"`
	UpdatedAt    time.Time `bson:"updatedAt,omitempty" json:"updatedAt,omitempty" db:"updated_at"`
	RefreshToken NullUUID  `bson:"refresh_token" json:"refresh_token" db:"refresh_token"`
}

type Users struct {
	List []User `bson:"list" json:"list"`
}
