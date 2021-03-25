package handlers

import (
	"net/http"
	"strconv"

	"github.com/ryanpback/raspberry_a_pi/helpers"
	"github.com/ryanpback/raspberry_a_pi/models"

	"github.com/gorilla/mux"
)

// UsersCreate creates a new user
func UsersCreate(w http.ResponseWriter, r *http.Request) {
	request, err := decode(r)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	if valid, err := helpers.ValidateUser(request, helpers.GetUserCreateFields()); !valid {
		respondFailedValidation(w, err)

		return
	}

	user, err := models.UserCreate(request)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	response := Response{
		"user": user,
	}
	respondWithJSON(w, http.StatusCreated, response)
}

// UsersShow retrieves a single user based on ID
func UsersShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	userID, _ := strconv.Atoi(id)

	user, err := models.UserFindByID(userID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())

		return
	}

	response := Response{
		"user": user,
	}
	respondWithJSON(w, http.StatusOK, response)
}

// UsersEdit will edit a user (passed by ID)
func UsersEdit(w http.ResponseWriter, r *http.Request) {
	request, err := decode(r)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	if valid, err := helpers.ValidateUser(request, helpers.GetUserEditFields()); !valid {
		respondFailedValidation(w, err)

		return
	}

	params := mux.Vars(r)
	id := params["id"]
	userID, _ := strconv.Atoi(id)

	user, err := models.UserFindByID(userID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())

		return
	}

	updatedUser, err := models.UserEdit(user, request)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	response := Response{
		"user": updatedUser,
	}
	respondWithJSON(w, http.StatusOK, response)
}
