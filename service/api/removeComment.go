package api

import (
    "net/http"
    "strconv"
    "github.com/julienschmidt/httprouter"
)

// removeCommentHandler handles the HTTP request to remove a comment.
func (rt *_router) removeCommentHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    // Extract the comment ID from the URL parameters and convert it to an integer.
    commentId, err := strconv.Atoi(ps.ByName("commentId"))
    if err != nil {
        // If the conversion fails, return an HTTP 400 Bad Request error.
        http.Error(w, "Invalid comment ID", http.StatusBadRequest)
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

    // If the operation is successful, respond with a 204 No Content status.
    // This indicates that the action has been successfully processed, but there is no content to return.
    w.WriteHeader(http.StatusNoContent)
}
