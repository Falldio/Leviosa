package db

import (
	"time"

	"Leviosa/pkg/log"

	"github.com/mmcdole/gofeed"
)

func newFeed(fd *gofeed.Feed) Feed {
	feed := &Feed{}
	if fd == nil {
		return *feed
	}
	feed.Url = fd.FeedLink
	feed.Title = fd.Title
	feed.Description = fd.Description
	if fd.Image != nil {
		feed.Image = fd.Image.URL
	} else {
		feed.Image = newImage(fd)
	}
	if fd.UpdatedParsed != nil {
		feed.Updated = fd.UpdatedParsed.UnixMicro()
	}
	feed.Created = time.Now().UnixMicro()

	err := dbm.Insert(feed)
	if err != nil {
		log.Logger.Info(err.Error())
		_, err = dbm.Update(feed)
		if err != nil {
			log.Logger.Error(err.Error())
		}
		return Feed{}
	}
	return *feed
}

func newPosts(feed *gofeed.Feed, feedId int64) []Post {
	if feed == nil {
		return []Post{}
	}
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
		if post.Content == "" {
			post.Content = post.Description
		}
		if item.UpdatedParsed != nil {
			post.Updated = item.PublishedParsed.UnixMicro()
		}
		var author string = feed.Title
		if item.Author != nil {
			author = item.Author.Name
		}
		post.Author = author
		err := dbm.Insert(post)
		if err != nil {
			log.Logger.Info(err.Error())
			post.Read = true
		}
		posts = append(posts, *post)
	}
	return posts
}

func newImage(feed *gofeed.Feed) string {
	if feed == nil || len(feed.Items) == 0 {
		return ""
	}
	icons, err := iconFinder.Find(feed.Items[0].Link)
	if err != nil || len(icons) == 0 {
		log.Logger.Info("Retrying with proxy")
		icons, err = iconFinderWithProxy.Find(feed.Items[0].Link)
		if err != nil || len(icons) == 0 {
			log.Logger.Info("Failed to find icon")
			return ""
		}
	}
	return icons[0].URL
}
