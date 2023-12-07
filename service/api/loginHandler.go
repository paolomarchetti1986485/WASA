package api

import (
	"WASA/service/api/reqcontext"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// loginHandler handles the HTTP request for user login.
// It decodes the user's credentials from the request and either logs them in or creates a new user.
func (rt *_router) loginHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	// Decode the user data from the request body.
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// If the JSON decoding fails, log a warning and return a bad request status.
		rt.baseLogger.WithError(err).Warning("Wrong JSON received")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if the user exists in the database.
	userID, err := rt.db.GetUserByUsername(user.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// If the user does not exist, create a new user.
			userID, err = rt.db.AddUser(user.Username)
			if err != nil {
				// If creating a new user fails, log an error and return an internal server error status.
				rt.baseLogger.WithError(err).Error("Failed to create new user")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		} else {
			// If there is a database error other than ErrNoRows, log an error and return an internal server error status.
			rt.baseLogger.WithError(err).Error("Database error")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	// At this point, the user either existed or has been created.
	// Generate a response, such as returning the user's ID.
	// You might also want to return a token or additional user info here.
	response := map[string]int{"userId": userID}
	w.Header().Set("Content-Type", "application/json") // Set response content type to JSON.
	w.WriteHeader(http.StatusOK)                       // Send an HTTP 200 OK status.
	err = json.NewEncoder(w).Encode(response)
	if err != nil { // Encode and send the response.
		rt.baseLogger.WithError(err).Error("Failed to send response")
	}
}
