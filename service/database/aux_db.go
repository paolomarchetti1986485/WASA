package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// GetUserByUsername retrieves a user's ID based on their username.
// It returns the user ID and any error encountered.
func (db *appdbimpl) GetUserByUsername(username string) (int, error) {
	var userID int
	// Query to find the user ID based on the username.
	err := db.c.QueryRow("SELECT UserID FROM Users WHERE Username = ?", username).Scan(&userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// No user was found with the given username.
			return 0, err
		}
		// An error occurred during the query execution.
		return 0, fmt.Errorf("error querying user by username: %w", err)
	}
	return userID, nil
}

// GetPhotoComments retrieves all comments for a given photo.
// It returns a slice of comments and any error encountered.
func (db *appdbimpl) GetPhotoComments(photoID int) ([]Comment, error) {
	var comments []Comment

	// SQL query to retrieve comments for a specific photo.
	query := "SELECT CommentID, UserID, CommentText, Timestamp FROM Comments WHERE PhotoID = ?"
	rows, err := db.c.Query(query, photoID)
	if err != nil {
		// An error occurred while querying the comments.
		return nil, fmt.Errorf("error querying comments: %w", err)
	}
	defer rows.Close()

	// Iterating over the query results.
	for rows.Next() {
		var comment Comment
		comment.PhotoID = photoID // Assigning the photo ID to each comment.
		// Scanning the row into the comment struct.
		if err := rows.Scan(&comment.CommentID, &comment.UserID, &comment.CommentText, &comment.Timestamp); err != nil {
			// An error occurred while scanning the comment.
			return nil, fmt.Errorf("error scanning comment: %w", err)
		}
		// Appending the comment to the slice of comments.
		comments = append(comments, comment)
	}

	// Checking for any errors that occurred during the iteration.
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error in rows: %w", err)
	}

	return comments, nil
}

// GetUserPhotos retrieves all photos uploaded by a specific user.
// It returns a slice of photos and any error encountered.
func (db *appdbimpl) GetUserPhotos(userId int) ([]Photo, error) {
	var photos []Photo

	// SQL query to retrieve photos uploaded by a user.
	rows, err := db.c.Query("SELECT PhotoID, UploadDateTime FROM Photos WHERE UserID = ?", userId)
	if err != nil {
		// An error occurred while querying the photos.
		return nil, fmt.Errorf("error querying photos: %w", err)
	}
	defer rows.Close()

	// Iterating over the query results.
	for rows.Next() {
		var photo Photo
		// Scanning the row into the photo struct.
		if err := rows.Scan(&photo.PhotoID, &photo.UploadDateTime); err != nil {
			// An error occurred while scanning the photo.
			return nil, fmt.Errorf("error scanning photo: %w", err)
		}

		// Fetching comments for each photo.
		photo.Comments, err = db.GetPhotoComments(photo.PhotoID)
		if err != nil {
			return nil, fmt.Errorf("error fetching photo comments: %w", err)
		}

		// Fetching likes for each photo.
		photo.Likes, err = db.GetPhotoLikes(photo.PhotoID)
		if err != nil {
			return nil, fmt.Errorf("error fetching photo likes: %w", err)
		}

		// Appending the photo to the slice of photos.
		photos = append(photos, photo)
	}

	// Checking for any errors that occurred during the iteration.
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error in rows: %w", err)
	}

	return photos, nil
}

// GetPhotoLikes retrieves all likes for a given photo.
// It returns a slice of likes and any error encountered.
func (db *appdbimpl) GetPhotoLikes(photoID int) ([]Like, error) {
	var likes []Like

	// SQL query to retrieve likes for a specific photo.
	query := "SELECT UserID, Timestamp FROM Likes WHERE PhotoID = ?"
	rows, err := db.c.Query(query, photoID)
	if err != nil {
		// An error occurred while querying the likes.
		return nil, fmt.Errorf("error querying likes: %w", err)
	}
	defer rows.Close()

	// Iterating over the query results.
	for rows.Next() {
		var like Like
		like.PhotoID = photoID // Assigning the photo ID to each like.
		// Scanning the row into the like struct.
		if err := rows.Scan(&like.UserID, &like.Timestamp); err != nil {
			// An error occurred while scanning the like.
			return nil, fmt.Errorf("error scanning like: %w", err)
		}
		// Appending the like to the slice of likes.
		likes = append(likes, like)
	}

	// Checking for any errors that occurred during the iteration.
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error in rows: %w", err)
	}

	return likes, nil
}

