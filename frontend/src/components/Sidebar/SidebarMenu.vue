<template>
  <primary-sidebar></primary-sidebar>
</template>

<script lang="ts" setup>
import PrimarySidebar from "./PrimarySidebar.vue";
import FeedEntry from "../Feed/FeedEntry.vue";
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
    };
  });
});
</script>

<style lang="scss" scoped>
#sidebar-menu {
  overflow-y: hidden;
}

.v-btn::before {
  background-color: transparent;
}

.left-align {
  text-align: left;
}
</style>
