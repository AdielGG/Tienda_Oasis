<template>
  <nav class="main-navbar">
    
      
    <section class="main-bar-buttons-container">

      <v-btn icon="mdi-facebook" color="primary" variant="text"></v-btn>
      <v-btn icon="mdi-twitter" color="primary" variant="text"></v-btn>
      <v-btn icon="mdi-instagram" color="primary" variant="text"></v-btn>
      <LoginBottom 
        v-if="this.$router.currentRoute.value.fullPath !== '/login'  
        && this.$router.currentRoute.value.fullPath !== '/register' 
        && !this.$store.state.logged"/>
      <UserBottom v-if="this.$store.state.logged"/>
      <UserMenu/>
    </section>  
  </nav>
</template>

<script>
import LoginBottom from './LoginBottom.vue';
import UserBottom from './UserBottom.vue';
import UserMenu from './UserMenu.vue';

export default {
  
  components:{
    LoginBottom,
    UserBottom,
    UserMenu
  },
  mounted(){
    console.log(this.path)
  }

}
</script>

<style>
.main-navbar {
  display: flex;
    width: 100%;

    background-color: #f5f5f5;
    flex-direction: row;
    flex-wrap: wrap;
    align-content: space-around;
    justify-content: flex-end;
    height: 3em;
}
.main-bar-buttons-container{
  display:flex;

}
</style>