import request from '@/utils/request'

const adminApi = {
  AdminList: '/v1/admins/',
  AddAdmin: '/v1/admins/',
  UpdateAdmin: '/v1/admins/',
  ShowAdmin: '/v1/admins/',
  DelAdmin: '/v1/admins/',
  LogoutAdmin: '/v1/admins/logout/',
  AdminUsernameExist: '/v1/admins/username/',
  AdminResetPassword: '/v1/admins/'
}

/**
 * 获取列表
 * parameter: {
 *     type: 0,
 * }
 * @param parameter
 * @returns {*}
 */
export function getAdminList (parameter) {
  return request({
    url: adminApi.AdminList,
    method: 'get',
    params: parameter
    // headers: {
    //   'Content-Type': 'application/json;charset=UTF-8'
    // }
  })
}

/**
 * 获取所有角色
 * parameter: {
 *     type: 0,
 * }
 * @param parameter
 * @returns {*}
 */
export function getRoleAll (parameter) {
  return request({
    url: adminApi.RoleAll,
    method: 'post',
    params: parameter
    // headers: {
    //   'Content-Type': 'application/json;charset=UTF-8'
    // }
  })
}

/**
 * 展示详情
 * @param id ID
 * @returns {*}
 */
export function showAdmin (id) {
  return request({
    url: adminApi.ShowAdmin + id,
    method: 'get'
    // headers: {
    //   'Content-Type': 'application/json;charset=UTF-8'
    // }
  })
}

/**
 * 新增
 * data: {
 *
 * }
 * @param data
 * @returns {*}
 */
export function addAdmin (data) {
  return request({
    url: adminApi.AddAdmin,
    method: 'post',
    data: data,
    headers: {
      'Content-Type': 'application/json;charset=UTF-8'
    }
  })
}

/**
 * 修改
 * @param id
 * @param data
 * @returns {*}
 */
export function updateAdmin (id, data) {
  return request({
    url: adminApi.UpdateAdmin + id,
    method: 'put',
    data: data,
    headers: {
      'Content-Type': 'application/json;charset=UTF-8'
    }
  })
}

/**
 * 删除
 * @param id ID
 * @returns {*}
 */
export function delAdmin (id) {
  return request({
    url: adminApi.DelAdmin + id,
    method: 'delete'
    // headers: {
    //   'Content-Type': 'application/json;charset=UTF-8'
    // }
  })
}

/**
 * 强制用户下线
 * @id user_id
 * @returns {*}
 */
export function logoutAdmin (id) {
  return request({
    url: adminApi.LogoutAdmin + id,
    method: 'post'
    // headers: {
    //   'Content-Type': 'application/json;charset=UTF-8'
    // }
  })
}

/**
 * 角色 key 是否存在
 * parameter: {
 *     username: "name",
 * }
 * @param parameter
 * @returns {*}
 */
export function adminUsernameExist (parameter) {
  return request({
    url: adminApi.AdminUsernameExist,
    method: 'get',
    params: parameter
    // headers: {
    //   'Content-Type': 'application/json;charset=UTF-8'
    // }
  })
}

/**
 * 重置密码
 * @param id
 * @param data
 * @returns {*}
 */
export function adminResetPassword (id, data) {
  return request({
    url: adminApi.AdminResetPassword + id + '/reset/',
    method: 'put',
    data: data,
    headers: {
      'Content-Type': 'application/json;charset=UTF-8'
    }
  })
}
