package models

import (
	"time"

	"gorm.io/gorm"
)

type URL struct {
	gorm.Model
	// ShortCode is the code used to access the original URL
	ShortCode string `gorm:"unique;index;type:varchar(32);check:short_code <> ''" json:"short_code" validate:"required,max=32"`
	// OriginalUrl is the full, original URL that was shortened
	OriginalUrl string `gorm:"not null;type:varchar(255);check:original_url <> '';index" json:"original_url" validate:"required,url,max=255"`
	// ExpirationDate is when the shortened URL will expire
	ExpirationDate time.Time `json:"expiration_date" validate:"required,gt=now"`
	// ClickCount tracks how many times the shortened URL has been accessed
	ClickCount int `gorm:"default:0" json:"click_count" validate:"gte=0"`
}
