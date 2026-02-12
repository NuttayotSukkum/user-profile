package domain

import (
	"time"
)

const (
	StatusActive   = "ACTIVE"
	StatusInactive = "INACTIVE"
)

type UserProfile struct {
	ID        string `gorm:"type:uuid;primaryKey"`
	FirstName string `gorm:"size:100;not null"`
	LastName  string `gorm:"size:100;not null"`
	Email     string `gorm:"size:150;uniqueIndex;not null"`
	Password  string `gorm:"size:150;not null"`
	Status    string `gorm:"size:50;not null"`
	Role      string `gorm:"size:50;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *UserProfile) TableName() string {
	return "user_profiles"
}
