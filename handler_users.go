package main

import (
	"fmt"
	"net/http"

	"github.com/bsouch/GoChatApp/internal/crypto"
	"github.com/bsouch/GoChatApp/internal/database"
	"github.com/google/uuid"
)

func (apiCfg apiConfig) handlerCreateUsers(w http.ResponseWriter, r *http.Request) {
	//Deserialise request body
	type parameters struct {
		UserName string `json:"UserName"`
		Password string `json:"Password"`
	}

	params := parameters{}
	jsonErr := jsonDeserialise(r, &params)
	if jsonErr != nil {
		errorJsonResponse(w, http.StatusBadRequest, fmt.Sprintf("Unable to deserialise request object in Create User. Error: %v", jsonErr))
	}

	hashedPassword, hashErr := crypto.HashPassword(params.Password)
	if hashErr != nil {
		errorJsonResponse(w, http.StatusBadRequest, fmt.Sprintf("Error hashing password: %v", hashErr))
	}

	userDbo := database.CreateUserParams{
		UserID:   uuid.New(),
		UserName: params.UserName,
		Password: hashedPassword,
	}

	dboUser, dbErr := apiCfg.DB.CreateUser(r.Context(), userDbo)
	if dbErr != nil {
		errorJsonResponse(w, http.StatusInternalServerError, fmt.Sprintf("Error creating new user: %v", dbErr))
	}

	jsonResponse(w, http.StatusOK, dboUserToUser(dboUser))
}
