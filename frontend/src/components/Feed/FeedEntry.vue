<template>
  <v-list-item @click="getPostList">
    <v-list-item-title class="left-align">{{ feed.title }}</v-list-item-title>
    <v-list-item-subtitle
      class="left-align"
      v-html="feed.description"
    ></v-list-item-subtitle>
  </v-list-item>
</template>

<script setup lang="ts">
import { useStore } from "vuex";
import { useRouter } from "vue-router";

const props = defineProps(["feed"]);

const store = useStore();
const router = useRouter();

const getPostList = () => {
  router.push("/feeds/" + props.feed.id);
};
</script>

<style scoped>
.left-align {
  text-align: left;
}
</style>
