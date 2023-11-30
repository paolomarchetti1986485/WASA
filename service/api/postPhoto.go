package api

import (
	"encoding/json"
	"net/http"
	
    "strconv"
    "time"
	"github.com/julienschmidt/httprouter"
)
func (rt *_router) uploadPhotoHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    userId, err := strconv.Atoi(ps.ByName("userId"))
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    // Assuming the photo data is sent as a base64 encoded string in the request body.
    // This needs to be adapted based on how the photo data is actually being sent.
    var photoData UnuploadedPhoto
    if err := json.NewDecoder(r.Body).Decode(&photoData); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Use the current time as the upload time. Adjust as necessary.
    photoID, err := rt.db.UploadPhoto(userId, time.Now())
    if err != nil {
        rt.baseLogger.WithError(err).Error("Failed to upload photo")
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]int{"photoId": photoID})
}
