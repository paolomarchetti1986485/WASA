package api
import (
	
    
    "WASA/service/database"
	
    "time"
)
// Username represents a user's username.
type Username string

// UserId represents a user's ID.
type UserId int

// Login represents the identifier sent by a user during the login.
type Login struct {
    Identifier string `json:"identifier"`
}

// User represents a user in WASAPhoto.
type User struct {
    ID       int   `json:"id"`
    Username string `json:"username"`
}

// CommentId represents a comment's ID.
type CommentId int

// Comment represents a comment on a photo.
type Comment struct {
    Text      string    `json:"text"`
    CommentId int `json:"commentId"`
    UserId   int    `json:"userId"`
    PhotoId  int    `json:"photoid"`
    Timestamp time.Time  `json:"timestamp"`
}

// LikeId represents a like's ID.
type LikeId int

// UnuploadedPhoto represents a photo that has not been uploaded yet.
type UnuploadedPhoto string

// PhotoId represents a photo's identifier.
type PhotoId int

// Photo represents a photo in WASAPhoto.
type Photo struct {
    ID        int    `json:"id"`
    DateTime  time.Time     `json:"dateTime"` // Consider changing this to time.Time for actual implementation
    Comments  []database.Comment  `json:"comments"`
    Likes     []database.Like   `json:"likes"`
    UserID   int     `json:"photoid"`
}

// Profile represents a user's full profile, including followers, following, and photos.
type Profile struct {
    Followers []database.User  `json:"followers"`
    Following []database.User  `json:"following"`
    Photos    []database.Photo `json:"photos"`
}

// Stream represents the stream of a user, which is a collection of photos.
type Stream []database.Photo

func (u User) ToDatabaseU() database.User {
	return database.User{
		UserID: u.ID,
	}
}

// Converts a Photo from the api package to a Photo of the database package
func (p Photo) ToDatabaseP() database.Photo {
	return database.Photo{
		Comments: p.Comments,
		Likes:    p.Likes,
		UserID:    p.UserID,
		PhotoID:  p.ID,
		UploadDateTime: p.DateTime,
	}
}

// Converts a PhotoId from the api package to a PhotoId of the database package
func (p Photo) ToDatabase() database.Photo {
	return database.Photo{
		PhotoID: p.ID,
	}
}

// Converts a PhotoId from the api package to a PhotoId of the database package
func (u User) ToDatabase() database.User {
	return database.User{
		Username: u.Username,
	}
}

// Converts a Comment from the api package to a Comment of the database package
func (c Comment) ToDatabase() database.Comment {
	return database.Comment{
	    CommentID: c.CommentId,
        PhotoID: c.PhotoId,
        UserID: c.UserId,
        CommentText: c.Text,
        Timestamp: c.Timestamp,
	}
}

// Converts a CommentId from the api package to a CommentId of the database package
func (c Comment) ToDatabaseCommentID() database.Comment {
	return database.Comment{
		CommentID: c.CommentId,
	}
}

