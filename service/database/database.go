package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	DeleteUserByID(userID int) error
	UploadPhoto(userID int, uploadDateTime time.Time, PhotoData []byte) (int, error)
	RemovePhoto(photoID int) error
	AddUser(username string) (int, error)
	UpdateUsername(userID int, newUsername string) error
	FollowUser(followerID, followingID int) error
	UnfollowUser(followerID, followingID int) error
	AddLike(photoID, userID int) error
	RemoveLike(photoID, userID int) error
	AddComment(photoID, userID int, commentText string) (int, error)
	RemoveComment(commentID int) error
	BanUser(bannedID, userID int) error
	UnbanUser(bannedID, userID int) error
	GetUserByUsername(username string) (int, error)
	GetUserStream(userId int) ([]Photo, error)
	GetUser(userId int) (User, error)
	GetUserPhotos(userId int) ([]Photo, error)
	GetUserFollowers(userId int) ([]User, error)
	GetUserFollowing(userId int) ([]User, error)
	Ping() error
	GetAllUsers() ([]User, error)
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Activate foreign keys for db

	_, errPramga := db.Exec(`PRAGMA foreign_keys= ON`)
	if errPramga != nil {
		return nil, fmt.Errorf("error setting pragmas: %w", errPramga)
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='users';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		err = createDatabase(db)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}

// Creates all the necessary sql tables for the WASAPhoto app.
func createDatabase(db *sql.DB) error {
	tables := [6]string{
		`CREATE TABLE IF NOT EXISTS Users (UserID INTEGER PRIMARY KEY, Username TEXT NOT NULL);`,
		`CREATE TABLE IF NOT EXISTS Photos (PhotoID INTEGER PRIMARY KEY, UserID INTEGER, UploadDateTime DATETIME, PhotoData BLOB, FOREIGN KEY (UserID) REFERENCES Users(UserID));`,
		`CREATE TABLE IF NOT EXISTS Comments (CommentID INTEGER PRIMARY KEY, PhotoID INTEGER, UserID INTEGER, CommentText TEXT NOT NULL, Timestamp DATETIME DEFAULT CURRENT_TIMESTAMP, FOREIGN KEY (PhotoID) REFERENCES Photos(PhotoID), FOREIGN KEY (UserID) REFERENCES Users(UserID));`,
		`CREATE TABLE IF NOT EXISTS Likes (PhotoID INTEGER, UserID INTEGER, Timestamp DATETIME DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY (PhotoID, UserID), FOREIGN KEY (PhotoID) REFERENCES Photos(PhotoID), FOREIGN KEY (UserID) REFERENCES Users(UserID));`,
		`CREATE TABLE IF NOT EXISTS Followers (FollowerID INTEGER, FollowingID INTEGER, PRIMARY KEY (FollowerID, FollowingID), FOREIGN KEY (FollowerID) REFERENCES Users(UserID), FOREIGN KEY (FollowingID) REFERENCES Users(UserID));`,
		`CREATE TABLE IF NOT EXISTS Banned (BannedUserID INTEGER, UserID INTEGER, BanDateTime DATETIME, PRIMARY KEY (BannedUserID, UserID), FOREIGN KEY (UserID) REFERENCES Users(UserID), FOREIGN KEY (BannedUserID) REFERENCES Users(UserID));`,
	}

	// Iteration to create all the needed sql schemas
	for i := 0; i < len(tables); i++ {

		sqlStmt := tables[i]
		_, err := db.Exec(sqlStmt)

		if err != nil {
			return err
		}
	}
	return nil
}
