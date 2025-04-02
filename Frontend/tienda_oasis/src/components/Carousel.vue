<template>
   <v-carousel show-arrows="hover" cycle>
  <v-carousel-item
    src="http://localhost:8080/img/programas.jpg"
    cover
  ></v-carousel-item>

  <v-carousel-item
    src="http://localhost:8080/img/ebooks.jpg"
    cover
  ></v-carousel-item>

  <v-carousel-item
    src="http://localhost:8080/img/cursos.png"
    cover
  ></v-carousel-item>
</v-carousel>
  </template>
<script>

import ProductItem from './ProductItem.vue';
export default{
  components:{
    ProductItem
  },
  data(){
    return{
      sliders: []
    }
  },
  methods:{
    loadProducts(){
      this.sliders= [
        
        {img:"../assets/programas.jpg",
        text: "Mejores Pro1ductos"

        },
        {img:"@/assets/cursos.png",
        text: "Mejores Pr2oductos"

        },
        {img:"@/assets/ebooks.jpg",
        text: "Mejores Pr3oductos"

        }

      ]
    }
  },
  mounted(){
    this.loadProducts()
  }
}
    
    
</script>