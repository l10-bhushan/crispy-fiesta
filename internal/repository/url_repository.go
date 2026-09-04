// Repository layer for URL api
package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/l10-bhushan/crispy-fiesta/internal/models"
)

// We create this structure, suppose this as an object that can talk to database
// This is a go design pattern, we create a structure that holds the db object
// That we can later use to talk to database
type URLRepository struct {
	db *pgxpool.Pool
}

// This is a common constructor pattern in go.
// We are initializing a URLRepository object and returning a pointer to it.
// We are using a pointer here because we don't wanna have multiple copies of URLRepository.
// return type for & is *
func NewURLRepository(db *pgxpool.Pool) *URLRepository {
	return &URLRepository{
		db: db,
	}
}

// Function to create an entry in urls table in database
// it takes context, shortcode, originalurl and we return the response and error
func (r *URLRepository) Create(ctx context.Context, shortCode string, originalURL string) (*models.CreateShortURLResponse, error) {

	// Creating an instance of CreateShortURLResponse struct
	// we are using & here coz this will return the address of the instance created
	urlResponse := &models.CreateShortURLResponse{
		ShortCode: shortCode,
		URL:       originalURL,
	}

	// Query for inserting data into our table and we are also expecting returning value
	query := `INSERT INTO urls (short_code, original_url) VALUES ($1, $2) RETURNING id, created_at, expires_at`

	// Using db.QueryRow to execute the query and using scan to store the returning values into our struct instance
	err := r.db.QueryRow(ctx, query, shortCode, originalURL).Scan(&urlResponse.Id, &urlResponse.CreatedAt, &urlResponse.ExpiresAt)
	if err != nil {
		return nil, err
	}

	return urlResponse, nil
}

// Function to fetch the data by shortCode
func (r *URLRepository) FindByShortCode(ctx context.Context, short_code string) (*models.CreateShortURLResponse, error) {
	// Creating the CreateShortURLResponse instance and storing it's address
	urlResponse := &models.CreateShortURLResponse{}

	// Query to fetch data using shortcode
	query := `SELECT id, short_code, original_url,created_at FROM urls WHERE short_code = $1`

	// QueryRow to fetch data and using scan to store data in urlResponse instance
	err := r.db.QueryRow(ctx, query, short_code).Scan(&urlResponse.Id, &urlResponse.ShortCode, &urlResponse.URL, &urlResponse.CreatedAt)
	if err != nil {
		return nil, err
	}

	return urlResponse, nil
}
