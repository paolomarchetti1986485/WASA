package database

func (db *appdbimpl) AddUser(username string) (int, error) {
    stmt, err := db.c.Prepare("INSERT INTO Users (Username) VALUES (?)")
    if err != nil {
        return 0, fmt.Errorf("prepare add user statement: %w", err)
    }
    defer stmt.Close()

    res, err := stmt.Exec(username)
    if err != nil {
        return 0, fmt.Errorf("execute add user statement: %w", err)
    }

    userID, err := res.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("retrieve last insert ID: %w", err)
    }

    return int(userID), nil
}
func (db *appdbimpl) UpdateUsername(userID int, newUsername string) error {
    stmt, err := db.c.Prepare("UPDATE Users SET Username = ? WHERE UserID = ?")
    if err != nil {
        return fmt.Errorf("prepare update username statement: %w", err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(newUsername, userID)
    if err != nil {
        return fmt.Errorf("execute update username statement: %w", err)
    }

    return nil
}
