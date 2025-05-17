<template>
  <AdminNavbar>
    <slot>
      <div class="text-center">
        <h1>Nuevo Usuario</h1>
      </div>
      <v-form @submit.prevent="addUser">
        <div class="input-pair">
          <v-text-field 
            label="Nombre" 
            v-model="nombre"
            :error-messages="nombreErrors"
            required
          ></v-text-field>
          <v-text-field 
            label="Apellido" 
            v-model="lastname"
            :error-messages="lastnameErrors"
            required
          ></v-text-field>
        </div>

        <div class="input-pair">
          <v-text-field 
            label="Usuario" 
            v-model="usuario"
            :error-messages="usuarioErrors"
            required
          ></v-text-field>
          
          <v-combobox 
            label="Rol"
            v-model="rol"
            :items="['Administrador', 'Cliente', 'Dependiente', 'Soporte']"
            :error-messages="rolErrors"
            required
          ></v-combobox>
        </div>
        
        <v-text-field 
          label="Correo" 
          v-model="correo"
          type="email"
          :error-messages="correoErrors"
          required
        ></v-text-field>
        
        <div class="input-pair">
          <v-text-field 
            label="Contraseña" 
            v-model="contraseña"
            type="password"
            :error-messages="contraseñaErrors"
            required
          ></v-text-field>
          <v-text-field 
            label="Confirmar Contraseña" 
            v-model="contraseñaConfirmacion"
            type="password"
            :error-messages="contraseñaConfirmacionErrors"
            required
          ></v-text-field>
        </div>
       
        <div class="text-center">
          <v-btn 
            color="primary" 
            type="submit"
            :loading="loading"
          >
            {{ loading ? 'Procesando...' : 'Agregar' }}
          </v-btn>
        </div>
      </v-form>

      <v-dialog v-model="dialog" width="auto">
        <v-card max-width="400" :title="ErrorTitle" :text="errorText">
          <template v-slot:actions>
            <v-btn class="ms-auto" text="Ok" @click="dialog = false"></v-btn>
          </template>
        </v-card>
      </v-dialog>
    </slot>
  </AdminNavbar>
</template>

<script>
import AdminNavbar from "../../components/menus/AdminNavbar.vue";
import { z } from 'zod';
import { userSchema } from '@/plugins/validationSchemas';
import axios from 'axios';

export default {
  data() {
    return {
      nombre: "",
      lastname: "",
      usuario: "",
      correo: "",
      rol: "",
      contraseña: "",
      contraseñaConfirmacion: "",
      dialog: false,
      errorText: '',
      ErrorTitle: '',
      loading: false
    };
  },
  components: { 
    AdminNavbar 
  },
  methods: {
    async addUser() {
      if (!this.validar()) {
        return;
      }

      this.loading = true;
      
      try {
        const response = await axios.post('/api/users', {
          nombre: this.nombre,
          lastname: this.lastname,
          usuario: this.usuario,
          correo: this.correo,
          rol: this.rol,
          contraseña: this.contraseña
        });

        // Éxito - redirigir o limpiar formulario
        this.$router.push('/admin/users');
        this.$toast.success('Usuario creado exitosamente');
        
      } catch (error) {
        this.dialog = true;
        this.errorText = error.response?.data?.message || 'Error al crear el usuario';
        this.ErrorTitle = 'Error del servidor';
        console.error(error);
      } finally {
        this.loading = false;
      }
    },
    validar() {
      try {
        userSchema.parse({
          nombre: this.nombre,
          lastname: this.lastname,
          usuario: this.usuario,
          correo: this.correo,
          rol: this.rol,
          contraseña: this.contraseña,
          contraseñaConfirmacion: this.contraseñaConfirmacion
        });
        
        return true;
        
      } catch (error) {
        if (error instanceof z.ZodError) {
          const firstError = error.errors[0];
          this.dialog = true;
          this.errorText = firstError.message;
          
          // Tipos de error comunes
          if (firstError.code === 'too_small') {
            this.ErrorTitle = 'Longitud mínima';
          } else if (firstError.code === 'too_big') {
            this.ErrorTitle = 'Longitud máxima';
          } else if (firstError.code === 'invalid_string') {
            this.ErrorTitle = 'Formato incorrecto';
          } else if (firstError.path.includes('contraseñaConfirmacion')) {
            this.ErrorTitle = 'Error de contraseña';
          } else {
            this.ErrorTitle = 'Error de validación';
          }
        }
        return false;
      }
    }
  }
}
</script>

<style>
.input-pair {
  display: flex;
  justify-content: space-between;
  gap: 16px;
}

.input-pair > * {
  flex: 1;
}
</style>