<template>
  <div>
    <v-btn
      class="btn primary-btn"
      @click="login"
    >
    Entrar
    </v-btn>
    <v-btn
      class="btn primary-btn"
      @click="register"
    >
    Resgistrarse
    </v-btn>
  </div>
</template>

<script>
export default {
  methods:{
    login(){
      this.$router.push('/login')
    },
    register(){
      this.$router.push('/register')
    }
  }
}
</script>

<style>

</style>