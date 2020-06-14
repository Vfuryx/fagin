import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/login',
    method: 'post',
    data
  })
}

export function getInfo() {
  return request({
    url: '/v1/us/info',
    method: 'get'
  })
}

export function logout() {
  return request({
    url: '/v1/user/logout',
    method: 'post'
  })
}

export function listUser() {
  return request({
    url: '/v1/user/',
    method: 'get'
  })
}

// 查询角色列表
export function userRolesList(query) {
  return request({
    url: '/v1/role/list',
    method: 'post'
  })
}

// 新增用户
export function addUser(data) {
  return request({
    url: '/v1/user/',
    method: 'post',
    data: data
  })
}

// 查询用户详细
export function getUser(id) {
  return request({
    url: '/v1/user/' + id,
    method: 'get'
  })
}

// 新增用户
export function updateUser(data) {
  return request({
    url: '/v1/user/' + data.id,
    method: 'put',
    data: data
  })
}

// 用户状态修改
export function changeUserStatus(id, status) {
  const data = {
    status
  }
  return request({
    url: '/v1/user/' + id + '/status/',
    method: 'put',
    data: data
  })
}
// 用户密码修改
export function resetUserPwd(id, pw) {
  const data = {
    password: pw
  }
  return request({
    url: '/v1/user/' + id + '/reset/',
    method: 'put',
    data: data
  })
}

// 删除用户
export function delUser(id) {
  return request({
    url: '/v1/user/' + id,
    method: 'delete'
  })
}

// 删除用户组
export function delUsers(ids) {
  return request({
    url: '/v1/user/',
    method: 'delete',
    data: {
      ids: ids
    }
  })
}
