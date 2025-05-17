<template>
    <table class="table">
        <colgroup>
            <col class="col-star" />
            <col />
            <col />
            <col />
            <col class="col-end" />
        </colgroup>
        <thead>
            <tr>
                <th class="row-start">Imagen</th>
                <th>Nombre</th>
                <th>Descripción</th>
                <th>Precio</th>
                <th>Cantidad</th>
                <th class="row-end">Acciones</th>
            </tr>
        </thead>
        <tbody>

            <tr>
                <td><img src="@/assets/background.jpg" alt="background" width="100px"></td>
                <td>Prducto1</td>
                <td>Descripción 1</td>
                <td>2333</td>
            <td>50</td>
            <td class="row-end">
                <v-btn variant="text" color="primary">
                    <v-icon
                        icon="mdi-pencil-outline"
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
            <td class="row-start"><img src="@/assets/background.jpg" alt="background" width="100px"></td>
            <td>Prducto1</td>
            <td>Descripción 1</td>
            <td>2333</td>
            <td>50</td>
            <td class="row-end">
                <v-btn variant="text" color="primary">
                    <v-icon
                        icon="mdi-pencil-outline"
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
            <td class="row-start"><img src="@/assets/background.jpg" alt="background" width="100px"></td>
            <td>Prducto1</td>
            <td>Descripción 1</td>
            <td>2333</td>
            <td>50</td>
            <td class="row-end">
                <v-btn variant="text" color="primary">
                    <v-icon
                    icon="mdi-pencil-outline"
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
            <td class="row-start"><img src="@/assets/background.jpg" alt="background" width="100px"></td>
            <td>Prducto1</td>
            <td>Descripción 1</td>
            <td>2333</td>
            <td>50</td>
            <td class="row-end">
                <v-btn variant="text" color="primary">
                    <v-icon
                        icon="mdi-pencil-outline"
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
</template>

<script>
export default {

}
</script>

<style>

</style>