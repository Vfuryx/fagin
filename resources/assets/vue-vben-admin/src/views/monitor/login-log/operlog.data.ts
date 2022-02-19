import { BasicColumn } from '/@/components/Table';
import { FormSchema } from '/@/components/Table';
import { DescItem } from '/@/components/Description';

export const columns: BasicColumn[] = [
  {
    title: '登录名称',
    dataIndex: 'login_name',
    width: 100,
  },
  {
    title: '登录地址',
    dataIndex: 'ip',
    width: 100,
  },
  {
    title: '登录地点',
    dataIndex: 'location',
    width: 100,
  },
  {
    title: '浏览器',
    dataIndex: 'browser',
    width: 100,
  },
  {
    title: '操作系统',
    dataIndex: 'os',
    width: 100,
  },
  {
    title: '登录日期',
    dataIndex: 'created_at',
    width: 150,
  },
];

export const searchFormSchema: FormSchema[] = [
  {
    field: 'login_name',
    label: '登录名称',
    component: 'Input',
    colProps: { span: 8 },
  },
  {
    field: 'time',
    label: '时间范围',
    component: 'RangePicker',
    colProps: { span: 11 },
    componentProps: {
      showTime: true,
    },
  },
];

export const formSchema: FormSchema[] = [];

export const detailSchema: DescItem[] = [];
