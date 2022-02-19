import { BasicPageParams, BasicFetchResult } from '/@/api/model/baseModel';

export type AccountParams = BasicPageParams & {
  username?: string;
  nickname?: string;
};

export type CreateAccountParams = {
  username: string;
  nick_name: string;
  password: string;
  roles: number[];
  department_id: number;
  email: string;
  phone: string;
  remark: string;
  status: number;
  sex: number;
  home_path?: string;
};

export type UpdateAccountParams = {
  username: string;
  nick_name: string;
  roles: number[];
  department_id: number;
  email: string;
  phone: string;
  remark: string;
  status: number;
  sex: number;
  home_path?: string;
};

export type RoleParams = {
  name?: string;
  status?: string;
};

export interface DeptListParams {
  name: string;
}

export interface DeptListItem {
  id: string;
  name: string;
  sort: string;
  created_at: string;
  remark: string;
  status: number;
}

export interface AccountListItem {
  id: string;
  username: string;
  email: string;
  nick_name: string;
  roles: any;
  create_at: string;
  remark: string;
  status: number;
}

export interface AccountGetResultModel {
  id: string;
  username: string;
  email: string;
  nick_name: string;
  phone: string;
  avatar: string;
  sex: number;
  roles: number[];
  create_at: string;
  remark: string;
  status: number;
}

export interface RoleListItem {
  id: string;
  name: string;
}

export type AccountListGetResultModel = BasicFetchResult<AccountListItem>;

export type DeptListGetResultModel = BasicFetchResult<DeptListItem>;

export type RoleListGetResultModel = RoleListItem[];
