import { BasicFetchResult } from '/@/api/model/baseModel';

export interface DeptListItem {
  id: string;
  name: string;
  sort: string;
  created_at: string;
  remark: string;
  status: number;
  parent_id: number;
}

export type CreateDeptParams = {
  name: string;
  sort: number;
  remark: string;
  status: number;
  parent_id: number;
};

export type DeptListGetResultModel = BasicFetchResult<DeptListItem>;
