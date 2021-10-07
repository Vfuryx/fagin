import {BasicFetchResult} from '/@/api/model/baseModel';

export type MenuParams = {
  menuName?: string;
  status?: string;
};

export type CreateMenuParams = {
  type: number;
  parent_id: number;
  name: string;
  title: string;
  sort: number;
  status: number;
  path: string;
  redirect?: string;
  frame_src?: string;
  current_active?: string;
  icon?: string;
  permission?: string;
  component?: string;
  method?: string;
  is_no_cache?: number;
  is_hide_child?: number;
  is_show?: number;
};


export type MenuItem = {
  type: number;
  parent_id: number;
  name: string;
  title: string;
  sort: number;
  status: number;
  path: string;
  redirect?: string;
  frame_src?: string;
  current_active?: string;
  icon?: string;
  permission?: string;
  component?: string;
  method?: string;
  is_no_cache?: number;
  is_hide_child?: number;
  is_show?: number;
};

export interface MenuListItem {
  id: string;
  sort: string;
  created_at: string;
  status: number;
  icon: string;
  component: string;
  permission: string;
}

export type MenuListGetResultModel = BasicFetchResult<MenuListItem>;
