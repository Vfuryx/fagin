<template>
  <div :class="prefixCls">
    <a-tabs v-model="currentTab" tab-position="left" size="small" @change="handleTabChange">
      <a-tab-pane v-for="v in icons" :tab="v.title" :key="v.key">
        <ul>
          <li
            v-for="(icon, key) in v.icons"
            :key="`${v.key}-${key}`"
            :class="{ 'active': selectedIcon==icon }"
            @click="handleSelectedIcon(icon)">
            <a-icon :type="icon" :style="{ fontSize: '36px' }" />
          </li>
        </ul>
      </a-tab-pane>
    </a-tabs>
  </div>
</template>

<script>
import icons from './icons'

export default {
  name: 'IconSelect',
  props: {
    prefixCls: {
      type: String,
      default: 'ant-pro-icon-selector'
    },
    // eslint-disable-next-line
    value: {
      type: String
    }
  },
  data () {
    return {
      selectedIcon: this.value || '',
      currentTab: 'directional',
      icons
    }
  },
  watch: {
    value (val) {
      this.selectedIcon = val
      this.autoSwitchTab()
    }
  },
  created () {
    if (this.value) {
      this.autoSwitchTab()
    }
  },
  methods: {
    handleSelectedIcon (icon) {
      this.selectedIcon = icon
      this.$emit('change', icon)
    },
    handleTabChange (activeKey) {
      this.currentTab = activeKey
    },
    autoSwitchTab () {
      icons.some(item => item.icons.some(icon => icon === this.value) && (this.currentTab = item.key))
    }
  }
}
</script>

<style lang="less" scoped>
@import "../index.less";

ul {
  list-style: none;
  padding: 0;
  overflow: auto;
  height: 300px;

  li {
    float: right;
    display: block;
    padding: 1px;
    margin: 3px 0;
    border-radius: @border-radius-base;

    &:hover, &.active {
      box-shadow: 0px 0px 5px 2px @primary-color;
      cursor: pointer;
      color: @white;
      background-color: @primary-color;
    }
  }
}
.ant-tabs .ant-tabs-left-content {
  padding-left: 5px;
  border-left: 1px solid #e8e8e8;
}

::-webkit-scrollbar {
  display: none; /*ChromeSafari*/
}
</style>
