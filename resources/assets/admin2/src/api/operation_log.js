import request from '@/utils/request'

const OperationLogApi = {
  OperationLogList: '/v1/operation/logs/',
  ShowOperationLog: '/v1/operation/logs/',
  DelOperationLog: '/v1/operation/logs/'
}

/**
 * 获取列表
 * parameter: {
 *     type: 0,
 * }
 * @param parameter
 * @returns {*}
 */
export function getOperationLogList (parameter) {
  return request({
    url: OperationLogApi.OperationLogList,
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
export function showOperationLog (id) {
  return request({
    url: OperationLogApi.ShowOperationLog + id,
    method: 'get'
    // headers: {
    //   'Content-Type': 'application/json;charset=UTF-8'
    // }
  })
}

/**
 * 删除
 * @param id ID
 * @returns {*}
 */
export function delOperationLog (id) {
  return request({
    url: OperationLogApi.DelOperationLog + id,
    method: 'delete'
    // headers: {
    //   'Content-Type': 'application/json;charset=UTF-8'
    // }
  })
}
