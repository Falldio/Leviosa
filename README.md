<h1 align="center">Leviosa</h1>

<div align="center">
  <strong>Yet another desktop RSS Feed Reader</strong><br>
  Just for fun and learning🚀~<br>
</div>

> **Please Note**: The project is still at an early dev stage. It doesn't cover all basic functions of a regular RSS Feed Reader yet.

## Why another RSS Feed Reader?

There're already many wonderful RSS Feed Readers on the mobile platform. However, I believe some serious contents are more suitable to be read with you sitting there in front of your computer, thinking and noting. I've seen a lot of desktop RSS Feed Readers, asking for a subscription fee, or running with a browser. Dude, I just want a simple, easy-to-use openbox app, and that's where Leviosa comes in.
## Features

+ A single, easy-to-use desktop application.
+ What you subscribe stays in your local device.
+ Cross-platform thanks to [Wails](https://wails.io/)!

![Preview](img/preview.png)

## Tools I Use

+ [Vuetify](https://vuetifyjs.com/): a wonderful Vue Component Framework.
+ [Wails](https://wails.io/): it allows us to build this cross-platform project using Golang.
+ [SQLite](https://www.sqlite.org/index.html): a lightweight database.
+ [gofeed](https://github.com/mmcdole/gofeed): a Go library for parsing RSS and Atom feeds.
+ [go-sqlite3](https://github.com/mattn/go-sqlite3): a sqlite3 driver that conforms to the built-in database/sql interface.
+ [gorp](https://github.com/go-gorp/gorp): it's really a time saver for mapping Go structs to database tables.

## Installation

No stable release yet. You can build it from source code.

## Build from Source

1. You must have [Go](https://go.dev/dl/) installed.
2. Install [Wails](https://wails.io/docs/gettingstarted/installation/).
3. Since we use `go-sqlite3`, which requires `cgo`, you have do do the following steps:
  1. Install [GCC](https://gcc.gnu.org/install/binaries.html) and set the proper `$PATH$` environment variable.
  2. Set the `CGO_ENABLED` environment variable to `1`.
4. Clone this repo.
5. Run `wails dev` if you wanna run it in dev mode, or `wails build` if you wanna build it.

## TODO

- [ ] Add Tag Management
- [ ] Add Unread Bubble
- [ ] Add Feed Management
- [ ] Add favorite list
- [ ] Improve UI