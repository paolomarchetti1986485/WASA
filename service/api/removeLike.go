package api

import (
    "net/http"
    "strconv"
    "github.com/julienschmidt/httprouter"
)

// removeLikeHandler handles the HTTP request to remove a 'like' from a photo.
func (rt *_router) removeLikeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    // Extract the like ID (which is being used as the user ID in this context) from the URL parameters and convert it to an integer.
    likeId, err := strconv.Atoi(ps.ByName("likeId"))
    if err != nil {
        // If the like ID is invalid, return an HTTP 400 Bad Request error.
        http.Error(w, "Invalid like ID", http.StatusBadRequest)
        return
    }

    // Extract the photo ID from the URL parameters and convert it to an integer.
    photoId, err := strconv.Atoi(ps.ByName("photoId"))
    if err != nil {
        // If the photo ID is invalid, return an HTTP 400 Bad Request error.
        http.Error(w, "Invalid photo ID", http.StatusBadRequest)
        return
    }

    // Call the RemoveLike method to remove the like from the photo.
    err = rt.db.RemoveLike(photoId, likeId)
    if err != nil {
        // If the RemoveLike operation fails, log the error and return an HTTP 500 Internal Server Error.
        rt.baseLogger.WithError(err).Error("Failed to remove like")
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    // If the operation is successful, respond with a 204 No Content status.
    // This indicates that the action has been successfully processed, but there is no content to return.
    w.WriteHeader(http.StatusNoContent)
}
