package models

import (
	"time"
)

type Referred struct {
	ID               int64
	AgentId          string
	JobUrl           string
	RequestReferer   string
	RequestIp        string
	RequestUserAgent string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        time.Time
}
