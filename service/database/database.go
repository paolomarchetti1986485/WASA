
/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
    UploadPhoto(userID int, uploadDateTime time.Time) (int, error)
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
	
	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='example_table';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE example_table (id INTEGER NOT NULL PRIMARY KEY, name TEXT);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}
	// Define a list of SQL statements to create your tables if they don't exist.
    createTableStatements := []string{
        `CREATE TABLE IF NOT EXISTS Users (UserID INTEGER PRIMARY KEY, Username TEXT NOT NULL);`,
        `CREATE TABLE IF NOT EXISTS Photos (PhotoID INTEGER PRIMARY KEY, UserID INTEGER, UploadDateTime TEXT, FOREIGN KEY (UserID) REFERENCES Users(UserID));`,
        `CREATE TABLE IF NOT EXISTS Comments (CommentID INTEGER PRIMARY KEY, PhotoID INTEGER, UserID INTEGER, CommentText TEXT NOT NULL, Timestamp DATETIME DEFAULT CURRENT_TIMESTAMP, FOREIGN KEY (PhotoID) REFERENCES Photos(PhotoID), FOREIGN KEY (UserID) REFERENCES Users(UserID));`,
        `CREATE TABLE IF NOT EXISTS Likes (PhotoID INTEGER, UserID INTEGER, Timestamp DATETIME DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY (PhotoID, UserID), FOREIGN KEY (PhotoID) REFERENCES Photos(PhotoID), FOREIGN KEY (UserID) REFERENCES Users(UserID));`,
        `CREATE TABLE IF NOT EXISTS Followers (FollowerID INTEGER, FollowingID INTEGER, PRIMARY KEY (FollowerID, FollowingID), FOREIGN KEY (FollowerID) REFERENCES Users(UserID), FOREIGN KEY (FollowingID) REFERENCES Users(UserID));`,
		`CREATE TABLE IF NOT EXISTS Banned (BannedUserID INTEGER, UserID INTEGER, BanDateTime DATETIME, PRIMARY KEY (BanUserID, UserID), FOREIGN KEY (UserID) REFERENCES Users(UserID), FOREIGN KEY (BannedUserID) REFERENCES Users(UserID));`,
        // Add other table creation statements here for `BannedUsers` if applicable.
    }

    // Iterate over each create table statement and execute it.
    for _, sqlStmt := range createTableStatements {
        _, err := db.Exec(sqlStmt)
        if err != nil {
            return nil, fmt.Errorf("error creating database table: %w", err)
        }
    }
	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
