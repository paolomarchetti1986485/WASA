package database

import (
    "fmt"
)

// GetUserStream retrieves a stream of photos for a user based on their followings.
// It takes the user's ID and returns a slice of photos and any error encountered.
func (db *appdbimpl) GetUserStream(userId int) ([]Photo, error) {
    var photos []Photo

    // SQL query to retrieve photos from users that the given user is following.
    // The photos are ordered by upload date and time in descending order.
    query := `
        SELECT p.PhotoID, p.UserID, p.UploadDateTime
        FROM Photos p
        INNER JOIN Followers f ON p.UserID = f.FollowingID
        WHERE f.FollowerID = ?
        ORDER BY p.UploadDateTime DESC`
    rows, err := db.c.Query(query, userId)
    if err != nil {
        // If querying the database fails, return an error.
        return nil, fmt.Errorf("error querying user stream: %w", err)
    }
    defer rows.Close() // Ensure the rows are closed after the function execution.

    // Iterating over the query results.
    for rows.Next() {
        var photo Photo
        // Scanning the row into the photo struct.
        if err := rows.Scan(&photo.PhotoID, &photo.UserID, &photo.UploadDateTime); err != nil {
            // If scanning the row fails, return an error.
            return nil, fmt.Errorf("error scanning photo: %w", err)
        }

        // Fetching comments for each photo.
        photo.Comments, err = db.GetPhotoComments(photo.PhotoID)
        if err != nil {
            // If fetching comments fails, return an error.
            return nil, fmt.Errorf("error fetching photo comments: %w", err)
        }

        // Fetching likes for each photo.
        photo.Likes, err = db.GetPhotoLikes(photo.PhotoID)
        if err != nil {
            // If fetching likes fails, return an error.
            return nil, fmt.Errorf("error fetching photo likes: %w", err)
        }

        // Appending the photo to the slice of photos.
        photos = append(photos, photo)
    }

    // Checking for any errors that occurred during the iteration.
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error in rows: %w", err)
    }

    // Returning the slice of photos and nil (no error).
    return photos, nil
}
