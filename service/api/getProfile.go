package api

import (
	"encoding/json"
	"net/http"
	
    "strconv"
	"github.com/julienschmidt/httprouter"
    "WASA/service/database"
)

func (rt *_router) getProfileHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    var followers []database.User
    var following []database.User
    var photos []database.Photo
    userId, err := strconv.Atoi(ps.ByName("userId"))
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    
    photos, err = rt.db.GetUserPhotos(userId)
    if err != nil {
        rt.baseLogger.WithError(err).Error("Failed to get user photos")
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    followers, err = rt.db.GetUserFollowers(userId)
    if err != nil {
        rt.baseLogger.WithError(err).Error("Failed to get user followers")
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    following, err = rt.db.GetUserFollowing(userId)
    if err != nil {
        rt.baseLogger.WithError(err).Error("Failed to get user following")
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    profile := Profile{
        Photos:    photos,
        Followers: followers,
        Following: following,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(profile)
}

