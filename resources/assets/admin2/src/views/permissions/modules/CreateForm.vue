<template>
  <a-modal
    :title="model.id > 0 ? '权限编辑' : '权限新增'"
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
          label="权限类型"
          prop="type"
          style="width: 100%;margin-right: 0">
          <a-select
            @change="typeChange"
            v-model="model.type">
            <a-select-option :value="0">后台菜单</a-select-option>
            <a-select-option :value="1">商家菜单</a-select-option>
          </a-select>
        </a-form-model-item>
        <a-form-model-item
          :label-col="{ span: 5 }"
          :wrapper-col="{span: 10}"
          label="所属分组"
          prop="gid"
          style="width: 100%;margin-right: 0">
          <a-select v-model="model.gid">
            <a-select-option :value="0" :key="0">
              未选择
            </a-select-option>
            <a-select-option v-for="(v,k) in list" :value="v.id||0" :key="k">
              {{ v.name }}
            </a-select-option>
          </a-select>
        </a-form-model-item>
        <a-form-model-item
          label="接口名称"
          :label-col="{ span: 5 }"
          :wrapper-col="{span: 16}"
          prop="name"
          style="width: 100%;margin-right: 0">
          <a-input v-model="model.name"></a-input>
        </a-form-model-item>
        <a-form-model-item
          label="路径"
          :label-col="{ span: 5 }"
          :wrapper-col="{span: 16}"
          prop="paths"
          style="width: 100%;margin-right: 0">
          <a-input v-model="model.paths"></a-input>
        </a-form-model-item>
        <a-form-model-item
          label="请求方式"
          :label-col="{ span: 5 }"
          :wrapper-col="{span: 18}"
          prop="method"
          style="width: 100%;margin-right: 0">
          <a-radio-group name="is_show" v-model="model.method">
            <a-radio value="GET">GET</a-radio>
            <a-radio value="POST">POST</a-radio>
            <a-radio value="PUT">PUT</a-radio>
            <a-radio value="DELETE">DELETE</a-radio>
          </a-radio-group>
        </a-form-model-item>
        <a-form-model-item
          label="显示排序"
          :label-col="{ span: 5 }"
          :wrapper-col="{span: 5}"
          prop="sort"
          style="width: 100%;margin-right: 0">
          <a-input-number v-model="model.sort" controls-position="right" :min="0" />
        </a-form-model-item>
        <a-form-model-item
          label="菜单状态"
          :label-col="{ span: 5 }"
          :wrapper-col="{span: 18}"
          prop="status"
          style="width: 100%;margin-right: 0">
          <a-radio-group name="status" v-model="model.status">
            <a-radio :value="1">开启</a-radio>
            <a-radio :value="0">关闭</a-radio>
          </a-radio-group>
        </a-form-model-item>
      </a-form-model>
    </a-spin>
  </a-modal>
</template>

<script>
import Vue from 'vue'
import { FormModel, TreeSelect, InputNumber } from 'ant-design-vue'
import store from '@/store'

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
        gid: 0,
        type: 0,
        name: '',
        paths: '',
        method: 'GET',
        sort: 0,
        status: 1
      },
      list: [],
      rules: {
        id: [
          { required: true, message: '不能空', trigger: 'blur' },
          { type: 'integer', min: 0, message: '不正确', trigger: 'blur' }
        ],
        gid: [
          { required: true, message: '不能空', trigger: 'blur' },
          { type: 'integer', min: 1, message: '请选择', trigger: 'blur' }
        ],
        type: [
          { required: true, message: '不能空', trigger: 'blur' },
          { type: 'enum', enum: [0, 1], message: '类型不正确', trigger: 'blur' }
        ],
        name: [
          { required: true, message: '不能为空', trigger: 'blur' },
          { min: 1, max: 32, message: '范围 1 - 32 字符', trigger: 'blur' }
        ],
        paths: [
          { required: true, message: '不能为空', trigger: 'blur' },
          { min: 1, max: 128, message: '范围 1 - 128 字符', trigger: 'blur' }
        ],
        method: [
          { required: true, message: '不能空', trigger: 'blur' },
          { type: 'enum', enum: ['GET', 'POST', 'PUT', 'DELETE'], message: '类型不正确', trigger: 'blur' }
        ],
        sort: [
          { required: true, message: '不能空', trigger: 'blur' },
          { type: 'integer', min: 0, message: '上级不正确', trigger: 'blur' }
        ],
        status: [
          { required: true, message: '不能空', trigger: 'blur' },
          { type: 'enum', enum: [0, 1], message: '类型不正确', trigger: 'blur' }
        ]
      }
    }
  },
  created () {
  },
  mounted () {
    this.getList()
  },
  methods: {
    handleOk () {
      return this.$emit('ok', this.$refs.ruleForm, this.model)
    },
    handleCancel () {
      return this.$emit('cancel', this.$refs.ruleForm)
    },
    typeChange () {
      this.model.gid = 0
      this.getList()
    },
    getList () {
      const that = this
      store
        .dispatch('getPermissionGroupAll', { type: this.model.type })
        .then(response => {
          if (response.code !== 0) {
            this.$notification.error({
              message: '错误',
              description: response.message
            })
          }
          that.list = response.data
        }).catch(() => {
        this.$notification.error({
          message: '错误',
          description: '请求信息失败，请重试'
        })
      })
    }
  }
}
</script>
