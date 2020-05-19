export default function ({ app, redirect, route }) {
    let redirectURL = '/login';
    let token = app.$cookies.get('token')
    // 判断是否跳到login
    if (route.path === '/login' && token) {
        redirect('/')
    }
    // 通过外部登录系统
    getTokenFromUrl(app, route.query)
    //需要进行权限判断的页面开头
    if (!token) {
        redirect(redirectURL)
    }
}

// 通过外部登录系统，兼容yzc的老系统
function getTokenFromUrl(app, query) {
    if (query.access_token) {
        app.$cookies.set('token', query.access_token)
    }
}