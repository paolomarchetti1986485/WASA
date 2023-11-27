package api

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"WASA/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)
func (rt *_router) unbanUserHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    userId, err := strconv.Atoi(ps.ByName("userId"))
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    bannedId, err := strconv.Atoi(ps.ByName("bannedId"))
    if err != nil {
        http.Error(w, "Invalid banned ID", http.StatusBadRequest)
        return
    }

    err = rt.db.UnbanUser(bannedId, userId)
    if err != nil {
        rt.baseLogger.WithError(err).Error("Failed to unban user")
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
