package database

import (
	"fmt"
)

// AddLike records a 'like' by a user on a photo in the database.
// It takes the photo's ID and the user's ID who is liking the photo.
// Returns an error if the operation fails.
func (db *appdbimpl) AddLike(photoID, userID int) error {
	// Prepare the SQL statement for inserting a new like into the Likes table.
	// The statement includes the photo ID and the user ID.
	stmt, err := db.c.Prepare("INSERT INTO Likes (PhotoID, UserID) VALUES (?, ?)")
	if err != nil {
		// If preparing the statement fails, return an error.
		return fmt.Errorf("prepare add like statement: %w", err)
	}
	defer stmt.Close() // Ensure the statement is closed after the function execution.

	// Execute the prepared statement with the provided photo ID and user ID.
	_, err = stmt.Exec(photoID, userID)
	if err != nil {
		// If executing the statement fails, return an error.
		return fmt.Errorf("execute add like statement: %w", err)
	}

	// Return nil if the operation is successful (no error).
	return nil
}

// RemoveLike removes a 'like' from a photo in the database.
// It takes the photo's ID and the user's ID who had liked the photo.
// Returns an error if the operation fails.
func (db *appdbimpl) RemoveLike(photoID, userID int) error {
	// Prepare the SQL statement for deleting a like from the Likes table.
	// The statement includes conditions for both the photo ID and the user ID.
	stmt, err := db.c.Prepare("DELETE FROM Likes WHERE PhotoID = ? AND UserID = ?")
	if err != nil {
		// If preparing the statement fails, return an error.
		return fmt.Errorf("prepare remove like statement: %w", err)
	}
	defer stmt.Close() // Ensure the statement is closed after the function execution.

	// Execute the prepared statement with the provided photo ID and user ID.
	_, err = stmt.Exec(photoID, userID)
	if err != nil {
		// If executing the statement fails, return an error.
		return fmt.Errorf("execute remove like statement: %w", err)
	}

	// Return nil if the operation is successful (no error).
	return nil
}
