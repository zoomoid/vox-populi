package biz

import "time"

type MetadataMixin struct {
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Revision  int       `json:"revision,omitempty"`
}
