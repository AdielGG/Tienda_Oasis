<template>
    <v-btn 
        class="user-btn"
        @click="Drawer"
        >
        
        <v-icon icon="mdi-account"></v-icon>
        
        {{ $store.state.user.username }}

    </v-btn>

</template>

<script>

export default {
    methods:{
        Drawer(){
            this.$store.commit('setDrawer', true)
            console.log(this.$store.state.drawer)
        }
    },
}
</script>

<style scoped>
 

 .user-btn{
    margin: 0.3em 0.5em 0 0.5em;
 }
 
</style>