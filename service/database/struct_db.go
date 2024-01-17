package database

import (
	"time"
)

// User represents a user in the WASAPhoto app.
type User struct {
	UserID   int    `json:"userId"`   // User's ID
	Username string `json:"username"` // User's username
}

// Photo represents a photo uploaded by a user.
type Photo struct {
	PhotoID        int       `json:"photoId"`        // Photo's ID
	UserID         int       `json:"userId"`         // ID of the User who posted the photo
	UploadDateTime time.Time `json:"uploadDateTime"` // Date & time in which the photo was uploaded
	Comments       []Comment
	Likes          []Like
	PhotoData      []byte `json:"-"` // Excluding from JSON representation
}

// Comment represents a comment made on a photo.
type Comment struct {
	CommentID   int       `json:"commentId"`   // ID of the comment
	PhotoID     int       `json:"photoId"`     // ID of the photo where the comment was posted
	UserID      int       `json:"userId"`      // ID of the user who posted the comment
	CommentText string    `json:"commentText"` // Text of the comment
	Timestamp   time.Time `json:"timestamp"`   // Date & Time in which the comment was posted
}

// Like represents a like made on a photo.
type Like struct {
	PhotoID   int       `json:"photoId"`   // ID of the photo where the like was placed
	UserID    int       `json:"userId"`    // ID of the user who placed the like
	Timestamp time.Time `json:"timestamp"` // Date & time in which the like was placed
}

// Follower represents a user following another user.
type Follower struct {
	FollowerID  int `json:"followerId"`  // ID of the follower, the person who the user is starting following
	FollowingID int `json:"followingId"` // ID of the following, the person who is starting following someone
}

// BannedUser represents a user who has been banned.
type BannedUser struct {
	BannedUserID int       `json:"bannedUserId"` // ID of the user we want to ban
	UserID       int       `json:"userId"`       // ID of the user who banned someone
	BanDateTime  time.Time `json:"banDateTime"`  // Date & time a user banned someone
}
