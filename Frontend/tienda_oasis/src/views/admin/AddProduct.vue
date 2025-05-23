<template>
    <v-overlay class="align-center justify-center" activator="parent">
        <div class="overlay-items-content">

            <div class="text-center">
                <h1>Nuevo Producto</h1>
            </div>
            <v-form>
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

                <div class="input-pair">
                    
                </div>
                <div class="input-pair">
                    <div class="file-container">
                        <div class="images">
                                <div class="image-container">
                                    <v-icon size="70px">mdi-archive</v-icon>
                                    
                                    <span >{{ fileName }}</span>
                                </div>
                        </div>
                        <form class="inputfile">
                            <input
                                class="fileinput"
                                type="file"                                
                                id="media"
                                accept="apllication/*"
                                @change="(event) => onFileChange(event)"
                            />
                            <div>
                                <section>
                                <v-icon size="x-large">mdi-cloud-upload</v-icon> 
                                Contenido
                                </section>
                            </div>
                        </form>
                    </div>
                    <div class="file-container">
                        <div class="images">
                            <div
                                
                                class="images-lists"
                            >
                                <div class="image-container">
                                    <img v-if="imageSrc != ''" :src="imageSrc" id="output" class="image-style" />
                                    <v-icon v-else size="x-large">mdi-image</v-icon>
                                </div>
                            </div>
                        </div>
                        <form class="inputfile">
                            <input
                                class="fileinput"
                                type="file"
                                id="media"
                                accept="image/*"
                                @change="(event) => onImgChange(event)"
                            />
                            <div>
                                <section>
                                <v-icon size="x-large">mdi-cloud-upload</v-icon> 
                                    Imagen
                                </section>
                            </div>
                        </form>
                    </div>
                </div>

                <div class="text-center aling-items-center">
                    
                      <v-btn 
                      style="margin-bottom: 1em;"
                      color="primary" 
                      @click="uploadImage"
                      :loading="loading"
                      >
                      {{ loading ? 'Procesando...' : 'Agregar' }}
                    </v-btn>
                </div>
            </v-form>
            
            <v-dialog v-model="dialog" width="auto">
                <v-card max-width="400" :title="ErrorTitle" :text="errorText">
                    <template v-slot:actions>
                        <v-btn class="ms-auto"  text="Ok" @click="dialog = false"></v-btn>
                    </template>
                </v-card>
            </v-dialog>
        </div>
    </v-overlay>
</template>

<script>
import AdminNavbar from '@/components/menus/AdminNavbar.vue';
import AddImg from '@/components/AddImg.vue';
import AddFiles from '@/components/AddFiles.vue';
import { z } from 'zod';
import { productSchema } from '@/plugins/validationSchemas';
import axios from 'axios';

export default {
    components: {
        AdminNavbar,
        AddImg,
        AddFiles
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
            loading: false,
            file: null,
            fileName: null,
            img: null,
            imageSrc: '',
        };
    },
    
    methods: {
        onFileChange(e) {
            this.file = e.target.files[0];
            this.fileName = this.file.name;
        },

/************************************************************************* */
        onImgChange(e) {
            this.img = e.target.files[0];
            this.imageSrc = URL.createObjectURL(this.img);
        },

/************************************************************************* */
        async uploadFile() {
            const formData = new FormData();
            formData.append('file', this.file);
            try {
            const response = await axios.post('http://localhost:8080/upload/programs', formData, {
                headers: {
                'Content-Type': 'multipart/form-data' // No es necesario si Axios lo maneja automáticamente
                }
            });
            console.log('Archivo subido:', response.data);
            } catch (error) {
            console.error('Error al subir el archivo:', error);
            }
        },

/************************************************************************* */
        async uploadImage() {
            const formData = new FormData();
            formData.append('file', this.img);
            try {
            const response = await axios.post('http://localhost:8080/upload/img', formData, {
                headers: {
                'Content-Type': 'image/*' 
                }
            });
            console.log('Archivo subido:', response.data);
            return response.data.image;
            } catch (error) {
            console.error('Error al subir el archivo:', error);
            }
        },

 /*************************************************************************** */       
        validar() {
            return true;
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
        }, 
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
    margin: 1em 2em;
}
.v-combobox {
    max-width: 45%;
}
.aling-items-center {
    display : flex;
    flex-direction: column;
    justify-content: center;
}
.overlay-items-content{
    background-color: rgb(0, 0, 0);
    min-width: max-content;
    border-radius: 1em;
}
</style>
<style scoped>
.container {
  margin-bottom: 40px;
  display: flex;
  flex-direction: column;
}


form.inputfile {
  
  width:max-content;
  border-radius: 5px;
  border: 1.5px dashed #a0a0a0;
  cursor: pointer;
}
form.inputfile div {
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}
form.inputfile input {
   
  color: #0c5645;
  background-color: #117a60;
  margin: 0;
  padding: 0;
  width: 8em;
  height: 4em;
  outline: none;
  opacity: 0;
  z-index: 3;
    cursor: pointer !important;
    position: relative;
}

form.inputfile section {
  cursor: pointer;
  z-index: 0;
  position: unset;
  margin-top: -55px;
}

.images {
  position: relative;
  margin: 4em 0 2em;

  display: flex;
  flex-direction: row;
  justify-content: center;
}

.images-lists {
  position: relative;
  margin-left: 10px;
  margin-right: 10px;
}

.image-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    width: 100%;
    margin: 1em;
    padding: 4px;
    border: 0.5px solid #a0a0a0;
    border-radius: 10px;
}

.image-style {
  height: 150px;
  width: 150px;
  object-fit: cover;
}

.cross-icon {
    position: relative;
    color: red;
  top: -4em;
  left: -8em;
  cursor: pointer;
}
.file-container{
    width: 50%;
    display: flex;
    flex-direction: column;
    align-items: center;
}
</style>