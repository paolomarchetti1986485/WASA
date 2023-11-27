package api

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"WASA/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)
func (rt *_router) followUserHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

    err = rt.db.FollowUser(followerID, followingID)
    if err != nil {
        rt.baseLogger.WithError(err).Error("Failed to follow user")
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    response := map[string]string{"message": "You are now following this user."}
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(response)
}
