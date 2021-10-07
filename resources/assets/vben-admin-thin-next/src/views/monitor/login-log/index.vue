<template>
  <div>
    <BasicTable @register="registerTable" />
  </div>
</template>
<script lang="ts" setup name="LoginLogManagement">
  import { onActivated } from '@vue/runtime-core';
  import { usePermission } from '/@/hooks/web/usePermission';
  import { PermCodeEnum } from '/@/enums/permCodeEnum';

  import { BasicTable, useTable } from '/@/components/Table';
  import { columns, searchFormSchema } from './operlog.data';

  import { getLoginLogListByPage } from '/@/api/system/login-log';

  const { hasPermission } = usePermission();
  const [registerTable, { reload }] = useTable({
    title: '登录日志列表',
    api: getLoginLogListByPage,
    columns,
    formConfig: {
      labelWidth: 120,
      schemas: searchFormSchema,
    },
    useSearchForm: hasPermission(PermCodeEnum.AdminMonitorLoginLogQuery),
    showTableSetting: true,
    bordered: true,
    showIndexColumn: false,
  });

  onActivated(() => {
    reload();
  });
</script>
