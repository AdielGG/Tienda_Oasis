import { createStore } from "vuex";
import defaultValues from './defaultValues';

export default createStore({
    state: {
        user: defaultValues.userDefault,
        drawer: false,
        logged: false,
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
        }
    },
    actions: {  

        }
});