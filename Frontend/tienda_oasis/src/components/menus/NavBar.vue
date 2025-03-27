<template>
    <nav>
        <ul class="nav-bar">
            <li >
                <a href="/" class="nav-item">Inicio</a>
            </li>
            <li>
                <a href="/products" class="nav-item">Productos</a>
            </li>
            <li>
                <a href="/about" class="nav-item">Nosotros</a>
            </li>
            <li>
                <a href="" class="nav-item">Ayuda</a>
            </li>
            <v-spacer></v-spacer>
            <LoginBottom v-if="!loged">

            </LoginBottom>

            <button 
                v-else
                class="user-btn"
                @click="drawer = !drawer">
                <img src="../assets/Icons/user.png" alt="user" class="icon-img">
                <span class="user-name">{{this.userd.name }}</span>
            </button>
            
                
        </ul>
    </nav>
    
  <v-card class="user-menu">
      <v-layout>
        <v-navigation-drawer
          v-model="drawer"
          temporary
          location="end"
        >
          <v-list-item
            @click="drawer = !drawer"
            
          >
          <img src="../assets/Icons/user.png" alt="user" class="icon-img">
          <span class="user-name">{{ this.userd.name }}</span>
        </v-list-item>
  
          <v-divider></v-divider>
  
          <v-list density="compact" nav>
            <v-list-item 
              prepend-icon="mdi-account" 
              title="Mi Cuenta" 
              value="acount" 
              @click="() => console.log(user)">
            </v-list-item>
            <!-- <v-list-item 
            v-if="user.rol === 'admin'"
            prepend-icon="mdi-forum" 
            title="Panel de Admin" 
            value="admon" 
            @click="() => this.$router.push('/admin')">
          </v-list-item> -->
            <v-list-item 
            prepend-icon="mdi-logout" 
            title="Cerrar Sesion" 
            value="logout" 
            @click="logout">
          </v-list-item>
          <v-divider></v-divider>
          <v-list-item 
            prepend-icon="mdi-logout" 
            title="Ocultar" 
            value="hidden" 
            @click="drawer = !drawer">
          </v-list-item>
          <v-divider></v-divider>
          </v-list>
        </v-navigation-drawer>
      </v-layout>
    </v-card>
</template>

<script>
import store from '@/store/store';
import LoginBottom from './LoginBottom.vue';
export default{
    data(){
        return {
            drawer: store.getters.getDrawer,
            rail: false,
            loged: store.state.logued,
            userd: {
                name: 'asd',
                lastname: '',
                id: '',
                username: '',
                password: '',
                age: '',
                email: '',
            },
        }
    },
    methods:{
        logout(){
            store.mutations.setUser(store.state,null)
            store.mutations.set
        },
      },
    components: {
        LoginBottom,
    },
    computed: {
        user(){
            return store.state.user;
        }
    },  
    updated() {
      this.user = store.state.user;
      console.log(this.user);
    },
    beforeMount() {
      this.user = store.state.user;
      console.log(this.user);

    }
  }


</script>

<style scoped>
*{
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

.nav-bar{
    
    box-sizing: border-box;
    border-radius: 10px;
    display: grid;
    grid-template-columns: 1fr 1fr 1fr 1fr 1fr 1fr;
    background-color: rgba(6, 184, 255, 0.829);
    width: 100%;
    height: 10%;
    justify-content: space-between;
}

li {
    margin: 10px;
    list-style: none;
    text-align: center;
  }


.nav-item{
    text-align: center;
    align-items: center;
    width: 10%;
    box-sizing: border-box;
    color: rgb(255, 255, 255);
    font-size: 2em;
    text-shadow: 0 0 8px black;
    font-family:Georgia, 'Times New Roman', Times, serif;
}

.nav-item:hover{
    box-sizing: border-box;
    transition-duration: 0.2s;
    font-size: 2.1em;
    background-color: rgba(0, 0, 0, 0);
    border-bottom: solid 1px rgb(255, 255, 255);
}
.user-btn{
  background-color: rgba(68, 198, 230, 0);
  border-radius: 1em;
  color: black;
  font-size: 1.2em;
  padding-top: 0.2em;
  transition: 0.3s;
  line-height: 0;
  padding-bottom: 0.1em;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.5);
}
.user-btn:hover{
  transition-duration: 0.5s;
  background-color: rgb(177, 231, 166);
  border: 1px solid rgb(0, 0, 0);
}
.icon-img{
  float: left;
  margin-top: 0.01em;
  
}
.user-name{
  display: inline-block;
  margin-top: 0.9em;
  margin: 0.9em 0.3em 0.1em 0.3em;
}

.v-list-item__content .user-name{
  margin: 0.1em 0.3em 0.1em 0.3em;
  padding: 0.4em;
}
.router-link-active{
  color: red;
}
.router-link-active:hover{
  color:rgb(255, 0, 0);
}
.user-menu{
  background-color: white !important;
  position: absolute;
  right: 0;
  top: 0;
  width: 300px;
}
.icon{
  float: left;
}
h1{
  font-family: 'Franklin Gothic Medium', 'Arial Narrow', Arial, sans-serif;
  text-align: center;
  font-size: 2em;
  color: red;
  text-shadow: 0 0 10px black;
  background:rgb(220, 225, 243);
  

}


</style>
