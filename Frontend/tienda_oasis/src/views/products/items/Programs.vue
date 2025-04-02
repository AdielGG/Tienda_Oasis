<template>
    <div class="item-container">
        
        <ProductItem title="NuevoPeoducto" price="5" :rating="3" description="Nuevo proffucto a la tienda" img="src/assets/background.jpg" @click="console.log('hola')"/>
        <ProductItem title="NuevoPeoducto" price="5" :rating="2" description="Nuevo proffucto a la tienda" img="src/assets/oasis.png"/>
        <ProductItem title="NuevoPeoducto" price="5" :rating="2" description="Nuevo proffucto a la tienda" img="src/assets/background.jpg"/>
        <ProductItem title="NuevoPeoducto" price="5" :rating="4.5" description="Nuevo proffucto a la tienda" img="src/assets/oasis.png"/>
        <ProductItem title="NuevoPeoducto" price="5" :rating="5" description="Nuevo proffucto a la tienda" img="src/assets/background.jpg"/>
        <ProductItem title="NuevoPeoducto" price="5" :rating="4" description="Nuevo proffucto a la tienda" img="src/assets/oasis.png"/>
        <ProductItem title="NuevoPeoducto" price="5" :rating="4" description="Nuevo proffucto a la tienda" img="src/assets/background.jpg"/>
        
    </div>
</template>

<script>
import ProductItem from '@/components/ProductItem.vue';
export default {
components:{
    ProductItem,
}
}
</script>

<style >
    div.item-container{
        display: flex;
        flex-direction: row-reverse;
        flex-wrap: wrap;
        align-content: flex-start;
        justify-content: flex-end; 
    }
    
</style>