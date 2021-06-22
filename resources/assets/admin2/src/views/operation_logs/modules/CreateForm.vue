<template>
  <a-modal
    :title="model.id > 0 ? '管理员编辑' : '管理员新增'"
    :width="860"
    :visible="visible"
    :confirmLoading="loading"
    @ok="handleOk"
    @cancel="handleCancel"
    cancel-text="取消"
  >
    <template slot="footer">
      <a-button @click="handleCancel">关闭</a-button>
    </template>

    <a-spin :spinning="loading">
      <a-form-model
        layout="inline"
        :model="model"
        :rules="rules"
        ref="ruleForm"
      >
        <!-- 检查是否有 id 并且大于0，大于0是修改。其他是新增，新增不显示主键ID -->
        <a-form-model-item
          :label-col="{ span: 4 }"
          :wrapper-col="{span: 5}"
          v-if="model && model.id > 0"
          label="日志编号"
          prop="id"
          style="width: 88%;margin-right: 0">
          <a-input v-model="model.id" disabled />
        </a-form-model-item>
        <a-form-model-item
          label="操作人员"
          :label-col="{ span: 7 }"
          :wrapper-col="{span: 17}"
          prop="user"
          style="width: 50%;margin-right: 0">
          <a-input v-model="model.user" disabled />
        </a-form-model-item>
        <a-form-model-item
          label="操作日期"
          :label-col="{ span: 7 }"
          :wrapper-col="{span: 17}"
          prop="created_at"
          style="width: 50%;margin-right: 0">
          <a-input v-model="model.created_at" disabled />
        </a-form-model-item>
        <a-form-model-item
          label="系统模块"
          :label-col="{ span: 7 }"
          :wrapper-col="{span: 17}"
          prop="module"
          style="width: 50%;margin-right: 0">
          <a-input v-model="model.module" disabled />
        </a-form-model-item>
        <a-form-model-item
          label="请求方式"
          :label-col="{ span: 7 }"
          :wrapper-col="{span: 17}"
          prop="method"
          style="width: 50%;margin-right: 0">
          <a-input v-model="model.method" disabled />
        </a-form-model-item>
        <a-form-model-item
          label="IP"
          :label-col="{ span: 7 }"
          :wrapper-col="{span: 17}"
          prop="ip"
          style="width: 50%;margin-right: 0">
          <a-input v-model="model.ip" disabled />
        </a-form-model-item>
        <a-form-model-item
          label="访问位置"
          :label-col="{ span: 7 }"
          :wrapper-col="{span: 17}"
          prop="location"
          style="width: 50%;margin-right: 0">
          <a-input v-model="model.location" disabled />
        </a-form-model-item>
        <a-form-model-item
          label="耗时"
          :label-col="{ span: 7 }"
          :wrapper-col="{span: 17}"
          prop="latency_time"
          style="width: 50%;margin-right: 0">
          <a-input v-model="model.latency_time" disabled />
        </a-form-model-item>
        <a-form-model-item
          label="body"
          :label-col="{ span: 4 }"
          :wrapper-col="{span: 20}"
          prop="input"
          style="width: 88%;margin-right: 0">
          <a-textarea v-model="model.input" disabled />
        </a-form-model-item>
        <a-form-model-item
          label="UserAgent"
          :label-col="{ span: 4 }"
          :wrapper-col="{span: 20}"
          prop="user_agent"
          style="width: 88%;margin-right: 0">
          <a-textarea v-model="model.user_agent" disabled />
        </a-form-model-item>
      </a-form-model>
    </a-spin>
  </a-modal>
</template>

<script>
import Vue from 'vue'
import { FormModel, TreeSelect, InputNumber } from 'ant-design-vue'

Vue.use(FormModel)
Vue.use(TreeSelect)
Vue.use(InputNumber)

export default {
  props: {
    visible: {
      type: Boolean,
      required: true
    },
    loading: {
      type: Boolean,
      default: () => false
    },
    model: {
      type: Object,
      default: () => null
    }
  },
  data () {
    return {
      form: {
        id: 0,
        user: '',
        method: '',
        path: '',
        ip: '',
        location: '',
        module: '',
        operation: '',
        input: '',
        latency_time: '',
        user_agent: '',
        created_at: ''
      },
      list: [],
      rules: {
      }
    }
  },
  created () {
  },
  mounted () {
  },
  methods: {
    handleOk () {
      return this.$emit('ok', this.$refs.ruleForm, this.model)
    },
    handleCancel () {
      return this.$emit('cancel', this.$refs.ruleForm)
    }
  }
}
</script>
