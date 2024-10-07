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
	// Start a transaction to ensure all deletes are performed atomically.
	tx, err := db.c.Begin()
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	// Funzione di rollback con gestione dell'errore
	rollback := func() {
		if rbErr := tx.Rollback(); rbErr != nil {
			fmt.Printf("rollback error: %v\n", rbErr)
		}
	}
	// Step 1: Delete all likes associated with the photo.
	_, err = tx.Exec("DELETE FROM Likes WHERE PhotoID = ?", photoID)
	if err != nil {
		rollback()
		return fmt.Errorf("delete likes: %w", err)
	}

	// Step 2: Delete all comments associated with the photo.
	_, err = tx.Exec("DELETE FROM Comments WHERE PhotoID = ?", photoID)
	if err != nil {
		rollback()
		return fmt.Errorf("delete comments: %w", err)
	}

	// Step 3: Delete the photo itself.
	_, err = tx.Exec("DELETE FROM Photos WHERE PhotoID = ?", photoID)
	if err != nil {
		rollback()
		return fmt.Errorf("delete photo: %w", err)
	}

	// Commit the transaction if all deletes were successful.
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	// Return nil if the operation is successful (no error).
	return nil
}

func (db *appdbimpl) GetPhotoData(photoID int) ([]byte, error) {
	var imageData []byte
	err := db.c.QueryRow("SELECT PhotoData FROM Photos WHERE PhotoID = ?", photoID).Scan(&imageData)
	if err != nil {
		return nil, fmt.Errorf("error querying photo data: %w", err)
	}
	return imageData, nil
}
