<template>
  <div>
    <BasicTable @register="registerTable">
      <template #action="{ record }">
        <TableAction
          :actions="[
            {
              icon: 'clarity:info-standard-line',
              tooltip: '查看详情',
              onClick: handleView.bind(null, record),
              auth: PermCodeEnum.AdminMonitorOperLogDetail,
            },
            // {
            //   icon: 'ant-design:delete-outlined',
            //   color: 'error',
            //   popConfirm: {
            //     title: '是否确认删除',
            //     confirm: handleDelete.bind(null, record),
            //   },
            // },
          ]"
        />
      </template>
    </BasicTable>
    <OperLogDrawer @register="registerDrawer" />
  </div>
</template>
<script lang="ts" setup name="OperLogManagement">
  import { onActivated } from 'vue';

  import { BasicTable, useTable, TableAction } from '/@/components/Table';
  import { usePermission } from '/@/hooks/web/usePermission';
  import { useDrawer } from '/@/components/Drawer';
  import OperLogDrawer from './OperLogDrawer.vue';
  import { columns, searchFormSchema } from './operlog.data';
  import { PermCodeEnum } from '/@/enums/permCodeEnum';

  import { getOperLogListByPage } from '/@/api/system/oper-log';

  const { hasPermission } = usePermission();
  const [registerDrawer, { openDrawer }] = useDrawer();
  const [registerTable, { reload }] = useTable({
    title: '操作日志列表',
    api: getOperLogListByPage,
    columns,
    formConfig: {
      labelWidth: 120,
      schemas: searchFormSchema,
    },
    useSearchForm: hasPermission(PermCodeEnum.AdminMonitorOperLogQuery),
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

  function handleView(record: Recordable) {
    openDrawer(true, {
      record,
    });
  }

  onActivated(() => {
    reload();
  });
</script>
