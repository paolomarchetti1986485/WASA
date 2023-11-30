package api

import (
	"net/http"
	
    "strconv"

	"github.com/julienschmidt/httprouter"
)
func (rt *_router) removeCommentHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    // Extracting comment ID from the URL parameters
    commentId, err := strconv.Atoi(ps.ByName("commentId"))
    if err != nil {
        http.Error(w, "Invalid comment ID", http.StatusBadRequest)
        return
    }

    // Call to database function to remove the comment
    err = rt.db.RemoveComment(commentId)
    if err != nil {
        rt.baseLogger.WithError(err).Error("Failed to remove comment")
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    // Sending back a successful response with no content
    w.WriteHeader(http.StatusNoContent)
}
