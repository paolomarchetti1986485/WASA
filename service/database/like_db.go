package database
import (
	"fmt"
)
func (db *appdbimpl) AddLike(photoID, userID int) error {
    stmt, err := db.c.Prepare("INSERT INTO Likes (PhotoID, UserID) VALUES (?, ?)")
    if err != nil {
        return fmt.Errorf("prepare add like statement: %w", err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(photoID, userID)
    if err != nil {
        return fmt.Errorf("execute add like statement: %w", err)
    }

    return nil
}
func (db *appdbimpl) RemoveLike(photoID, userID int) error {
    stmt, err := db.c.Prepare("DELETE FROM Likes WHERE PhotoID = ? AND UserID = ?")
    if err != nil {
        return fmt.Errorf("prepare remove like statement: %w", err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(photoID, userID)
    if err != nil {
        return fmt.Errorf("execute remove like statement: %w", err)
    }

    return nil
}
