<template>
    <v-menu location="end">
        <template v-slot:activator="{ props }">
            <v-list-item density="compact" v-bind="props" prepend-icon="mdi-cog" title="Settings" value="settings"></v-list-item>
        </template>
        <v-list>
            <v-list-item v-for="(item, index) in items" :key="index" :value="index" @click-once="handleClick(index)">
                <v-list-item-title>{{ item.name }}</v-list-item-title>
            </v-list-item>
        </v-list>
    </v-menu>
    <add-r-s-s-feed-dialog v-model="store.state.addRSSFeedDialog"></add-r-s-s-feed-dialog>
</template>

<script lang="ts" setup>
import AddRSSFeedDialog from "./AddRSSFeedDialog.vue"
import { useStore } from 'vuex'
import { ExportFeeds, ImportFeeds } from "../../wailsjs/go/main/App"

const store = useStore()

const items = [
    { name: "Add New RSS Feed" },
    { name: "Export Feeds" },
    { name: "Import Feeds" }
] as { name: string }[]

const handleClick = (index: number) => {
    switch (index) {
        case 0:
            store.commit("setAddRSSFeedDialog", true)
            break
        case 1:
            ExportFeeds()
            break
        case 2:
            ImportFeeds().then((result) =>
                store.commit("addFeeds", result)
            )
            break
    }
}
</script>

<style scoped></style>