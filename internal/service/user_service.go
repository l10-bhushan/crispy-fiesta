package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/l10-bhushan/crispy-fiesta/internal/apperrors"
	"github.com/l10-bhushan/crispy-fiesta/internal/models"
	"github.com/l10-bhushan/crispy-fiesta/internal/repository"
	"github.com/l10-bhushan/crispy-fiesta/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo        *repository.UserRepository
	authService *AuthService
}

func NewUserService(repo *repository.UserRepository, authService *AuthService) *UserService {
	return &UserService{
		repo:        repo,
		authService: authService,
	}
}

// Service to fetch all users
func (s *UserService) FetchAll(ctx context.Context) ([]models.UserResponse, error) {

	users, err := s.repo.FetchAll(ctx)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) FetchUserInformation(ctx context.Context, userId uuid.UUID) (*models.UserResponse, error) {
	// // Validating the email
	// validEmail := utils.ValidateEmail(email)

	// // If email is not valid return error
	// if !validEmail {
	// 	return nil, apperrors.ErrorInvalidInput
	// }

	// Return the err, if any error in db
	user, err := s.repo.FetchUserInformation(ctx, userId)
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

// Service to delete User
func (s *UserService) DeleteUser(ctx context.Context, id string) error {

	err := s.repo.DeleteUser(ctx, id)

	if err != nil {
		return err
	}

	return nil
}

// Service for user login
func (s *UserService) Login(ctx context.Context, loginRequest models.LoginRequest) (*string, error) {

	// validating the email for correct format
	if !utils.ValidateEmail(loginRequest.Email) {
		return nil, apperrors.ErrorInvalidInput
	}

	// We are using FindByEmail to fetch user data based on email
	userResponse, err := s.repo.FindByEmail(ctx, loginRequest.Email)
	if err != nil {
		if errors.Is(err, apperrors.ErrorNotFound) {
			return nil, apperrors.ErrorUnauthorized
		}
		return nil, err
	}

	// Checking if the password matches, if not we return error
	if err := bcrypt.CompareHashAndPassword([]byte(userResponse.Password), []byte(loginRequest.Password)); err != nil {
		return nil, apperrors.ErrorUnauthorized
	}

	token, err := s.authService.GenerateToken(userResponse.Id)
	if err != nil {
		return nil, err
	}
	// Returning the user information if password matches
	return &token, nil

}
