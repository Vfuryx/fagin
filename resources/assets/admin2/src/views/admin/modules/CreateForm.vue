<template>
  <a-modal
    :title="model.id > 0 ? '管理员编辑' : '管理员新增'"
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
          :label-col="{ span: 4 }"
          :wrapper-col="{span: 5}"
          v-if="model && model.id > 0"
          label="主键ID"
          prop="id"
          style="width: 88%;margin-right: 0">
          <a-input v-model="model.id" disabled />
        </a-form-model-item>
        <a-form-model-item
          label="用户名称"
          :label-col="{ span: 7 }"
          :wrapper-col="{span: 17}"
          prop="username"
          style="width: 50%;margin-right: 0">
          <a-input v-model="model.username"></a-input>
        </a-form-model-item>
        <a-form-model-item
          label="手机号码"
          :label-col="{ span: 7 }"
          :wrapper-col="{span: 17}"
          prop="phone"
          style="width: 50%;margin-right: 0">
          <a-input v-model="model.phone"></a-input>
        </a-form-model-item>
        <a-form-model-item
          label="邮箱"
          :label-col="{ span: 7 }"
          :wrapper-col="{span: 17}"
          prop="email"
          style="width: 50%;margin-right: 0">
          <a-input v-model="model.email"></a-input>
        </a-form-model-item>
        <a-form-model-item
          label="用户昵称"
          :label-col="{ span: 7 }"
          :wrapper-col="{span: 17}"
          prop="nick_name"
          style="width: 50%;margin-right: 0">
          <a-input v-model="model.nick_name"></a-input>
        </a-form-model-item>
        <a-form-model-item
          v-if="model && model.id === 0"
          label="用户密码"
          :label-col="{ span: 7 }"
          :wrapper-col="{span: 17}"
          prop="password"
          style="width: 50%;margin-right: 0">
          <a-input v-model="model.password"></a-input>
        </a-form-model-item>
        <a-form-model-item
          :label-col="{ span: 7 }"
          :wrapper-col="{span: 17}"
          label="用户性别"
          prop="sex"
          style="width: 50%;margin-right: 0">
          <a-select v-model="model.sex">
            <a-select-option v-for="(v,k) in sexOptions" :value="v.value||0" :key="k">
              {{ v.label }}
            </a-select-option>
          </a-select>
        </a-form-model-item>
        <a-form-model-item
          :label-col="{ span: 4 }"
          :wrapper-col="{span: 16}"
          label="角色"
          prop="roles"
          style="width: 88%;margin-right: 0">
          <a-tree-select
            :replace-fields="{children:'children', title:'name', key:'id', value: 'id' }"
            v-model="model.roles"
            :tree-data="list"
            tree-checkable
            show-checked-strategy="SHOW_PARENT"></a-tree-select>
        </a-form-model-item>
        <a-form-model-item
          label="管理员状态"
          :label-col="{ span: 4 }"
          :wrapper-col="{span: 10}"
          prop="status"
          style="width: 88%;margin-right: 0">
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
    return {
      // 性别状态字典
      sexOptions: [
        { label: '未公开', value: 0 },
        { label: '男', value: 1 },
        { label: '女', value: 2 }
      ],
      form: {
        id: 0,
        username: '',
        phone: '',
        email: '',
        nick_name: '',
        sex: 0,
        roles: [],
        status: 1
      },
      list: [],
      rules: {
        id: [
          { required: true, message: '不能空', trigger: 'blur' },
          { type: 'integer', min: 0, message: '不正确', trigger: 'blur' }
        ],
        username: [
          { required: true, message: '不能为空', trigger: 'blur' },
          { min: 1, max: 32, message: '范围 1 - 32 字符', trigger: 'blur' },
          { validator: this.validateUsername, trigger: 'blur' }
        ],
        nick_name: [
          { required: true, message: '不能为空', trigger: 'blur' },
          { min: 1, max: 32, message: '范围 1 - 32 字符', trigger: 'blur' }
        ],
        phone: [
          { required: true, message: '手机号码不能为空', trigger: 'blur' },
          { pattern: /^1[3|4|5|6|7|8|9][0-9]\d{8}$/, message: '请输入正确的手机号码', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '密码不能为空', trigger: 'blur' },
          { min: 8, max: 32, message: '长度在 8 到 32 个字符', trigger: 'blur' }
        ],
        email: [
          { required: true, message: '邮箱地址不能为空', trigger: 'blur' },
          { type: 'email', message: '\'请输入正确的邮箱地址', trigger: ['blur', 'change'] }
        ],
        sex: [
          { required: true, message: '不能空', trigger: 'blur' },
          { type: 'enum', enum: [0, 1, 2], message: '类型不正确', trigger: 'blur' }
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
    getList () {
      const that = this
      store
        .dispatch('getRoleAll', { type: this.model.type })
        .then(response => {
          if (response.code !== 0) {
            this.$notification.error({
              message: '错误',
              description: response.message
            })
          }
          that.list = response.data.roles
        }).catch(() => {
        this.$notification.error({
          message: '错误',
          description: '请求信息失败，请重试'
        })
      })
    },
    validateUsername (rule, value, callback, source, options) {
      store
        .dispatch('adminUsernameExist', { 'username': value, 'id': this.model.id })
        .then(response => {
          if (response.code !== 0) {
            return callback(new Error('用户名已存在'))
          }
          return callback()
        }).catch(() => {
        return callback(new Error('用户名已存在'))
      })
    }
  }
}
</script>
