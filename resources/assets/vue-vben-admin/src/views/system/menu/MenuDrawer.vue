<template>
  <BasicDrawer
    v-bind="$attrs"
    @register="registerDrawer"
    showFooter
    :title="getTitle"
    width="50%"
    @ok="handleSubmit"
  >
    <BasicForm @register="registerForm" />
  </BasicDrawer>
</template>
<script lang="ts" setup>
  import { ref, computed, unref, toRaw } from 'vue';
  import { BasicForm, useForm } from '/@/components/Form/index';
  import { formSchema } from './menu.data';
  import { BasicDrawer, useDrawerInner } from '/@/components/Drawer';
  import { createMenu, updateMenu, getMenuAll, getMenuDetail } from '/@/api/system/menu';

  const isUpdate = ref(true);
  const id = ref(0);

  const emit = defineEmits(['success', 'register']);

  const [registerForm, { resetFields, setFieldsValue, updateSchema, validate }] = useForm({
    labelWidth: 100,
    schemas: formSchema,
    showActionButtonGroup: false,
    baseColProps: { lg: 12, md: 24 },
  });

  const [registerDrawer, { setDrawerProps, closeDrawer }] = useDrawerInner(async (data) => {
    resetFields();
    setDrawerProps({ confirmLoading: false });
    isUpdate.value = !!data?.isUpdate;

    if (unref(isUpdate)) {
      id.value = data.record.id;

      const detail = await getMenuDetail(id.value);
      setFieldsValue({
        ...detail,
      });
    }

    const treeData = await getMenuAll();
    treeData.unshift({ icon: '', id: 0, parent_id: 0, title: '顶级菜单' });
    updateSchema({
      field: 'parent_id',
      componentProps: { treeData },
    });
  });

  const getTitle = computed(() => (!unref(isUpdate) ? '新增菜单' : '编辑菜单'));

  async function handleSubmit() {
    try {
      const values = await validate();
      setDrawerProps({ confirmLoading: true });
      // TODO custom api
      if (!values) return;
      const param = toRaw({
        component: values.component || '',
        icon: values.icon || '',
        is_no_cache: values.is_no_cache || 0,
        is_hide_child: values.is_hide_child || 0,
        is_show: values.is_show || 0,
        method: values.method,
        name: values.name || '',
        parent_id: values.parent_id || 0,
        path: values.path || '',
        redirect: values.redirect || '',
        frame_src: values.frame_src || '',
        current_active: values.current_active || '',
        permission: values.permission || '',
        sort: values.sort || 0,
        status: values.status || 0,
        title: values.title || '',
        type: values.type || 0,
      });

      if (!unref(isUpdate)) {
        await createMenu(param);
      } else {
        await updateMenu(id.value, param);
      }

      closeDrawer();
      emit('success');
    } finally {
      setDrawerProps({ confirmLoading: false });
    }
  }
</script>
