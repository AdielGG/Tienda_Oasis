
<template>
      
    <div class="file-container">
        <div class="images">
            <div
                v-if="imageSrc != ''"
                :key="1"
                class="images-lists"
            >
                <div class="image-container">
                    <v-icon size="70px">mdi-archive
                    </v-icon>
                    <span>
                        {{ form.media.name }}
                    </span>
                </div>
            </div>
        </div>
        <form>
        <input
            class="fileinput"
            type="file"
            id="media"
            accept="apllication/*"
            @change="(event) => handelFileUpload(event)"
        />
        <div>
            <section>
            <v-icon size="x-large">mdi-cloud-upload</v-icon> 
            Contenido
            </section>
        </div>
        </form>
    </div>
</template>
<script setup>
import { ref } from "vue";
import store from '@/store/store';

const form = ref({
  media: {},
});

const fileSrc = ref("");
const fileName = ref("")

const handelFileUpload = (e) => {
  var file = e.target.files[0] || e.dataTransfer.files[0]; 
  

  const src = URL.createObjectURL(file);
  fileSrc.value = src;
  console.log(file)
  
  form.value.media = e.target.files[0];
  this.$store.commit('setFile', e.target.files[0]);
  console.log(form.value.media, "file upload");

  console.log("files already uploaded", fileSrc.value);
};


</script>

<style scoped>
.container {
  margin-bottom: 40px;
  display: flex;
  flex-direction: column;
}


form {
  
  width:max-content;
  border-radius: 5px;
  border: 1.5px dashed #a0a0a0;
  cursor: pointer;
}
form div {
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}
form input {
   
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

form section {
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
    width: 100%;
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