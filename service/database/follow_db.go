package database

func (db *appdbimpl) FollowUser(followerID, followingID int) error {
    stmt, err := db.c.Prepare("INSERT INTO Followers (FollowerID, FollowingID) VALUES (?, ?)")
    if err != nil {
        return fmt.Errorf("prepare follow user statement: %w", err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(followerID, followingID)
    if err != nil {
        return fmt.Errorf("execute follow user statement: %w", err)
    }

    return nil
}
func (db *appdbimpl) UnfollowUser(followerID, followingID int) error {
    stmt, err := db.c.Prepare("DELETE FROM Followers WHERE FollowerID = ? AND FollowingID = ?")
    if err != nil {
        return fmt.Errorf("prepare unfollow user statement: %w", err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(followerID, followingID)
    if err != nil {
        return fmt.Errorf("execute unfollow user statement: %w", err)
    }

    return nil
}
