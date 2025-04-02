<template>
    <v-card
      class="product-item mx-auto my-12"
      :disabled="loading"
      :loading="loading"
      max-width="250"
    >
      <template v-slot:loader="{ isActive }">
        <v-progress-linear
          :active="isActive"
          color="#0030F0"
          height="4"
          indeterminate
        ></v-progress-linear>
      </template>
  
      <v-img
        height="150"
        :src="img"
        cover
      ></v-img>
  
      <v-card-item>
        <v-card-title>{{ title }}</v-card-title>
  
        <v-card-subtitle>
          <v-icon
            color="error"
            icon="mdi-fire-circle"
            size="small"
          ></v-icon>
        </v-card-subtitle>
      </v-card-item>
  
      <v-card-text>
        <Rating :rating="rating"></Rating>
  
        <div class="my-4 text-subtitle-1">
          $: {{ price }}
        </div>
  
        <div>{{ description }}</div>
      </v-card-text>
  
      <v-divider class="color-blue mx-4 mb-1"></v-divider>
  
  
      <v-card-actions>
        <v-btn
          class="reserve-btn"
          text="Reserve"
          block
          
          @click="reserve"
        ></v-btn>
      </v-card-actions>
    </v-card>
  </template>
  <script>
    import Rating from '@/components/Rating.vue';
    export default{
      props:{
        img: null,
        title: null,
        description: null,
        rating: null,
        price: null

      },
      data(){
        return{

          selection: 1,
          loading: false
        }
      },
      methods:{

        reserve() {
          this.loading = true
          setTimeout(() => (loading.value = false), 200)
          
        }  
      },
      components:{
        Rating
      }
    }
  </script>

  <style>
    .v-card.product-item{
      margin: auto;
      box-shadow: 0 0 3px #0032FA;
    }
    .v-divider.color-blue{
      color: #0030F0;
      opacity: 1;
    }
    button.reserve-btn{
      background-color: white;
      color: #0030F0;
      border-radius: 0.3em;
      border: #0030F0 solid 1px;
    }

</style>