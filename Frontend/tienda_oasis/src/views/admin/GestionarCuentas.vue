<template>
    <table>
        <tr>
            <th>
                Nombre
            </th>
            <th>
                Apellido
            </th>
            <th>
                Usuario
            </th>
            <th>
                Rol
            </th>
            <th>
                Acciones
            </th>
        </tr>
        <tr v-for="user in users" :key="user.id">
            <td>
                {{ user.name }}
            </td> 
            <td>
                {{ user.lastname }}
            </td> 
            <td>
                {{ user.username }}
            </td> 
            <td>
                {{ user.role }}
            </td> 
            <td>
                <button>
                    editar
                </button>
                <button>
                    eliminar
                </button>
            </td> 

        </tr>
    </table>   
</template>

<script>
import axios from 'axios';

export default{
    data(){
        return{
            users: null
        }
    },
    mounted() {
        this.getUsers()
    },
    methosd: {
        getUsers(){
            this.users = async ()=> {
               return  await axios.get("localhost:8080/users") 
            }
                
        }
    }
}


</script>