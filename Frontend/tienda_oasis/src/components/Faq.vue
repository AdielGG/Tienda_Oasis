<template>
  <div class="faq-item">
        <h3 @click="toggle"> {{ question }}</h3>
        <transition name="fade">
            <p v-if="isVisible"> {{ answer }}</p>    

        </transition>
  </div>
</template>

<script>
 export default {
  name: 'Faq',
  props: {
    question: String,
    answer: String
  },
  data() {
    return {
      isVisible: false
    }
  },
  methods: {
    toggle() {
      this.isVisible = !this.isVisible
    }
  }
}
</script>

<style scoped>
    h3{
        cursor: pointer;
        color: #007BFF !important;
    }
    p{
        margin: 10px 0;
        color: #333;
    }

    .fade-enter-active,
    .fade-leave-active {
      transition: opacity .5s;
    }
    .fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
      opacity: 0;
    }

    .faq-item{
        width: 80%;
        margin:10px 0;
        text-align: center;
    }

</style>