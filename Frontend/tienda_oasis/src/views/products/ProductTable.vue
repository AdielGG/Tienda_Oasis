<template>
    <v-card class="product-container">
      <v-tabs
        v-model="tab"
        align-tabs="center"
        color="deep-purple-accent-4"
      >
        <v-tab value="programs">Programas</v-tab>
        <v-tab value="curses">Cursos</v-tab>
        <v-tab value="books">Libros</v-tab>
      </v-tabs>
  
      <v-tabs-window v-model="tab">
        <v-tabs-window-item value="programs">
          <Programs></Programs>
        </v-tabs-window-item>
        <v-tabs-window-item value="curses">
          <Curses></Curses>
        </v-tabs-window-item>
        <v-tabs-window-item value="books">
          <Books></Books>
        </v-tabs-window-item>
      </v-tabs-window>
    </v-card>
  </template>
  <script setup>
    import { ref } from 'vue'
    import Programs from './items/Programs.vue';
    import Books from './items/Books.vue';
    import Curses from './items/Videos.vue'
  
    const tab = ref(null)
  </script>

  <style>
    .v-card{
      position:sticky;
      z-index: auto !important;
    }
  </style>