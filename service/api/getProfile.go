package api

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"WASA/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getProfileHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    userId, err := strconv.Atoi(ps.ByName("userId"))
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    user, err := rt.db.GetUser(userId)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            http.Error(w, "User not found or does not exist", http.StatusNotFound)
        } else {
            rt.baseLogger.WithError(err).Error("Failed to get user")
            w.WriteHeader(http.StatusInternalServerError)
        }
        return
    }

    photos, err := rt.db.GetUserPhotos(userId)
    if err != nil {
        rt.baseLogger.WithError(err).Error("Failed to get user photos")
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    followers, err := rt.db.GetUserFollowers(userId)
    if err != nil {
        rt.baseLogger.WithError(err).Error("Failed to get user followers")
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    following, err := rt.db.GetUserFollowing(userId)
    if err != nil {
        rt.baseLogger.WithError(err).Error("Failed to get user following")
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    profile := Profile{
        User:      user,
        Photos:    photos,
        Followers: followers,
        Following: following,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(profile)
}

