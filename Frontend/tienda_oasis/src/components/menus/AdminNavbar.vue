<template>
    <v-card>
        <v-layout>
            <v-navigation-drawer
                class="bg-blue-accent-4"
                theme="dark"
                permanent
            >
                <v-list color="transparent">
                    <v-list-item
                        prepend-icon="mdi-view-dashboard"
                        title="Dashboard"
                    ></v-list-item>
                    <v-divider></v-divider>
                    <v-list-item
                        link
                        @click="adminUsers"
                        prepend-icon="mdi-account-group"
                        title="Usuarios"
                    ></v-list-item>
                    <v-list-item
                        link
                        @click="adminProducts"
                        prepend-icon="mdi-package-variant-closed"
                        title="Productos"
                    ></v-list-item>
                    <v-list-item
                        link
                        @click="adminSugestions"
                        prepend-icon="mdi-message-text-outline"
                        title="Sugerencias"
                    ></v-list-item>
                </v-list>

                <template v-slot:append>
                    <div class="pa-2">
                        <v-btn block class="bt-store"> Tienda </v-btn>
                        <v-btn block> Logout </v-btn>
                    </div>
                </template>
            </v-navigation-drawer>
            <v-main>
                <div class="title-bar">
                    <v-breadcrumbs :items="items">
                        <template v-slot:divider>
                            <v-icon icon="mdi-forward" size="x-small"></v-icon>
                        </template>
                    </v-breadcrumbs>
                </div>
                <slot> </slot>
            </v-main>
        </v-layout>
    </v-card>
</template>

<script>
export default {
    data() {
        return {
            items: [
                {
                    title: "Panel de Administración",
                    disabled: false,
                    href: "#",
                },
            ],
        };
    },
    methods: {
        adminProducts() {
            this.$router.push("/admin/products");
        },
        adminUsers() {
            this.$router.push("/admin/users");
        },
        
        adminSugestions() {
            this.$router.push("/admin/sugerencias");
        },
    },
};
</script>

<style>
.title-bar {
    display: flex;
    background: #bbdefb;
    align-items: left;
    font-size: 2em;
    font-weight: bold;
    margin: 0;
    padding: 0;
}
.title-bar .v-breadcrumbs {
    margin: 0;
    padding: 0;
}
.bt-store {
    color: black;
    background-color: #fff;
    margin-bottom: 0.5em;
}

</style>
