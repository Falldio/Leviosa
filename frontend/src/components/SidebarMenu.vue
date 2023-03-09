<template>
    <v-navigation-drawer id="sidebar-menu" v-model="drawer" :rail="rail" :disable-resize-watcher="disableResizeWatcher"
        permanent>
        <v-list density="compact" nav>
        </v-list>
        <template v-slot:append>
            <v-list density="compact" nav>
                <settings-menu></settings-menu>
            </v-list>
        </template>
    </v-navigation-drawer>
    <v-navigation-drawer permanent>
        <v-list lines="two">
            <feed-entry v-for="item in items" :key="item.id" :feed="item" />
        </v-list>
    </v-navigation-drawer>
</template>

<script lang='ts' setup>
import SettingsMenu from "./SettingsMenu.vue"
import FeedEntry from "./FeedEntry.vue"
import { computed, ref } from 'vue'
import { useStore } from 'vuex'

const store = useStore()

const items = computed(() => {
    return store.state.feeds.map((feed: any) => {
        return {
            id: feed.id,
            title: feed.title,
            description: feed.description,
            icon: feed.image
        }
    })
})

const drawer = ref(true)
const disableResizeWatcher = ref(true)
const rail = ref(true)

</script>

<style lang="scss" scoped>
#sidebar-menu {
    overflow-y: hidden;
}

.v-btn::before {
    background-color: transparent;
}
</style>