package database
import (
	"fmt"
    "time"
)


func (db *appdbimpl) UploadPhoto(userID int, uploadDateTime time.Time, photoData []byte) (int, error) {
    stmt, err := db.c.Prepare("INSERT INTO Photos (UserID, UploadDateTime, PhotoData) VALUES (?, ?, ?)")
    if err != nil {
        return 0, fmt.Errorf("prepare upload photo statement: %w", err)
    }
    defer stmt.Close()

    res, err := stmt.Exec(userID, uploadDateTime, photoData)
    if err != nil {
        return 0, fmt.Errorf("execute upload photo statement: %w", err)
    }

    photoID, err := res.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("retrieve last insert ID: %w", err)
    }

    return int(photoID), nil
}


func (db *appdbimpl) RemovePhoto(photoID int) error {
    stmt, err := db.c.Prepare("DELETE FROM Photos WHERE PhotoID = ?")
    if err != nil {
        return fmt.Errorf("prepare remove photo statement: %w", err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(photoID)
    if err != nil {
        return fmt.Errorf("execute remove photo statement: %w", err)
    }

    return nil
}
