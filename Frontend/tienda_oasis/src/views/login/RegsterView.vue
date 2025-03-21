<template>
    <v-container>
      <form>
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
  import axios from 'axios'
import store from '@/store/store';
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
              if(!this.validar()){
                  this.loading = false;
                  return;
              }
            axios.post('http://localhost:8080/register', {
                name: this.name,
                lastname: this.lastname,
                age: this.age.toString(),
                email: this.email,
                username: this.username,
                password: this.password
            }).then(async response => {
                this.loading = false;
                console.log(response.data);
                if(response.data.OK === 'Logued in successfully'){
                    const {data} = await axios.get('http://localhost:8080/user/' + this.username)
                    store.state.user = data.user
                    store.state.logued = true;
                    this.$router.push('/')
                }
            }).catch(error => {
                
                console.log(error);
                this.loading = false;
                this.dialog = true;
 
                this.errorText = error.message === 'Network Error' 
                                ? 'No se Pudo conectar con el servidor' 
                                : 'Usuario o contraseña incorrectos';

                this.ErrorTitle = error.name === "AxiosError" 
                                ? 'Error del Servidor' 
                                : 'Error';
            })
              
          },
          validar(){
              if(this.username === '' || this.password === ''){
                  this.dialog = true;
                  this.errorText = 'El usuario o la contraseña no puede estar vacio';
                  this.ErrorTitle = 'Campos requeridos';
                  return false;
              }
              if(this.password !== this.cpassword){
                  this.dialog = true;
                  this.errorText = 'Las contraseñas no coinciden';
                  this.ErrorTitle = 'Contraseñas no coinciden';
                  return false;
              }
              if(this.age < 18){
                  this.dialog = true;
                  this.errorText = 'Debes ser mayor de edad';
                  this.ErrorTitle = 'Debes ser mayor de edad';
                  return false;
              }
              if(this.username.length < 4 || this.password.length < 4){
                  this.dialog = true;
                  this.errorText = 'El usuario o la contraseña debe tener al menos 4 caracteres';
                  this.ErrorTitle = 'Longitud minimo';
                  return false;
              }
              if(this.username.length > 20 || this.password.length > 20){
                  this.dialog = true;
                  this.errorText = 'El usuario o la contraseña debe tener al menos 20 caracteres';
                  this.ErrorTitle = 'Longitud maximo';
                  return false;
              }
              for(let i = 0; i < this.username.length; i++){
                  if(!this.username[i].match(/[a-zA-Z]/)){
                      this.dialog = true;
                      this.errorText = 'El usuario solo puede contener letras';
                      this.ErrorTitle = 'Caracteres no permitidos';
                      return false;
                  }
              }
              return true;
          }     
      }
  }
  </script>
  
  <style scpoed>
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