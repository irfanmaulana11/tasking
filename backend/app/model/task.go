package model

import "time"

type Task struct {
	ID             string     `json:"id"`
	Title          string     `json:"title"`
	Description    string     `json:"description"`
	Assignee       string     `json:"assignee"`
	AssignedLeader string     `json:"assigned_leader,omitempty"`
	Status         string     `json:"status"`
	Progress       int        `json:"progress"`
	ProgressBy     *string    `json:"progress_by,omitempty"`
	Deadline       *time.Time `json:"deadline,omitempty"`
	CreatedBy      string     `json:"created_by"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}
