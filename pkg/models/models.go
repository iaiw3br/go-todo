package models

import "time"

type TodoList struct {
	Id          int
	Title       string
	IsCompleted bool
	Created     time.Time
}
