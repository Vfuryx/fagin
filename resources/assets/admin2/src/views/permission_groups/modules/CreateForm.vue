<template>
  <a-modal
    :title="model && model.id > 0 ? '权限组编辑' : '权限组新增'"
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
        <!-- 检查是否有 id 并且大于0，大于0是修改。其他是新增，新增不显示主键ID -->
        <a-form-model-item
          :label-col="{ span: 5 }"
          :wrapper-col="{span: 5}"
          v-if="model && model.id > 0"
          label="主键ID"
          prop="id"
          style="width: 100%;margin-right: 0">
          <a-input v-model="model.id" disabled />
        </a-form-model-item>
        <a-form-model-item
          :label-col="{ span: 5 }"
          :wrapper-col="{span: 10}"
          label="菜单类型"
          prop="type"
          style="width: 100%;margin-right: 0">
          <a-select v-model="model.type">
            <a-select-option :value="0">后台菜单</a-select-option>
            <a-select-option :value="1">商家菜单</a-select-option>
          </a-select>
        </a-form-model-item>
        <a-form-model-item
          label="名称"
          :label-col="{ span: 5 }"
          :wrapper-col="{span: 16}"
          prop="name"
          style="width: 100%;margin-right: 0">
          <a-input v-model="model.name"></a-input>
        </a-form-model-item>
        <a-form-model-item
          label="显示排序"
          :label-col="{ span: 5 }"
          :wrapper-col="{span: 5}"
          prop="sort"
          style="width: 100%;margin-right: 0">
          <a-input-number v-model="model.sort" controls-position="right" :min="0" />
        </a-form-model-item>
      </a-form-model>
    </a-spin>
  </a-modal>
</template>

<script>
import Vue from 'vue'
// import pick from 'lodash.pick'
import { FormModel, InputNumber } from 'ant-design-vue'

Vue.use(FormModel)
Vue.use(InputNumber)

// 表单字段
// const fields = [
//   'id', 'name', 'sort'
// ]

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
    },
    list: {
      type: Array,
      default: () => []
    }
  },
  data () {
    this.formLayout = {
      labelCol: {
        xs: { span: 24 },
        sm: { span: 7 }
      },
      wrapperCol: {
        xs: { span: 24 },
        sm: { span: 13 }
      }
    }
    return {
      // form: this.$form.createForm(this),
      labelCol: { span: 10 },
      wrapperCol: { span: 14 },
      form: {
        id: 0,
        name: '',
        sort: 0,
        type: 0
      },
      rules: {
        id: [
          { required: true, message: '不能空', trigger: 'blur' },
          { type: 'integer', min: 0, message: '不正确', trigger: 'blur' }
        ],
        sort: [
          { required: true, message: '不能空', trigger: 'blur' },
          { type: 'integer', min: 0, message: '上级不正确', trigger: 'blur' }
        ],
        name: [
          { required: true, message: '不能为空', trigger: 'blur' },
          { min: 1, max: 32, message: '范围 1 - 32 字符', trigger: 'blur' }
        ]
      }
    }
  },
  created () {
    // // 防止表单未注册
    // fields.forEach(v => this.form.getFieldDecorator(v))
    //
    // // 当 model 发生改变时，为表单设置值
    // this.$watch('model', () => {
    //   this.model && this.form.setFieldsValue(pick(this.model, fields))
    // })
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
