package api

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"WASA/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// getHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) loginHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        rt.baseLogger.WithError(err).Warning("Wrong JSON received")
        w.WriteHeader(http.StatusBadRequest)
        return
    } else if !IsValidUsernameIdentifier(user.IdUser) {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    // Check if user exists (assuming you have a function for this)
    userID, err := rt.db.GetUserByUsername(user.Username)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            // User does not exist, create a new user
            userID, err = rt.db.AddUser(user.Username)
            if err != nil {
                rt.baseLogger.WithError(err).Error("Failed to create new user")
                w.WriteHeader(http.StatusInternalServerError)
                return
            }
        } else {
            rt.baseLogger.WithError(err).Error("Database error")
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
    }

    // At this point, the user either existed or has been created
    // Generate a response (you might want to return a token or user info)
    response := map[string]int{"userId": userID}
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}
