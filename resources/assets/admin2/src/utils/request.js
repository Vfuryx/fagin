import axios from 'axios'
import store from '@/store'
import storage from 'store'
import notification from 'ant-design-vue/es/notification'
import { VueAxios } from './axios'
import { ACCESS_TOKEN } from '@/store/mutation-types'

// 创建 axios 实例
const request = axios.create({
  // API 请求的默认前缀
  baseURL: process.env.VUE_APP_API_BASE_URL,
  timeout: 6000 // 请求超时时间
})

// 异常拦截处理器
const errorHandler = (error) => {
  if (error.response) {
    const data = error.response.data
    // 从 localstorage 获取 token
    const token = storage.get(ACCESS_TOKEN)
    if (error.response.status === 403) {
      notification.error({
        message: 'Forbidden',
        description: data.message
      })
    }
    if (error.response.status === 401 && !(data.result && data.result.isLogin)) {
      notification.error({
        message: 'Unauthorized',
        description: 'Authorization verification failed'
      })
      if (!isEmpty(token)) {
        store.dispatch('Logout').then(() => {
          setTimeout(() => {
            window.location.reload()
          }, 1500)
        })
      }
    } else {
      notification.error({
        message: 'Forbidden',
        description: error.response.statusText + ',error_code：' + error.response.status
      })
    }
  } else {
    notification.error({
      message: 'Forbidden',
      description: '网络异常，请检查！'
    })
  }
  return Promise.reject(error)
}

// request interceptor
request.interceptors.request.use(config => {
  const token = storage.get(ACCESS_TOKEN)
  // 如果 token 存在
  // 让每个请求携带自定义 token 请根据实际情况自行修改
  if (token) {
    config.headers['Authorization'] = 'Bearer ' + token
  }
  return config
}, errorHandler)

// response interceptor
request.interceptors.response.use((response) => {
  if (response.status !== 200) {
    return notification.error({
      message: 'Forbidden',
      description: response.statusText
    })
  }
  // 如果出现401 代表token 失效
  if (response.data.code === 401) {
    notification.error({
      message: 'Forbidden',
      description: response.data.msg
    })
    storage.remove(ACCESS_TOKEN)
    this.$router.push('/auth/login')
  }

  // 如果出现402 代表接口无权限 失效
  if (response.data.code === 402) {
    return notification.error({
      message: 'Forbidden',
      description: response.data.msg
    })
  }

  // 429 代表请求太频繁
  if (response.data.code === 429) {
    return notification.error({
      message: 'Forbidden',
      description: '您请求太频繁了，请休息一会'
    })
  }

  // 令牌无效
  if (response.data.code === 20003 || response.data.code === 20012) {
    notification.error({
      message: 'Forbidden',
      description: '登录失效'
    })
    store.dispatch('Logout').then(() => {
      setTimeout(() => {
        window.location.reload()
      }, 1500)
    })
  }

  // 刷新了token 则重新存放
  if (!isEmpty(response.headers.authorization)) {
    const token = response.headers.authorization.split(' ')[1]
    storage.set(ACCESS_TOKEN, token, 7 * 24 * 60 * 60 * 1000)
  }

  return response.data
}, errorHandler)

const installer = {
  vm: {},
  install (Vue) {
    Vue.use(VueAxios, request)
  }
}

// 判断是否为空
export function isEmpty (str) {
  return str === '' || str === null || str === undefined
}

export default request

export {
  installer as VueAxios,
  request as axios
}
