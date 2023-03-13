<template>
  <div id="post-content">
    <h1>{{ post.title }}</h1>
    <div v-html="post.content"></div>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { GetPost } from "../../../wailsjs/go/main/App";
import { useRoute } from "vue-router";
import { db } from "../../../wailsjs/go/models";
const route = useRoute();
const post = ref({} as db.Post);

GetPost(+route.params.postId).then((result) => {
  post.value = result;
  post.value.content = htmlWithTargetBlank(post.value.content);
});

const htmlWithTargetBlank = (html: string) => {
  const doc = new DOMParser().parseFromString(html, "text/html");
  const links = doc.querySelectorAll("a");
  links.forEach((link) => {
    link.setAttribute("target", "_blank");
  });
  return doc.body.innerHTML;
};
</script>

<style scoped>
@import url("../../styles/post.css");

#post-content {
  padding: 20px;
}
</style>
