package utils

import "github.com/google/uuid"

func GenUUID() string {
	newUUID, _ := uuid.NewUUID()
	return newUUID.String()
}
