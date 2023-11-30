package api

import (
	"net/http"
	
    "strconv"

	"github.com/julienschmidt/httprouter"
)
func (rt *_router) unfollowUserHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    followerID, err := strconv.Atoi(ps.ByName("userId"))
    if err != nil {
        http.Error(w, "Invalid follower ID", http.StatusBadRequest)
        return
    }

    followingID, err := strconv.Atoi(ps.ByName("followerId"))
    if err != nil {
        http.Error(w, "Invalid following ID", http.StatusBadRequest)
        return
    }

    err = rt.db.UnfollowUser(followerID, followingID)
    if err != nil {
        rt.baseLogger.WithError(err).Error("Failed to unfollow user")
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
