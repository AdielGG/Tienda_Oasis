<template>
  <div class="rating">
      <img v-for="i in full" :key="i" src="../assets/icons/star.svg"/>
      <img v-if="half" src="../assets/icons/star-half-full.svg"/>
      <img v-for="i in outLine" :key="i" src="../assets/icons/star-outline.svg"/>
    <span class="rating-text">{{ rating}}{{ half ? '' : '.0'}}</span>

  </div>
</template>

<script>
export default {
    name: 'Rating',
    props: {
        rating: Number
    },
    data(){
        return{
            full: [],
            half: false,
            outLine: [],
            
        }
    },
    methods:{
        rate(){
            for(let i = 0; i < this.rating - this.rating%1  ; i++){
            this.full.push(i)
            }
            for(let i = 0; i < 5 - this.rating - this.rating%1 ; i++){
                this.outLine.push(i)
            }
            this.half = this.rating%1 !== 0

        }

    },
    mounted(){
        this.rate()
    }
}
</script>

<style>
.rating img{
    width: 1em;
    vertical-align: middle;
}
.rating-text{
    vertical-align: middle;
    margin-left: 0.5em;
    color: #0030F0;
}


</style>