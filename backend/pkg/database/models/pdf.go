package models

import (
	"time"
)

type PDF struct {
	Id          string    `json:"id" gorm:"column:id;type:char(36);primaryKey"`
	Uploader    string    `json:"uploader" gorm:"column:uploader;type:char(36);not null"`
	Author      string    `json:"author" gorm:"column:author;type:varchar(32);not null"`
	Title       string    `json:"title" gorm:"column:title;type:varchar(32);not null"`
	Description string    `json:"description" gorm:"column:description;type:longtext;not null"`
	Url         string    `json:"url" gorm:"column:url;type:varchar(150);not null"`
	CoverUrl    string    `json:"cover_url" gorm:"column:cover_url;type:varchar(150);not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
}

func (PDF) TableName() string {
	return "pdfs"
}
