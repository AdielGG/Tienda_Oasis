<template>
    <table class="table">
        <colgroup>
            <col class="col-star" />
            <col />
            <col />
            <col class="col-end" />
        </colgroup>
        <thead>
            <tr>
                <th class="row-start">Nombre</th>
                <th>Usuario</th>
                <th>Correo</th>
                <th class="row-end">Acciones</th>
            </tr>
        </thead>
        <tbody>
            <tr>
                <td class="row-start">Antonio Malanga Grande</td>
                <td>antonio</td>
                <td>an.grande@correo.co</td>
                <td class="row-end">
                    <v-btn variant="text" color="primary">
                        <v-icon
                            icon="mdi-account-edit-outline"
                            size="x-large"
                        ></v-icon>
                    </v-btn>
                    <v-btn variant="text" color="red">
                        <v-icon
                            icon="mdi-trash-can-outline"
                            size="x-large"
                        ></v-icon>
                    </v-btn>
                </td>
            </tr>
            <tr>
                <td class="row-start">Antonio Malanga Grande</td>
                <td>antonio</td>
                <td>an.grande@correo.co</td>
                <td class="row-end">
                    <v-btn variant="text" color="primary">
                        <v-icon
                            icon="mdi-account-edit-outline"
                            size="x-large"
                        ></v-icon>
                    </v-btn>
                    <v-btn variant="text" color="red">
                        <v-icon
                            icon="mdi-trash-can-outline"
                            size="x-large"
                        ></v-icon>
                    </v-btn>
                </td>
            </tr>
            <tr>
                <td class="row-start">Antonio Malanga Grande</td>
                <td>antonio</td>
                <td>an.grande@correo.co</td>
                <td class="row-end">
                    <v-btn variant="text" color="primary">
                        <v-icon
                            icon="mdi-account-edit-outline"
                            size="x-large"
                        ></v-icon>
                    </v-btn>
                    <v-btn variant="text" color="red">
                        <v-icon
                            icon="mdi-trash-can-outline"
                            size="x-large"
                        ></v-icon>
                    </v-btn>
                </td>
            </tr>
        </tbody>
    </table>
    <div class="text-center">
        <v-pagination
            v-model="page"
            :length="15"
            :total-visible="5"
        ></v-pagination>
    </div>
</template>
<script>
export default {
    data() {
        return {
            page: 1,
        };
    },
};
</script>
<style>
.table {
    border-collapse: collapse;
    table-layout: fixed;
    width: 97%;
    margin: 1em;
}
.table td,
.table th {
    text-align: center;
    font-size: 1.2em;
}
.table th {
    color: white;
}
.row-start {
    border-start-start-radius: 10px;
    border-end-start-radius: 10px;
}
.row-end {
    border-start-end-radius: 10px;
    border-end-end-radius: 10px;
}
.table thead {
    background: blue;
}
.table tbody tr:nth-child(odd) {
    background: #e3f2fd;
}
.text-center {
    align-items: center;
    text-align: center;
}
</style>
