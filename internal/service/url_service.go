package service

import (
	"fmt"
	"url-shortner/internal/config"
	"url-shortner/internal/models"
	"url-shortner/internal/repository"
)

type URLService interface {
	CreateShortURL(originalUrl string) (string, error)
	GetOriginalURL(shortURL string) (string, error)
}

type urlService struct {
	repo repository.URLRepository
}

func (s *urlService) CreateShortURL(longURL string) (string, error) {
	// If the long URL does not exist, generate a new short code
	shortCode := generateShortCode()

	// Create a new URL model with the long URL and short code
	// TODO: Set expiration time and remove expired urls from the database periodically
	newURL := &models.URL{
		OriginalUrl: longURL,
		ShortCode:   shortCode,
	}

	// Save the new URL to the repository
	err := s.repo.Create(newURL)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", config.APP_URL, shortCode), nil
}

func (s *urlService) GetOriginalURL(shortCode string) (string, error) {
	// Find the URL with the given short code in the repository
	url, err := s.repo.FindByShortURL(shortCode)
	if err != nil {
		return "", err
	}

	if url == nil {
		// If the short code does not exist, return an error
		return "", fmt.Errorf("Short code not found")
	}

	// Return the long URL associated with the short code
	return url.OriginalUrl, nil
}

func NewURLService(repo repository.URLRepository) URLService {
	return &urlService{repo: repo}
}
