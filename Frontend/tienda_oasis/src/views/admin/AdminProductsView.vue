<template>
    <AdminNavbar>
        <slot> Hellos </slot>
    </AdminNavbar>
</template>
<script>
import AdminNavbar from "@/components/menus/AdminNavbar.vue";

export default {
    name: "AdminProductsView",
    data() {
        return {
            products: [],
        };
    },
    components: {
        AdminNavbar,
    },
    methods: {},
    created() {},
};
</script>
