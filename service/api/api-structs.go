// Username represents a user's username.
type Username string

// UserId represents a user's ID.
type UserId int

// Login represents the identifier sent by a user during the login.
type Login struct {
    Identifier Username `json:"identifier"`
}

// User represents a user in WASAPhoto.
type User struct {
    ID       UserId   `json:"id"`
    Username Username `json:"username"`
}

// CommentId represents a comment's ID.
type CommentId int

// Comment represents a comment on a photo.
type Comment struct {
    Text      string    `json:"text"`
    CommentId CommentId `json:"commentId"`
    UserId    UserId    `json:"userId"`
}

// LikeId represents a like's ID.
type LikeId UserId

// UnuploadedPhoto represents a photo that has not been uploaded yet.
type UnuploadedPhoto string

// PhotoId represents a photo's identifier.
type PhotoId int

// Photo represents a photo in WASAPhoto.
type Photo struct {
    ID        PhotoId    `json:"id"`
    DateTime  string     `json:"dateTime"` // Consider changing this to time.Time for actual implementation
    Comments  []Comment  `json:"comments"`
    Likes     []LikeId   `json:"likes"`
}

// Profile represents a user's full profile, including followers, following, and photos.
type Profile struct {
    Followers []User  `json:"followers"`
    Following []User  `json:"following"`
    Photos    []Photo `json:"photos"`
}

// Stream represents the stream of a user, which is a collection of photos.
type Stream []Photo
