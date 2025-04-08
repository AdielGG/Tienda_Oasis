<template>
  <v-container>
    <form>
        <v-card class="main-card">
            <div class="logo-container">
                <img src="../../assets/oasis.png" alt="Logo"  class="logo">
            </div>

            <v-card-title>                
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
            </v-card-title>
            
            <section class="actions-container">
                <button class="btn btn-primary" @click="login" :disabled="loading">
                    <span v-if="loading">Cargando...</span>
                    <span v-else>Entrar</span>

                </button>
                <button class="btn btn-secondary" @click="() => { this.$router.push('/register')}" :disabled="loading">
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
export default {
    data() {
        return {
            username: '',
            password: '',
            loading: false,
            dialog: false,
            errorText: '',
            ErrorTitle: ''
        }
    },
    methods: {
        async login() {
            this.loading = true;

            if(!this.validar()){
                this.loading = false;
                return;
            }

            await axios.post('http://localhost:8080/login', {
                username: this.username,
                password: this.password

            }).then(response => {

                
                this.$store.commit('setUser', response.data.user)
                this.$store.commit('setLogged', true)
                this.loading = false;
                
                console.log(this.$store.state.user.name)

                if(this.$store.state.user.role === 'admin'){
                    this.$router.push('/admin')
                }
                else{
                    this.$router.push('/')
                }

            }).catch(error => {
                
                console.log(error);

                this.loading = false;
                
                this.errorText = error.message === 'Network Error' 
                                ? 'No se Pudo conectar con el servidor' 
                                : 'Usuario o contraseña incorrectos';
                
                this.ErrorTitle = error.name === "AxiosError" 
                                ? 'Error del Servidor' 
                                : 'Error';
                
                this.dialog = true;
                            })
                            
            
        },
        validar(){
            if(this.username === '' || this.password === ''){

                this.dialog = true;
                this.errorText = 'El usuario o la contraseña no puede estar vacio';
                this.ErrorTitle = 'Campos requeridos';

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
    },
    
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
    font-family: 'Trebuchet MS', 'Lucida Sans Unicode', 'Lucida Grande', 'Lucida Sans', Arial, sans-serif;
    transition: 0.3s;
}
.btn-primary{
    background-color: #007bff;

}

.btn-primary:hover{
    background-color: #3b90ff;
    transition: background-color 0.3s ease;
}
.btn-primary:active{
    background-color: #6070ff;
    transition: background-color 0.3s ease;
}
.btn-secondary:hover{
    color:#007bff;
    border: solid #007bff 1px;
    transition: 0.3s;
}

</style>