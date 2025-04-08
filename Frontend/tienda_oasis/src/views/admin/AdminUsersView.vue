<template>
    <AdminNavBar>
        <slot>
            <div class="text-center">
                <h1>Usuarios</h1>
            </div>
            <v-fab
                key="absolute"
                app
                color='primary'
                location="bottom center"
                size="large"
                @click="AddUser"
            >
                <v-icon>mdi-plus</v-icon>
                Nuevo Usuario
            </v-fab>
            <UsersTable> </UsersTable>
        </slot>
    </AdminNavBar>
</template>

<script>
import axios from "axios";
import UsersTable from "@/components/UsersTable.vue";
import AdminNavBar from "@/components/menus/AdminNavbar.vue";
export default {
    data() {
        return {
            users: null,
        };
    },
    mounted() {
        // this.getUsers();
    },
    components: {
        AdminNavBar,
        UsersTable,
    },
    methods: {
        
        AddUser() {
            this.$router.push("/admin/users/add");
        },
    },
};
</script>
