<template>
    <AdminNavbar>
        <slot>
            <div class="text-center">
                <h1>Nuevo Producto</h1>
            </div>
            <v-form>
                <div class="input-pair">
                    <v-text-field label="Nombre" v-model="nombre"></v-text-field>
                    <v-text-field label="Descripción" v-model="descripcion"></v-text-field>

                </div>
                <div class="input-pair">

                    <v-text-field label="Precio" v-model="precio"></v-text-field>
                    <v-text-field label="Cantidad" v-model="cantidad"></v-text-field>
                </div>
                <div class="text-center aling-items-center">
                    <img src="../../assets/background.jpg" cover width="300px"/>
                    <v-file-input
                    :rules="rules"
                    accept="image/png, image/jpeg, image/bmp"
                    label="Imagen"
                    placeholder="Selecciona una imagen"
                    prepend-icon="mdi-camera"
                    v-model="imagen"
                    width="50%"
                    ></v-file-input>
                    <v-btn color="primary" @click="addProduct">Agregar</v-btn>
                </div>
            </v-form>
        </slot>
    </AdminNavbar>
</template>

<script>
import AdminNavbar from '@/components/menus/AdminNavbar.vue';
export default {
    components: {
        AdminNavbar
    },
    data() {
        return {
            nombre: "",
            descripcion: "",
            precio: "",
            imagen: "@/assets/background.jpg",
            cantidad: "",
            rules: [
                v => !!v || "Imagen es requerida",
                v => v.size <= 2000000 || "El tamaño de la imagen no puede exceder de 2MB",
            ],
        };
    },
    methods: {
        addProduct() {
            console.log(this.nombre);
            console.log(this.descripcion);
            console.log(this.precio);
            console.log(this.cantidad);
            console.log(this.imagen);
        },
    },
}
</script>

<style>
.input-pair {
    display: flex;
    justify-content: space-between;
    align-items: center;
}
.v-text-field {
    margin: 0 2em;
}
.v-combobox {
    max-width: 45%;
}
.aling-items-center {
    display : flex;
    flex-direction: column;
    justify-content: center;
}
</style>