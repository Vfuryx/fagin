import { defHttp } from '/@/utils/http/axios';

import {
  LoginLogDetail,
  LoginLogPageListGetResultModel,
  LoginLogParams,
} from '/@/api/system/model/loginLogModel';

enum Api {
  LoginLogList = '/login-logs',
  LoginLogDetail = '/login-logs/',
}

export const getLoginLogListByPage = (params?: LoginLogParams) =>
  defHttp.get<LoginLogPageListGetResultModel>({ url: Api.LoginLogList, params });

export const getLoginLogDetail = (id: number) =>
  defHttp.get<LoginLogDetail>({ url: Api.LoginLogDetail + id.toString() });
