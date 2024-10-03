package auth

import (
	"crypto/rand"
	"encoding/base64"
	"slices"
)

type AuthState struct {
	ReturnTo string `json:"returnTo"`
	Checksum string `json:"checksum"`
}

func NewAuthState(returnTo string) (*AuthState, error) {
	checksum, err := generateChecksum()
	if err != nil {
		return nil, err
	}
	state := AuthState{
		ReturnTo: returnTo,
		Checksum: checksum,
	}

	return &state, nil
}

func ValidateReturnTo(returnTo string) bool {
	if returnTo == "" {
		return false
	}
	if !slices.Contains(validLoginRedirects, returnTo) {
		return false
	}
	return true
}

func generateChecksum() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	checksum := base64.StdEncoding.EncodeToString(b)
	return checksum, nil
}
