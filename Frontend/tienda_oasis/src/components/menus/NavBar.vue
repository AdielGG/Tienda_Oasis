import store from '@/store/store';
<template>
    <nav class="navbar">
        <v-btn class="nav-toggle" @click="toggleMenu"> 
            <v-icon size="x-large">mdi-menu</v-icon>
        </v-btn>

        <div :class="class1">
            <ul class="nav nav-fill tm-main-nav-ul">
                <router-link to="/" class="nav-item">Inicio</router-link>
                <router-link to="/products" class="nav-item">Productos</router-link>
                <router-link to="/about" class="nav-item">Nosotros</router-link>
                <router-link to="/help" class="nav-item">Ayuda</router-link>
            </ul>
        </div>
    </nav>
</template>

<script>
export default {
    data() {
        return {
            class1: "tm-main-nav",
            active_nav: false,
        }
    },
    methods: {
        toggleMenu() {
            this.class1 = this.active_nav ? "tm-main-nav tm-main-nav_active" : "tm-main-nav";
            this.active_nav = !this.active_nav;
        }
    },
    mounted() {
        this.active_nav = this.$store.state.menu;
    },
    

}
</script>

<style >
li {
    list-style: none;
}
a {
    text-decoration: none;
}
.navbar {
    width: 100%;
    background-color: rgba(36, 36, 36, 0.664);
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
}
.nav {
    display: flex;
    align-content: center;
}
.nav-item {
    color: #a8a9aa;
    margin-left: 1em;
    font-size: 2em;
}
.nav-item:hover {
    color: rgb(0, 174, 255);
}
.router-link-active {
    color: rgb(0, 174, 255);
}

.tm-main-nav .nav-link {
    color: white;
    font-size: 1.3rem;
    padding: 15px 20px;
}
.navbar {
    height: auto;
    border-bottom: 10px solid #076599;
}
.tm-main-nav .nav-link.active,
.tm-main-nav .nav-link:hover,
.tm-main-nav .nav-link:focus {
    color: #9ac9f7;
}

.navbar button.v-btn.nav-toggle {
    display: none !important;
    margin-left: auto;
    margin-right: 1em;
    font-size: 2em;
}
@media (max-width: 768px) {
    .navbar button.v-btn.nav-toggle{
        display: flex !important;
        align-content: center;
        
        font-size: 2em;
        color: white;
        background-color: transparent;
        border: none;
        cursor: pointer;
        box-shadow: none;
    }
    .nav-toggle:hover {
        color: white;
        background-color: transparent;
        border: none;
    }
    .navbar {
        height: 4em;
        display: flex;
        flex-direction: column;
        align-content: flex-start;
        flex-wrap: wrap;
        justify-content: center;
        border-bottom: 8px solid #076599;
    }
    .tm-main-nav {
        flex-direction: column;

        width: 100%;
        height: auto;

        background-color: rgba(36, 36, 36, 0.938);
        border-bottom: 8px solid #076599;

        position: absolute;
        top: 100%;
        right: 0;

        z-index: 1000;

        padding: 1em;
        transition: right 0.4s ease-in-out;
    }
    .tm-main-nav-ul {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
    }
    .nav-item {
        color: #a8a9aa;
        margin-left: 0;
        margin-right: 0;
        font-size: 1.5em;
    }
    .router-link-active {
    color: rgb(0, 174, 255);
    }
    .nav-item:hover {
        color: rgb(0, 174, 255);
    }
    .tm-main-nav_active {
        right: 100% ;

    }
}
</style>
