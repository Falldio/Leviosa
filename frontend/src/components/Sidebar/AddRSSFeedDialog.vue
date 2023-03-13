<template>
  <v-dialog>
    <v-card>
      <v-card-title>
        <span class="headline">Add New RSS Feed</span>
      </v-card-title>
      <v-card-text>
        <v-text-field label="URL" v-model="url"></v-text-field>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="blue darken-1" text @click="close">Cancel</v-btn>
        <v-btn color="blue darken-1" text @click="add">Add</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts" setup>
import { useStore } from "vuex";
import { ref } from "vue";
import { AddRSSFeed } from "../../../wailsjs/go/main/App";

const store = useStore();

const url = ref("");

const close = () => {
  store.commit("setAddRSSFeedDialog", false);
};

const add = () => {
  AddRSSFeed(url.value).then((result) => {
    if (result.title == "") {
      close();
      return;
    }
    store.commit("addFeed", result);
    close();
  });
};
</script>

<style scoped></style>
