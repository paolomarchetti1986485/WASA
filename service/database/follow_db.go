package database

import (
	"fmt"
)

// FollowUser creates a record in the database to represent a user following another user.
// It takes the IDs of the follower and the followed user.
// Returns an error if the operation fails.
func (db *appdbimpl) FollowUser(follower int, followed int) error {
	// Execute the SQL statement to insert a new record into the 'followers' table.
	// The statement includes the follower's ID and the followed user's ID.
	_, err := db.c.Exec("INSERT INTO followers (FollowerID,FollowingID) VALUES (?, ?)",
		follower, followed)
	if err != nil {
		// If executing the statement fails, return the error.
		return err
	}

	// Return nil if the operation is successful (no error).
	return nil
}

// UnfollowUser removes a record from the database that represents a user following another user.
// It takes the IDs of the follower and the followed user.
// Returns an error if the operation fails.
func (db *appdbimpl) UnfollowUser(followerID, followingID int) error {
	// Prepare the SQL statement for deleting a record from the 'Followers' table.
	// The statement includes conditions for both the follower's ID and the followed user's ID.
	stmt, err := db.c.Prepare("DELETE FROM Followers WHERE FollowerID = ? AND FollowingID = ?")
	if err != nil {
		// If preparing the statement fails, return an error.
		return fmt.Errorf("prepare unfollow user statement: %w", err)
	}
	defer stmt.Close() // Ensure the statement is closed after the function execution.

	// Execute the prepared statement with the provided followerID and followingID.
	_, err = stmt.Exec(followerID, followingID)
	if err != nil {
		// If executing the statement fails, return an error.
		return fmt.Errorf("execute unfollow user statement: %w", err)
	}

	// Return nil if the operation is successful (no error).
	return nil
}
