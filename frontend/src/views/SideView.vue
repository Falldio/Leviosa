<template>
  <primary-sidebar></primary-sidebar>
</template>

<script lang="ts" setup>
import PrimarySidebar from "../components/Sidebar/PrimarySidebar.vue";
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
