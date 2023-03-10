<template>
    <primary-sidebar></primary-sidebar>
    <v-navigation-drawer permanent>
        <v-list lines="two">
            <feed-entry v-for="item in items" :key="item.id" :feed="item" />
        </v-list>
    </v-navigation-drawer>
</template>

<script lang='ts' setup>
import PrimarySidebar from "./PrimarySidebar.vue"
import FeedEntry from "./FeedEntry.vue"
import { computed } from 'vue'
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