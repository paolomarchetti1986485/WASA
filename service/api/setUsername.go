package api

import (
	"encoding/json"
	"net/http"
	
    "strconv"

	"github.com/julienschmidt/httprouter"
)
// UsernameUpdateRequest represents the request structure to update a username.
type UsernameUpdateRequest struct {
    Username string `json:"username"`
}

// UsernameUpdateResponse represents the response structure after updating a username.
type UsernameUpdateResponse struct {
    Username string `json:"username"`
    Message  string   `json:"message"`
}

func (rt *_router) setUsernameHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    userId, err := strconv.Atoi(ps.ByName("userId"))
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    var req UsernameUpdateRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Validate the new username
    if !IsValidUsername(req.Username) {
        http.Error(w, "Invalid username", http.StatusBadRequest)
        return
    }

    // Assuming UpdateUsername is your existing database function
    err = rt.db.UpdateUsername(userId, string(req.Username))
    if err != nil {
        rt.baseLogger.WithError(err).Error("Failed to update username")
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    resp := UsernameUpdateResponse{
        Username: req.Username,
        Message:  "Username has been updated successfully",
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}
