package model

import (
	"time"
)

type Categories struct {
	CategoryID int    `json:"category_id"`
	Category   string `json:"category"`
}

type Book struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Tags      string    `json:"tags"`
	CreatedAt time.Time `json:"createdAt"`
	Category Categories	`json:"category" gorm:"foreignKey:CategoryID"`
}

type User struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"user_name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
