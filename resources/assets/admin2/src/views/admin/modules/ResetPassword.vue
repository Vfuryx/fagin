<template>
  <a-modal
    title="修改密码"
    :width="860"
    :visible="visible"
    :confirmLoading="loading"
    @ok="handleOk"
    @cancel="handleCancel"
    ok-text="确认"
    cancel-text="取消"
  >
    <a-spin :spinning="loading">
      <a-form-model
        layout="inline"
        :model="model"
        :rules="rules"
        ref="ruleForm"
      >
        <a-form-model-item
          label="用户密码"
          :label-col="{ span: 7 }"
          :wrapper-col="{span: 17}"
          prop="password"
          style="width: 50%;margin-right: 0">
          <a-input v-model="model.password"></a-input>
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
    id: {
      type: Number,
      default: () => 0
    }
  },
  data () {
    return {
      model: {
        password: '',
        id: 0
      },
      rules: {
        password: [
          { required: true, message: '密码不能为空', trigger: 'blur' },
          { min: 8, max: 32, message: '长度在 8 到 32 个字符', trigger: 'blur' }
        ]
      }
    }
  },
  created () {
  },
  mounted () {
  },
  methods: {
    handleOk () {
      this.model.id = this.id
      return this.$emit('ok', this.$refs.ruleForm, this.model)
    },
    handleCancel () {
      return this.$emit('cancel', this.$refs.ruleForm)
    }
  }
}
</script>
