package dto

import "time"

type Todo struct {
	Name  string `json:"name"`
	Check int    `json:"check"`
}

type TodoResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Check     int       `json:"check"`
	CreatedAt time.Time `json:"created_at"`
}
