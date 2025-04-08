<template>
    <AdminNavBar>
        <slot>
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
        this.getUsers();
    },
    components: {
        AdminNavBar,
        UsersTable,
    },
    methosd: {
        getUsers() {
            this.users = async () => {
                return await axios.get("localhost:8080/users");
            };
        },
    },
};
</script>
