<template>
  <div class="contenedor-panel">
    <h2>Enviar Queja o Sugerencia</h2>
    <form @submit.prevent="enviarQuejaOSugerencia">
      <div class="form-grupo">
        <label for="categoria">Categoría:</label>
        <select id="categoria" v-model="categoria" class="input">
          <option value="soporte">Soporte Técnico</option>
          <option value="mejora">Sugerencia de Mejora</option>
          <option value="facturacion">Facturación</option>
        </select>
      </div>

      <div class="form-grupo">
        <label for="mensaje">Mensaje:</label>
        <textarea id="mensaje" v-model="mensaje" required class="input"></textarea>
      </div>

      <button type="submit" class="btn-enviar">Enviar</button>
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      mensaje: "",
      categoria: ""
    };
  },
  methods: {
    enviarQuejaOSugerencia() {
      console.log(this.mensaje);
    }
  }
};
</script>

<style scoped>
.contenedor-panel {
  width: 60%; /* Ajusta el ancho del contenedor */
  margin: 0 auto; /* Centra el contenedor horizontalmente */
  padding: 20px;
  background-color: #ffffff;
  border-radius: 8px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

h2 {
  text-align: center;
  color: #333;
}

.form-grupo {
  margin-bottom: 15px;
}

label {
  display: block;
  margin-bottom: 5px;
  color: #555;
}

.input {
  width: 100%;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 4px;
  box-sizing: border-box;
}

.btn-enviar {
  display: block;
  width: 100%;
  padding: 10px;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  text-transform: uppercase;
}

.btn-enviar:hover {
  background-color: #0056b3;
}
</style>