<template>
  <div class="menu">
    <div class="menu-left">
      <v-btn
        class="ma-0"
        variant="text"
        icon="mdi-ship-wheel"
        size="small"
        color="red-lighten-2"
      ></v-btn>
    </div>
    <div v-for="(item, index) in state.menuItems" class="menu-main">
      <div class="menu-main-item">
        <v-tooltip :text="item.name" location="top">
          <template v-slot:activator="{ props }">
            <v-btn
              variant="text"
              v-bind="props"
              :icon="item.icon"
              size="small"
              color="red-lighten-2"
              :loading="item.loading"
              @click="item.action"
            ></v-btn
          ></template>
        </v-tooltip>
      </div>
    </div>
    <div class="menu-right">
      <v-btn
        class="ma-0"
        variant="text"
        icon="mdi-dots-horizontal"
        size="small"
        color="red-lighten-2"
      ></v-btn>
    </div>
  </div>
  <add-r-s-s-feed-dialog
    v-model="store.state.addRSSFeedDialog"
  ></add-r-s-s-feed-dialog>
  <v-snackbar v-model="snackbar" :timeout="2000">
    {{ snackbarText }}

    <template v-slot:actions>
      <v-btn color="blue" variant="text" @click="snackbar = false">
        Close
      </v-btn>
    </template>
  </v-snackbar>
</template>

<script lang="ts" setup>
import { reactive, ref } from "vue";
import AddRSSFeedDialog from "./AddRSSFeedDialog.vue";
import { useStore } from "vuex";
import {
  ExportFeeds,
  ImportFeeds,
  FetchUpdatesForAllFeeds,
} from "../../../wailsjs/go/main/App";

const store = useStore();
let snackbar = ref(false);
let snackbarText = ref("Null");

type btnFunction = () => void;
interface item {
  name: string;
  icon: string;
  action: btnFunction;
  loading: boolean;
}
const addAction = (): void => {
  store.commit("setAddRSSFeedDialog", true);
};

const exportAction = (): void => {
  ExportFeeds();
};
const importAction = (): void => {
  ImportFeeds().then((result) => store.commit("addFeeds", result));
};

const fetchUpdatesAction = async (): Promise<void> => {
  state.menuItems[3].loading = true;
  try {
    await FetchUpdatesForAllFeeds();
  } catch (error) {
    snackbarText.value = "Error:" + error;
    snackbar.value = true;
    console.log(error);
  } finally {
    snackbarText.value = "FetchUpdate Success!";
    snackbar.value = true;
    state.menuItems[3].loading = false;
  }
};
const menuItems: item[] = [
  {
    name: "Add New RSS Feed",
    icon: "mdi-book-plus",
    action: addAction,
    loading: false,
  },
  {
    name: "Export Feeds",
    icon: "mdi-export-variant",
    action: exportAction,
    loading: false,
  },
  {
    name: "Import Feeds",
    icon: "mdi-file-import",
    action: importAction,
    loading: false,
  },
  {
    name: "Fetch Updates",
    icon: "mdi-reload",
    action: fetchUpdatesAction,
    loading: false,
  },
];

const state = reactive({
  menuItems,
});
const output = (something: any) => {
  console.log(something);
};
</script>

<style scoped lang="less">
.menu {
  display: flex;
  height: 42px;
  border-style: solid;
  border-color: rgb(223, 237, 237);
  border-top-width: 0px;
  border-bottom-width: 1px;
  border-left: 0px;
  border-right: 0px;
  align-items: center;
  justify-content: space-between;
  .menu-left {
    float: left;
    width: calc(20%-5px);
  }
  .menu-main {
    .menu-main-item {
      width: 40px;
      v-btn {
        cursor: pointer;
      }
    }
  }
  .menu-right {
    float: right;
    width: calc(20%-5px);
  }
}
</style>
