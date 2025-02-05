package api

import (
	"WASA/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

// DeleteUserHandler handles the API request to delete a user by ID.
func (rt *_router) DeleteUserHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Extracting the user ID from the URL parameter.
	userIDStr := ps.ByName("userId")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Calling the database function to delete the user.
	err = rt.db.DeleteUserByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Sending a successful response.
	w.WriteHeader(http.StatusNoContent)
}
