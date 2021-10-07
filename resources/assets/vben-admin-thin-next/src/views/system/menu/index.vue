<template>
  <div>
    <BasicTable @register="registerTable" @fetch-success="onFetchSuccess">
      <template #toolbar>
        <a-button
          type="primary"
          v-if="hasPermission(PermCodeEnum.AdminSystemMenuCreate)"
          @click="handleCreate"
        >
          新增菜单
        </a-button>
      </template>
      <template #action="{ record }">
        <TableAction
          :actions="[
            {
              icon: 'clarity:note-edit-line',
              onClick: handleEdit.bind(null, record),
              auth: PermCodeEnum.AdminSystemMenuUpdate,
            },
            {
              icon: 'ant-design:delete-outlined',
              color: 'error',
              auth: PermCodeEnum.AdminSystemMenuDelete,
              popConfirm: {
                title: '是否确认删除',
                confirm: handleDelete.bind(null, record),
              },
            },
          ]"
        />
      </template>
    </BasicTable>
    <MenuDrawer @register="registerDrawer" @success="handleSuccess" />
  </div>
</template>
<script lang="ts" setup name="MenuManagement">
  import { nextTick } from 'vue';
  import { onActivated } from '@vue/runtime-core';
  import { usePermission } from '/@/hooks/web/usePermission';
  import { BasicTable, useTable, TableAction } from '/@/components/Table';
  import { useDrawer } from '/@/components/Drawer';
  import { PermCodeEnum } from '/@/enums/permCodeEnum';
  import MenuDrawer from './MenuDrawer.vue';
  import { columns, searchFormSchema } from './menu.data';
  import { deleteMenu, getMenuList } from '/@/api/system/menu';

  const { hasPermission } = usePermission();
  const [registerDrawer, { openDrawer }] = useDrawer();
  const [registerTable, { reload, expandAll }] = useTable({
    title: '菜单列表',
    api: getMenuList,
    columns,
    formConfig: {
      labelWidth: 120,
      schemas: searchFormSchema,
    },
    isTreeTable: true,
    pagination: false,
    striped: false,
    useSearchForm: hasPermission(PermCodeEnum.AdminSystemMenuQuery),
    showTableSetting: true,
    bordered: true,
    showIndexColumn: false,
    canResize: false,
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
    await deleteMenu(record.id);
    await reload();
  }

  function handleSuccess() {
    reload();
  }

  function onFetchSuccess() {
    // 演示默认展开所有表项
    nextTick(expandAll);
  }

  onActivated(() => {
    reload();
  });
</script>
