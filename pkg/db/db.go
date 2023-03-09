package db

import (
	"database/sql"
	"log"
	"os"
	"sort"

	"github.com/mmcdole/gofeed"
	"gopkg.in/gorp.v2"
)

var (
	fp  *gofeed.Parser
	dbm *gorp.DbMap
)

const (
	dbDriver = "sqlite3"
	dbPath   = "leviosa.db"
)

func InitDb() *gorp.DbMap {
	fp = gofeed.NewParser()
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		os.Create(dbPath)
	}
	database, _ := sql.Open("sqlite3", dbPath)

	dbm = &gorp.DbMap{Db: database, Dialect: gorp.SqliteDialect{}}
	dbm.AddTableWithName(Feed{}, "feeds").SetKeys(true, "Id").AddIndex("FeedUrlIndex", "Btree", []string{"Url"}).SetUnique(true)
	dbm.AddTableWithName(Post{}, "posts").SetKeys(true, "Id").AddIndex("PostUrlIndex", "Btree", []string{"Url"}).SetUnique(true)

	// a.dbMap.TraceOn("[gorp]", log.New(os.Stdout, "leviosa:", log.Lmicroseconds))

	dbm.CreateTablesIfNotExists()

	err = dbm.CreateIndex()
	if err != nil {
		log.Println("Index already exists")
	}

	return dbm
}

func AddRSSFeed(url string) Feed {
	fd, err := fp.ParseURL(url)
	if err != nil {
		return Feed{}
	}
	feed := newFeed(fd)
	posts := newPosts(fd, feed.Id)
	for _, p := range posts {
		if !p.Read {
			feed.Unread++
		}
	}
	return feed
}

func GetFeeds() []Feed {
	var feeds []Feed
	dbm.Select(&feeds, "select * from feeds order by updated desc")
	sort.Slice(feeds, func(i, j int) bool {
		return feeds[i].Updated > feeds[j].Updated
	})
	return feeds
}

func GetPosts(feedId int64) []Post {
	var posts []Post
	dbm.Select(&posts, "select id, title, description from posts where feed_id = ? order by updated desc", feedId)
	return posts
}

func GetPost(postId int64) Post {
	var post Post
	dbm.SelectOne(&post, "select * from posts where id = ?", postId)
	return post
}
