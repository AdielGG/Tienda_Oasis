<template>
            
        <nav class="navbar navbar-expand-md tm-main-nav-container">
            <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#tmMainNav" aria-controls="tmMainNav" aria-expanded="false" aria-label="Toggle navigation">
                    <i class="fa fa-bars"></i>
            </button>

            <div class="tm-main-nav">
                <ul class="nav nav-fill tm-main-nav-ul">
                    <router-link to="/" class="nav-item">Inicio</router-link>
                    <router-link to="/products" class="nav-item">Productos</router-link>
                    <router-link to="/about" class="nav-item">Nosotros</router-link>
                    <router-link to="/about" class="nav-item">Ayuda</router-link>
                    <router-link to="/faq" class="nav-item">FAQ</router-link>
                </ul>
            </div>
        </nav>
        
    
</template>

<script>

</script>


<style scoped>
li{
    list-style: none;
}
a{
    text-decoration: none;
}
.navbar{
    width: 100%;
    background-color: rgba(36, 36, 36, 0.664);
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
}
.nav{
    display: flex;
    align-content: center;
}
.nav-item{
    margin-left: 1em;
    font-size: 2em;
}
.nav-item:hover{
    color: rgb(0, 174, 255);
}
.router-link-active{
    color: rgb(0, 174, 255);
}

.tm-main-nav .nav-link {
	color: white;
	font-size: 1.3rem;
	padding: 15px 20px;
    
}
.navbar{
    height: auto;
    border-bottom: 10px solid #076599;

}
.tm-main-nav .nav-link.active,
.tm-main-nav .nav-link:hover,
.tm-main-nav .nav-link:focus {
	color: #9AC9F7;
}
</style>