<template>
  <BasicDrawer
    v-bind="$attrs"
    @register="registerDrawer"
    showFooter
    :title="getTitle"
    width="45%"
    @ok="handleSubmit"
  >
    <BasicForm @register="registerForm">
      <template #menu="{ model, field }">
        <BasicTree
          v-model:value="model[field]"
          :treeData="treeData"
          :replaceFields="{ title: 'title', key: 'id' }"
          checkable
          toolbar
          title="菜单分配"
        />
      </template>
    </BasicForm>
  </BasicDrawer>
</template>
<script lang="ts" setup>
  import { ref, computed, unref, toRaw, defineEmits } from 'vue';
  import { BasicForm, useForm } from '/@/components/Form/index';
  import { formSchema } from './role.data';
  import { BasicDrawer, useDrawerInner } from '/@/components/Drawer';
  import { BasicTree, TreeItem } from '/@/components/Tree';

  import { createRole, getMenuAll, getRoleDetail, updateRole } from '/@/api/system/role';
  import { CreateRoleParams, UpdateRoleParams } from '/@/api/system/model/roleModel';

  const isUpdate = ref(true);
  const treeData = ref<TreeItem[]>([]);
  const id = ref(0);

  const emit = defineEmits(['success', 'register']);

  const [registerForm, { resetFields, setFieldsValue, validate }] = useForm({
    labelWidth: 90,
    schemas: formSchema,
    showActionButtonGroup: false,
  });

  const [registerDrawer, { setDrawerProps, closeDrawer }] = useDrawerInner(async (data) => {
    resetFields();
    setDrawerProps({ confirmLoading: false });

    // 需要在setFieldsValue之前先填充treeData，否则Tree组件可能会报key not exist警告
    if (unref(treeData).length === 0) {
      treeData.value = (await getMenuAll()) as any as TreeItem[];
    }
    isUpdate.value = !!data?.isUpdate;

    if (unref(isUpdate)) {
      id.value = data.record.id;

      const detail = await getRoleDetail(id.value);
      setFieldsValue({
        ...detail,
      });
    }
  });

  const getTitle = computed(() => (!unref(isUpdate) ? '新增角色' : '编辑角色'));

  async function handleSubmit() {
    try {
      const values = await validate();
      setDrawerProps({ confirmLoading: true });
      if (!values) return;
      // TODO custom api
      console.log(values);

      if (!unref(isUpdate)) {
        await createRole(
          toRaw<CreateRoleParams>({
            name: values.name,
            key: values.key,
            remark: values.remark,
            status: values.status,
            menu_ids: values.menu_ids,
          }),
        );
      } else {
        await updateRole(
          id.value,
          toRaw<UpdateRoleParams>({
            name: values.name,
            remark: values.remark,
            status: values.status,
            menu_ids: values.menu_ids,
          }),
        );
      }

      closeDrawer();
      emit('success');
    } finally {
      setDrawerProps({ confirmLoading: false });
    }
  }
</script>
