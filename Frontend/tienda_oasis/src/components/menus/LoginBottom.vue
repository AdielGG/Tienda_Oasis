<template>
    <section>

        <v-btn  
            @click="login"
            class="nav-btn second">
            Iniciar Sesion
        </v-btn>
        <v-btn 
            @click="register"
            class="nav-btn primary">
            Registrarse
        </v-btn>
    </section>
</template>

<script>
import router from '@/router';

export default {
    methods:{
        login(){
            router.push('/login')
        },
        register(){
            router.push('/register')
        }
    }
}
</script>

<style scoped>
section{
    display: flex;
    justify-content: space-between;
    width: 100%;
    height: 100%;
}

.nav-btn{
    margin: 0.3em 0.5em;
}

.primary{
    background-color: rgba(0, 89, 255, 0.979);
    color: rgb(0, 0, 0);
}

.primary:hover{
    transition-duration: 0.3s;
    background-color: rgb(0, 119, 255);
    border: 1px solid rgb(0, 0, 0);
}

.second{
    background-color: rgba(255, 255, 255, 0.979);
    color: rgb(0, 0, 0);
}
.second:hover{
    transition-duration: 0.3s;
    background-color: rgb(217, 223, 222);
    border: 1px solid rgb(0, 0, 0);
}
</style>