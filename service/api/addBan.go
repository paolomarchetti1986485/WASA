package api

import (
	"WASA/service/api/reqcontext"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

// banUserHandler is an HTTP handler for banning a user.
// It retrieves user IDs from the request parameters and performs the ban operation.
func (rt *_router) banUserHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Extract the user ID of the admin or the user who is performing the ban from the URL parameters.
	userId, err := strconv.Atoi(ps.ByName("userId"))
	if err != nil {
		// If the user ID is not valid, return an HTTP 400 Bad Request error.
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Extract the user ID of the user who is to be banned from the URL parameters.
	bannedId, err := strconv.Atoi(ps.ByName("bannedId"))
	if err != nil {
		// If the banned user's ID is not valid, return an HTTP 400 Bad Request error.
		http.Error(w, "Invalid banned ID", http.StatusBadRequest)
		return
	}

	// Call the BanUser method to ban the user.
	err = rt.db.BanUser(bannedId, userId)
	if err != nil {
		// If the BanUser operation fails, log the error and return an HTTP 500 Internal Server Error.
		rt.baseLogger.WithError(err).Error("Failed to ban user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Prepare a success response message.
	response := map[string]string{"message": "User has been successfully banned."}
	w.Header().Set("Content-Type", "application/json") // Set the content type to JSON.
	w.WriteHeader(http.StatusCreated)                  // Send a 201 HTTP status code indicating the resource was successfully created.
	err = json.NewEncoder(w).Encode(response)          // Encode and send the response message.
	if err != nil {
		rt.baseLogger.WithError(err).Error("Failed to send response")
	}

}
