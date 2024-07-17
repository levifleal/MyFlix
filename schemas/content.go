package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Content struct {
	gorm.Model
	Title        string
	Desc         string
	ContentType  string
	Genre        string
	ReleaseDate time.Time
}

type ContentResponse struct {
	Id           uint      `json:"id"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	DeletedAt    time.Time `json:"deletedAt,omitempty"`
	Title        string    `json:"title"`
	Desc         string    `json:"desc"`
	ContentType  string    `json:"contentType"`
	Genre        string    `json:"genre"`
	RealeaseDate time.Time `json:"realeaseDate"`
}
