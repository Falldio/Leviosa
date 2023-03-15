<template>
  <div class="menu">
    <div class="menu-left">
      <v-btn
        class="ma-0"
        variant="text"
        icon="mdi-text-search"
        size="small"
        color="red-lighten-2"
      ></v-btn>
    </div>
    <div v-for="(item, index) in mainitems" class="menu-main">
      <div class="menu-main-item">
        <v-tooltip :text="item.name" location="top">
          <template v-slot:activator="{ props }">
            <v-btn
              variant="text"
              v-bind="props"
              :icon="item.icon"
              size="small"
              color="red-lighten-2"
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
</template>

<script lang="ts" setup>
import { useStore } from "vuex";
// import { ExportFeeds, ImportFeeds } from "../../../wailsjs/go/main/App";
const store = useStore();
type btnFunction = () => void;

interface item {
  name: string;
  icon: string;
  action: btnFunction;
}
const mainitems: item[] = [
  // { name: "Add New RSS Feed", icon: "mdi-book-plus", action: addAction },
  // { name: "Export Feeds", icon: "mdi-export-variant", action: exportAction },
  // { name: "Import Feeds", icon: "mdi-file-import", action: importAction },
];

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
