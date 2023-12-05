package api

import (
    "net/http"
    "strconv"
    "github.com/julienschmidt/httprouter"
)

// unfollowUserHandler handles the HTTP request to unfollow a user.
func (rt *_router) unfollowUserHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    // Extract the ID of the user who wants to unfollow (followerID) from the URL parameters and convert it to an integer.
    followerID, err := strconv.Atoi(ps.ByName("userId"))
    if err != nil {
        // If the follower ID is invalid, return an HTTP 400 Bad Request error.
        http.Error(w, "Invalid follower ID", http.StatusBadRequest)
        return
    }

    // Extract the ID of the user to be unfollowed (followingID) from the URL parameters and convert it to an integer.
    followingID, err := strconv.Atoi(ps.ByName("followerId"))
    if err != nil {
        // If the following ID is invalid, return an HTTP 400 Bad Request error.
        http.Error(w, "Invalid following ID", http.StatusBadRequest)
        return
    }

    // Call the UnfollowUser method to remove the follower.
    err = rt.db.UnfollowUser(followerID, followingID)
    if err != nil {
        // If the UnfollowUser operation fails, log the error and return an HTTP 500 Internal Server Error.
        rt.baseLogger.WithError(err).Error("Failed to unfollow user")
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    // If the operation is successful, respond with a 204 No Content status.
    // This indicates that the action has been successfully processed, but there is no content to return.
    w.WriteHeader(http.StatusNoContent)
}
