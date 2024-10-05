package api

import (
	"WASA/service/api/reqcontext"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// getProfileHandler handles HTTP requests to retrieve a user's profile information.
func (rt *_router) getProfileHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	token := extractBearerToken(r)
	if !validToken(token) {
		http.Error(w, "Unauthorized: Invalid or missing token", http.StatusUnauthorized)
		return
	}

	// Extract the user ID from the URL parameters and convert it to an integer.
	userId, err := strconv.Atoi(ps.ByName("userId"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Retrieve the user's profile information.
	profile, err := rt.db.GetUserProfile(userId)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Failed to get user profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json and encode the profile into JSON.
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(profile)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Failed to send response")
	}
}
