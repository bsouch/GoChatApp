package main

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID    uuid.UUID
	UserName  string
	Password  []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}

/*
func dboUserToUser(dboUser database.User) User {
	return User{
		UserID:   dboUser.UserID,
		UserName: dboUser.UserName,
	}
}
*/
