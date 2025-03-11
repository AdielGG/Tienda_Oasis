<template>
    <RouterView />
</template>

<script>
import NavBar from './components/menus/NavBar.vue';
  export default {
    data(){
      return {

      }
    },
    components:[
      NavBar,
    ]
    
  }
</script>
<style scoped>
  .header{
    width: 100%;
    display: inline-flex;
    align-items: center;
    position: relative;
    /* align-content: center; */
    /* display: inline-block; */
    flex-direction: column;
    flex-wrap: wrap;
  }
  body{
    display: block;
  }
  .main-container{
    display: block;
    width: auto;
    height: auto;
    align-content: center;
    align-items: center;
    
  }
  .title{
    font-family: 'Courier New', Courier, monospace;
  }
</style>