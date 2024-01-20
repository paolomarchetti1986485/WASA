package api

import (
	"WASA/service/api/reqcontext"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// addCommentHandler is an HTTP handler for adding a comment to a photo.
// It extracts user and photo IDs from the request parameters and decodes the comment from the request body.
func (rt *_router) addCommentHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	token := extractBearerToken(r)
	if !validToken(token) {
		http.Error(w, "Unauthorized: Invalid or missing token", http.StatusUnauthorized)
		return
	}
	// Extract the user ID from the URL parameters.
	userId, err := strconv.Atoi(ps.ByName("userId"))
	if err != nil {
		// If the user ID is not valid, return an HTTP 400 Bad Request error.
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Extract the photo ID from the URL parameters.
	photoId, err := strconv.Atoi(ps.ByName("photoId"))
	if err != nil {
		// If the photo ID is not valid, return an HTTP 400 Bad Request error.
		http.Error(w, "Invalid photo ID", http.StatusBadRequest)
		return
	}

	// Decode the comment from the request body.
	var comment Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		// If decoding fails, return an HTTP 400 Bad Request error.
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Call the AddComment method to add the comment to the database.
	commentID, err := rt.db.AddComment(photoId, userId, comment.Text)
	if err != nil {
		// If the AddComment operation fails, log the error and return an HTTP 500 Internal Server Error.
		rt.baseLogger.WithError(err).Error("Failed to add comment")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if token != userId {
		http.Error(w, "Forbidden: You do not have permission to perform this action", http.StatusForbidden)
		return
	}
	// Prepare the response containing the comment ID.
	w.Header().Set("Content-Type", "application/json") // Set the content type to JSON.
	w.WriteHeader(http.StatusCreated)                  // Send a 201 HTTP status code indicating the resource was successfully created.
	err = json.NewEncoder(w).Encode(map[string]int{"commentId": commentID})
	if err != nil {
		rt.baseLogger.WithError(err).Error("Failed to send response") // Encode and send the response with the comment ID.
	}

}
