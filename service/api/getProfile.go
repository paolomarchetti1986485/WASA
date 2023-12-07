package api

import (
	"WASA/service/api/reqcontext"
	"WASA/service/database"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

// getProfileHandler handles HTTP requests to retrieve a user's profile information.
func (rt *_router) getProfileHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Initialize slices to store followers, following users, and photos.
	var followers []database.User
	var following []database.User
	var photos []database.Photo

	// Extract the user ID from the URL parameters and convert it to an integer.
	userId, err := strconv.Atoi(ps.ByName("userId"))
	if err != nil {
		// If the conversion fails, return an HTTP 400 Bad Request error.
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Retrieve the photos uploaded by the user.
	photos, err = rt.db.GetUserPhotos(userId)
	if err != nil {
		// If the operation fails, log the error and return an HTTP 500 Internal Server Error.
		rt.baseLogger.WithError(err).Error("Failed to get user photos")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Retrieve the followers of the user.
	followers, err = rt.db.GetUserFollowers(userId)
	if err != nil {
		// If the operation fails, log the error and return an HTTP 500 Internal Server Error.
		rt.baseLogger.WithError(err).Error("Failed to get user followers")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Retrieve the users that the user is following.
	following, err = rt.db.GetUserFollowing(userId)
	if err != nil {
		// If the operation fails, log the error and return an HTTP 500 Internal Server Error.
		rt.baseLogger.WithError(err).Error("Failed to get user following")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Aggregate the retrieved information into a profile struct.
	profile := Profile{
		Photos:    photos,
		Followers: followers,
		Following: following,
	}

	// Set the Content-Type header to application/json and encode the profile into JSON.
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(profile)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Failed to send response")
	}
}
