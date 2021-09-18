package model

import "time"

type Model struct {
	ID        int64      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_on"`
	UpdatedAt time.Time  `josn:"modified_on"`
	DeletedAt *time.Time `gorm:"column:deleted" json:"deleted_on"`
}
