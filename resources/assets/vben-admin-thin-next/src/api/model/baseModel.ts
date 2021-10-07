export interface BasicPageParams {
  page: number;
  page_size: number;
}

export interface BasicFetchResult<T extends any> {
  items: T[];
  total: number;
}
