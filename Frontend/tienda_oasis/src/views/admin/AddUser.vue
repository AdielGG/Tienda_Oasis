<template>
  <AdminNavbar>
    <slot>
      <div class="text-center">
        <h1>Nuevo Usuario</h1>
      </div>
      <v-form>
        <div class="input-pair">

          <v-text-field label="Nombre" v-model="nombre"></v-text-field>
          <v-text-field label="Apellido" v-model="lastname"></v-text-field>
        </div>

        <div class="input-pair">
          <v-text-field label="Usuario" v-model="usuario"></v-text-field>
          
          <v-combobox label="Rol" 
          :items="['Administrador', 'Cliente','Dependiente','Soporte']"
          ></v-combobox>
          
        </div>
        <v-text-field label="Correo" v-model="correo"></v-text-field>
        <div class="input-pair">

          <v-text-field label="Contraseña" v-model="contraseña"></v-text-field>
          <v-text-field label="Confirmar Contraseña" v-model="contraseñaConfirmacion"></v-text-field>
        </div>
       
        <div class="text-center">
          <v-btn color="primary" @click="addUser">Agregar</v-btn>
        </div>
      </v-form>
    </slot>
  </AdminNavbar>
</template>

<script>
import AdminNavbar from "../../components/menus/AdminNavbar.vue";

export default {
  data() {
    return {
      nombre: "",
      usuario: "",
      correo: "",
      contraseña: "",
      contraseñaConfirmacion: "",
    };
  },
  components: { 
    AdminNavbar 
  },
  methods: {
    addUser() {
      console.log(this.nombre);
      console.log(this.usuario);
      console.log(this.correo);
      console.log(this.contraseña);
      console.log(this.contraseñaConfirmacion);
    },
  },

}
</script>

<style>

</style>