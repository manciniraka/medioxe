package entity

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	FullName  string    `json:"full_name" gorm:"column:full_name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
