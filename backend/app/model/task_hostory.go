package model

import "time"

type TaskHistory struct {
	ID        uint      `json:"id"`
	TaskID    string    `json:"task_id"`
	ActionBy  string    `json:"action_by"`
	Action    string    `json:"action"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
}
