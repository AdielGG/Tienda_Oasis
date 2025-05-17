<template>

    <v-card class="user-menu">
        <v-layout>
          <v-navigation-drawer
            v-model="drawer"
            temporary
            location="end"
          >
            <v-list-item
              
              style="text-align: center;"
            >
         
            <span class="user-name">{{ $store.state.user.username }}</span>
         
          </v-list-item>
    
            <v-divider></v-divider>
    
            <v-list density="compact" nav>
         
              <li>
                <button
                >
                  <img src="../../assets/icons/account-edit.svg" width="22">
                  <span> Editar Perfil </span>
                </button>
                
              </li>
              
              <li v-if="$store.state.user.role === 'admin'">
                <button
                >
                  <img src="../../assets/icons/shield-crown.svg" width="22">
                  <span> Admin Panel </span>
                </button>
                
              </li>
              <li>
                <button
                  @click="logout"
                >
                  <img src="../../assets/icons/logout.svg" color="blue" width="22">
                  <span> Cerrar Sesion </span>
                </button>
                
              </li>
              <v-divider></v-divider>
          
              <li>
                <button
                  @click="Drawer"
                >
                  <img src="../../assets/icons/arrow-right-circle-outline.svg"  width="22">
                  <span> Ocultar</span>
                </button>
                
              </li>
              
              <v-divider></v-divider>
          
            </v-list>
          </v-navigation-drawer>
        </v-layout>
      </v-card>
        

</template>

<script>
import defaultValues from '@/store/defaultValues'
import store from '@/store/store'

export default {
    
    methods:{
        logout(){
            this.$store.commit('setDrawer', false)
            this.$store.commit('setLogged', false)
            this.$store.commit('setUser', defaultValues.userDefault)
            this.$router.push('/')
        
        },
        Drawer(){
            this.$store.commit('setDrawer', false)

        },
        administrar(){
          this.$store.commit('setDrawer', false)
          this.$router.push('/admin')
        }
    },
    computed: {
      drawer(){
        return this.$store.state.drawer;
      },
    },
      
    updated() {
      this.user = store.state.user;
    },

    beforeMount() {
      this.user = store.state.user;

    }
  }     


</script>

<style>
.user-menu{
  position: relative;
  z-index: 99;
}

.user-menu li{
  list-style: none;
  
}
.user-menu button{
  width: 100%;
  height: 100%;
  text-align: left;
}
.user-menu button:hover{
  display: block;
  background-color: #f5f5f5;
  border-radius: 0.2em;
}
.user-menu button span{
  vertical-align: middle;
  
}
.user-menu button img{
  border: none;
  vertical-align: middle;
  margin: 0.5em ;
}
</style>