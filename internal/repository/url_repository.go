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

// Fetch all the url data from the db
// Here, we are not using pointer return coz we can return nil when an error for slice
// in the above example if we don't use pointer we cannot return nil, coz retuning nil for an empty
// struct won't compile
func (r *URLRepository) FetchAllData(ctx context.Context) ([]models.CreateShortURLResponse, error) {

	// Query to fetch all the records from the urls table
	query := `SELECT * FROM urls ORDER BY DESC`

	// Executing the query using r.db.Query, returns pgx.rows
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	// closing the connection after return
	defer rows.Close()

	// Creating an array instance
	var urls []models.CreateShortURLResponse

	// looping through the rows
	for rows.Next() {
		var url models.CreateShortURLResponse

		// Populating the data into respective fields
		err := rows.Scan(&url.Id, &url.ShortCode, &url.URL, &url.CreatedAt)
		if err != nil {
			return nil, err
		}

		// appending url to urls
		urls = append(urls, url)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return urls, nil
}

// Fetch url data using the id
func (r *URLRepository) GetById(ctx context.Context, id string) (*models.CreateShortURLResponse, error) {

	// Create an instance and return the address the address is store in url
	url := &models.CreateShortURLResponse{}

	// Query to fetch url data from the urls table
	query := `SELECT id, short_code , original_url , created_at, expires_at FROM urls WHERE id = $1`

	// Querying from the db
	err := r.db.QueryRow(ctx, query, id).Scan(&url.Id, &url.ShortCode, &url.URL, &url.CreatedAt, &url.ExpiresAt)
	if err != nil {
		return nil, err
	}

	return url, nil
}
