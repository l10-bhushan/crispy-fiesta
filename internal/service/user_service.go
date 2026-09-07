package service

import (
	"context"

	"github.com/l10-bhushan/crispy-fiesta/internal/apperrors"
	"github.com/l10-bhushan/crispy-fiesta/internal/models"
	"github.com/l10-bhushan/crispy-fiesta/internal/repository"
	"github.com/l10-bhushan/crispy-fiesta/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) FindByEmail(ctx context.Context, email string) (*models.UserResponse, error) {
	// Validating the email
	validEmail := utils.ValidateEmail(email)

	// If email is not valid return error
	if !validEmail {
		return nil, apperrors.ErrorInvalidInput
	}

	// Return the err, if any error in db
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	// Return the user
	return user, nil
}

func (s *UserService) RegisterUser(ctx context.Context, userRequest models.RegisterUser) (*models.UserResponse, error) {

	// Checking if email is valid
	if !utils.ValidateEmail(userRequest.Email) {
		return nil, apperrors.ErrorInvalidInput
	}

	// Checking if user already exists or not
	existingUser, err := s.repo.FindByEmail(ctx, userRequest.Email)
	if err == nil && existingUser != nil {
		return nil, apperrors.ErrorConflict
	}

	// Hashing the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, apperrors.ErrorInternal
	}

	userRequest.Password = string(hashedPassword)

	createdUser, err := s.repo.RegisterUser(ctx, userRequest)
	if err != nil {
		return nil, apperrors.HandleDBErrors(err)
	}

	return createdUser, nil
}
