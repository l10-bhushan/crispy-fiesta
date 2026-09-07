package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/l10-bhushan/crispy-fiesta/internal/apperrors"
	"github.com/l10-bhushan/crispy-fiesta/internal/models"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Fetch All users
func (r *UserRepository) FetchAll(ctx context.Context) ([]models.UserResponse, error) {

	// Query to fetch all the users
	query := `SELECT id , first_name, last_name, email, created_at FROM users`

	// Executing the query
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, apperrors.HandleDBErrors(err)
	}

	// Closing the connection
	defer rows.Close()

	// Creating a slice to store the rows returned from rows
	// Here, we can use models directly without using & because we can return nil for slices
	var users []models.UserResponse
	for rows.Next() {
		var user models.UserResponse

		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, apperrors.HandleDBErrors(err)
		}

		users = append(users, user)
	}

	return users, nil
}

// Fetch the user information based on email
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.UserResponse, error) {
	// Below we use & because there is a difference between returning value and returning pointer
	// Suppose, in the below instance we are creating a models.UserResponse and assigning the address
	// to user variable
	// So, while returning we can return either the struct pointer or nil. i.e users or nil
	// but if we use user := models.UserResponse. while returning we will have to pass an empty struct
	// i.e return models.UserResponse, error
	user := &models.UserResponse{}
	query := `SELECT id, first_name, last_name, email, created_at FROM users WHERE email = $1`

	err := r.db.QueryRow(ctx, query, email).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, apperrors.HandleDBErrors(err)
	}
	return user, nil
}

func (r *UserRepository) RegisterUser(ctx context.Context, userRequest models.RegisterUser) (*models.UserResponse, error) {

	// Creating an instance of models.UserResponse and storing its address in user
	user := &models.UserResponse{}

	// Query to create user data
	query := `INSERT INTO users (first_name, last_name, email , password) VALUES ($1, $2, $3, $4) RETURNING id, first_name, last_name, email, created_at`

	// Executing the query
	err := r.db.QueryRow(ctx, query, userRequest.FirstName, userRequest.LastName, userRequest.Email, userRequest.Password).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt)

	if err != nil {
		return nil, apperrors.HandleDBErrors(err)
	}

	return user, nil
}

// Query to Delete user
func (r *UserRepository) DeleteUser(ctx context.Context, id string) error {

	// Query to delete User
	query := "DELETE FROM users WHERE id = $1"

	// Executing the query
	_, err := r.db.Exec(ctx, query)
	if err != nil {
		return apperrors.HandleDBErrors(err)
	}

	return nil
}
