package model

import (
	"time"

	"gorm.io/gorm"
)

type URL struct {
	gorm.Model
	// ShortUrl is the shortened version of the original URL
	ShortUrl string `gorm:"unique;type:varchar(32);check:original_url <> ''" json:"short_url" validate:"required,max=32"`
	// OriginalUrl is the full, original URL that was shortened
	OriginalUrl string `gorm:"not null;type:varchar(255);check:original_url <> '';index" json:"original_url" validate:"required,url,max=255"`
	// ExpirationDate is when the shortened URL will expire
	ExpirationDate time.Time `json:"expiration_date" validate:"required,gt=now"`
	// ClickCount tracks how many times the shortened URL has been accessed
	ClickCount int `json:"click_count" validate:"gte=0"`
}
