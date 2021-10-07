import { defHttp } from '/@/utils/http/axios';
import {
  RoleParams,
  CreateRoleParams,
  RoleListGetResultModel,
  UpdateRoleParams,
  MenuAllGetResultModel,
} from '/@/api/system/model/roleModel';

enum Api {
  RoleList = '/roles',
  GetRoleDetail = '/roles/',
  CreateRole = '/roles',
  UpdateRole = '/roles/',
  DeleteRole = '/roles/',
  SetRoleStatus = '/roles/',
  GetMenuAll = '/roles/menus',
}

export const getRoleListByPage = (params?: RoleParams) =>
  defHttp.get<RoleListGetResultModel>({ url: Api.RoleList, params });

export const getRoleList = (params?: RoleParams) =>
  defHttp.get<RoleListGetResultModel>({ url: Api.RoleList, params });

export const getRoleDetail = (id: number) =>
  defHttp.get({ url: Api.GetRoleDetail + id.toString() });

export const createRole = (params: CreateRoleParams) =>
  defHttp.post({ url: Api.CreateRole, params });

export const updateRole = (id: number, params: UpdateRoleParams) =>
  defHttp.put({ url: Api.UpdateRole + id.toString(), params });

export const deleteRole = (id: number) => defHttp.delete({ url: Api.DeleteRole + id.toString() });

export const setRoleStatus = (id: number, status: number) =>
  defHttp.put({ url: Api.SetRoleStatus + id.toString() + '/status', params: { status } });

export const getMenuAll = () => defHttp.get<MenuAllGetResultModel>({ url: Api.GetMenuAll });
