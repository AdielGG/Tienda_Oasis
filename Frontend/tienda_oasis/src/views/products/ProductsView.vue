<template>
  <main>
    <nav-bar ></nav-bar>
    <ProductTable></ProductTable>
  </main>
</template>

<script>
import NavBar from '../../components/menus/NavBar.vue'
import ProductTable from '@/views/products/ProductTable.vue';
export default {
  components: { NavBar,ProductTable },

}
</script>

<style>
    body{
        background-color: rgb(156, 180, 180) !important;
    }
  
    main{
        
    }



</style>