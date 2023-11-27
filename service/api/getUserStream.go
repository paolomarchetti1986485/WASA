package api

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"WASA/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getStreamHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    userId, err := strconv.Atoi(ps.ByName("userId"))
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    stream, err := rt.db.GetUserStream(userId)
    if err != nil {
        rt.baseLogger.WithError(err).Error("Failed to get user stream")
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(stream)
}

