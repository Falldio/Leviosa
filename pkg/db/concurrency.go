package db

import (
	"Leviosa/pkg/log"
	"encoding/json"
	"os"
	"sort"
)

const (
	workerNum = 5  // size of the worker pool
	chanSize  = 20 // size of the channel
)

func FetchUpdatesForAllFeeds() {
	var feeds []Feed
	dbm.Select(&feeds, "select id, url from feeds")
	jobs := make(chan int64, chanSize)
	for i := 0; i < workerNum; i++ {
		go fetchUpdatesBatch(jobs)
	}

	for _, f := range feeds {
		jobs <- f.Id
	}
	close(jobs)
	log.Logger.Info("Fetching Updates Finished!")
}

func fetchUpdatesBatch(jobs <-chan int64) {
	for feedId := range jobs {
		FetchUpdatesForFeed(feedId)
	}
}

func FetchUpdatesForFeed(feedId int64) {
	feed := Feed{}
	dbm.SelectOne(&feed, "select * from feeds where id = ?", feedId)
	log.Logger.Debug("Fetching updates for feed: " + feed.Url)
	fd := parseUrl(feed.Url)
	if fd == nil {
		log.Logger.Error("Failed to parse feed: " + feed.Url)
		return
	}
	posts := newPosts(fd, feedId)
	if len(posts) != 0 {
		sort.Slice(posts, func(i, j int) bool {
			return posts[i].Updated > posts[j].Updated
		})
		log.Logger.Debug("New posts: " + posts[0].Title)
		feed.Updated = posts[0].Updated
	}
	_, err := dbm.Update(&feed)
	if err != nil {
		log.Logger.Error(feed.Title)
		log.Logger.Error(err.Error())
	}
}

func ImportFeeds(path string) []Feed {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Logger.Error(err.Error())
	}
	var feeds []Feed
	err = json.Unmarshal(file, &feeds)
	if err != nil {
		log.Logger.Error(err.Error())
	}

	jobs := make(chan string, chanSize)
	results := make(chan Feed, chanSize)

	for i := 0; i < workerNum; i++ {
		go importFeedsBatch(jobs, results)
	}

	for _, feed := range feeds {
		jobs <- feed.Url
	}
	close(jobs)

	ans := []Feed{}
	for feed := range results {
		ans = append(ans, feed)
	}

	return ans
}

func importFeedsBatch(jobs <-chan string, results chan<- Feed) {
	for feedUrl := range jobs {
		results <- AddRSSFeed(feedUrl)
	}
}
