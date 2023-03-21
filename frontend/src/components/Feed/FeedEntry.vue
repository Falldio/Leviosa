<template>
  <v-menu location="end" v-model="menu">
    <template v-slot:activator="{ props }">
      <v-list-item
        v-bind:="props"
        @click.left.prevent="getPostList"
        @click.right.prevent="menu = !menu"
      >
        <div class="content">
          <div class="content-top">
            <v-list-item-title class="topleft-align">
              {{ feed.title }}
            </v-list-item-title>
            <div class="topright-align">
              <v-icon
                v-if="feed.pinned"
                color="orange-darken-2"
                icon="mdi-pin"
              ></v-icon>
              <!-- <v-list>
              <v-list-item
                v-for="(item, index) in items"
                :key="index"
                @click="onItemClick(item)"
              >
                <v-list-item-title>{{ item }}</v-list-item-title>
              </v-list-item>
            </v-list> -->
            </div>
          </div>
          <div class="content-bottom">
            <v-list-item-subtitle
              class="bottomleft-align"
              v-html="feed.description"
            ></v-list-item-subtitle>
          </div>
        </div>
      </v-list-item>
    </template>
    <div class="menu-content"></div>
    <v-list>
      <v-list-item
        ><v-btn
          variant="text"
          color="orange-darken-2"
          v-if="feed.pinned"
          prepend-icon="mdi-pin-off"
          @click="setFeedPinned(feed, false)"
          >UnPin</v-btn
        ><v-btn
          variant="text"
          color="orange-darken-2"
          v-else
          prepend-icon="mdi-pin"
          @click="setFeedPinned(feed, true)"
          >Pin</v-btn
        ></v-list-item
      >

      <v-list-item
        ><v-btn
          variant="text"
          color="error"
          prepend-icon="mdi-delete-outline"
          @click="deleteFeed(feed)"
        >
          Delete</v-btn
        ></v-list-item
      >

      <v-list-item
        ><v-btn
          variant="text"
          color="error"
          prepend-icon="mdi-reload"
          @click="updateFeed(feed)"
        >
          Update</v-btn
        ></v-list-item
      >
    </v-list>
  </v-menu>
  <v-snackbar v-model="snackbar" :timeout="2000">
    {{ snackbarText }}

    <template v-slot:actions>
      <v-btn color="blue" variant="text" @click="snackbar = false">
        Close
      </v-btn>
    </template>
  </v-snackbar>
</template>

<script setup lang="ts">
import { ref, reactive, isReactive } from "vue";
import { useStore } from "vuex";
import { useRouter } from "vue-router";
import {
  SetPinned,
  GetFeeds,
  FetchUpdates,
  UnsubscribeFeed,
} from "../../../wailsjs/go/main/App";
const store = useStore();
const router = useRouter();
const props = defineProps<{
  feed: any;
  index: number;
}>();

const feed = reactive(props.feed);
let menu = ref(false);
let snackbar = ref(false);
let snackbarText = ref("Null");

const getPostList = () => {
  menu.value = false;
  router.push("/feeds/" + props.feed.id);
};

const setFeedPinned = (feed: any, state: boolean) => {
  console.log(SetPinned(feed.id, state));
  store.state.feeds[props.index].pinned = state;
};

const deleteFeed = (feed: any) => {
  UnsubscribeFeed(feed.id);

  store.state.feeds.splice(props.index, 1);
  // emits("delete-event", feed.id, props.index);
};

const updateFeed = async (feed: any): Promise<void> => {
  try {
    await FetchUpdates(feed.id);
  } catch (error) {
    snackbarText.value = "Error:" + error;
    snackbar.value = true;
    console.log(error);
  } finally {
    snackbarText.value = "UpdateFeed Success!";
    snackbar.value = true;
  }
};

const output = (something: any) => {
  console.log(something);
};
</script>

<style scoped lang="less">
.content {
  .content-top {
    display: flex;
    justify-content: space-between;
    align-items: center;
    .topleft-align {
      text-align: left;
    }
    .topright-align {
    }
  }
  .content-bottom {
    .bottomleft-align {
      text-align: left;
    }
  }
}
.menu-content {
}
</style>
