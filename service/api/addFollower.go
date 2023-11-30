package api

import (
	"fmt"
	"net/http"
	"strconv"
    "log"
	"github.com/julienschmidt/httprouter"
)

// addFollower handles the request to follow a user.
func (rt *_router) followUserHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    log.Println("followUserHandler called")
    // Extract followerID and followingID from the request parameters
    followerIDStr := ps.ByName("userId")
    followingIDStr := ps.ByName("followerId")

    // Convert IDs from string to int
    followerID, err := strconv.Atoi(followerIDStr)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "Invalid follower ID")
        return
    }

    followingID, err := strconv.Atoi(followingIDStr)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "Invalid following ID")
        return
    }

    // Users can't follow themselves
    if followerID == followingID {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "Users can't follow themselves")
        return
    }

    // Add the follower using the database function
    err = rt.db.FollowUser(followerID, followingID)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "Error adding follower: %s", err)
        return
    }

    // If successful, respond with a 204 No Content status
    w.WriteHeader(http.StatusNoContent)
}
