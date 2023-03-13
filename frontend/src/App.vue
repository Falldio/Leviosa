<script lang="ts" setup>
import SideView from "./views/SideView.vue";
import FeedView from "./views/FeedView.vue";
import { useStore } from "vuex";
import { GetFeeds } from "../wailsjs/go/main/App";
import { useRouter, useRoute } from "vue-router";
import { onUnmounted } from "vue";

const store = useStore();
const router = useRouter();
const route = useRoute();

GetFeeds().then((result) => {
  store.commit("setFeeds", result);
});

const handleBackButton = () => {
  router.back();
};

window.addEventListener("popstate", handleBackButton);

onUnmounted(() => {
  window.removeEventListener("popstate", handleBackButton);
});
</script>

<template>
  <v-app>
    <SideView />
    <FeedView />
    <v-main>
      <router-view></router-view>
    </v-main>
  </v-app>
</template>

<style></style>
