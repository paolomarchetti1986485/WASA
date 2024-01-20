package api

import (
	"WASA/service/api/reqcontext"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// getStreamHandler handles the HTTP request to retrieve the user's stream.
// The stream typically consists of photos or content from users that the requesting user follows.
func (rt *_router) getStreamHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	token := extractBearerToken(r)
	if !validToken(token) {
		http.Error(w, "Unauthorized: Invalid or missing token", http.StatusUnauthorized)
		return
	}
	// Extract the user ID from the URL parameters and convert it to an integer.
	userId, err := strconv.Atoi(ps.ByName("userId"))
	if err != nil {
		// If the conversion fails, return an HTTP 400 Bad Request error.
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Retrieve the user's stream (e.g., photos from followed users) from the database.
	stream, err := rt.db.GetUserStream(userId)
	if err != nil {
		// If the operation fails, log the error and return an HTTP 500 Internal Server Error.
		rt.baseLogger.WithError(err).Error("Failed to get user stream")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if token != userId {
		http.Error(w, "Forbidden: You do not have permission to perform this action", http.StatusForbidden)
		return
	}
	// Set the Content-Type header to application/json and encode the stream into JSON.
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(stream)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Failed to send response")
	}
}
