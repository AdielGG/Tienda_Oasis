<template>
    <header class="tm-site-header">
       
        <NavBar></NavBar>
    </header> 
    <div class="about-container">
        <h1>Sobre Nuestra Tienda Digital</h1>

        <div class="about-section">
            <h2>Nuestra Misión</h2>
            <p>
                En <strong>TiendaDigital</strong>, nos dedicamos a ofrecer los
                mejores productos digitales con una experiencia de compra
                sencilla, segura y accesible para todos. Creemos en el poder de
                lo digital para simplificar la vida.
            </p>
        </div>

        <div class="about-section">
            <h2>¿Por Qué Elegirnos?</h2>
            <ul>
                <li>✅ Productos digitales de alta calidad</li>
                <li>✅ Descarga inmediata después de la compra</li>
                <li>✅ Soporte técnico 24/7</li>
                <li>✅ Precios competitivos</li>
                <li>✅ Actualizaciones gratuitas</li>
            </ul>
        </div>

        <div class="about-section team-section">
            <h2>Nuestro Equipo</h2>
            <p>
                Somos un equipo apasionado por la tecnología y la innovación,
                comprometido con brindar la mejor experiencia a nuestros
                clientes.
            </p>
            <div class="team-grid">
                <div class="team-member">
                    <div class="member-avatar">👨‍💻</div>
                    <h3>Brayan Castellano</h3>
                    <p>Fundador & CEO</p>
                </div>
                <div class="team-member">
                    <div class="member-avatar">👩‍🎨</div>
                    <h3>Marcos Javier</h3>
                    <p>Diseñador Principal</p>
                </div>
                <div class="team-member">
                    <div class="member-avatar">👨‍💻</div>
                    <h3>Adiel Alejandro</h3>
                    <p>Desarrollador Senior</p>
                </div>
            </div>
        </div>

        <div class="about-section">
            <h2>Contacto</h2>
            <p>
                ¿Tienes preguntas? Escríbenos a
                <a href="mailto:info@tiendadigital.com"
                    >info@tiendadigital.com</a
                >
                o síguenos en nuestras redes sociales.
            </p>
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
