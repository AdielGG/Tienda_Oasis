<script setup>
import Header from "@/components/Header.vue";
import Carousel from "@/components/Carousel.vue";
import ProductItem from "@/components/ProductItem.vue";
import Footer from "@/components/Footer.vue";
</script>

<template>
    <main>
        <Header></Header>
        <v-divider class="color-blue mx-4 mb-1 margin-top-2"></v-divider>

        <div class="title">
            <h1>Nosotros Ofrecemos</h1>
        </div>

        <Carousel></Carousel>

        <v-divider class="color-blue mx-4 mb-1 margin-top-2"></v-divider>

        <div class="title">
            <h1>Productos Recientes</h1>
        </div>

        <div class="content-items">
            <ProductItem
                title="NuevoPeoducto"
                price="5"
                :rating="3"
                description="Nuevo proffucto a la tienda"
                img="src/assets/background.jpg"
                @click="console.log('hola')"
            />
            <ProductItem
                title="NuevoPeoducto"
                price="5"
                :rating="2"
                description="Nuevo proffucto a la tienda"
                img="src/assets/background.jpg"
            />
            <ProductItem
                title="NuevoPeoducto"
                price="5"
                :rating="4.5"
                description="Nuevo proffucto a la tienda"
                img="src/assets/oasis.png"
            />
            <ProductItem
                title="NuevoPeoducto"
                price="5"
                :rating="5"
                description="Nuevo proffucto a la tienda"
                img="src/assets/background.jpg"
            />
            <ProductItem
                title="NuevoPeoducto"
                price="5"
                :rating="4"
                description="Nuevo proffucto a la tienda"
                img="src/assets/background.jpg"
            />
        </div>
        <v-divider class="color-blue mx-4 mb-1"></v-divider>

        <div class="title">
            <h1>Más Vendidos</h1>
        </div>

        <div class="content-items">
            <ProductItem
                title="NuevoPeoducto"
                price="5"
                :rating="3"
                description="Nuevo proffucto a la tienda"
                img="src/assets/background.jpg"
                @click="console.log('hola')"
            />
            <ProductItem
                title="NuevoPeoducto"
                price="5"
                :rating="2"
                description="Nuevo proffucto a la tienda"
                img="src/assets/background.jpg"
            />
            <ProductItem
                title="NuevoPeoducto"
                price="5"
                :rating="5"
                description="Nuevo proffucto a la tienda"
                img="src/assets/background.jpg"
            />
            <ProductItem
                title="NuevoPeoducto"
                price="5"
                :rating="4"
                description="Nuevo proffucto a la tienda"
                img="src/assets/oasis.png"
            />
            <ProductItem
                title="NuevoPeoducto"
                price="5"
                :rating="4"
                description="Nuevo proffucto a la tienda"
                img="src/assets/background.jpg"
            />
        </div>

        <Footer></Footer>
    </main>
</template>

<style scoped>
.margin-top-2 {
    margin-top: 2em;
}
.title {
    text-align: center;
    font-size: 2em;
    margin-top: 2em;
    color: black;
    font-family: "Courier New", Courier, monospace;
}

.content-items {
    display: flex;
    flex-direction: row-reverse;
    flex-wrap: wrap;
    align-content: flex-start;
    justify-content: flex-end;
}
</style>
