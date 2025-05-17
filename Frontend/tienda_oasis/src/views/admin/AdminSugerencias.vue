<template>
    <AdminNavBar>
      <slot>
        <div class="text-center">
          <h1>Quejas y Sugerencias</h1>
        </div>
        <v-fab
          key="absolute"
          app
          color="primary"
          location="bottom center"
          size="large"
          @click="AddSugerencia"
        >
          <v-icon>mdi-message-text-plus</v-icon>
          Nueva Queja/Sugerencia
        </v-fab>
        <div class="contenedor-tabla">
          <table v-if="sugerencias.length > 0">
            <thead>
              <tr>
                <th>ID</th>
                <th>Categoría</th>
                <th>Mensaje</th>
                <th>Fecha</th>
                <th>Acciones</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="sugerencia in sugerencias" :key="sugerencia.id">
                <td>{{ sugerencia.id }}</td>
                <td>{{ sugerencia.categoria }}</td>
                <td>{{ sugerencia.mensaje }}</td>
                <td>{{ sugerencia.fecha }}</td>
                <td>
                  <button @click="marcarComoHecho(sugerencia.id)">Hecho</button>
                  <button @click="eliminarSugerencia(sugerencia.id)">Eliminar</button>
                </td>
              </tr>
            </tbody>
          </table>
          <div v-else class="mensaje-vacio">
            <v-icon color="primary" size="large">mdi-message-alert</v-icon>
            <p>No hay quejas o sugerencias disponibles.</p>
          </div>
        </div>
      </slot>
    </AdminNavBar>
  </template>
  
  <script>
  import axios from "axios";
  import AdminNavBar from "@/components/menus/AdminNavbar.vue";
  
  export default {
    data() {
      return {
        sugerencias: []
      };
    },
    components: {
      AdminNavBar
    },
    async mounted() {
      try {
        const respuesta = await axios.get("/api/sugerencias");
        this.sugerencias = respuesta.data;
      } catch (error) {
        console.error("Error al cargar las sugerencias:", error);
      }
    },
    methods: {
      AddSugerencia() {
        this.$router.push("/admin/sugerencias/add");
      },
      async marcarComoHecho(id) {
        // Lógica para marcar sugerencia como hecha
      },
      async eliminarSugerencia(id) {
        // Lógica para eliminar sugerencia
      }
    }
  };
  </script>
  
  <style scoped>
  .text-center {
    text-align: center;
    margin-bottom: 20px;
  }
  
  .contenedor-tabla {
    margin-top: 20px;
  }
  
  table {
    width: 100%;
    border-collapse: collapse;
    margin-top: 20px;
  }
  
  th, td {
    border: 1px solid #ccc;
    padding: 8px;
    text-align: left;
  }
  
  th {
    background-color: #f4f4f4;
  }
  
  button {
    margin: 0 5px;
    padding: 5px 10px;
    border: none;
    border-radius: 4px;
    cursor: pointer;
  }
  
  button:first-of-type {
    background-color: #28a745; /* Verde para "hecho" */
    color: white;
  }
  
  button:last-of-type {
    background-color: #dc3545; /* Rojo para "eliminar" */
    color: white;
  }
  
  button:hover {
    opacity: 0.9;
  }
  .mensaje-vacio {
  text-align: center;
  margin-top: 20px;
  color: #888;
}

.mensaje-vacio v-icon {
  margin-bottom: 10px;
}
  </style>