<template>
  <BasicModal v-bind="$attrs" @register="registerModal" :title="getTitle" @ok="handleSubmit">
    <BasicForm @register="registerForm" />
  </BasicModal>
</template>
<script lang="ts" setup>
  import { ref, toRaw } from 'vue';
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { BasicForm, useForm } from '/@/components/Form/index';
  import { accountPasswordFormSchema } from './account.data';
  import { setAccountPassword } from '/@/api/system/account';

  const id = ref(0);

  const emit = defineEmits(['success', 'register']);

  const [registerForm, { resetFields, validate }] = useForm({
    labelWidth: 100,
    schemas: accountPasswordFormSchema,
    showActionButtonGroup: false,
    actionColOptions: {
      span: 23,
    },
  });

  const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
    resetFields();
    setModalProps({ confirmLoading: false });
    id.value = data.record.id;
  });

  const getTitle = '修改密码';

  async function handleSubmit() {
    try {
      const values = await validate();
      setModalProps({ confirmLoading: true });
      // TODO custom api
      console.log(values);
      await setAccountPassword(id.value, values.password);
      closeModal();
      emit('success');
    } finally {
      setModalProps({ confirmLoading: false });
    }
  }
</script>
