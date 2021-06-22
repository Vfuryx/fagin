import request from '@/utils/request'

const permissionApi = {
  GroupPermission: '/v1/group_permissions',
  PermissionGroupList: '/v1/permission/groups/',
  PermissionGroupAll: '/v1/permission/groups/all',
  AddPermissionGroup: '/v1/permission/groups/',
  UpdatePermissionGroup: '/v1/permission/groups/',
  ShowPermissionGroup: '/v1/permission/groups/',
  DelPermissionGroup: '/v1/permission/groups/',

  PermissionList: '/v1/permissions/',
  AddPermission: '/v1/permissions/',
  UpdatePermission: '/v1/permissions/',
  ShowPermission: '/v1/permissions/',
  DelPermission: '/v1/permissions/'
}

/**
 * 获取列表
 * parameter: {
 *     type: 0,
 * }
 * @param parameter
 * @returns {*}
 */
export function getPermissionGroupList (parameter) {
  return request({
    url: permissionApi.PermissionGroupList,
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
export function getPermissionGroupAll (parameter) {
  return request({
    url: permissionApi.PermissionGroupAll,
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
export function showPermissionGroup (id) {
  return request({
    url: permissionApi.ShowPermissionGroup + id,
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
export function addPermissionGroup (parameter) {
  return request({
    url: permissionApi.AddPermissionGroup,
    method: 'post',
    data: parameter,
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
export function updatePermissionGroup (id, data) {
  return request({
    url: permissionApi.UpdatePermissionGroup + id,
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
export function delPermissionGroup (id) {
  return request({
    url: permissionApi.DelPermissionGroup + id,
    method: 'delete'
    // headers: {
    //   'Content-Type': 'application/json;charset=UTF-8'
    // }
  })
}

/**
 * 获取列表
 * parameter: {
 *     type: 0,
 * }
 * @param parameter
 * @returns {*}
 */
export function getPermissionList (parameter) {
  return request({
    url: permissionApi.PermissionList,
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
export function showPermission (id) {
  return request({
    url: permissionApi.ShowPermission + id,
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
export function addPermission (parameter) {
  return request({
    url: permissionApi.AddPermission,
    method: 'post',
    data: parameter,
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
export function updatePermission (id, data) {
  return request({
    url: permissionApi.UpdatePermission + id,
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
export function delPermission (id) {
  return request({
    url: permissionApi.DelPermission + id,
    method: 'delete'
    // headers: {
    //   'Content-Type': 'application/json;charset=UTF-8'
    // }
  })
}

// 分组权限
export function getGroupPermission (parameter) {
  return request({
    url: permissionApi.GroupPermission,
    method: 'get',
    params: parameter
  })
}
