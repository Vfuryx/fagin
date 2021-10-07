<template>
  <BasicDrawer
    v-bind="$attrs"
    @register="registerDrawer"
    :showFooter="false"
    :title="getTitle"
    width="45%"
  >
    <Description @register="registerDescription" />
  </BasicDrawer>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { detailSchema } from './operlog.data';
  import { BasicDrawer, useDrawerInner } from '/@/components/Drawer';

  import { getOperLogDetail } from '/@/api/system/oper-log';
  import Description from '/@/components/Description/src/Description.vue';
  import { useDescription } from '/@/components/Description';

  const getTitle = '操作日志详情';

  const detail = ref({});
  const [registerDrawer, { setDrawerProps }] = useDrawerInner(async (data) => {
    setDrawerProps({ confirmLoading: false });
    detail.value = await getOperLogDetail(data.record.id);
    console.log(detail.value);
  });

  const [registerDescription] = useDescription({
    title: '',
    column: 1,
    data: detail,
    schema: detailSchema,
  });
</script>
