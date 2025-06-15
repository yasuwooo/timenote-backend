package model

import (
	"time"

	"github.com/google/uuid"
)

type ActivityLog struct {
	ID        uuid.UUID  `json:"id"`
	UserID    uuid.UUID  `json:"user_id"`
	Title     string     `json:"title"`
	Memo      string     `json:"memo"`
	StartTime time.Time  `json:"start_time"`
	EndTime   time.Time  `json:"end_time"`
	Notes     string     `json:"notes"`
	ParentID  *uuid.UUID `json:"parent_id"`
	CreatedAt time.Time  `json:"created_at"`
	Path      string     `json:"path"`
	Tags      []Tag      `json:"tags"`
}

type Tag struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at"`
}

type ActivityLogTag struct {
	ID         uuid.UUID `json:"id"`
	ActivityID uuid.UUID `json:"activity_id"`
	TagID      uuid.UUID `json:"tag_id"`
	CreatedAt  time.Time `json:"created_at"`
}
