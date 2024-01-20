package api

import (
	"WASA/service/api/reqcontext"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// deletePhotoHandler handles HTTP requests to delete a specific photo.
// It extracts the photo ID from the request parameters and performs the delete operation.
func (rt *_router) deletePhotoHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	token := extractBearerToken(r)
	if !validToken(token) {
		http.Error(w, "Unauthorized: Invalid or missing token", http.StatusUnauthorized)
		return
	}
	// Extract the photo ID from the URL parameters and convert it to an integer.
	photoId, err := strconv.Atoi(ps.ByName("photoId"))
	if err != nil {
		// If the conversion fails, return an HTTP 400 Bad Request error.
		http.Error(w, "Invalid photo ID", http.StatusBadRequest)
		return
	}
	userId, err := strconv.Atoi(ps.ByName("userId"))
	if err != nil {
		// If the user ID is invalid, return an HTTP 400 Bad Request error.
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	// Call the RemovePhoto method to delete the photo from the database.
	err = rt.db.RemovePhoto(photoId)
	if err != nil {
		// If the RemovePhoto operation fails, log the error and return an HTTP 500 Internal Server Error.
		rt.baseLogger.WithError(err).Error("Failed to delete photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if token != userId {
		http.Error(w, "Forbidden: You do not have permission to perform this action", http.StatusForbidden)
		return
	}
	// If the operation is successful, respond with a 204 No Content status.
	// This indicates that the action has been successfully processed, but there is no content to return.
	w.WriteHeader(http.StatusNoContent)
}
