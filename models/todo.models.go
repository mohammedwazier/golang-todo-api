package models

import "time"

type Todo struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Check     int       `json:"check" gorm:"default:0"`
	CreatedAt time.Time `json:"created_at"`
}

func (Todo) TableName() string {
	return "todo"
}
