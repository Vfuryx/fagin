import { defHttp } from '/@/utils/http/axios';
import {
  CreateMenuParams,
  MenuItem,
  MenuParams,
} from '/@/api/system/model/menuModel';

enum Api {
  MenuAll = '/menus/all',
  MenuList = '/menus',
  MenuDetail = '/menus/',
  CreateMenu = '/menus',
  UpdateMenu = '/menus/',
  DeleteMenu = '/menus/',
}

export const getMenuAll = () => defHttp.get({ url: Api.MenuAll });

export const getMenuList = (params?: MenuParams) => defHttp.get({ url: Api.MenuList, params });

export const getMenuDetail = (id: number) =>
  defHttp.get<MenuItem>({ url: Api.MenuDetail + id.toString() });

export const createMenu = (params: CreateMenuParams) =>
  defHttp.post({ url: Api.CreateMenu, params });

export const updateMenu = (id: number, params: CreateMenuParams) =>
  defHttp.put({ url: Api.UpdateMenu + id.toString(), params });

export const deleteMenu = (id: number) => defHttp.delete({ url: Api.DeleteMenu + id.toString() });
