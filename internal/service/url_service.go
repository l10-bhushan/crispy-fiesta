// Service layer for URL api
package service

import (
	"context"
	"errors"

	"github.com/l10-bhushan/crispy-fiesta/internal/models"
	"github.com/l10-bhushan/crispy-fiesta/internal/repository"
	"github.com/l10-bhushan/crispy-fiesta/internal/utils"
)

// Creating the URLService to hold out repository structure that talks to the database
type URLService struct {
	repo *repository.URLRepository
}

// Constructor function to create instance of URLService
func NewURlService(repo *repository.URLRepository) *URLService {
	return &URLService{
		repo: repo,
	}
}

// Create function
func (s *URLService) Create(ctx context.Context, originalURL string) (*models.CreateShortURLResponse, error) {
	// Validating the URL using our custom validator
	err := utils.ValidateURL(originalURL)
	if err != nil {
		return nil, err
	}

	// Generating short_code
	shortCode, err := utils.GenerateShortCode()
	if err != nil {
		return nil, err
	}
	data, err := s.repo.Create(ctx, shortCode, originalURL)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Fetch function
func (s *URLService) GetOriginalURL(ctx context.Context, shortCode string) (*models.CreateShortURLResponse, error) {
	if shortCode == "" {
		return nil, errors.New("short code cannot be empty")
	}

	data, err := s.repo.FindByShortCode(ctx, shortCode)
	if err != nil {
		return nil, err
	}

	return data, nil
}
