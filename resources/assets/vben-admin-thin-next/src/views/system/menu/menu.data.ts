import { BasicColumn } from '/@/components/Table';
import { FormSchema } from '/@/components/Table';
import { h } from 'vue';
import { Tag } from 'ant-design-vue';
import { Icon } from '/@/components/Icon';

export const columns: BasicColumn[] = [
  {
    title: '菜单名称',
    dataIndex: 'title',
    width: 200,
    align: 'left',
  },
  {
    title: '图标',
    dataIndex: 'icon',
    width: 55,
    customRender: ({ record }) => {
      return h(Icon, { icon: record.icon });
    },
  },
  {
    title: '权限标识',
    dataIndex: 'permission',
    width: 200,
  },
  {
    title: '路径',
    dataIndex: 'path',
  },
  {
    title: '组件',
    dataIndex: 'component',
  },
  {
    title: '排序',
    dataIndex: 'sort',
    width: 50,
  },
  {
    title: '状态',
    dataIndex: 'status',
    width: 60,
    customRender: ({ record }) => {
      const status = record.status;
      const enable = ~~status !== 0;
      const color = enable ? 'green' : 'red';
      const text = enable ? '启用' : '停用';
      return h(Tag, { color: color }, () => text);
    },
  },
  {
    title: '创建时间',
    dataIndex: 'created_at',
    width: 168,
  },
];

const isDir = (type: number) => type === 0;
const isMenu = (type: number) => type === 1;
const isButton = (type: number) => type === 2;

export const searchFormSchema: FormSchema[] = [
  {
    field: 'title',
    label: '菜单名称',
    component: 'Input',
    colProps: { span: 8 },
  },
  {
    field: 'status',
    label: '状态',
    component: 'Select',
    componentProps: {
      options: [
        { label: '启用', value: 1 },
        { label: '停用', value: 0 },
      ],
    },
    colProps: { span: 8 },
  },
];

export const formSchema: FormSchema[] = [
  {
    field: 'type',
    label: '菜单类型',
    component: 'RadioButtonGroup',
    defaultValue: 0,
    componentProps: {
      options: [
        { label: '目录', value: 0 },
        { label: '菜单', value: 1 },
        { label: '按钮', value: 2 },
      ],
    },
    colProps: { lg: 24, md: 24 },
  },
  {
    field: 'name',
    label: '名称',
    component: 'Input',
    required: ({ values }) => !isButton(values.type),
    ifShow: ({ values }) => !isButton(values.type),
  },
  {
    field: 'title',
    label: '标题',
    component: 'Input',
    required: true,
  },
  {
    field: 'parent_id',
    label: '上级菜单',
    component: 'TreeSelect',
    defaultValue: 0,
    required: ({ values }) => !isDir(values.type),
    componentProps: {
      replaceFields: {
        title: 'title',
        key: 'id',
        value: 'id',
      },
      getPopupContainer: () => document.body,
    },
  },

  {
    field: 'icon',
    label: '图标',
    component: 'IconPicker',
    required: true,
    // ifShow: ({ values }) => !isButton(values.type),
  },
  {
    field: 'path',
    label: '路由地址',
    component: 'Input',
    required: true,
    ifShow: true,
  },
  {
    field: 'component',
    label: '组件路径',
    component: 'Input',
    required: ({ values }) => !isButton(values.type),
    ifShow: ({ values }) => !isButton(values.type),
  },
  {
    field: 'redirect',
    label: '重定向',
    component: 'Input',
    required: false,
    ifShow: ({ values }) => !isButton(values.type),
  },
  {
    field: 'frame_src',
    label: 'iframe 地址',
    component: 'Input',
    required: false,
    ifShow: ({ values }) => !isButton(values.type),
  },
  {
    field: 'current_active',
    label: '激活菜单',
    component: 'Input',
    required: false,
    ifShow: ({ values }) => isMenu(values.type),
  },
  {
    field: 'permission',
    label: '权限标识',
    component: 'Input',
    required: true,
    ifShow: ({ values }) => !isDir(values.type),
  },
  {
    field: 'method',
    label: '方法',
    component: 'RadioButtonGroup',
    defaultValue: 'GET',
    componentProps: {
      options: [
        { label: 'GET', value: 'GET' },
        { label: 'POST', value: 'POST' },
        { label: 'PUT', value: 'PUT' },
        { label: 'DELETE', value: 'DELETE' },
      ],
    },
    ifShow: ({ values }) => isButton(values.type),
  },
  {
    field: 'sort',
    label: '排序',
    component: 'InputNumber',
    defaultValue: 100,
    required: true,
  },
  {
    field: 'status',
    label: '状态',
    component: 'RadioButtonGroup',
    defaultValue: 1,
    componentProps: {
      options: [
        { label: '启用', value: 1 },
        { label: '禁用', value: 0 },
      ],
    },
  },
  {
    field: 'is_no_cache',
    label: '是否缓存',
    component: 'RadioButtonGroup',
    defaultValue: 0,
    componentProps: {
      options: [
        { label: '是', value: 0 },
        { label: '否', value: 1 },
      ],
    },
    ifShow: ({ values }) => isMenu(values.type),
  },
  {
    field: 'is_hide_child',
    label: '隐藏子菜单',
    component: 'RadioButtonGroup',
    defaultValue: 0,
    componentProps: {
      options: [
        { label: '是', value: 1 },
        { label: '否', value: 0 },
      ],
    },
    ifShow: ({ values }) => isMenu(values.type),
  },
  {
    field: 'is_show',
    label: '是否显示',
    component: 'RadioButtonGroup',
    defaultValue: 0,
    componentProps: {
      options: [
        { label: '是', value: 1 },
        { label: '否', value: 0 },
      ],
    },
    ifShow: ({ values }) => !isButton(values.type),
  },
];
