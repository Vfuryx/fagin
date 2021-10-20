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
<script lang="ts" setup name="AccountDrawer">
  import { ref, computed, unref, toRaw } from 'vue';
  import { BasicForm, useForm } from '/@/components/Form/index';
  import { BasicDrawer, useDrawerInner } from '/@/components/Drawer';
  import { accountFormSchemaFun } from './account.data';
  import { getDeptList, getAccount, createAccount, updateAccount } from '/@/api/system/account';
  import { CreateAccountParams, UpdateAccountParams } from '/@/api/system/model/accountModel';

  const isUpdate = ref(true);
  const id = ref(0);

  const emit = defineEmits(['success', 'register']);

  const accountFormSchema = accountFormSchemaFun(id);

  const [registerForm, { setFieldsValue, updateSchema, resetFields, validate }] = useForm({
    labelWidth: 100,
    schemas: accountFormSchema,
    showActionButtonGroup: false,
    baseColProps: { lg: 12, md: 24 },
  });

  const [registerDrawer, { setDrawerProps, closeDrawer }] = useDrawerInner(async (data) => {
    resetFields();
    setDrawerProps({ confirmLoading: false });

    isUpdate.value = !!data?.isUpdate;

    if (unref(isUpdate)) {
      id.value = data.record.id;
      // 获取用户信息
      const account = await getAccount(id.value);
      setFieldsValue({
        ...account,
      });
    } else {
      id.value = 0;
    }

    const treeData = await getDeptList();
    updateSchema([
      {
        field: 'password',
        show: !unref(isUpdate),
      },
      {
        field: 'department_id',
        componentProps: { treeData },
      },
    ]);
  });

  const getTitle = computed(() => (!unref(isUpdate) ? '新增账号' : '编辑账号'));

  async function handleSubmit() {
    try {
      const values = await validate();

      setDrawerProps({ confirmLoading: true });
      // TODO custom api
      if (!values) return;
      console.log(values);

      if (!unref(isUpdate)) {
        await createAccount(
          toRaw<CreateAccountParams>({
            username: values.username,
            password: values.password,
            roles: values.roles,
            department_id: values.department_id,
            nick_name: values.nick_name,
            email: values.email,
            phone: values.phone,
            remark: values.remark,
            status: values.status,
            sex: values.sex,
          }),
        );
      } else {
        await updateAccount(
          id.value,
          toRaw<UpdateAccountParams>({
            username: values.username,
            nick_name: values.nick_name,
            roles: values.roles,
            department_id: values.department_id,
            email: values.email,
            phone: values.phone,
            remark: values.remark,
            status: values.status,
            sex: values.sex,
          }),
        );
      }

      closeDrawer();
      emit('success', { isUpdate: unref(isUpdate), values: { ...values, id: id.value } });
    } finally {
      setDrawerProps({ confirmLoading: false });
    }
  }
</script>
