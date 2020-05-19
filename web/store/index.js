
export const state = () => ({
    user: {},
    isLogin: false,
})

export const actions = {

    setUser({ commit }, res) {
        commit('userStatus', res.data)
    },

    async nuxtServerInit({ commit }, { app }) {
        if (app.$cookies.get('token')) {
            const { data: res } = await this.$axios.get("user/info");
            if (res.code !== 200) return console.log(res.msg);
            commit("userStatus", res.data);
        }
    }
}

export const mutations = {
    userStatus(state, user) {
        if (user) {
            state.user = user
            state.isLogin = true
        } else {
            state.user = {}
            state.isLogin = false
        }
    },

}
