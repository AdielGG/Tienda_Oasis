<script setup>
import NavBar from '@/components/menus/NavBar.vue';
</script>

<template>
  <main>
    <video src="../assets/video.mp4" autoplay="true" muted="true" loop="true" ></video>
    
    <header class="header">

      <h1> <b>Tienda Oasis</b> </h1>
    </header>
    
    <NavBar/>


  </main>
</template>


<style scoped>
 
  .header{
    background-color: rgba(13, 20, 20, 0);
    width: 100%;
    text-align: center;
    

  }
  .header h1{
    
    color: rgb(255, 238, 0);
    font-size: 4em;
    font-family: 'Franklin Gothic Medium', 'Arial Narrow', Arial, sans-serif;
  }

  main{
    width: 100%;
    height: 100%;
    display: block;
  }
  
  video {
  position: fixed;
  right: 0;
  bottom: 0;
  min-width: 105%;
  min-height: 100%;
  transform: translateX(calc((100% - 100vw) / 2));
  z-index: -2;
}
</style>