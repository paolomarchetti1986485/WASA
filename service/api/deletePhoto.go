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
	// Extract the photo ID from the URL parameters and convert it to an integer.
	photoId, err := strconv.Atoi(ps.ByName("photoId"))
	if err != nil {
		// If the conversion fails, return an HTTP 400 Bad Request error.
		http.Error(w, "Invalid photo ID", http.StatusBadRequest)
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

	// If the operation is successful, respond with a 204 No Content status.
	// This indicates that the action has been successfully processed, but there is no content to return.
	w.WriteHeader(http.StatusNoContent)
}
