package api

import (
	"WASA/service/api/reqcontext"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

// uploadPhotoHandler handles the HTTP request to upload a photo.
func (rt *_router) uploadPhotoHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	token := extractBearerToken(r)
	if !validToken(token) {
		http.Error(w, "Unauthorized: Invalid or missing token", http.StatusUnauthorized)
		return
	}
	// Extract and validate the user ID from URL parameters.
	userId, err := strconv.Atoi(ps.ByName("userId"))
	if err != nil {
		// If the user ID is invalid, return an HTTP 400 Bad Request error.
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Retrieve the file from the form data.
	file, _, err := r.FormFile("file")
	if err != nil {
		// If retrieving the file fails, return an error.
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close() // Ensure the file is closed after processing.

	// Read the contents of the file.
	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		// If reading the file content fails, return an internal server error.
		http.Error(w, "Error reading the file", http.StatusInternalServerError)
		return
	}

	// Upload the photo to the database using the current time as the upload timestamp.
	photoID, err := rt.db.UploadPhoto(userId, time.Now(), fileContent) // fileContent is of type []byte.
	if err != nil {
		// If the upload operation fails, log the error and return an internal server error.
		rt.baseLogger.WithError(err).Error("Failed to upload photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if token != userId {
		http.Error(w, "Forbidden: You do not have permission to perform this action", http.StatusForbidden)
		return
	}
	// Prepare and send a successful response back to the client.
	w.Header().Set("Content-Type", "application/json") // Set the content type to JSON.
	w.WriteHeader(http.StatusCreated)                  // Send a 201 HTTP status code indicating the resource was successfully created.
	err = json.NewEncoder(w).Encode(map[string]int{"photoId": photoID})
	if err != nil { // Encode and send the photo ID in the response.
		rt.baseLogger.WithError(err).Error("Failed to send response")
	}
}
