import { BasicColumn } from '/@/components/Table';
import { FormSchema } from '/@/components/Table';
import { DescItem } from '/@/components/Description';
import { h } from 'vue';
import { CodeEditor } from '/@/components/CodeEditor';

export const columns: BasicColumn[] = [
  {
    title: '系统模块',
    dataIndex: 'module',
    width: 100,
  },
  {
    title: '请求方式',
    dataIndex: 'method',
    width: 50,
  },
  {
    title: '操作人员',
    dataIndex: 'user',
    width: 100,
  },
  {
    title: '请求地址',
    dataIndex: 'path',
    width: 200,
  },
  {
    title: '操作日期',
    dataIndex: 'created_at',
    width: 100,
  },
];

export const searchFormSchema: FormSchema[] = [
  {
    field: 'path',
    label: '请求地址',
    component: 'Input',
    colProps: { span: 8 },
  },
  {
    field: 'method',
    label: '请求方式',
    component: 'Select',
    componentProps: {
      options: [
        { label: 'GET', value: 'GET' },
        { label: 'POST', value: 'POST' },
        { label: 'PUT', value: 'PUT' },
        { label: 'DELETE', value: 'DELETE' },
      ],
    },
    colProps: { span: 5 },
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

export const detailSchema: DescItem[] = [
  {
    field: 'id',
    label: '日志编号',
    labelMinWidth: 56,
  },
  {
    field: 'created_at',
    label: '操作日期',
  },
  {
    field: 'ip',
    label: 'IP地址',
  },
  {
    field: 'latency_time',
    label: '请求耗时',
  },
  {
    field: 'location',
    label: '访问位置',
  },
  {
    field: 'method',
    label: '请求方式',
  },
  {
    field: 'module',
    label: '系统模块',
  },
  {
    field: 'operation',
    label: '操作类型',
  },
  {
    field: 'path',
    label: '请求路径',
  },
  {
    field: 'user',
    label: '操作人员',
  },
  {
    field: 'user_agent',
    label: '用户代理',
  },
  {
    field: 'input',
    label: '参数',
    render: (_, data) => {
      return h(CodeEditor, {
        value: data.input,
        mode: 'application/json',
        readonly: true,
      });
    },
  },
];
