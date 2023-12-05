package database

import (
    "fmt"
)

// BanUser adds a ban record for a user in the database.
// It takes the IDs of the banned user and the user who is performing the ban.
// Returns an error if the operation fails.
func (db *appdbimpl) BanUser(bannedID, userID int) error {
    // Prepare the SQL statement for inserting a new ban record.
    // The statement includes the banned user's ID, the user's ID who is issuing the ban,
    // and the current timestamp for when the ban is issued.
    stmt, err := db.c.Prepare("INSERT INTO Banned (BannedUserID, UserID, BanDateTime) VALUES (?, ?, CURRENT_TIMESTAMP)")
    if err != nil {
        // If preparing the statement fails, return an error.
        return fmt.Errorf("prepare ban user statement: %w", err)
    }
    defer stmt.Close() // Ensure the statement is closed after the function execution.

    // Execute the prepared statement with the bannedID and userID.
    _, err = stmt.Exec(bannedID, userID)
    if err != nil {
        // If executing the statement fails, return an error.
        return fmt.Errorf("execute ban user statement: %w", err)
    }

    return nil // Return nil if the operation is successful.
}

// UnbanUser removes a ban record for a user in the database.
// It takes the IDs of the banned user and the user who is performing the unban operation.
// Returns an error if the operation fails.
func (db *appdbimpl) UnbanUser(bannedID, userID int) error {
    // Prepare the SQL statement for deleting a ban record.
    // The statement includes conditions for both the banned user's ID and the user's ID who issued the ban.
    stmt, err := db.c.Prepare("DELETE FROM Banned WHERE BannedUserID = ? AND UserID = ?")
    if err != nil {
        // If preparing the statement fails, return an error.
        return fmt.Errorf("prepare unban user statement: %w", err)
    }
    defer stmt.Close() // Ensure the statement is closed after the function execution.

    // Execute the prepared statement with the bannedID and userID.
    _, err = stmt.Exec(bannedID, userID)
    if err != nil {
        // If executing the statement fails, return an error.
        return fmt.Errorf("execute unban user statement: %w", err)
    }

    return nil // Return nil if the operation is successful.
}
