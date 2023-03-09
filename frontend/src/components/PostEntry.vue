<template>
    <v-list-item @click="loadPostDetail">
        <v-list-item-title class="left-align">{{ post.title }}</v-list-item-title>
        <v-list-item-subtitle class="left-align" v-html="post.description"></v-list-item-subtitle>
    </v-list-item>
</template>

<script setup lang="ts">
import { useRouter, useRoute } from 'vue-router'
const router = useRouter()
const route = useRoute()

const props = defineProps(['post'])

const loadPostDetail = () => {
    router.push('/feeds/' + route.params.feedId + '/posts/' + props.post.id)
}

</script>

<style scoped>

</style>