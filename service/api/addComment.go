package api

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"WASA/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)
func (rt *_router) addCommentHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

    var comment Comment
    if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    commentID, err := rt.db.AddComment(photoId, userId, comment.Text)
    if err != nil {
        rt.baseLogger.WithError(err).Error("Failed to add comment")
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]int{"commentId": commentID})
}
