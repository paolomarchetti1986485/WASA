package api

import (
	"WASA/service/api/reqcontext"
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// getPhotoHandler handles HTTP requests to retrieve a photo's image data.
func (rt *_router) getPhotoHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	token := extractBearerToken(r)
	if !validToken(token) {
		http.Error(w, "Unauthorized: Invalid or missing token", http.StatusUnauthorized)
		return
	}

	// Extract the photo ID from the URL parameters and convert it to an integer.
	photoID, err := strconv.Atoi(ps.ByName("photoId"))
	if err != nil {
		http.Error(w, "Invalid photo ID", http.StatusBadRequest)
		return
	}

	// Retrieve the photo data from the database.
	imageData, err := rt.db.GetPhotoData(photoID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Image not found", http.StatusNotFound)
		} else {
			rt.baseLogger.WithError(err).Error("Failed to get photo data")
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	// Set the Content-Type header to image/jpeg and write the image data to the response.
	w.Header().Set("Content-Type", "image/jpeg")
	_, err = w.Write(imageData)
	if err != nil {
		// Se la scrittura della risposta fallisce, logga l'errore
		rt.baseLogger.WithError(err).Error("Failed to write image data to response")
	}
}
