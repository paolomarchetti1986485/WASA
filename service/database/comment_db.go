package database

import (
	"database/sql"
	"fmt"
)

// AddComment adds a new comment to a photo in the database.
// It takes the photo's ID, the user's ID who is commenting, and the comment text.
// Returns the ID of the newly added comment and any error encountered.
func (db *appdbimpl) AddComment(photoID, userID int, commentText string) (int, error) {
	// Prepare the SQL statement for inserting a new comment into the Comments table.
	// The statement includes the photo ID, user ID, and comment text.
	stmt, err := db.c.Prepare("INSERT INTO Comments (PhotoID, UserID, CommentText) VALUES (?, ?, ?)")
	if err != nil {
		// If preparing the statement fails, return an error.
		return 0, fmt.Errorf("prepare add comment statement: %w", err)
	}
	defer stmt.Close() // Ensure the statement is closed after the function execution.

	// Execute the prepared statement with the provided photoID, userID, and commentText.
	res, err := stmt.Exec(photoID, userID, commentText)
	if err != nil {
		// If executing the statement fails, return an error.
		return 0, fmt.Errorf("execute add comment statement: %w", err)
	}

	// Retrieve the ID of the last inserted comment.
	commentID, err := res.LastInsertId()
	if err != nil {
		// If retrieving the last insert ID fails, return an error.
		return 0, fmt.Errorf("retrieve last insert ID: %w", err)
	}

	// Return the comment ID and nil (no error).
	return int(commentID), nil
}

// RemoveComment removes a comment from the database based on its ID.
// It takes the ID of the comment to be removed.
// Returns an error if the operation fails.
func (db *appdbimpl) RemoveComment(commentID int) error {
	// Prepare the SQL statement for deleting a comment from the Comments table.
	// The statement includes a condition for the specific comment ID.
	stmt, err := db.c.Prepare("DELETE FROM Comments WHERE CommentID = ?")
	if err != nil {
		// If preparing the statement fails, return an error.
		return fmt.Errorf("prepare remove comment statement: %w", err)
	}
	defer stmt.Close() // Ensure the statement is closed after the function execution.

	// Execute the prepared statement with the provided commentID.
	_, err = stmt.Exec(commentID)
	if err != nil {
		// If executing the statement fails, return an error.
		return fmt.Errorf("execute remove comment statement: %w", err)
	}

	// Return nil (no error) if the operation is successful.
	return nil
}

// GetCommentById retrieves a comment by its ID.
func (db *appdbimpl) GetCommentById(commentID int) (Comment, error) {
	var comment Comment
	query := "SELECT CommentID, PhotoID, UserID, CommentText, Timestamp FROM Comments WHERE CommentID = ?"
	err := db.c.QueryRow(query, commentID).Scan(&comment.CommentID, &comment.PhotoID, &comment.UserID, &comment.CommentText, &comment.Timestamp)
	if err != nil {
		if err == sql.ErrNoRows {
			return comment, fmt.Errorf("comment not found")
		}
		return comment, fmt.Errorf("error retrieving comment: %w", err)
	}
	return comment, nil
}
