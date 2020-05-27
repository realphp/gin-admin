
export const state = () => ({
    userInfo: {},
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
            state.userInfo = user
            state.isLogin = true
        } else {
            state.userInfo = {}
            state.isLogin = false
        }
    },

}


export const getters = {
    userInfo(state) {
        return state.userInfo
    },

}
