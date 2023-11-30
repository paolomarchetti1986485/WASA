package database
import (
	"fmt"
)
func (db *appdbimpl) GetUserStream(userId int) ([]Photo, error) {
    var photos []Photo

    query := `
        SELECT p.PhotoID, p.UserID, p.UploadDateTime
        FROM Photos p
        INNER JOIN Followers f ON p.UserID = f.FollowingID
        WHERE f.FollowerID = ?
        ORDER BY p.UploadDateTime DESC
        LIMIT 1000`
    rows, err := db.c.Query(query, userId)
    if err != nil {
        return nil, fmt.Errorf("error querying user stream: %w", err)
    }
    defer rows.Close()

    for rows.Next() {
        var photo Photo
        if err := rows.Scan(&photo.PhotoID, &photo.UserID, &photo.UploadDateTime); err != nil {
            return nil, fmt.Errorf("error scanning photo: %w", err)
        }
        photos = append(photos, photo)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error in rows: %w", err)
    }

    return photos, nil
}
