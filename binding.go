/*
This file includes necessary functions for the underlying SQLite operations, which
should be called from the frontend for feature implementation.

Functions in this package are not ganruanteed to be stable since the app is still at an
early dev stage.

Usage for thr frontend developers:

# When you run `wails dev` or `wails generate module`, corresponding JS bindings of the functions
will be generated inside `/frontend/wailsjs/main/App.js`.

# You can simply import these methods, and call them just like regular JS functions.

# One may refer to https://wails.io/docs/howdoesitwork#calling-bound-go-methods for more info.
*/
package main

import (
	"encoding/json"
	"os"

	"Leviosa/pkg/db"
	"Leviosa/pkg/log"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// Add a single new RSS Feed subscrition
func (a *App) AddRSSFeed(url string) db.Feed {
	return db.AddRSSFeed(url)
}

// Get all subscribed RSS Feeds in the database
func (a App) GetFeeds() []db.Feed {
	return db.GetFeeds()
}

// Given a specific RSS Feed `Id`, get all posts (without the content info)
func (a App) GetPosts(feedId int64) []db.Post {
	return db.GetPosts(feedId)
}

// Given a specific Post `Id`, get the single post with all info
func (a App) GetPost(postId int64) db.Post {
	return db.GetPost(postId)
}

// Export all RSS Feed subscriptions from the database and export them in a JSON format
func (a App) ExportFeeds() {
	options := runtime.SaveDialogOptions{
		Title:                "Choose a JSON file to save to",
		DefaultFilename:      "feeds.json",
		CanCreateDirectories: true,
	}
	exportPath, err := runtime.SaveFileDialog(a.ctx, options)
	if err != nil {
		log.Logger.Error(err.Error())
	}
	if exportPath == "" {
		return
	}
	var feeds = db.GetFeeds()
	file, _ := json.Marshal(feeds)
	os.WriteFile(exportPath, file, 0644)
}

// Import external subscription list, should be a valid JSON file
func (a App) ImportFeeds() []db.Feed {
	options := runtime.OpenDialogOptions{
		Title:   "Choose a JSON file to import",
		Filters: []runtime.FileFilter{{DisplayName: "JSON Files (*.json)", Pattern: "*.json;*.JSON"}},
	}
	importPath, err := runtime.OpenFileDialog(a.ctx, options)
	if err != nil {
		log.Logger.Error(err.Error())
	}
	if importPath == "" {
		return []db.Feed{}
	}
	file, err := os.ReadFile(importPath)
	if err != nil {
		log.Logger.Error(err.Error())
	}
	var feeds []db.Feed
	err = json.Unmarshal(file, &feeds)
	if err != nil {
		log.Logger.Error(err.Error())
	}
	for _, feed := range feeds {
		db.AddRSSFeed(feed.Url)
	}
	return feeds
}

// Fetch possible updates for all feeds
// If you wish to fecth update for a single RSS Feed
// consider using FetchUpdates()
func (a App) FetchUpdatesForAllFeeds() {
	db.FetchUpdatesForAllFeeds()
}

// Fetch updates for a single RSS Feed
func (a App) FetchUpdates(feedId int64) {
	db.FetchUpdates(feedId)
}

// Star/Unstar a post
func (a App) SetStarred(postId int64, starred bool) {
	db.SetStarred(postId, starred)
}

// Pin/Unpin a feed
func (a App) SetPinned(feedId int64, pinned bool) {
	db.SetPinned(feedId, pinned)
}

// Unsubscribe a feed
func (a App) UnsubscribeFeed(feedId int64) {
	db.UnsubscribeFeed(feedId)
}
