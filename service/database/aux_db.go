package database
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
func (db *appdbimpl) GetUserPhotos(userId int) ([]Photo, error) {
    var photos []Photo

    rows, err := db.c.Query("SELECT PhotoID, UploadDateTime FROM Photos WHERE UserID = ?", userId)
    if err != nil {
        return nil, fmt.Errorf("error querying photos: %w", err)
    }
    defer rows.Close()

    for rows.Next() {
        var photo Photo
        if err := rows.Scan(&photo.ID, &photo.DateTime); err != nil {
            return nil, fmt.Errorf("error scanning photo: %w", err)
        }
        photos = append(photos, photo)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error in rows: %w", err)
    }

    return photos, nil
}
func (db *appdbimpl) GetUser(userId int) (User, error) {
    var user User

    err := db.c.QueryRow("SELECT UserID, Username FROM Users WHERE UserID = ?", userId).Scan(&user.ID, &user.Username)
    if err != nil {
        return User{}, fmt.Errorf("error querying user: %w", err)
    }

    return user, nil
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
        if err := rows.Scan(&photo.ID, &photo.DateTime); err != nil {
            return nil, fmt.Errorf("error scanning photo: %w", err)
        }
        photos = append(photos, photo)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error in rows: %w", err)
    }

    return photos, nil
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
        if err := rows.Scan(&follower.ID, &follower.Username); err != nil {
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
        var followee User
        if err := rows.Scan(&followee.ID, &followee.Username); err != nil {
            return nil, fmt.Errorf("error scanning followee: %w", err)
        }
        following = append(following, followee)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error in rows: %w", err)
    }

    return following, nil
}
