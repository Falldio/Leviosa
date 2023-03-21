<template>
  <div>
    <v-navigation-drawer permanent>
      <FeedMenu></FeedMenu>
      <v-list lines="two">
        <feed-entry
          v-for="(item, index) in items"
          :key="item"
          :index="index"
          :feed="item"
        />
      </v-list>
    </v-navigation-drawer>
  </div>
</template>
<script setup lang="ts">
import FeedEntry from "../components/Feed/FeedEntry.vue";
import FeedMenu from "../components/Feed/FeedMenu.vue";
import { computed } from "vue";
import { useStore } from "vuex";

const store = useStore();
const items = computed(() => {
  return store.state.feeds.map((feed: any) => {
    return {
      id: feed.id,
      title: feed.title,
      description: feed.description,
      icon: feed.image,
      unread: feed.unread,
      pinned: feed.pinned,
    };
  });
});
</script>
<style lang="less"></style>
