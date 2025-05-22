import { createStore } from "vuex";
import defaultValues from './defaultValues';
import { set } from "zod";

export default createStore({
    state: {
        user: defaultValues.userDefault,
        drawer: false,
        logged: false,
        menu: false,
        img_product: null,
        file: null,
    },
    mutations: {
        setUser(state, user){            
            state.user = user;
        },
        setDrawer(state, drawer){
            state.drawer = drawer;
        },
        setLogged(state, logged){
            state.logged = logged;
        },
        setImgProduct(state, img_product){
            state.img_product = img_product;
        },
        setFile(state, file){
            state.file = file;
        }

    },
    actions: {  

        }
});