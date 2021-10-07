import { defHttp } from '/@/utils/http/axios';

import {
  OperLogDetail,
  OperLogPageListGetResultModel,
  OperLogParams,
} from '/@/api/system/model/operLogModel';

enum Api {
  OperLogList = '/operation/logs',
  OperLogDetail = '/operation/logs/',
}

export const getOperLogListByPage = (params?: OperLogParams) =>
  defHttp.get<OperLogPageListGetResultModel>({ url: Api.OperLogList, params });

export const getOperLogDetail = (id: number) =>
  defHttp.get<OperLogDetail>({ url: Api.OperLogDetail + id.toString() });
