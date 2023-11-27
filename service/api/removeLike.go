package api

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"WASA/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)
func (rt *_router) removeLikeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

    err = rt.db.RemoveLike(photoId, userId)
    if err != nil {
        rt.baseLogger.WithError(err).Error("Failed to remove like")
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
