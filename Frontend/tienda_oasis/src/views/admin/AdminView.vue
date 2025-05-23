<template>
    <AdminNavbar>
        <slot> </slot>
    </AdminNavbar>
</template>

<script>
import AdminNavbar from "@/components/menus/AdminNavbar.vue";
export default {
    components: {
        AdminNavbar,
    },
};
</script>

<style></style>
