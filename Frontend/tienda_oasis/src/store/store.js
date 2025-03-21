export default {
    namespaced: true,
    state: {
        user: null,
        logued: false
    },
    mutations: {
        setUser(state, user){
            state.user = user;
        }
    }
}