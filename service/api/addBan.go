package api

import (
	"encoding/json"
	"net/http"
	
    "strconv"
	"github.com/julienschmidt/httprouter"
)
func (rt *_router) banUserHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

    err = rt.db.BanUser(bannedId, userId)
    if err != nil {
        rt.baseLogger.WithError(err).Error("Failed to ban user")
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    response := map[string]string{"message": "User has been successfully banned."}
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(response)
}
