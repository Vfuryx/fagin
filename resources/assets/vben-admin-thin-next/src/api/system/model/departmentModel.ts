export interface DeptListItem {
  id: number;
  name: string;
  sort: number;
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

export type DeptListGetResultModel = DeptListItem[];
