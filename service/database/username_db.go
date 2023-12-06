package database

import (
    "fmt"
)

// AddUser adds a new user to the database with the provided username.
// Returns the new user's ID and any error encountered.
func (db *appdbimpl) AddUser(username string) (int, error) {
    // Prepare the SQL statement for inserting a new user into the Users table.
    // The statement includes the username.
    stmt, err := db.c.Prepare("INSERT INTO Users (Username) VALUES (?)")
    if err != nil {
        // If preparing the statement fails, return an error.
        return 0, fmt.Errorf("prepare add user statement: %w", err)
    }
    defer stmt.Close() // Ensure the statement is closed after the function execution.

    // Execute the prepared statement with the provided username.
    res, err := stmt.Exec(username)
    if err != nil {
        // If executing the statement fails, return an error.
        return 0, fmt.Errorf("execute add user statement: %w", err)
    }

    // Retrieve the ID of the newly inserted user.
    userID, err := res.LastInsertId()
    if err != nil {
        // If retrieving the last insert ID fails, return an error.
        return 0, fmt.Errorf("retrieve last insert ID: %w", err)
    }

    // Return the user ID and nil (no error).
    return int(userID), nil
}

// UpdateUsername updates the username of an existing user in the database.
// Takes the user's ID and the new username.
// Returns an error if the operation fails.
func (db *appdbimpl) UpdateUsername(userID int, newUsername string) error {
    // Prepare the SQL statement for updating a user's username in the Users table.
    // The statement includes the new username and the user's ID as conditions.
    stmt, err := db.c.Prepare("UPDATE Users SET Username = ? WHERE UserID = ?")
    if err != nil {
        // If preparing the statement fails, return an error.
        return fmt.Errorf("prepare update username statement: %w", err)
    }
    defer stmt.Close() // Ensure the statement is closed after the function execution.

    // Execute the prepared statement with the new username and the user's ID.
    _, err = stmt.Exec(newUsername, userID)
    if err != nil {
        // If executing the statement fails, return an error.
        return fmt.Errorf("execute update username statement: %w", err)
    }

    // Return nil if the operation is successful (no error).
    return nil
}
// DeleteUserByID deletes a user from the database based on their ID.
func (db *appdbimpl) DeleteUserByID(userID int) error {
    // SQL query to delete the user.
    _, err := db.c.Exec("DELETE FROM Users WHERE UserID = ?", userID)
    if err != nil {
        return fmt.Errorf("error deleting user: %w", err)
    }

    return nil
}