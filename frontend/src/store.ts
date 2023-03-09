import { createStore, createLogger } from 'vuex'
import { db } from '../wailsjs/go/models'

export const store = createStore({
    state () {
        return {
            addRSSFeedDialog: false,
            feeds: [] as db.Feed[],
            postDetail: "" as string,
        }
    },
    mutations: {
        addFeed(state, feed: db.Feed) {
            state.feeds.unshift(feed);
        },
        addFeeds(state, feeds: db.Feed[]) {
            state.feeds.unshift(...feeds);
        },
        setAddRSSFeedDialog(state, value: boolean) {
            state.addRSSFeedDialog = value;
        },
        setFeeds(state, feeds: db.Feed[]) {
            state.feeds = feeds;
        },
        setPostDetail(state, postDetail: string) {
            state.postDetail = postDetail;
        }
    },
    actions: {
    },
    plugins: [createLogger()]
})