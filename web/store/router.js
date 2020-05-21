// import { asyncRouterHandle } from '@/utils/asyncRouter';



const formatRouter = (routes) => {
    routes && routes.map(item => {
        item.meta.hidden = item.hidden
        if (item.children && item.children.length > 0) {
            formatRouter(item.children)
        }
    })
}


export const state = () => ({
    asyncRouters: []
})

export const mutations = {
    // 设置动态路由
    setAsyncRouter(state, asyncRouters) {
        state.asyncRouters = asyncRouters
    }
}

export const actions = {
    // 从后台获取动态路由
    async SetAsyncRouter({ commit }) {
        const baseRouter = [{
            path: '/layout',
            name: 'layout',
            component: "view/layout/index.vue",
            meta: {
                title: "底层layout"
            },
            children: []
        }]
        const asyncRouterRes = await this.$axios.post("menu/list");
        const asyncRouter = asyncRouterRes.data.Menus
        formatRouter(asyncRouter)

        asyncRouter.push({
            path: "404",
            name: "404",
            hidden: true,
            meta: {
                title: "迷路了*。*",
            },
            component: 'view/error/index.vue'
        })
        baseRouter[0].children = asyncRouter
        // baseRouter.push({
        //     path: '*',
        //     redirect: '/layout/404'

        // })
        // asyncRouterHandle(baseRouter)
        commit('setAsyncRouter', baseRouter)
        return true
    }
}

export const getters = {
    // 获取动态路由
    asyncRouters(state) {
        return state.asyncRouters
    }
}