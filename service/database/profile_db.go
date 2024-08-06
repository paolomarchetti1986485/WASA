package database

import (
	"fmt"
)

// Profile represents the user's profile information.
type Profile struct {
	Username  string  `json:"username"`
	Photos    []Photo `json:"photos"`
	Followers []User  `json:"followers"`
	Following []User  `json:"following"`
}

// GetUserProfile retrieves the profile of a user including their username, photos, followers, and following.
func (db *appdbimpl) GetUserProfile(userId int) (Profile, error) {
	var profile Profile

	// Retrieve the username
	err := db.c.QueryRow("SELECT Username FROM Users WHERE UserID = ?", userId).Scan(&profile.Username)
	if err != nil {
		return profile, fmt.Errorf("error querying username: %w", err)
	}

	// Retrieve the user's photos
	photosQuery := `
        SELECT p.PhotoID, p.UserID, p.UploadDateTime
        FROM Photos p
        WHERE p.UserID = ?
        ORDER BY p.UploadDateTime DESC`
	photosRows, err := db.c.Query(photosQuery, userId)
	if err != nil {
		return profile, fmt.Errorf("error querying user photos: %w", err)
	}
	defer photosRows.Close()

	for photosRows.Next() {
		var photo Photo
		if err := photosRows.Scan(&photo.PhotoID, &photo.UserID, &photo.UploadDateTime); err != nil {
			return profile, fmt.Errorf("error scanning photo: %w", err)
		}

		photo.Comments, err = db.GetPhotoComments(photo.PhotoID)
		if err != nil {
			return profile, fmt.Errorf("error fetching photo comments: %w", err)
		}

		photo.Likes, err = db.GetPhotoLikes(photo.PhotoID)
		if err != nil {
			return profile, fmt.Errorf("error fetching photo likes: %w", err)
		}

		profile.Photos = append(profile.Photos, photo)
	}

	if err := photosRows.Err(); err != nil {
		return profile, fmt.Errorf("error in rows: %w", err)
	}

	// Retrieve the user's followers
	followersQuery := `
        SELECT u.UserID, u.Username
        FROM Users u
        INNER JOIN Followers f ON u.UserID = f.FollowerID
        WHERE f.FollowingID = ?`
	followersRows, err := db.c.Query(followersQuery, userId)
	if err != nil {
		return profile, fmt.Errorf("error querying user followers: %w", err)
	}
	defer followersRows.Close()

	for followersRows.Next() {
		var follower User
		if err := followersRows.Scan(&follower.UserID, &follower.Username); err != nil {
			return profile, fmt.Errorf("error scanning follower: %w", err)
		}
		profile.Followers = append(profile.Followers, follower)
	}

	if err := followersRows.Err(); err != nil {
		return profile, fmt.Errorf("error in rows: %w", err)
	}

	// Retrieve the users the user is following
	followingQuery := `
        SELECT u.UserID, u.Username
        FROM Users u
        INNER JOIN Followers f ON u.UserID = f.FollowingID
        WHERE f.FollowerID = ?`
	followingRows, err := db.c.Query(followingQuery, userId)
	if err != nil {
		return profile, fmt.Errorf("error querying user following: %w", err)
	}
	defer followingRows.Close()

	for followingRows.Next() {
		var following User
		if err := followingRows.Scan(&following.UserID, &following.Username); err != nil {
			return profile, fmt.Errorf("error scanning following: %w", err)
		}
		profile.Following = append(profile.Following, following)
	}

	if err := followingRows.Err(); err != nil {
		return profile, fmt.Errorf("error in rows: %w", err)
	}

	return profile, nil
}
