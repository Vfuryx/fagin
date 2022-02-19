<template>
  <BasicModal v-bind="$attrs" @register="registerModal" :title="getTitle" @ok="handleSubmit">
    <BasicForm @register="registerForm" />
  </BasicModal>
</template>
<script lang="ts" setup>
  import { ref, computed, unref, toRaw } from 'vue';
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { BasicForm, useForm } from '/@/components/Form/index';
  import { formSchema } from './dept.data';

  import { createDept, getDeptList, updateDept } from '/@/api/system/department';

  const isUpdate = ref(true);
  const id = ref(0);

  const [registerForm, { resetFields, setFieldsValue, updateSchema, validate }] = useForm({
    labelWidth: 100,
    schemas: formSchema,
    showActionButtonGroup: false,
  });

  const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
    resetFields();
    setModalProps({ confirmLoading: false });
    isUpdate.value = !!data?.isUpdate;

    if (unref(isUpdate)) {
      id.value = data.record.id;
      setFieldsValue({
        ...data.record,
      });
    }
    const treeData = await getDeptList();
    console.log(treeData);
    treeData.unshift({
      id: 0,
      parent_id: 0,
      name: '顶级部门',
      sort: 999,
      created_at: '',
      remark: '',
      status: 1,
    });
    updateSchema({
      field: 'parent_id',
      componentProps: { treeData },
    });
  });

  const emit = defineEmits(['success', 'register']);

  const getTitle = computed(() => (!unref(isUpdate) ? '新增部门' : '编辑部门'));

  async function handleSubmit() {
    try {
      const values = await validate();
      setModalProps({ confirmLoading: true });
      // TODO custom api
      if (!values) return;
      const param = toRaw({
        name: values.name || '',
        remark: values.remark || '',
        status: values.status || 0,
        parent_id: values.parent_id || 0,
        sort: values.sort || 0,
      });

      if (!unref(isUpdate)) {
        await createDept(param);
      } else {
        await updateDept(id.value, param);
      }

      console.log(values);
      closeModal();
      emit('success');
    } finally {
      setModalProps({ confirmLoading: false });
    }
  }
</script>
