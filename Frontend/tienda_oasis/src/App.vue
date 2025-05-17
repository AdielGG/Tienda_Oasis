<template>
  <MainBar/>
  <RouterView/>
</template>

<script>
import MainBar from './components/menus/MainBar.vue';

export default {
    data(){
      return {

      }
    },
    components:{
      MainBar,
    },
    
    
}
</script>
<style >
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
    width: 100%;
    height: 100%;
    border: none;
    margin: 0;
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