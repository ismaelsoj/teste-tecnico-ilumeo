package models

import "time"

type Historico struct {
	ID               int64     `json:"id"`
	Origin           string    `json:"origin"`
	ResponseStatusID int       `json:"response_status_id"`
	CreatedAt        time.Time `json:"created_at"`
}
