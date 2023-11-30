package api

import (
	
	"net/http"
    "strconv"

	"github.com/julienschmidt/httprouter"
)
func (rt *_router) deletePhotoHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    photoId, err := strconv.Atoi(ps.ByName("photoId"))
    if err != nil {
        http.Error(w, "Invalid photo ID", http.StatusBadRequest)
        return
    }

    err = rt.db.RemovePhoto(photoId)
    if err != nil {
        rt.baseLogger.WithError(err).Error("Failed to delete photo")
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
