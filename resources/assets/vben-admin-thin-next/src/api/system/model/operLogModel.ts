import { BasicFetchResult, BasicPageParams } from '/@/api/model/baseModel';

export type OperLogParams = BasicPageParams & {
  roleName?: string;
  status?: string;
};

export interface OperLogListItem {
  id: string;
  module: string;
  method: string;
  user: string;
  path: string;
  status: number;
}

export interface OperLogDetail {
  id: string;
  orderNo: string;
  createTime: string;
  status: number;
  icon: string;
  component: string;
  permission: string;
}

export type OperLogPageListGetResultModel = BasicFetchResult<OperLogListItem>;
