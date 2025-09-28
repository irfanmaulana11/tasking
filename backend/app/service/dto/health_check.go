package dto

import "time"

type HealthCheck struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}
