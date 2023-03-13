import { createRouter, createMemoryHistory } from "vue-router";

const routes = [
  { path: "/", component: () => import("./views/Home.vue") },
  {
    path: "/feeds/:feedId",
    component: () => import("./components/Post/PostList.vue"),
  },
  {
    path: "/feeds/:feedId/posts/:postId",
    component: () => import("./components/Post/PostDetail.vue"),
  },
] as { path: string; component: any }[];

export const router = createRouter({
  history: createMemoryHistory(),
  routes: routes,
});
