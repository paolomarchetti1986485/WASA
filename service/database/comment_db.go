package database

func (db *appdbimpl) AddComment(photoID, userID int, commentText string) (int, error) {
    stmt, err := db.c.Prepare("INSERT INTO Comments (PhotoID, UserID, CommentText) VALUES (?, ?, ?)")
    if err != nil {
        return 0, fmt.Errorf("prepare add comment statement: %w", err)
    }
    defer stmt.Close()

    res, err := stmt.Exec(photoID, userID, commentText)
    if err != nil {
        return 0, fmt.Errorf("execute add comment statement: %w", err)
    }

    commentID, err := res.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("retrieve last insert ID: %w", err)
    }

    return int(commentID), nil
}
func (db *appdbimpl) RemoveComment(commentID int) error {
    stmt, err := db.c.Prepare("DELETE FROM Comments WHERE CommentID = ?")
    if err != nil {
        return fmt.Errorf("prepare remove comment statement: %w", err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(commentID)
    if err != nil {
        return fmt.Errorf("execute remove comment statement: %w", err)
    }

    return nil
}
