<template>
  <header class="tm-site-header">
        <section class="site-title">

            <p class="tm-site-description">Tu Tienda Favorita</p>
            
        </section>
        <NavBar></NavBar>
    </header> 
</template>

<script>
import NavBar from './menus/NavBar.vue'
export default {
    components:{
        NavBar
    }
}
</script>

<style>
.site-title{
    font-family: 'Open Sans', Helvetica, Arial, sans-serif;
	font-size: 3em;
	padding: 3em 2em 0 2em;
    color: white;
	text-shadow: 0px 3px 2px rgba(62,62,62,1);
}
.tm-site-header{
    background: url('../assets/portada.jpg') no-repeat;
	height: 20em;
    background-size: cover;
    position: relative;
    
    display: flex;
    align-content: space-between;
    justify-content: flex-start;
    align-items: baseline;
    flex-wrap: wrap;
    flex-direction: row;
}
@media (max-width: 768px) {
    .site-title{
        font-family: 'Open Sans', Helvetica, Arial, sans-serif;
		font-size: 2em;
		padding: 1em 2em 0 2em;

    }
    
    
}

</style>