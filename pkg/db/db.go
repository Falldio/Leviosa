package db

import (
	"Leviosa/pkg/config"
	"Leviosa/pkg/log"
	"database/sql"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"

	"github.com/mmcdole/gofeed"
	"go.deanishe.net/favicon"
	"gopkg.in/gorp.v2"
)

var (
	fp                  *gofeed.Parser
	dbm                 *gorp.DbMap
	proxyTransport      *http.Transport
	iconFinder          *favicon.Finder
	iconFinderWithProxy *favicon.Finder
)

const (
	dbDriver = "sqlite3"
	dbPath   = "leviosa.db"
)

const (
	SearchModeAnd = iota
	SearchModeOr
)

func init() {
	proxy, _ := url.Parse(config.Config.ProxyUrl)
	proxyTransport = &http.Transport{Proxy: http.ProxyURL(proxy)}
	proxyClient := &http.Client{Transport: proxyTransport}
	iconFinderWithProxy = favicon.New(favicon.WithClient(proxyClient))
	iconFinder = favicon.New()
}

func InitDb() *gorp.DbMap {
	fp = gofeed.NewParser()
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		os.Create(dbPath)
	}
	database, _ := sql.Open("sqlite3", dbPath)

	log.Logger.Info("Initializing database")

	dbm = &gorp.DbMap{Db: database, Dialect: gorp.SqliteDialect{}}
	dbm.AddTableWithName(Feed{}, "feeds").SetKeys(true, "Id").AddIndex("FeedUrlIndex", "Btree", []string{"Url"}).SetUnique(true)
	dbm.AddTableWithName(Post{}, "posts").SetKeys(true, "Id").AddIndex("PostUrlIndex", "Btree", []string{"Url"}).SetUnique(true)
	dbm.AddTableWithName(Tag{}, "tags").SetKeys(true, "Id").AddIndex("TagNameIndex", "Btree", []string{"Name"}).SetUnique(true)
	dbm.AddTableWithName(FeedTag{}, "feed_tags").AddIndex("FeedTagIdIndex", "Btree", []string{"FeedId", "TagId"})

	// a.dbMap.TraceOn("[gorp]", log.New(os.Stdout, "leviosa:", log.Lmicroseconds))

	dbm.CreateTablesIfNotExists()

	err = dbm.CreateIndex()
	if err != nil {
		log.Logger.Info(err.Error())
	}

	return dbm
}

func AddRSSFeed(feedUrl string) Feed {
	log.Logger.Debug("Adding feed: " + feedUrl)
	fd := parseUrl(feedUrl)
	if fd == nil {
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

func FetchUpdatesForAllFeeds() {
	var feeds []Feed
	dbm.Select(&feeds, "select id, url from feeds")
	for _, f := range feeds {
		log.Logger.Debug("Fetching updates for feed: " + f.Url)
		fd := parseUrl(f.Url)
		newPosts(fd, f.Id)
	}
}

func FetchUpdates(feedId int64) {
	feed := Feed{}
	dbm.SelectOne(&feed, "select url from feeds where id = ?", feedId)
	log.Logger.Debug("Fetching updates for feed: " + feed.Url)
	fd := parseUrl(feed.Url)
	newPosts(fd, feedId)
}

func SetStarred(postId int64, starred bool) {
	post := Post{}
	dbm.SelectOne(&post, "select * from posts where id = ?", postId)
	post.Starred = starred
	dbm.Update(&post)
}

func SetPinned(feedId int64, pinned bool) {
	feed := Feed{}
	dbm.SelectOne(&feed, "select * from feeds where id = ?", feedId)
	feed.Pinned = pinned
	dbm.Update(&feed)
}

func UnsubscribeFeed(feedId int64) {
	feed := Feed{}
	dbm.SelectOne(&feed, "select * from feeds where id = ?", feedId)
	dbm.Delete(&feed)
	posts := []Post{}
	dbm.Select(&posts, "select * from posts where feed_id = ?", feedId)
	for _, p := range posts {
		dbm.Delete(&p)
	}
}

func AddTags(feedId int64, tags ...string) {
	for _, t := range tags {
		tag := Tag{}
		err := dbm.SelectOne(&tag, "select * from tags where name = ?", t)
		if err != nil {
			tag.Name = t
			dbm.Insert(&tag)
		}
		feedTag := FeedTag{FeedId: feedId, TagId: tag.Id}
		dbm.Insert(&feedTag)
	}
}

func GetTags() []Tag {
	var tags []Tag
	dbm.Select(&tags, "select * from tags")
	return tags
}

func DeleteTagFromFeed(feedId int64, tagName string) {
	dbm.Exec("delete from feed_tags where feed_id = ? and tag_id in (select id from tags where name = ?)", feedId, tagName)
}

func SearchFeedsByTags(mode int, tags ...string) []Feed {
	var tagIds []interface{}
	for _, t := range tags {
		tag := Tag{}
		err := dbm.SelectOne(&tag, "select * from tags where name = ?", t)
		if err != nil {
			log.Logger.Info("Tag not found: " + t)
			return []Feed{}
		}
		tagIds = append(tagIds, tag.Id)
	}
	var feeds []Feed
	if mode == SearchModeAnd {
		dbm.Select(&feeds, "select f.* from feeds f, feed_tags ft where f.id = ft.feed_id and ft.tag_id in ("+strings.TrimRight(strings.Repeat("?,", len(tagIds)), ",")+") group by f.id having count(*) = ?", tagIds, len(tagIds))
	} else if mode == SearchModeOr {
		dbm.Select(&feeds, "select f.* from feeds f, feed_tags ft where f.id = ft.feed_id and ft.tag_id in ("+strings.TrimRight(strings.Repeat("?,", len(tagIds)), ",")+") group by f.id", tagIds)
	}
	return feeds
}

func parseUrl(feedUrl string) *gofeed.Feed {
	fd, err := fp.ParseURL(feedUrl)
	if err != nil {
		log.Logger.Error(err.Error())
		log.Logger.Info("Retrying with proxy")
		fp.Client.Transport = proxyTransport
		fd, err = fp.ParseURL(feedUrl)
		fp.Client.Transport = nil
		if err != nil {
			log.Logger.Error(err.Error())
			return nil
		}
	}
	return fd
}

func SetRead(postId int64, read bool) {
	post := Post{}
	dbm.SelectOne(&post, "select * from posts where id = ?", postId)
	post.Read = read
	dbm.Update(&post)
}
