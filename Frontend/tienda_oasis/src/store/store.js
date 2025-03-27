export default {
    namespaced: true,
    state: {
        user: null,
        logued: false,
        drawer: false,
    },
    mutations: {
        setUser(state, user){
            state.user = user;
        },
        setLogued(state,logued){
            state.logued = false
        },
        setDrawer(state,drawer){
            state.drawer = drawer
        }
    },
    getters:{
        getUser(state){
            return state.user
        },
        getLogued(state){
            return state.logued
        },
        getDrawer(state){
            return state.drawer
        }
    }
}
