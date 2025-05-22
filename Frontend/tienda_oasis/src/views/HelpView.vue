<template>
    <header class="tm-site-header">
       
        <NavBar></NavBar>
    </header> 
    <div class="about-container">
        <h1>Esta es la PAgina de ayuda</h1>

        <div class="about-section">
            <h2>En Desarrollo ...</h2>
        </div>
    </div>
    <Footer></Footer>
</template>

<script>
import Footer from "@/components/Footer.vue";
import NavBar from "@/components/menus/NavBar.vue";
export default {
    components: {
        NavBar,
        Footer,
    },
};
</script>

<style>
@media (min-width: 1024px) {
    .about {
        min-height: 100vh;
        display: flex;
        align-items: center;
    }
}
.about-container {
    max-width: 1000px;
    margin: 0 auto;
    padding: 2rem;
    line-height: 1.6;
    font-size: 1.5em;
}
.about-container h1 {
    background-color: #076599;
    color: #ffffff;
}

.about-section {
    margin-bottom: 3rem;
    padding: 1.5rem;
    background: #076599;
    border-radius: 8px;
}

.about-section h2 {
    color: #000000;
    border-bottom: 2px solid #000000;
    padding-bottom: 0.5rem;
}

.team-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
    gap: 1.5rem;
    margin-top: 1.5rem;
}

.team-member {
    text-align: center;
    padding: 1rem;
    background: rgb(0, 0, 0);
    border-radius: 8px;
    box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
}

.member-avatar {
    font-size: 2.5rem;
    margin-bottom: 1rem;
}

ul {
    padding-left: 1.5rem;
}

a {
    color: #000000;
    text-decoration: none;
}

a:hover {
    text-decoration: underline;
}
</style>
