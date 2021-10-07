import {
  DeptListItem,
  DeptListGetResultModel,
  CreateDeptParams,
} from '/@/api/system/model/departmentModel';
import { defHttp } from '/@/utils/http/axios';

enum Api {
  DeptList = '/departments',
  CreateDept = '/departments',
  UpdateDept = '/departments/',
  DeleteDept = '/departments/',
}

export const getDeptList = (params?: DeptListItem) =>
  defHttp.get<DeptListGetResultModel>({ url: Api.DeptList, params });

export const createDept = (params: CreateDeptParams) =>
  defHttp.post({ url: Api.CreateDept, params });

export const updateDept = (id: number, params: CreateDeptParams) =>
  defHttp.put({ url: Api.UpdateDept + id.toString(), params });

export const deleteDept = (id: number) => defHttp.delete({ url: Api.DeleteDept + id.toString() });
