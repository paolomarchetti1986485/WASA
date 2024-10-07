package api

import (
	"WASA/service/api/reqcontext"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) isUserBannedHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userIdStr := ps.ByName("userId")
	bannedIdStr := ps.ByName("bannedId")

	log.Printf("Received userId: %s, bannedId: %s", userIdStr, bannedIdStr)

	// Verifica che i parametri non siano vuoti
	if userIdStr == "" || bannedIdStr == "" {
		log.Println("Missing userId or bannedId")
		http.Error(w, "Missing user ID or banned ID", http.StatusBadRequest)
		return
	}

	// Converti i parametri in interi
	userId, err := strconv.Atoi(userIdStr)
	if err != nil || userId <= 0 {
		log.Printf("Invalid userId: %s", userIdStr) // Log dell'errore
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	bannedId, err := strconv.Atoi(bannedIdStr)
	if err != nil || bannedId <= 0 {
		log.Printf("Invalid bannedId: %s", bannedIdStr) // Log dell'errore
		http.Error(w, "Invalid banned ID", http.StatusBadRequest)
		return
	}

	// Aggiungi un log per confermare i parametri corretti
	log.Printf("Checking ban status for userId: %d and bannedId: %d", userId, bannedId)

	// Controlla se l'utente è stato bannato
	isBanned, err := rt.db.IsUserBanned(userId, bannedId)
	if err != nil {
		log.Printf("Error checking ban status: %s", err)
		http.Error(w, "Failed to check ban status", http.StatusInternalServerError)
		return
	}

	// Restituisci il risultato in formato JSON
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(map[string]bool{"isBanned": isBanned})
	if err != nil {
		rt.baseLogger.WithError(err).Error("Failed to send response")
	}
}
