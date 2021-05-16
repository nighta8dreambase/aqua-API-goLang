package services

import (
	"github.com/satori/go.uuid"
	"log"
)

func GenUUIDv4() uuid.UUID{
	var err error
	id := uuid.NewV4()
	if err != nil {
		log.Fatalf("uuid.NewV4() failed with %s\n", err)
	}
	return id
}