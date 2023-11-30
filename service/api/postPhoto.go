package api

import (
	"encoding/json"
	"net/http"
	"time"
    "strconv"
	"github.com/julienschmidt/httprouter"
    "io/ioutil"      
)

func (rt *_router) uploadPhotoHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    // Extract and validate userId from URL parameters
    userId, err := strconv.Atoi(ps.ByName("userId"))
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    // Parse the multipart form in the request
    err = r.ParseMultipartForm(10 << 20) // Limit upload size (e.g., 10 MB)
    if err != nil {
        http.Error(w, "Could not parse multipart form", http.StatusBadRequest)
        return
    }

    // Get the file from the form data
    file, _, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Error retrieving the file", http.StatusBadRequest)
        return
    }
    defer file.Close()

    // Read the file content
    fileContent, err := ioutil.ReadAll(file)
    if err != nil {
        http.Error(w, "Error reading the file", http.StatusInternalServerError)
        return
    }

    // Pass the file content to the UploadPhoto function
    photoID, err := rt.db.UploadPhoto(userId, time.Now(), fileContent) // fileContent is []byte
    if err != nil {
        rt.baseLogger.WithError(err).Error("Failed to upload photo")
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    // Send response back to client
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]int{"photoId": photoID})
}
