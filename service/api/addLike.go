package api

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/julienschmidt/httprouter"
)

// addLikeHandler handles HTTP requests to add a 'like' to a photo.
// It extracts likeId and photoId from the request parameters and performs the like operation.
func (rt *_router) addLikeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    // Extract the likeId from the URL parameters and convert it to an integer.
    // In this context, likeId is used interchangeably with userId.
    likeId, err := strconv.Atoi(ps.ByName("likeId"))
    if err != nil {
        // If the conversion fails, return an HTTP 400 Bad Request error.
        http.Error(w, "Invalid like ID", http.StatusBadRequest)
        return
    }

    // Extract the photoId from the URL parameters and convert it to an integer.
    photoId, err := strconv.Atoi(ps.ByName("photoId"))
    if err != nil {
        // If the conversion fails, return an HTTP 400 Bad Request error.
        http.Error(w, "Invalid photo ID", http.StatusBadRequest)
        return
    }

    // Call the AddLike method to add the like to the photo.
    err = rt.db.AddLike(photoId, likeId)
    if err != nil {
        // If the AddLike operation fails, log the error and return an HTTP 500 Internal Server Error.
        rt.baseLogger.WithError(err).Error("Failed to add like")
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    // Prepare the response containing the likeId.
    response := map[string]LikeId{"likeId": LikeId(likeId)}
    w.Header().Set("Content-Type", "application/json") // Set the content type to JSON.
    w.WriteHeader(http.StatusCreated) // Send a 201 HTTP status code indicating the resource was successfully created.
    json.NewEncoder(w).Encode(response) // Encode and send the response with the likeId.
}
