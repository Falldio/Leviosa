<template>
    <v-navigation-drawer width="200" id="sidebar-menu" class="text-sm-left" :disable-resize-watcher="true"
        permanent>
        <v-list density="compact" nav>
            <v-list-item prepend-icon="mdi-layers" title="Unread"></v-list-item>
            <v-list-item prepend-icon="mdi-star-box" title="Stars"></v-list-item>
            <v-list-item prepend-icon="mdi-reload" title="Fetch Updates"></v-list-item>
            <v-list-group value="Tags">
                <template v-slot:activator="{props}">
                    <v-list-item title="Tags"></v-list-item>
                </template>
                <v-list-item title="Golang"></v-list-item>
            </v-list-group>
        </v-list>
        <template v-slot:append>
            <v-list density="compact" nav>
                <settings-menu></settings-menu>
            </v-list>
        </template>
    </v-navigation-drawer>
</template>

<script setup lang="ts">
import SettingsMenu from './SettingsMenu.vue'
</script>

<style scoped>

</style>