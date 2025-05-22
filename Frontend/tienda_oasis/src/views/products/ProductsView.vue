<template>
    <header class="tm-site-header">
       
        <NavBar></NavBar>
    </header> 
    <ProductTable></ProductTable>
    
</template>

<script setup>
import NavBar from '../../components/menus/NavBar.vue'
import ProductTable from '@/views/products/ProductTable.vue';
</script>

<style>

    body{
        background-color: rgb(255, 255, 255) !important;
    }
  
</style>