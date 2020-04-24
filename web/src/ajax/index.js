import axios from 'axios'

// 创建一个axios实例
const service = axios.create({
    baseURL: 'http://localhost:5800', //process.env.VUE_APP_API,
    timeout: 10000
})

// 请求拦截器
service.interceptors.request.use(
    config => {
        // 在请求发出前做一些处理
        if (!/^https:\/\/|http:\/\//.test(config.url)) {
            const token = sessionStorage.getItem('token')
            if (token && token !== 'undefined') {
                config.headers.Authorization = token
            }
        }
        return config
    },
    error => {
        // 发送失败
        console.log(error)
        Promise.reject(error)
    }
)

// 设置响应拦截器
service.interceptors.response.use(
    response => {
        const res = response.data
        /*
        需要统一处理的业务状态码
        const { code } = res
        if (code !== 401) {
          // 业务错误信息统一的错误处理
          console.log(res.msg)
        }*/
        return res
    },
    error => {
        // http请求异常统一处理
        if (error && error.response) {
            switch (error.response.status) {
                case 400:
                    error.message = '请求错误'
                    break
                case 401:
                    error.message = '未授权，请登录'
                    break
                case 403:
                    error.message = '拒绝访问'
                    break
                case 404:
                    error.message = `请求地址出错: ${error.response.config.url}`
                    break
                case 408:
                    error.message = '请求超时'
                    break
                case 500:
                    error.message = '服务器内部错误'
                    break
                case 501:
                    error.message = '服务未实现'
                    break
                case 502:
                    error.message = '网关错误'
                    break
                case 503:
                    error.message = '服务不可用'
                    break
                case 504:
                    error.message = '网关超时'
                    break
                case 505:
                    error.message = 'HTTP版本不受支持'
                    break
                default:
                    break
            }
        }
        return Promise.reject(error)
    }
)

export default service
