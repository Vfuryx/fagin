import { isAccountExist, setAccountStatus } from '/@/api/system/account';
import { BasicColumn } from '/@/components/Table';
import { FormSchema } from '/@/components/Table';
import { h, VNode } from 'vue';
import { Switch, Tag } from 'ant-design-vue';
import { PermCodeEnum } from '/@/enums/permCodeEnum';
import { useMessage } from '/@/hooks/web/useMessage';

export const columns: BasicColumn[] = [
  {
    title: '用户名',
    dataIndex: 'username',
    width: 120,
  },
  {
    title: '昵称',
    dataIndex: 'nick_name',
    width: 120,
  },
  {
    title: '邮箱',
    dataIndex: 'email',
    width: 180,
  },
  {
    title: '角色',
    dataIndex: 'roles',
    width: 180,
    customRender: ({ record }) => {
      const result: VNode[] = [];
      record.roles.forEach((v) => {
        result.push(h(Tag, { color: 'blue', style: 'margin-right: 8px;' }, () => v.name));
      });
      return h('div', {}, result);
    },
  },

  {
    title: '状态',
    dataIndex: 'status',
    width: 120,
    auth: PermCodeEnum.AdminSystemAccountUpdateStatus,
    customRender: ({ record }) => {
      if (record.username == 'admin') {
        return h('div', {}, {});
      }
      if (!Reflect.has(record, 'pendingStatus')) {
        record.pendingStatus = false;
      }
      return h(Switch, {
        checked: record.status === 1,
        checkedChildren: '已启用',
        unCheckedChildren: '已禁用',
        loading: record.pendingStatus,
        onChange(checked: boolean) {
          record.pendingStatus = true;
          const newStatus = checked ? 1 : 0;
          const { createMessage } = useMessage();
          setAccountStatus(record.id, newStatus)
            .then(() => {
              record.status = newStatus;
              createMessage.success(`已成功修改账号状态`);
            })
            .catch(() => {
              createMessage.error('修改账号状态失败');
            })
            .finally(() => {
              record.pendingStatus = false;
            });
        },
      });
    },
  },
  {
    title: '备注',
    dataIndex: 'remark',
  },
  {
    title: '创建时间',
    dataIndex: 'create_at',
    width: 180,
  },
];

export const searchFormSchema: FormSchema[] = [
  {
    field: 'username',
    label: '用户名',
    component: 'Input',
    colProps: { span: 8 },
  },
  {
    field: 'nick_name',
    label: '昵称',
    component: 'Input',
    colProps: { span: 8 },
  },
];

export const accountPasswordFormSchema: FormSchema[] = [
  {
    field: 'password',
    label: '密码',
    component: 'InputPassword',
    defaultValue: '',
    required: true,
    ifShow: true,
  },
];

export const accountFormSchemaFun = (id: any): FormSchema[] => {
  return [
    {
      field: 'username',
      label: '用户名',
      component: 'Input',
      helpMessage: ['不能输入带有admin的用户名'],
      defaultValue: '',
      rules: [
        {
          required: true,
          message: '请输入用户名',
        },
        {
          validator(_, value) {
            return new Promise((resolve, reject) => {
              isAccountExist(value, id.value)
                .then(() => resolve())
                .catch((err) => {
                  reject(err.message || '验证失败');
                });
            });
          },
        },
      ],
    },
    {
      field: 'password',
      label: '密码',
      component: 'InputPassword',
      defaultValue: '',
      required: ({}) => !id.value,
      ifShow: true,
    },
    {
      field: 'nick_name',
      label: '昵称',
      component: 'Input',
      defaultValue: '',
      required: true,
    },
    {
      field: 'email',
      label: '邮箱',
      component: 'Input',
      required: true,
    },
    {
      field: 'phone',
      label: '手机',
      component: 'Input',
      required: true,
    },
    {
      field: 'roles',
      label: '角色',
      component: 'ApiSelect',
      required: true,
    },
    {
      field: 'department_id',
      label: '所属部门',
      component: 'TreeSelect',
      componentProps: {
        replaceFields: {
          title: 'name',
          key: 'id',
          value: 'id',
        },
        getPopupContainer: () => document.body,
      },
      required: true,
    },
    {
      field: 'home_path',
      label: '首页路径',
      component: 'TreeSelect',
      defaultValue: null,
      required: true,
      componentProps: {
        replaceFields: {
          title: 'title',
          key: 'id',
          value: 'path',
        },
        getPopupContainer: () => document.body,
      },
    },
    {
      field: 'remark',
      label: '备注',
      component: 'InputTextArea',
    },
    {
      field: 'status',
      label: '状态',
      component: 'RadioButtonGroup',
      defaultValue: 1,
      componentProps: {
        options: [
          { label: '启用', value: 1 },
          { label: '停用', value: 0 },
        ],
      },
      required: true,
    },
    {
      field: 'sex',
      label: '性别',
      component: 'RadioButtonGroup',
      defaultValue: 0,
      componentProps: {
        options: [
          { label: '未知', value: 0 },
          { label: '男', value: 1 },
          { label: '女', value: 2 },
        ],
      },
      required: true,
    },
  ];
};
