package api

import (
	"WASA/service/api/reqcontext"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// removeCommentHandler handles the HTTP request to remove a comment.
func (rt *_router) removeCommentHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	token := extractBearerToken(r)
	if !validToken(token) {
		http.Error(w, "Unauthorized: Invalid or missing token", http.StatusUnauthorized)
		return
	}
	// Extract the comment ID from the URL parameters and convert it to an integer.
	commentId, err := strconv.Atoi(ps.ByName("commentId"))
	if err != nil {
		// If the conversion fails, return an HTTP 400 Bad Request error.
		http.Error(w, "Invalid comment ID", http.StatusBadRequest)
		return
	}
	// Extract the user ID of the admin or the user who is performing the ban from the URL parameters.
	userId, err := strconv.Atoi(ps.ByName("userId"))
	if err != nil {
		// If the user ID is not valid, return an HTTP 400 Bad Request error.
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	// Retrieve the comment to check its author.
	comment, err := rt.db.GetCommentById(commentId)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Failed to retrieve comment")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Ensure the user trying to delete the comment is the author of the comment.
	if token != comment.UserID {
		http.Error(w, "Forbidden: You do not have permission to perform this action", http.StatusForbidden)
		return
	}
	// Call the RemoveComment method in the database to delete the comment.
	err = rt.db.RemoveComment(commentId)
	if err != nil {
		// If the RemoveComment operation fails, log the error and return an HTTP 500 Internal Server Error.
		rt.baseLogger.WithError(err).Error("Failed to remove comment")
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
