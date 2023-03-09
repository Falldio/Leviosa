import { createRouter, createMemoryHistory } from 'vue-router'

const routes = [
    { path: '/', component: () => import('./views/Home.vue')},
    { path: '/feeds/:feedId', component: () => import('./views/PostList.vue')},
    { path: '/feeds/:feedId/posts/:postId', component: () => import('./views/PostDetail.vue')},
] as { path: string, component: any}[]

export const router = createRouter({
    history: createMemoryHistory(),
    routes: routes
})