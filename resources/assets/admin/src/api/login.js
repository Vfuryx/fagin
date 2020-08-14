import request from '@/utils/request'

// 获取验证码
export function getCodeImg() {
  return request({
    url: '/captcha',
    method: 'post'
  })
}
