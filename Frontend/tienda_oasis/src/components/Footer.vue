<template>
    <div class="footer">
        <section class="top-section">
            <section class="links">
                <h1>Enlaces</h1>
                <router-link to="/">Inicio</router-link>
                <router-link to="/products">Productos</router-link>
                <router-link to="/about">Sobre Nosotros</router-link>
                <router-link to="/help">Ayuda</router-link>
            </section>
            <section class="links">
                <h1>Redes Sociales</h1>
                <a href="https://www.facebook.com/tiendaoasis/" target="_blank">
                    <v-icon>mdi-facebook</v-icon>
                    FaceBook
                </a>
                <a
                    href="https://www.instagram.com/tienda_oasis/"
                    target="_blank"
                >
                    <v-icon>mdi-instagram</v-icon>
                    Instagram
                </a>
                <a href="https://twitter.com/tienda_oasis" target="_blank">
                    <v-icon>mdi-twitter</v-icon>
                    Twitter
                </a>
            </section>
            <section class="links">
                <h1>Contacto</h1>
                <p>Email: info@tiendadigital.com</p>
                <p>Teléfono: +56 91 12345678</p>
                <p>Dirección: Av. de la Independencia, 123, Madrid</p>
                <p>Código Postal: 28020</p>
            </section>
            <section class="links">
                <h1>Descripción</h1>
                <p>
                    Tienda Oasis es una tienda online de programas, cursos y
                    libros de calidad. Dispone de una gran variedad de
                    productos, desde programas hasta cursos y libros. Tienda
                    Oasis es una tienda online de programas, cursos y libros de
                    calidad. Dispone de una gran variedad de productos, desde
                    programas hasta cursos y libros.
                </p>
            </section>
        </section>
        <section class="copyright">
            <p>Copyright © 2025 Tienda Oasis</p>
        </section>
    </div>
</template>

<script>
export default {};
</script>

<style>
.footer {
    color: white;
    display: block;
    flex-direction: column;
    justify-content: space-between;
    align-items: center;
    align-content: center;
    flex-wrap: wrap;
}
.links {
    color: white;
    margin: 1em;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: flex-start;
    flex-wrap: nowrap;
    max-width: 33%;
}

.footer {
    width: 100%;
    height: auto;
    background-color: rgb(0, 0, 0);
}
.copyright {
    text-align: center;
    margin: 2em;
}
.top-section {
    width: 100%;
    display: flex;
    flex-direction: row;
    flex-wrap: nowrap;
    justify-content: space-around;
    /* align-content: space-between; */
    align-items: flex-start;
}
.links a {
    color: white;
}
.links a:hover {
    background: black;
    color: white;
}
</style>
