<template>
    <nav>
        <ul class="nav-bar">
            <li>
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
            <v-btn class="nav-btn second">
                Iniciar Sesion
            </v-btn>
            <v-btn class="nav-btn primary">
                Registrarse
            </v-btn>
        </ul>
    </nav>
</template>

<script>

</script>

<style scoped>

.nav-bar{
    border-radius: 10px;
    display: flex;
    background-color: rgb(187, 35, 35);
}

li {
    margin: 10px;
    list-style: none;
}

.nav-item{
    color: black;
    font-size: 2em;
    font-family:Georgia, 'Times New Roman', Times, serif;
}

.nav-item:hover{
    background-color: rgba(0, 0, 0, 0);
    border-bottom: solid 1px black;
}
.nav-btn{
    width: auto;
    margin-top: 1em;
    margin-right: 1em;
}

.primary{
    background-color: rgb(238, 255, 0);
    color: rgb(0, 0, 0);
}

</style>
