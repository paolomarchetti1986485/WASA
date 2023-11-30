package database
import (
	"fmt"
	
)
func (db *appdbimpl) BanUser(bannedID, userID int) error {
    // Prepare SQL statement for inserting a new ban record
    stmt, err := db.c.Prepare("INSERT INTO BannedUsers (BannedUserID, UserID, BanDateTime) VALUES (?, ?, CURRENT_TIMESTAMP)")
    if err != nil {
        return fmt.Errorf("prepare ban user statement: %w", err)
    }
    defer stmt.Close()

    // Execute the statement
    _, err = stmt.Exec(bannedID, userID)
    if err != nil {
        return fmt.Errorf("execute ban user statement: %w", err)
    }

    return nil
}
func (db *appdbimpl) UnbanUser(bannedID, userID int) error {
    // Prepare SQL statement for deleting a ban record
    stmt, err := db.c.Prepare("DELETE FROM BannedUsers WHERE BannedUserID = ? AND UserID = ?")
    if err != nil {
        return fmt.Errorf("prepare unban user statement: %w", err)
    }
    defer stmt.Close()

    // Execute the statement
    _, err = stmt.Exec(bannedID, userID)
    if err != nil {
        return fmt.Errorf("execute unban user statement: %w", err)
    }

    return nil
}
