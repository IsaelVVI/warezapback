package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Opening struct {
	Id        uint `gorm:"primaryKey;autoIncrement"`
	Role      string
	Company   string
	Location  string
	Remote    bool
	Link      string
	Salary    int64
	CreatedAt time.Time `gorm:autoCreateTime`
	UpdatedAt time.Time `gorm:autoUpdateTime`
	DeletedAt gorm.DeletedAt
}

type OpeningResponse struct {
	Id        uint      `json:"id"`
	Role      string    `json:"role"`
	Company   string    `json:"company"`
	Location  string    `json:"location"`
	Remote    bool      `json:"remote"`
	Link      string    `json:"link"`
	Salary    int64     `json:"salary"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}
