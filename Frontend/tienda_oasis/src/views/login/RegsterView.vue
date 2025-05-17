<template>
    <v-container>
      <form @submit.prevent="create">
          <v-card class="main-card">
              <div class="logo-container">
                  <img src="../../assets/oasis.png" alt="Logo"  class="logo">
              </div>
  
              <v-card-title >
                  
                  <v-row>
                      <v-col cols="12">
                          <v-text-field
                              v-model="name"
                              label="Nombre"
                              required
                          ></v-text-field>
                      </v-col>
                  </v-row>
                  <v-row>
                      <v-col cols="12">
                          <v-text-field
                              v-model="lastname"
                              label="Apellido"
                              required
                          ></v-text-field>
                      </v-col>
                  </v-row>
                  <v-row>
                      <v-col cols="12">
                          <v-text-field
                              v-model="age"
                              label="Edad"
                              type="number"
                              required
                          ></v-text-field>
                      </v-col>
                  </v-row>
                  <v-row>
                      <v-col cols="12">
                          <v-text-field
                              v-model="email"
                              label="Correo Electronico"
                              type="email"
                              required
                          ></v-text-field>
                      </v-col>
                  </v-row>
                  <v-row>
                      <v-col cols="12">
                          <v-text-field
                              v-model="username"
                              label="Usuario"
                              required
                          ></v-text-field>
                      </v-col>
                  </v-row>
              </v-card-title>
              <v-card-title>
                  <v-row>
                      <v-col cols="12">
                          <v-text-field
                              v-model="password"
                              label="Contraseña"
                              type="password"
                              required
                              
                          ></v-text-field>
                      </v-col>
                  </v-row>
                  <v-row>
                      <v-col cols="12">
                          <v-text-field
                              v-model="cpassword"
                              label="ConfirmarContraseña"
                              type="password"
                              required
                          ></v-text-field>
                      </v-col>
                  </v-row>
              </v-card-title>
              <section class="actions-container">
                  <button class="btn btn-primary" @click="create" :disabled="loading">
                      <span v-if="loading">Cargando...</span>
                      <span v-else>Entrar</span>
  
                  </button>
                  <button class="btn btn-secondary" @click="() => { router.push('/login')}" :disabled="loading">
                      <span>Registrarse</span>
  
                  </button>                
                  <v-dialog
                      v-model="dialog"
                      width="auto"
                      >
                      <v-card
                          max-width="400"
                          prepend-icon="mdi-update"
                          :text="errorText"
                          :title="ErrorTitle"
                      >
                          <template v-slot:actions>
                          <v-btn
                              class="ms-auto"
                              text="Ok"
                              @click="dialog = false"
                          ></v-btn>
                          </template>
                      </v-card>
                      </v-dialog>
                  
              </section>
          </v-card>
      </form>
    </v-container>
  </template>
  
  <script>
import axios from 'axios';
import { z } from 'zod';
import { registerSchema } from '@/plugins/validationSchemas';

export default {
    data() {
        return {
            username: '',
            password: '',
            cpassword: '',
            name: '',
            lastname: '',
            age: '',
            email: '',
            loading: false,
            dialog: false,
            errorText: '',
            ErrorTitle: ''
        }
    },
    methods: {
        async create() {
            this.loading = true;
            
            if(!this.validar()) {
                this.loading = false;
                return false;
            }

            try {
                const response = await axios.post('http://localhost:8080/register', {
                    name: this.name,
                    lastname: this.lastname,
                    age: this.age.toString(),
                    email: this.email,
                    username: this.username,
                    password: this.password
                });

                this.loading = false;
                console.log(response.data);
                
                if(response.data.OK === 'Logued in successfully') {
                    this.$store.commit('setUser', response.data.user);
                    this.$store.commit('setLogged', true);
                    
                    if(this.$store.state.user.role === 'admin') {
                        this.$router.push('/admin');
                    } else {
                        this.$router.push('/');
                    }
                }
                
                return true;
                
            } catch (error) {
                console.log(error);
                this.loading = false;
                this.dialog = true;
                
                this.errorText = error.message === 'Network Error' 
                                ? 'No se pudo conectar con el servidor' 
                                : error.response?.data?.message || 'Error al registrar usuario';
                
                this.ErrorTitle = error.name === "AxiosError" 
                                ? 'Error del Servidor' 
                                : 'Error';
                
                return false;
            }
        },
        validar() {
            try {
                // Convertir age a número para la validación
                const ageNumber = parseInt(this.age);
                
                registerSchema.parse({
                    name: this.name,
                    lastname: this.lastname,
                    age: isNaN(ageNumber) ? 0 : ageNumber, // Manejar caso no numérico
                    email: this.email,
                    username: this.username,
                    password: this.password,
                    cpassword: this.cpassword
                });
                
                return true;
                
            } catch (error) {
                if (error instanceof z.ZodError) {
                    // Manejar el primer error encontrado
                    const firstError = error.errors[0];
                    
                    this.dialog = true;
                    this.errorText = firstError.message;
                    
                    // Asignar título según el tipo de error
                    if (firstError.code === 'too_small') {
                        this.ErrorTitle = 'Longitud mínimo';
                    } else if (firstError.code === 'too_big') {
                        this.ErrorTitle = 'Longitud máximo';
                    } else if (firstError.code === 'invalid_string') {
                        this.ErrorTitle = 'Caracteres no permitidos';
                    } else if (firstError.code === 'invalid_type') {
                        this.ErrorTitle = 'Tipo de dato incorrecto';
                    } else if (firstError.path.includes('cpassword')) {
                        this.ErrorTitle = 'Contraseñas no coinciden';
                    } else {
                        this.ErrorTitle = 'Error de validación';
                    }
                } else {
                    // Otros tipos de errores
                    this.dialog = true;
                    this.errorText = 'Error desconocido';
                    this.ErrorTitle = 'Error';
                }
                
                return false;
            }
        }
    }
}
</script>
  
  <style scoped>
  .logo-container{
      display: flex;
      justify-content: center;
      margin-bottom: 5em;
  
  }
  
  
  .logo{
      float: right;
      width: 50%;
  
  }
  
  .title{
      font-family: 'Trebuchet MS', 'Lucida Sans Unicode', 'Lucida Grande', 'Lucida Sans', Arial, sans-serif;
      font-size: 2em;
      color:#2766aa;
  
  }
  
  .main-card {
      display:block;
      margin: auto;
      width: 20em;
      height: 50%;
      border-radius: 1em;
      box-shadow: 0 0 10px rgba(0, 0, 0, 0.5);
  }
  
  .main-card .v-card__title {
      width: 50%;
      margin: auto;
  }
  
  .actions-container {
      display: block;
      justify-content: center;
      align-items: center;
      margin-top: 1em;
  }
  
  .btn {
      display: block;
      margin: 1em;
      width: 90%;
      border-radius: 1em;
      height: 2em;
  }
  .btn-primary{
      background-color: #007bff;
      font-family: 'Trebuchet MS', 'Lucida Sans Unicode', 'Lucida Grande', 'Lucida Sans', Arial, sans-serif;
  
  }
  
  .btn-primary:hover{
      background-color: #3b90ff;
      transition: background-color 0.3s ease;
  }
  .btn-primary:active{
      background-color: #6070ff;
      transition: background-color 0.3s ease;
  }
  
  </style>