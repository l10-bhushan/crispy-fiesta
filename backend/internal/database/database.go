package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Function to help us create a pool
func NewPool(ctx context.Context, databaseURL string) (*pgxpool.Pool, error) {
	// Creating a new pool
	// Creates the pool configuration/connection pool
	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create a pool : %v", err)
	}

	// Checking the connection to the pool
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to connect to db : %v", err)
	}

	// If everything goes well, we will return the pool
	return pool, nil
}
