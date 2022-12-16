package models

import (
	"time"

	"gorm.io/gorm"
)

type Referred struct {
	gorm.Model
	ID        int64
	Agent     string
	Url       string
	Ip        string
	Origin    string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
