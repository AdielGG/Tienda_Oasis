<template>
    <div class="item-container">
        
        <ProductItem title="NuevoPeoducto" price="5.50" rating="4.5" description="Nuevo proffucto a la tienda" img="src/assets/background.jpg"/>
        <ProductItem title="NuevoPeoducto" price="5.50" rating="4.5" description="Nuevo proffucto a la tienda" img="src/assets/oasis.png"/>
        <ProductItem title="NuevoPeoducto" price="5.50" rating="4.5" description="Nuevo proffucto a la tienda" img="src/assets/background.jpg"/>
        <ProductItem title="NuevoPeoducto" price="5.50" rating="4.5" description="Nuevo proffucto a la tienda" img="src/assets/oasis.png"/>
        <ProductItem title="NuevoPeoducto" price="5.50" rating="4.5" description="Nuevo proffucto a la tienda" img="src/assets/background.jpg"/>
        <ProductItem title="NuevoPeoducto" price="5.50" rating="4.5" description="Nuevo proffucto a la tienda" img="src/assets/oasis.png"/>
        <ProductItem title="NuevoPeoducto" price="5.50" rating="4.5" description="Nuevo proffucto a la tienda" img="src/assets/background.jpg"/>
        
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