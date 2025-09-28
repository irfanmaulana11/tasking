package dto

import "time"

type Task struct {
	ID             string        `json:"id"`
	Title          string        `json:"title"`
	Description    string        `json:"description"`
	Assignee       string        `json:"Assignee"`
	AssignedLeader string        `json:"assigned_leader"`
	Status         string        `json:"status"`
	Progress       int           `json:"progress"`
	ProgressBy     *string       `json:"progress_by,omitempty"`
	Deadline       *time.Time    `json:"deadline,omitempty"`
	CreatedBy      string        `json:"created_by"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
	TaskHistory    []TaskHistory `json:"task_history"`
}

type TaskHistory struct {
	ActionBy  string    `json:"action_by"`
	Action    string    `json:"action"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
}

type TaskProgress struct {
	ID       string
	Status   string
	Note     string `json:"note"`
	Progress int    `json:"progress"`
}
