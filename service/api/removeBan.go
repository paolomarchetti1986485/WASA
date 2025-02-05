package api

import (
	"WASA/service/api/reqcontext"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"strconv"
)

// unbanUserHandler handles the HTTP request to unban a user.
func (rt *_router) unbanUserHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	token := extractBearerToken(r)
	if !validToken(token) {
		http.Error(w, "Unauthorized: Invalid or missing token", http.StatusUnauthorized)
		return
	}
	// Extract the user ID (of the admin or moderator) from the URL parameters and convert it to an integer.
	userId, err := strconv.Atoi(ps.ByName("userId"))
	if err != nil {
		// If the user ID is invalid, return an HTTP 400 Bad Request error.
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Extract the banned user's ID from the URL parameters and convert it to an integer.
	bannedId, err := strconv.Atoi(ps.ByName("bannedId"))
	if err != nil {
		// If the banned user's ID is invalid, return an HTTP 400 Bad Request error.
		http.Error(w, "Invalid banned ID", http.StatusBadRequest)
		return
	}

	// Call the UnbanUser method to unban the user.
	err = rt.db.UnbanUser(bannedId, userId)
	if err != nil {
		// If the UnbanUser operation fails, log the error and return an HTTP 500 Internal Server Error.
		rt.baseLogger.WithError(err).Error("Failed to unban user")
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
