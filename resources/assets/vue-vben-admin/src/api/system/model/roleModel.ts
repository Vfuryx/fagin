import { BasicFetchResult } from '/@/api/model/baseModel';

export type RoleParams = {
  roleName?: string;
  status?: string;
};

export type CreateRoleParams = {
  name: string;
  key: string;
  remark: string;
  status: number;
  menu_ids?: number[];
};

export type UpdateRoleParams = {
  name: string;
  remark: string;
  status: number;
  menu_ids?: number[];
};

export interface RoleListItem {
  id: number;
  sort: string;
  created_at: string;
  status: number;
  icon: string;
  component: string;
  permission: string;
}

export interface MenuAllItem {
  id: string;
  sort: string;
  created_at: string;
  status: number;
  icon: string;
  component: string;
  permission: string;
}

export type RoleListGetResultModel = BasicFetchResult<RoleListItem>;

export type RolePageListGetResultModel = BasicFetchResult<RoleListItem>;

export type MenuAllGetResultModel = BasicFetchResult<MenuAllItem>;
