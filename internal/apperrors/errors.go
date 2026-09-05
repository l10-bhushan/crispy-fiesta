package apperrors

import (
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

var (
	ErrorNotFound     = errors.New("resource not found")
	ErrorInternal     = errors.New("internal server error")
	ErrorInvalidInput = errors.New("invalid input")
	ErrorConflict     = errors.New("resource already exists")
)

func HandleDBErrors(err error) error {

	if errors.Is(err, pgx.ErrNoRows) {
		return ErrorNotFound
	}

	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr) {

		switch pgErr.Code {
		case "23505":
			return ErrorConflict
		case "23503":
			return ErrorInvalidInput
		case "23502":
			return ErrorInvalidInput
		case "23514":
			return ErrorInvalidInput
		}
	}

	return ErrorInternal
}
