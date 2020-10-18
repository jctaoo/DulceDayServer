package helpers

import uuid "github.com/satori/go.uuid"

func GenerateRandomUserID() string {
	return uuid.NewV4().String()
}