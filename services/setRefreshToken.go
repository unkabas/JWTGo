package services

import (
	"crypto/rand"
	"encoding/base64"
)

func SetRefresh() (string, error) {
	buffer := make([]byte, 64)

	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}
	token := base64.URLEncoding.EncodeToString(buffer)
	if len(token) > 64 {
		token = token[:64]
	}
	return token, nil
}
