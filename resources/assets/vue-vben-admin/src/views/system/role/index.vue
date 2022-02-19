<template>
  <div>
    <BasicTable @register="registerTable">
      <template #toolbar>
        <a-button
          type="primary"
          v-if="hasPermission(PermCodeEnum.AdminSystemRoleCreate)"
          @click="handleCreate"
        >
          新增角色
        </a-button>
      </template>
      <template #action="{ record }">
        <TableAction
          :actions="[
            {
              icon: 'clarity:note-edit-line',
              onClick: handleEdit.bind(null, record),
              auth: PermCodeEnum.AdminSystemRoleUpdate,
            },
            {
              icon: 'ant-design:delete-outlined',
              color: 'error',
              auth: PermCodeEnum.AdminSystemRoleDelete,
              popConfirm: {
                title: '是否确认删除',
                confirm: handleDelete.bind(null, record),
              },
            },
          ]"
        />
      </template>
    </BasicTable>
    <RoleDrawer @register="registerDrawer" @success="handleSuccess" />
  </div>
</template>
<script lang="ts" setup name="RoleManagement">
  import { onActivated } from 'vue';

  import { BasicTable, useTable, TableAction } from '/@/components/Table';
  import { useDrawer } from '/@/components/Drawer';
  import { usePermission } from '/@/hooks/web/usePermission';
  import { PermCodeEnum } from '/@/enums/permCodeEnum';

  import RoleDrawer from './RoleDrawer.vue';
  import { columns, searchFormSchema } from './role.data';

  import { deleteRole, getRoleListByPage } from '/@/api/system/role';

  const { hasPermission } = usePermission();
  const [registerDrawer, { openDrawer }] = useDrawer();
  const [registerTable, { reload }] = useTable({
    title: '角色列表',
    api: getRoleListByPage,
    columns,
    formConfig: {
      labelWidth: 120,
      schemas: searchFormSchema,
    },
    useSearchForm: hasPermission(PermCodeEnum.AdminSystemRoleQuery),
    showTableSetting: true,
    bordered: true,
    showIndexColumn: false,
    actionColumn: {
      width: 80,
      title: '操作',
      dataIndex: 'action',
      slots: { customRender: 'action' },
      fixed: undefined,
    },
  });

  function handleCreate() {
    openDrawer(true, {
      isUpdate: false,
    });
  }

  function handleEdit(record: Recordable) {
    openDrawer(true, {
      record,
      isUpdate: true,
    });
  }

  async function handleDelete(record: Recordable) {
    await deleteRole(record.id);
    await reload();
  }

  function handleSuccess() {
    reload();
  }

  onActivated(() => {
    reload();
  });
</script>
