/*
This file includes necessary DTOs for the underlying SQLite operations.

They are defined in the format of Golang structs, and will be passed to the frontend as results of
functions in `db.go`

Usage for thr frontend developers:

# When you run `wails dev` or `wails generate module`, corresponding TS objects
will be generated inside `wailsjs/models.ts`.

# You can simply import the `db` namespace, and use them just like regular TS/JS objects.

# One may refer to https://wails.io/docs/howdoesitwork#calling-bound-go-methods for more info.
*/
package db

// A RSS Feed without any post info. It can be identified with the `Url` field but we use `Id` for query efficiency reasons.
type Feed struct {
	Id          int64  `db:"id, primarykey, autoincrement" json:"id"`
	Url         string `db:"url" json:"url"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	Image       string `db:"image" json:"image"`
	Updated     int64  `db:"updated" json:"updated"`
	Created     int64  `db:"created" json:"created"`
	Unread      int64  `db:"-"`
	Pinned      int64  `db:"pinned" json:"pinned"`
}

// A Post always belongs to a certain RSS Feed.
type Post struct {
	Id          int64  `db:"id, primarykey, autoincrement" json:"id"`
	FeedId      int64  `db:"feed_id" json:"feed_id"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	Url         string `db:"url" json:"url"`
	Content     string `db:"content" json:"content"`
	Updated     int64  `db:"updated" json:"updated"`
	Author      string `db:"author" json:"author"`
	Read        bool   `db:"read" json:"read"`
	Starred     bool   `db:"starred" json:"starred"`
}
