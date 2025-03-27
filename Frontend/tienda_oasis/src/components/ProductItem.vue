<template>
    <v-card
      :disabled="loading"
      :loading="loading"
      class="product-item mx-auto my-12"
      max-width="250"
    >
      <template v-slot:loader="{ isActive }">
        <v-progress-linear
          :active="isActive"
          color="deep-purple"
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
        <v-row
          align="center"
          class="mx-0"
        >
          <v-rating
            :model-value="rating"
            color="amber"
            density="compact"
            size="small"
            half-increments
            readonly
          ></v-rating>
  
          <div class="text-grey ms-4">
            {{ rating }}
          </div>
        </v-row>
  
        <div class="my-4 text-subtitle-1">
          $: {{ price }}
        </div>
  
        <div>{{ description }}</div>
      </v-card-text>
  
      <v-divider class="mx-4 mb-1"></v-divider>
  
  
      <v-card-actions>
        <v-btn
          color="deep-purple-lighten-2"
          text="Reserve"
          block
          border
          @click="reserve"
        ></v-btn>
      </v-card-actions>
    </v-card>
  </template>
  <script>
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
      }
    }
  </script>

  <style>
    .product-item{
      margin: auto;
      background-color: rgb(255, 255, 255);
      box-shadow: 0 0  3px rgb(87, 87, 87);
    }

</style>