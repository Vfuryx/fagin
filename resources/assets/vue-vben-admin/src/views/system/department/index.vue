<template>
  <div>
    <BasicTable @register="registerTable" @fetch-success="onFetchSuccess">
      <template #toolbar>
        <a-button
          type="primary"
          v-if="hasPermission(PermCodeEnum.AdminSystemDepartmentCreate)"
          @click="handleCreate"
        >
          新增部门
        </a-button>
      </template>
      <template #action="{ record }">
        <TableAction
          :actions="[
            {
              icon: 'clarity:note-edit-line',
              onClick: handleEdit.bind(null, record),
              auth: PermCodeEnum.AdminSystemDepartmentUpdate,
            },
            {
              icon: 'ant-design:delete-outlined',
              color: 'error',
              auth: PermCodeEnum.AdminSystemDepartmentDelete,
              popConfirm: {
                title: '是否确认删除',
                confirm: handleDelete.bind(null, record),
              },
            },
          ]"
        />
      </template>
    </BasicTable>
    <DeptModal @register="registerModal" @success="handleSuccess" />
  </div>
</template>
<script lang="ts" setup name="DepartmentManagement">
  import { nextTick, onActivated } from 'vue';

  import { usePermission } from '/@/hooks/web/usePermission';
  import { BasicTable, useTable, TableAction } from '/@/components/Table';
  import { useModal } from '/@/components/Modal';
  import { PermCodeEnum } from '/@/enums/permCodeEnum';

  import DeptModal from './DeptModal.vue';
  import { columns, searchFormSchema } from './dept.data';

  import { deleteDept, getDeptList } from '/@/api/system/department';

  const { hasPermission } = usePermission();

  const [registerModal, { openModal }] = useModal();
  const [registerTable, { reload, expandAll }] = useTable({
    title: '部门列表',
    api: getDeptList,
    columns,
    formConfig: {
      labelWidth: 120,
      schemas: searchFormSchema,
    },
    isTreeTable: true,
    pagination: false,
    striped: false,
    useSearchForm: hasPermission(PermCodeEnum.AdminSystemDepartmentQuery),
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
    openModal(true, {
      isUpdate: false,
    });
  }

  function handleEdit(record: Recordable) {
    openModal(true, {
      record,
      isUpdate: true,
    });
  }

  async function handleDelete(record: Recordable) {
    await deleteDept(record.id);
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
