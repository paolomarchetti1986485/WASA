package api

import (
    "fmt"
    "net/http"
    "strconv"
    "log"
    "github.com/julienschmidt/httprouter"
)

// followUserHandler handles the HTTP request to add a follower to a user.
// It extracts the follower and following user IDs from the request parameters and performs the follow operation.
func (rt *_router) followUserHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    log.Println("followUserHandler called")
    // Extract the ID of the user who wants to follow (followerID) and the ID of the user to be followed (followingID) from the URL parameters.
    followerIDStr := ps.ByName("userId")
    followingIDStr := ps.ByName("followerId")

    // Convert the followerID from string to int.
    followerID, err := strconv.Atoi(followerIDStr)
    if err != nil {
        // If the conversion fails, return an HTTP 400 Bad Request error.
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "Invalid follower ID")
        return
    }

    // Convert the followingID from string to int.
    followingID, err := strconv.Atoi(followingIDStr)
    if err != nil {
        // If the conversion fails, return an HTTP 400 Bad Request error.
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "Invalid following ID")
        return
    }

    // Check if the followerID and followingID are the same. Users cannot follow themselves.
    if followerID == followingID {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "Users can't follow themselves")
        return
    }

    // Call the FollowUser method to add the follower.
    err = rt.db.FollowUser(followerID, followingID)
    if err != nil {
        // If the FollowUser operation fails, return an HTTP 500 Internal Server Error.
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "Error adding follower: %s", err)
        return
    }

    // If the operation is successful, respond with a 204 No Content status.
    w.WriteHeader(http.StatusNoContent)
}
