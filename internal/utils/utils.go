package utils

import (
	"crypto/rand"
	"errors"
	"net/url"
)

// Utility function to validate URL
// We will use the in-build net/url package it will return error if any or nil if successfull
func ValidateURL(value string) error {

	// Using url.ParseRequestURI that will check if URL is correct
	parsedURL, err := url.ParseRequestURI(value)
	if err != nil {
		return errors.New("Invalid URL")
	}

	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return errors.New("Invalid URL")
	}

	return nil
}

func GenerateShortCode() (string, error) {
	// Specifying the length of short code
	length := 8

	const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Generates a 8 bytes array derived from the length
	bytes := make([]byte, length)
	// [0, 0, 0, 0, 0, 0, 0, 0 ]

	// rand.Read fills the bytes array with random numbers
	if _, err := rand.Read(bytes); err != nil {
		return "", errors.New("failed to created short code")
	}

	// Filling in the array with charachters
	for i := range bytes {
		bytes[i] = characters[int(bytes[i])%len(characters)]
	}

	return string(bytes), nil
}
