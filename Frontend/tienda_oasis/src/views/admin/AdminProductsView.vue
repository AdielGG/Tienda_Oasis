<template>
    <AdminNavbar>
        <slot> 
            <div class="text-center">
                <h1>Productos</h1>
            </div>
            <v-fab
                key="absolute"
                app
                color='primary'
                location="bottom center"
                size="large"
                @click="AddProduct"
            >
                <v-icon>mdi-plus</v-icon>
                <AddProduct></AddProduct>
            </v-fab>
            <ProductTable> </ProductTable>
            
        </slot>
    </AdminNavbar>
</template>
<script>
import AdminNavbar from "@/components/menus/AdminNavbar.vue";
import ProductTable from "@/components/ProductsTable.vue";
import AddProduct from "./AddProduct.vue";

export default {
    name: "AdminProductsView",
    data() {
        return {
            products: [],
        };
    },
    components: {
        AdminNavbar,
        ProductTable,
        AddProduct
    },
    methods: {
        AddProduct() {
            this.$router.push("/admin/products/add");
        },
    },
    created() {},
};
</script>
<style></style>
