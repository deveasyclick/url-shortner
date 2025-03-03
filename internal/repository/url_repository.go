package repository

import (
	"url-shortner/internal/models"

	"gorm.io/gorm"
)

type URLRepository interface {
	Create(url *models.URL) error
	FindByShortURL(shortURL string) (*models.URL, error)
	FindByOriginalURL(longURL string) (*models.URL, error)
	IncrementClicks(shortCode string) error
}

type urlRepository struct {
	db *gorm.DB
}

func (r *urlRepository) Create(url *models.URL) error {
	result := r.db.Create(url)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *urlRepository) FindByShortURL(shortCode string) (*models.URL, error) {
	var url models.URL
	result := r.db.Where("short_code = ?", shortCode).First(&url)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // Return nil, nil if no record is found
		}
		return nil, result.Error // Return the error if it's not a "not found" error
	}
	return &url, nil
}

func (r *urlRepository) FindByOriginalURL(longURL string) (*models.URL, error) {
	var url models.URL
	result := r.db.Where("original_url = ?", longURL).First(&url)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // Return nil, nil if no record is found
		}
		return nil, result.Error // Return the error if it's not a "not found" error
	}
	return &url, nil
}

func (r *urlRepository) IncrementClicks(shortCode string) error {
	result := r.db.Model(&models.URL{}).Where("short_code = ?", shortCode).UpdateColumn("click_count", gorm.Expr("click_count + ?", 1))
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewURLRepository(db *gorm.DB) URLRepository {
	return &urlRepository{db: db}
}
