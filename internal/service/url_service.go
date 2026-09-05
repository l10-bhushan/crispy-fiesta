// Service layer for URL api
package service

import (
	"context"

	"github.com/l10-bhushan/crispy-fiesta/internal/apperrors"
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

// Create Service
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

// Redirect service
func (s *URLService) Redirect(ctx context.Context, shortCode string) (*models.CreateShortURLResponse, error) {
	if shortCode == "" {
		return nil, apperrors.ErrorInvalidInput
	}

	data, err := s.repo.FindByShortCode(ctx, shortCode)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Fetch all Service
func (s *URLService) FetchAllData(ctx context.Context) ([]models.CreateShortURLResponse, error) {
	data, err := s.repo.FetchAllData(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Fetch data by Id service
func (s *URLService) FetchByID(ctx context.Context, id string) (*models.CreateShortURLResponse, error) {
	if id == "" {
		return nil, apperrors.ErrorInvalidInput
	}
	data, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Delete data by Id service
func (s *URLService) DeleteById(ctx context.Context, id string) error {

	if id == "" {
		return apperrors.ErrorInvalidInput
	}

	err := s.repo.DeleteById(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
