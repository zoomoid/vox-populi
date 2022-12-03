package biz

import "time"

type SoftDeleteMixin struct {
	IsDeleted bool       `json:"is_deleted,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
