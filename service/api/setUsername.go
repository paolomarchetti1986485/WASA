package api

import (
	"WASA/service/api/reqcontext"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// UsernameUpdateRequest represents the request structure to update a username.
type UsernameUpdateRequest struct {
	Username string `json:"username"`
}

// UsernameUpdateResponse represents the response structure after updating a username.
type UsernameUpdateResponse struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

// setUsernameHandler handles the HTTP request to update a user's username.
func (rt *_router) setUsernameHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Extract the user ID from the URL parameters and convert it to an integer.
	userId, err := strconv.Atoi(ps.ByName("userId"))
	if err != nil {
		// If the user ID is invalid, return an HTTP 400 Bad Request error.
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Decode the request body into the UsernameUpdateRequest struct.
	var req UsernameUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// If decoding fails, return an HTTP 400 Bad Request error.
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the new username using the IsValidUsername function.
	if !IsValidUsername(req.Username) {
		// If the username is invalid, return an HTTP 400 Bad Request error.
		http.Error(w, "Invalid username", http.StatusBadRequest)
		return
	}

	// Call the UpdateUsername method to update the user's username in the database.
	err = rt.db.UpdateUsername(userId, req.Username)
	if err != nil {
		if err.Error() == "username already exists" {
			// Se il nome utente è già in uso, restituisci un HTTP 409 Conflict
			http.Error(w, "username already exists", http.StatusConflict)
			return
		}
		// Se c'è un altro tipo di errore, restituisci un HTTP 500 Internal Server Error
		rt.baseLogger.WithError(err).Error("Failed to update username")
		http.Error(w, "Failed to update username", http.StatusInternalServerError)
		return
	}

	// Prepare a response indicating the successful username update.
	resp := UsernameUpdateResponse{
		Username: req.Username,
		Message:  "Username has been updated successfully",
	}
	w.Header().Set("Content-Type", "application/json") // Set the content type to JSON.
	err = json.NewEncoder(w).Encode(resp)              // Encode and send the response.
	if err != nil {
		rt.baseLogger.WithError(err).Error("Failed to send response")
	}
}
