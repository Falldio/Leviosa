<script lang="ts" setup>
import SidebarMenu from './components/SidebarMenu.vue'
import { useStore } from 'vuex'
import { GetFeeds } from '../wailsjs/go/main/App'
import { useRouter, useRoute } from 'vue-router'
import { onUnmounted } from 'vue'

const store = useStore()
const router = useRouter()
const route = useRoute()

GetFeeds().then((result) => {
  store.commit("setFeeds", result)
})

const handleBackButton = () => {
  router.back()
}

window.addEventListener('popstate', handleBackButton)

onUnmounted(() => {
  window.removeEventListener('popstate', handleBackButton)
})


</script>

<template>
  <v-app>
    <sidebar-menu />
    <v-main>
      <router-view></router-view>
    </v-main>
  </v-app>
</template>


<style>
::-webkit-scrollbar {
  width: 20px;
}

::-webkit-scrollbar-track {
  background-color: transparent;
}

::-webkit-scrollbar-thumb {
  background-color: #d6dee1;
  border-radius: 20px;
  border: 6px solid transparent;
  background-clip: content-box;

}

::-webkit-scrollbar-thumb:hover {
  background-color: #a8bbbf;
}
</style>
