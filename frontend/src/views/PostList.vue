<template>
    <v-list>
        <post-entry v-for="post in posts" :key="post.id" :post="post"></post-entry>
    </v-list>
</template>

<script setup lang="ts">
import PostEntry from '../components/PostEntry.vue'
import { useRoute } from 'vue-router'
import { ref, watch } from 'vue'
import { db } from '../../wailsjs/go/models'
import { GetPosts } from '../../wailsjs/go/main/App'

const route = useRoute()
const posts = ref([] as db.Post[])

GetPosts(+route.params.feedId).then((result) => {
    posts.value = result
})

watch(
    () => route.params.feedId,
    (feedId) => {
        GetPosts(+feedId).then((result) => {
            posts.value = result
        })
    }
)
</script>

<style scoped></style>