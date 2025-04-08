<template>
  <main>
    <header>
        <v-breadcrumbs :items="items" color="blue">
            <template v-slot:prepend>
                <v-icon icon="$vuetify" size="large" color="yellow"></v-icon>
            </template>
        </v-breadcrumbs>
    </header>  

    <v-sheet>
        <v-btn @click="$router.push('/')">
          Home
        </v-btn>
        <v-btn>
            Gestionar Productos
        </v-btn>
        <v-btn>
            Gestionar Cuentas
        </v-btn>
        <v-btn>
            Quejas y Sugerencias
        </v-btn>
    </v-sheet>
  </main>

</template>

<script setup>
  const items = [
    {
      title: 'AdminPanel',
      disabled: false,
      href: '#',
    },
  ]
</script>


<style>

</style>