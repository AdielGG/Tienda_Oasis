<template>
    <AdminNavbar>
        <slot>
            <div class="text-center">
                <h1>Nuevo Producto</h1>
            </div>
            <v-form @submit.prevent="addProduct">
                <div class="input-pair">
                    <v-text-field 
                      label="Nombre" 
                      v-model="nombre"
                      :error-messages="nombreErrors"
                    ></v-text-field>
                    <v-text-field 
                      label="Descripción" 
                      v-model="descripcion"
                      :error-messages="descripcionErrors"
                    ></v-text-field>
                </div>
                <div class="input-pair">
                    <v-text-field 
                      label="Precio" 
                      v-model="precio"
                      type="number"
                      step="0.01"
                      :error-messages="precioErrors"
                    ></v-text-field>
                    <v-text-field 
                      label="Cantidad" 
                      v-model="cantidad"
                      type="number"
                      :error-messages="cantidadErrors"
                    ></v-text-field>
                </div>
                <div class="text-center aling-items-center">
                    <img 
                      v-if="imagen" 
                      :src="previewImage" 
                      width="300px"
                      style="max-height: 200px; object-fit: contain;"
                    />
                    <v-file-input
                      accept="image/png, image/jpeg, image/bmp"
                      label="Imagen"
                      placeholder="Selecciona una imagen"
                      prepend-icon="mdi-camera"
                      v-model="imagen"
                      @change="onFileChange"
                      :error-messages="imagenErrors"
                    ></v-file-input>
                    <v-btn 
                      color="primary" 
                      @click="addProduct"
                      :loading="loading"
                    >
                      {{ loading ? 'Procesando...' : 'Agregar' }}
                    </v-btn>
                </div>
            </v-form>
            
            <v-dialog v-model="dialog" width="auto">
                <v-card max-width="400" :title="ErrorTitle" :text="errorText">
                    <template v-slot:actions>
                        <v-btn class="ms-auto" text="Ok" @click="dialog = false"></v-btn>
                    </template>
                </v-card>
            </v-dialog>
        </slot>
    </AdminNavbar>
</template>

<script>
import AdminNavbar from '@/components/menus/AdminNavbar.vue';
import { z } from 'zod';
import { productSchema } from '@/plugins/validationSchemas';

export default {
    components: {
        AdminNavbar
    },
    data() {
        return {
            nombre: "",
            descripcion: "",
            precio: "",
            imagen: null,
            cantidad: "",
            dialog: false,
            errorText: '',
            ErrorTitle: '',
            loading: false
        };
    },
    methods: {
        async addProduct() {
            if (!this.validar()) {
                return;
            }

            this.loading = true;
            
            try {
                const formData = new FormData();
                formData.append('nombre', this.nombre);
                formData.append('descripcion', this.descripcion);
                formData.append('precio', parseFloat(this.precio));
                formData.append('cantidad', parseInt(this.cantidad));
                formData.append('imagen', this.imagen);

                const response = await axios.post('/api/products', formData, {
                    headers: {
                        'Content-Type': 'multipart/form-data'
                    }
                });

                // Éxito - redirigir o limpiar formulario
                this.$router.push('/admin/products');
                
            } catch (error) {
                this.dialog = true;
                this.errorText = 'Error al agregar el producto';
                this.ErrorTitle = 'Error del servidor';
                console.error(error);
            } finally {
                this.loading = false;
            }
        },
        validar() {
            try {
                // Convertir a números para validación
                const precioNum = parseFloat(this.precio);
                const cantidadNum = parseInt(this.cantidad);

                productSchema.parse({
                    nombre: this.nombre,
                    descripcion: this.descripcion,
                    precio: isNaN(precioNum) ? 0 : precioNum,
                    cantidad: isNaN(cantidadNum) ? -1 : cantidadNum,
                    imagen: this.imagen
                });
                
                return true;
                
            } catch (error) {
                if (error instanceof z.ZodError) {
                    const firstError = error.errors[0];
                    this.dialog = true;
                    this.errorText = firstError.message;
                    
                    // Tipos de error comunes
                    if (firstError.code === 'too_small') {
                        this.ErrorTitle = 'Valor muy pequeño';
                    } else if (firstError.code === 'too_big') {
                        this.ErrorTitle = 'Valor muy grande';
                    } else if (firstError.code === 'invalid_type') {
                        this.ErrorTitle = 'Tipo de dato incorrecto';
                    } else {
                        this.ErrorTitle = 'Error de validación';
                    }
                }
                return false;
            }
        }
    }
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