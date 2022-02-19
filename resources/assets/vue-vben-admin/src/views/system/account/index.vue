<template>
  <PageWrapper dense contentFullHeight contentClass="flex">
    <DeptTree class="w-1/4 xl:w-1/5" @select="handleSelect" />
    <BasicTable @register="registerTable" class="w-3/4 xl:w-4/5" :searchInfo="searchInfo">
      <template #toolbar>
        <a-button
          type="primary"
          v-if="hasPermission(PermCodeEnum.AdminSystemAccountCreate)"
          @click="handleCreate"
          >新增账号</a-button
        >
      </template>
      <template #action="{ record }">
        <TableAction
          :actions="[
            {
              icon: 'clarity:info-standard-line',
              tooltip: '查看用户详情',
              onClick: handleView.bind(null, record),
              auth: PermCodeEnum.AdminSystemAccountDetail,
            },
            {
              icon: 'clarity:note-edit-line',
              tooltip: '编辑用户资料',
              onClick: handleEdit.bind(null, record),
              auth: PermCodeEnum.AdminSystemAccountUpdate,
            },
            {
              icon: 'teenyicons:password-outline',
              tooltip: '编辑用户密码',
              onClick: handleEditPassword.bind(null, record),
              auth: PermCodeEnum.AdminSystemAccountUpdatePWD,
            },
            {
              icon: 'carbon:logout',
              tooltip: '用户强制下线',
              auth: PermCodeEnum.AdminSystemAccountLogout,
              popConfirm: {
                title: '是否确认强制下线',
                confirm: handleLogout.bind(null, record),
              },
            },
            {
              icon: 'ant-design:delete-outlined',
              color: 'error',
              tooltip: '删除此账号',
              auth: PermCodeEnum.AdminSystemAccountDelete,
              popConfirm: {
                title: '是否确认删除',
                confirm: handleDelete.bind(null, record),
              },
            },
          ]"
        />
      </template>
    </BasicTable>
    <AccountDrawer @register="registerDrawer" @success="handleSuccess" />
    <AccountPasswordModal @register="registerModal" @success="handleSuccess" />
  </PageWrapper>
</template>
<script lang="ts" setup name="AccountManagement">
  import { reactive, onActivated } from 'vue';

  import { useGo } from '/@/hooks/web/usePage';
  import { BasicTable, useTable, TableAction } from '/@/components/Table';
  import { PageWrapper } from '/@/components/Page';
  import { useDrawer } from '/@/components/Drawer';
  import { useModal } from '/@/components/Modal';
  import { usePermission } from '/@/hooks/web/usePermission';
  import { PermCodeEnum } from '/@/enums/permCodeEnum';

  import DeptTree from './DeptTree.vue';
  import AccountDrawer from './AccountDrawer.vue';
  import AccountPasswordModal from './AccountPasswordModal.vue';
  import { columns, searchFormSchema } from './account.data';

  import { getAccountList, deleteAccount, logoutAccount } from '/@/api/system/account';

  const { hasPermission } = usePermission();
  const go = useGo();
  const [registerModal, { openModal }] = useModal();
  const [registerDrawer, { openDrawer }] = useDrawer();
  const searchInfo = reactive<Recordable>({});
  const [registerTable, { reload }] = useTable({
    title: '账号列表',
    api: getAccountList,
    rowKey: 'id',
    columns,
    formConfig: {
      labelWidth: 120,
      schemas: searchFormSchema,
      autoSubmitOnEnter: true,
    },
    useSearchForm: hasPermission(PermCodeEnum.AdminSystemAccountQuery),
    showTableSetting: true,
    bordered: true,
    actionColumn: {
      width: 185,
      title: '操作',
      dataIndex: 'action',
      slots: { customRender: 'action' },
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
  function handleEditPassword(record: Recordable) {
    openModal(true, { record });
  }

  async function handleDelete(record: Recordable) {
    await deleteAccount(record.id);
    await reload();
  }

  async function handleLogout(record: Recordable) {
    await logoutAccount(record.id);
    await reload();
  }

  function handleSuccess() {
    reload();
  }

  function handleSelect(deptId = '') {
    searchInfo.deptId = deptId;
    reload();
  }

  function handleView(record: Recordable) {
    go('/system/account_detail/' + record.id);
  }

  onActivated(() => {
    reload();
  });
</script>
