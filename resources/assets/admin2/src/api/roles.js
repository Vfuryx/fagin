import request from '@/utils/request'

const roleApi = {
  RoleAll: '/v1/role/all',
  RoleKeyExist: '/v1/role/key',
  RoleList: '/v1/role/',
  AddRole: '/v1/role/',
  UpdateRole: '/v1/role/',
  ShowRole: '/v1/role/',
  DelRole: '/v1/role/'
}

/**
 * 获取列表
 * parameter: {
 *     type: 0,
 * }
 * @param parameter
 * @returns {*}
 */
export function getRoleList (parameter) {
  return request({
    url: roleApi.RoleList,
    method: 'get',
    params: parameter
    // headers: {
    //   'Content-Type': 'application/json;charset=UTF-8'
    // }
  })
}

/**
 * 获取所有
 * parameter: {
 *     type: 0,
 * }
 * @param parameter
 * @returns {*}
 */
export function getRoleAll (parameter) {
  return request({
    url: roleApi.RoleAll,
    method: 'get',
    params: parameter
    // headers: {
    //   'Content-Type': 'application/json;charset=UTF-8'
    // }
  })
}

/**
 * 角色 key 是否存在
 * parameter: {
 *     key: "0",
 * }
 * @param parameter
 * @returns {*}
 */
export function roleKeyExist (parameter) {
  return request({
    url: roleApi.RoleKeyExist,
    method: 'get',
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
export function showRole (id) {
  return request({
    url: roleApi.ShowRole + id,
    method: 'get'
    // headers: {
    //   'Content-Type': 'application/json;charset=UTF-8'
    // }
  })
}

/**
 * 新增
 * parameter: {
 *
 * }
 * @param parameter
 * @returns {*}
 */
export function addRole (data) {
  return request({
    url: roleApi.AddRole,
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
export function updateRole (id, data) {
  return request({
    url: roleApi.UpdateRole + id,
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
export function delRole (id) {
  return request({
    url: roleApi.DelRole + id,
    method: 'delete'
    // headers: {
    //   'Content-Type': 'application/json;charset=UTF-8'
    // }
  })
}
