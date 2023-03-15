<template>
  <v-navigation-drawer
    width="250"
    id="sidebar-menu"
    class="text-sm-left"
    :disable-resize-watcher="true"
    permanent
  >
    <settings-menu></settings-menu>
    <div class="bar">
      <v-list>
        <div class="bar-items" v-for="(item, index) in barItems" :key="index">
          <v-list-item
            :value="item.title"
            active-color="red-lighten-2"
            @click="output(item.title)"
            v-if="!item.children"
          >
            <template v-slot:prepend>
              <v-icon :icon="item.icon"></v-icon>
            </template>
            <template v-slot:append>
              <div class="bar-items-infonum">{{ item.infonum }}</div>
            </template>
            <div class="bar-items-title">{{ item.title }}</div>
          </v-list-item>
          <v-list-group
            v-if="item.children"
            :value="item.title"
            active-color="red-lighten-2"
            class="bar-items-group"
          >
            <template v-slot:activator="{ props }">
              <v-list-item
                v-bind="props"
                active-color="red-lighten-2"
                class="bar-items-group-title"
              >
                <template v-slot:prepend>
                  <v-icon :icon="item.icon"></v-icon>
                </template>
                <div class="bar-items-title">{{ item.title }}</div>
              </v-list-item>
            </template>
            <div v-for="(jitem, index) in item.children">
              <v-list-item
                :value="jitem.title"
                active-color="red-lighten-2"
                @click="output(jitem)"
                class="bar-items-group-jitem"
              >
                <!-- <template v-slot:prepend>
                  <v-icon :icon="jitem.icon"></v-icon>
                </template> -->
                <template v-slot:append>
                  <div class="bar-items-infonum">{{ jitem.infonum }}</div>
                </template>
                <div class="bar-items-title">{{ jitem.title }}</div>
              </v-list-item>
            </div>
          </v-list-group>
        </div>
      </v-list>
    </div>
  </v-navigation-drawer>
</template>

<script setup lang="ts">
import SettingsMenu from "./SettingsMenu.vue";
import { reactive } from "vue";
interface item {
  title: string;
  icon: string;
  infonum?: number;
  children?: item[];
}
const barItems: item[] = reactive([
  { title: "Unread", icon: "mdi-book-open-page-variant-outline", infonum: 15 },
  { title: "Stars", icon: "mdi-star-box" },
  // { title: "Comics", icon: "mdi-thought-bubble" },
  {
    title: "Tags",
    icon: "mdi-tag",
    children: [
      { title: "Youtube", icon: "mdi-youtube" },
      { title: "Bilibili", icon: "mdi-video-account" },
    ],
  },
]);

const output = (something: any) => {
  console.log(something);
};
</script>

<style scoped lang="less">
.bar {
  .bar-items {
    .bar-items-infonum {
      font-size: small;
      color: rgb(181, 181, 181);
    }
    .bar-items-title {
      font-size: 13px;
      font-weight: 400;
      align-items: left;
    }
    .bar-items-group {
      .bar-items-group-title {
        font-size: 13px;
      }
    }
  }
}
</style>
