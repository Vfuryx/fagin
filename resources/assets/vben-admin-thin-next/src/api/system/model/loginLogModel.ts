import { BasicFetchResult, BasicPageParams } from '/@/api/model/baseModel';

export type LoginLogParams = BasicPageParams & {
  login_name?: string;
  time?: any;
};

export interface LoginLogListItem {
  id: string;
  module: string;
  method: string;
  user: string;
  path: string;
  status: number;
}

export interface LoginLogDetail {
  id: string;
  orderNo: string;
  createTime: string;
  status: number;
  icon: string;
  component: string;
  permission: string;
}

export type LoginLogPageListGetResultModel = BasicFetchResult<LoginLogListItem>;
