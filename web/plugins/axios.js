
export default function ({ app, $axios }) {
  $axios.onRequest(config => {
    // config.headers.Authorization = 'Bearer ' + app.$cookies.get('token')
    config.headers.Authorization = app.$cookies.get('token')
    return config
  })
}


