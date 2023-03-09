package db

import (
	"log"
	"time"

	"github.com/mmcdole/gofeed"
)

func newFeed(fd *gofeed.Feed) Feed {
	feed := &Feed{}
	feed.Url = fd.FeedLink
	feed.Title = fd.Title
	feed.Description = fd.Description
	if fd.Image != nil {
		feed.Image = fd.Image.URL
	}
	if fd.UpdatedParsed != nil {
		feed.Updated = fd.UpdatedParsed.UnixMicro()
	}
	feed.Created = time.Now().UnixMicro()
	err := dbm.Insert(feed)
	if err != nil {
		log.Println("Feed already exists")
		return Feed{}
	}
	return *feed
}

func newPosts(feed *gofeed.Feed, feedId int64) []Post {
	var posts []Post
	for _, item := range feed.Items {
		post := &Post{
			FeedId:      feedId,
			Title:       item.Title,
			Description: item.Description,
			Url:         item.Link,
			Content:     item.Content,
			Read:        false,
			Starred:     false,
		}
		if item.UpdatedParsed != nil {
			post.Updated = item.UpdatedParsed.UnixMicro()
		}
		var author string = feed.Title
		if item.Author != nil {
			author = item.Author.Name
		}
		post.Author = author
		err := dbm.Insert(post)
		if err != nil {
			log.Println("Post already exists")
			post.Read = true
		}
		posts = append(posts, *post)
	}
	return posts
}
