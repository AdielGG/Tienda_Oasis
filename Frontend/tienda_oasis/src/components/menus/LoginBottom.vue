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
    width: auto;
    margin-top: 1em;
    margin-right: 1em;
}

.primary{
    background-color: rgba(0, 255, 0, 0.979);
    color: rgb(0, 0, 0);
}

.primary:hover{
    transition-duration: 0.3s;
    background-color: rgb(43, 255, 0);
    border: 1px solid rgb(0, 0, 0);
}

.second{
    background-color: rgba(255, 255, 255, 0.979);
    color: rgb(0, 0, 0);
}
.second:hover{
    transition-duration: 0.3s;
    background-color: rgb(255, 255, 255);
    border: 1px solid rgb(0, 0, 0);
}
</style>