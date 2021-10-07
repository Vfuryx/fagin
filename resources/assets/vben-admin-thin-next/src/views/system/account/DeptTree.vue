<template>
  <div class="bg-white m-4 mr-0 overflow-hidden">
    <BasicTree
      title="部门列表"
      toolbar
      search
      :clickRowToExpand="false"
      :treeData="treeData"
      :replaceFields="{ key: 'id', title: 'name' }"
      @select="handleSelect"
    />
  </div>
</template>
<script lang="ts" setup>
  import { defineEmits, onMounted, ref } from 'vue';

  import { BasicTree, TreeItem } from '/@/components/Tree';
  import { getDeptList } from '/@/api/system/account';
  import { onActivated } from '@vue/runtime-core';

  const treeData = ref<TreeItem[]>([]);

  const emit = defineEmits(['select']);

  async function fetch() {
    treeData.value = (await getDeptList()) as unknown as TreeItem[];
  }

  function handleSelect(keys) {
    emit('select', keys[0]);
  }

  onMounted(() => {
    fetch();
  });
  onActivated(() => {
    fetch();
  });
</script>
