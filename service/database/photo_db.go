package database

import (
    "fmt"
    "time"
)

// UploadPhoto adds a new photo to the database.
// It takes the user's ID who is uploading the photo, the upload date and time, and the photo data.
// Returns the new photo's ID and any error encountered.
func (db *appdbimpl) UploadPhoto(userID int, uploadDateTime time.Time, photoData []byte) (int, error) {
    // Prepare the SQL statement for inserting a new photo into the Photos table.
    // The statement includes the user ID, upload date/time, and the photo data.
    stmt, err := db.c.Prepare("INSERT INTO Photos (UserID, UploadDateTime, PhotoData) VALUES (?, ?, ?)")
    if err != nil {
        // If preparing the statement fails, return an error.
        return 0, fmt.Errorf("prepare upload photo statement: %w", err)
    }
    defer stmt.Close() // Ensure the statement is closed after the function execution.

    // Execute the prepared statement with the provided user ID, upload date/time, and photo data.
    res, err := stmt.Exec(userID, uploadDateTime, photoData)
    if err != nil {
        // If executing the statement fails, return an error.
        return 0, fmt.Errorf("execute upload photo statement: %w", err)
    }

    // Retrieve the ID of the newly inserted photo.
    photoID, err := res.LastInsertId()
    if err != nil {
        // If retrieving the last insert ID fails, return an error.
        return 0, fmt.Errorf("retrieve last insert ID: %w", err)
    }

    // Return the photo ID and nil (no error).
    return int(photoID), nil
}

// RemovePhoto deletes a photo from the database based on its ID.
// It takes the ID of the photo to be removed.
// Returns an error if the operation fails.
func (db *appdbimpl) RemovePhoto(photoID int) error {
    // Prepare the SQL statement for deleting a photo from the Photos table.
    // The statement includes a condition for the specific photo ID.
    stmt, err := db.c.Prepare("DELETE FROM Photos WHERE PhotoID = ?")
    if err != nil {
        // If preparing the statement fails, return an error.
        return fmt.Errorf("prepare remove photo statement: %w", err)
    }
    defer stmt.Close() // Ensure the statement is closed after the function execution.

    // Execute the prepared statement with the provided photo ID.
    _, err = stmt.Exec(photoID)
    if err != nil {
        // If executing the statement fails, return an error.
        return fmt.Errorf("execute remove photo statement: %w", err)
    }

    // Return nil if the operation is successful (no error).
    return nil
}
