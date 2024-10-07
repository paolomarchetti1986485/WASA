package api

import (
	"WASA/service/api/reqcontext"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) searchUserHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Estrazione del parametro 'username' dalla query string
	username := r.URL.Query().Get("username")

	// Se il parametro username è vuoto, restituisce un errore 400
	if username == "" {
		http.Error(w, "Missing username parameter", http.StatusBadRequest)
		return
	}

	// Logica per cercare gli utenti basata sul prefisso del nome utente
	users, err := rt.db.SearchUsersByUsernamePrefix(username)
	if err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}

	// Se non ci sono utenti, restituisci un 404
	if len(users) == 0 {
		http.Error(w, "No users found", http.StatusNotFound)
		return
	}

	// Restituisci la lista degli utenti trovati
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Failed to send response")
	}
}
