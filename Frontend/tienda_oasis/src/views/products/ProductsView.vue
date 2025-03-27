<template>
    
    <NavBar/>
    <ProductTable></ProductTable>

</template>

<script setup>
import NavBar from '../../components/menus/NavBar.vue'
import ProductTable from '@/views/products/ProductTable.vue';
</script>

<style>

    body{
        background-color: rgb(156, 180, 180) !important;
    }
  
</style>