package api

import (
	"encoding/json"
	"net/http"
	
    "strconv"

	"github.com/julienschmidt/httprouter"
)
func (rt *_router) addLikeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    userId, err := strconv.Atoi(ps.ByName("userId"))
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    photoId, err := strconv.Atoi(ps.ByName("photoId"))
    if err != nil {
        http.Error(w, "Invalid photo ID", http.StatusBadRequest)
        return
    }

    err = rt.db.AddLike(photoId, userId)
    if err != nil {
        rt.baseLogger.WithError(err).Error("Failed to add like")
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    like := LikeId(userId) // Assuming LikeId is a type alias for UserId
    response := map[string]LikeId{"likeId": like}
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(response)
}
