import {
  AccountParams,
  AccountListGetResultModel,
  DeptListGetResultModel,
  RoleParams,
  RoleListGetResultModel,
  DeptListParams,
  AccountGetResultModel,
  CreateAccountParams,
  UpdateAccountParams,
} from '/@/api/system/model/accountModel';
import { defHttp } from '/@/utils/http/axios';

enum Api {
  AccountList = '/accounts',
  CreateAccount = '/accounts',
  UpdateAccount = '/accounts/',
  DeleteAccount = '/accounts/',
  GetAccount = '/accounts/',
  DeptList = '/accounts/departments',
  IsAccountExist = '/accounts/exists',
  GetAllRoleList = '/accounts/roles',
  SetAccountStatus = '/accounts/',
  setAccountPassword = '/accounts/',
  LogoutAccount = '/accounts/',
  getRoleRoute = '/accounts/role-route',
}

export const getAccountList = (params: AccountParams) =>
  defHttp.get<AccountListGetResultModel>({ url: Api.AccountList, params });

export const getAccount = (id: number) =>
  defHttp.get<AccountGetResultModel>({ url: Api.GetAccount + id.toString() });

export const createAccount = (params: CreateAccountParams) =>
  defHttp.post({ url: Api.CreateAccount, params });

export const updateAccount = (id: number, params: UpdateAccountParams) =>
  defHttp.put({ url: Api.UpdateAccount + id.toString(), params });

export const deleteAccount = (id: number) =>
  defHttp.delete({ url: Api.DeleteAccount + id.toString() });

export const getDeptList = (params?: DeptListParams) =>
  defHttp.get<DeptListGetResultModel>({ url: Api.DeptList, params });

export const getAllRoleList = (params?: RoleParams) =>
  defHttp.get<RoleListGetResultModel>({ url: Api.GetAllRoleList, params });

export const isAccountExist = (username: string, id?: number) =>
  defHttp.get({ url: Api.IsAccountExist, params: { username, id } }, { errorMessageMode: 'none' });

export const setAccountStatus = (id: number, status: number) =>
  defHttp.put({ url: Api.SetAccountStatus + id.toString() + '/status', params: { status } });

export const setAccountPassword = (id: number, password: number) =>
  defHttp.put({ url: Api.setAccountPassword + id.toString() + '/pwd', params: { password } });

export const logoutAccount = (id: number) =>
  defHttp.post({ url: Api.LogoutAccount + id.toString() + '/logout' });

export const getRoleRoute = (roles: number[]) =>
  defHttp.get({ url: Api.getRoleRoute, params: { roles: roles } });
