package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	Password  string    `json:"password"`
}
