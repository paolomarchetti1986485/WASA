package database
import (
	"database/sql"
	"fmt"
	
)
func (db *appdbimpl) GetUserByUsername(username string) (int, error) {
    var userID int
    err := db.c.QueryRow("SELECT UserID FROM Users WHERE Username = ?", username).Scan(&userID)
    if err != nil {
        if err == sql.ErrNoRows {
            // User not found
            return 0, err
        }
        // Other error
        return 0, fmt.Errorf("error querying user by username: %w", err)
    }
    return userID, nil
}
func (db *appdbimpl) GetPhotoComments(photoID int) ([]Comment, error) {
    var comments []Comment

    query := "SELECT CommentID, UserID, CommentText, Timestamp FROM Comments WHERE PhotoID = ?"
    rows, err := db.c.Query(query, photoID)
    if err != nil {
        return nil, fmt.Errorf("error querying comments: %w", err)
    }
    defer rows.Close()

    for rows.Next() {
        var comment Comment
        comment.PhotoID = photoID // Set the PhotoID for each comment
        if err := rows.Scan(&comment.CommentID, &comment.UserID, &comment.CommentText, &comment.Timestamp); err != nil {
            return nil, fmt.Errorf("error scanning comment: %w", err)
        }
        comments = append(comments, comment)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error in rows: %w", err)
    }

    return comments, nil
}


func (db *appdbimpl) GetUserPhotos(userId int) ([]Photo, error) {
    var photos []Photo

    rows, err := db.c.Query("SELECT PhotoID, UploadDateTime FROM Photos WHERE UserID = ?", userId)
    if err != nil {
        return nil, fmt.Errorf("error querying photos: %w", err)
    }
    defer rows.Close()

    for rows.Next() {
        var photo Photo
        if err := rows.Scan(&photo.PhotoID, &photo.UploadDateTime); err != nil {
            return nil, fmt.Errorf("error scanning photo: %w", err)
        }

        // Fetch comments for each photo
        photo.Comments, err = db.GetPhotoComments(photo.PhotoID)
        if err != nil {
            return nil, fmt.Errorf("error fetching photo comments: %w", err)
        }

        // Fetch likes for each photo
        photo.Likes, err = db.GetPhotoLikes(photo.PhotoID)
        if err != nil {
            return nil, fmt.Errorf("error fetching photo likes: %w", err)
        }

        photos = append(photos, photo)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error in rows: %w", err)
    }

    return photos, nil
}

func (db *appdbimpl) GetPhotoLikes(photoID int) ([]Like, error) {
    var likes []Like

    query := "SELECT UserID, Timestamp FROM Likes WHERE PhotoID = ?"
    rows, err := db.c.Query(query, photoID)
    if err != nil {
        return nil, fmt.Errorf("error querying likes: %w", err)
    }
    defer rows.Close()

    for rows.Next() {
        var like Like
        like.PhotoID = photoID // Set the PhotoID for each like
        if err := rows.Scan(&like.UserID, &like.Timestamp); err != nil {
            return nil, fmt.Errorf("error scanning like: %w", err)
        }
        likes = append(likes, like)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error in rows: %w", err)
    }

    return likes, nil
}

func (db *appdbimpl) GetUser(userId int) (User, error) {
    var user User

    err := db.c.QueryRow("SELECT UserID, Username FROM Users WHERE UserID = ?", userId).Scan(&user.UserID, &user.Username)
    if err != nil {
        return User{}, fmt.Errorf("error querying user: %w", err)
    }

    return user, nil
}

func (db *appdbimpl) GetUserFollowers(userId int) ([]User, error) {
    var followers []User

    rows, err := db.c.Query("SELECT u.UserID, u.Username FROM Users u INNER JOIN Followers f ON u.UserID = f.FollowerID WHERE f.FollowingID = ?", userId)
    if err != nil {
        return nil, fmt.Errorf("error querying followers: %w", err)
    }
    defer rows.Close()

    for rows.Next() {
        var follower User
        if err := rows.Scan(&follower.UserID, &follower.Username); err != nil {
            return nil, fmt.Errorf("error scanning follower: %w", err)
        }
        followers = append(followers, follower)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error in rows: %w", err)
    }

    return followers, nil
}
func (db *appdbimpl) GetUserFollowing(userId int) ([]User, error) {
    var following []User

    rows, err := db.c.Query("SELECT u.UserID, u.Username FROM Users u INNER JOIN Followers f ON u.UserID = f.FollowingID WHERE f.FollowerID = ?", userId)
    if err != nil {
        return nil, fmt.Errorf("error querying following: %w", err)
    }
    defer rows.Close()

    for rows.Next() {
        var follower User
        if err := rows.Scan(&follower.UserID, &follower.Username); err != nil {
            return nil, fmt.Errorf("error scanning followee: %w", err)
        }
        following = append(following, follower)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error in rows: %w", err)
    }

    return following, nil
}
