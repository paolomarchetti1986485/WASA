package api

import (
	"WASA/service/api/reqcontext"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// GetAllUsersHandler handles the API request to get all users.
func (rt *_router) GetAllUsersHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params, ctx reqcontext.RequestContext) {
	users, err := rt.db.GetAllUsers()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if len(users) == 0 {
		http.Error(w, "No users found", http.StatusNotFound)
		return
	}

	// Respond with the list of users in JSON format.
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
	}

}
