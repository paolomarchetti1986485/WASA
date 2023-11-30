package database
import (
	"fmt"
)

func (db *appdbimpl) FollowUser(follower int, followed int) error {

	_, err := db.c.Exec("INSERT INTO followers (FollowerID,FollowingID) VALUES (?, ?)",
		follower, followed)
	if err != nil {
		return err
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