// GetUser retrieves the user details based on the user ID.
// It returns a User struct and any error encountered.
func (db *appdbimpl) GetUser(userId int) (User, error) {
	var user User

	// SQL query to retrieve user details based on the user ID.
	err := db.c.QueryRow("SELECT UserID, Username FROM Users WHERE UserID = ?", userId).Scan(&user.UserID, &user.Username)
	if err != nil {
		// An error occurred while querying the user.
		return User{}, fmt.Errorf("error querying user: %w", err)
	}

	return user, nil
}

// GetUserFollowers retrieves all followers of a specific user.
// It returns a slice of users (followers) and any error encountered.
func (db *appdbimpl) GetUserFollowers(userId int) ([]User, error) {
	var followers []User

	// SQL query to retrieve followers of a user.
	rows, err := db.c.Query("SELECT u.UserID, u.Username FROM Users u INNER JOIN Followers f ON u.UserID = f.FollowerID WHERE f.FollowingID = ?", userId)
	if err != nil {
		// An error occurred while querying the followers.
		return nil, fmt.Errorf("error querying followers: %w", err)
	}
	defer rows.Close()

	// Iterating over the query results.
	for rows.Next() {
		var follower User
		// Scanning the row into the follower struct.
		if err := rows.Scan(&follower.UserID, &follower.Username); err != nil {
			// An error occurred while scanning the follower.
			return nil, fmt.Errorf("error scanning follower: %w", err)
		}
		// Appending the follower to the slice of followers.
		followers = append(followers, follower)
	}

	// Checking for any errors that occurred during the iteration.
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error in rows: %w", err)
	}

	return followers, nil
}

// GetUserFollowing retrieves all users that a specific user is following.
// It returns a slice of users (following) and any error encountered.
func (db *appdbimpl) GetUserFollowing(userId int) ([]User, error) {
	var following []User

	// SQL query to retrieve users that the specified user is following.
	rows, err := db.c.Query("SELECT u.UserID, u.Username FROM Users u INNER JOIN Followers f ON u.UserID = f.FollowingID WHERE f.FollowerID = ?", userId)
	if err != nil {
		// An error occurred while querying the users being followed.
		return nil, fmt.Errorf("error querying following: %w", err)
	}
	defer rows.Close()

	// Iterating over the query results.
	for rows.Next() {
		var follower User
		// Scanning the row into the followee struct.
		if err := rows.Scan(&follower.UserID, &follower.Username); err != nil {
			// An error occurred while scanning the followee.
			return nil, fmt.Errorf("error scanning followee: %w", err)
		}
		// Appending the followee to the slice of users being followed.
		following = append(following, follower)
	}

	// Checking for any errors that occurred during the iteration.
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error in rows: %w", err)
	}

	return following, nil
}
func (db *appdbimpl) GetAllUsers() ([]User, error) {
	var users []User

	// SQL query to retrieve all users.
	query := "SELECT UserID, Username FROM Users"
	rows, err := db.c.Query(query)
	if err != nil {
		// An error occurred while querying the users.
		return nil, fmt.Errorf("error querying users: %w", err)
	}
	defer rows.Close()

	// Iterating over the query results.
	for rows.Next() {
		var user User
		// Scanning the row into the user struct.
		if err := rows.Scan(&user.UserID, &user.Username); err != nil {
			// An error occurred while scanning the user.
			return nil, fmt.Errorf("error scanning user: %w", err)
		}
		// Appending the user to the slice of users.
		users = append(users, user)
	}

	// Checking for any errors that occurred during the iteration.
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error in rows: %w", err)
	}

	return users, nil
}
func (db *appdbimpl) SearchUsersByUsernamePrefix(prefix string) ([]User, error) {
	rows, err := db.c.Query("SELECT UserID, Username FROM Users WHERE Username LIKE ? LIMIT 100", prefix+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.UserID, &user.Username); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	// Aggiungi il controllo per rows.Err() qui
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return users, nil
}
