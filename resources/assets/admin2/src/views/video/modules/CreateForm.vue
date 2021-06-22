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
          label="标题"
          :label-col="{ span: 4 }"
          :wrapper-col="{span: 5}"
          prop="title"
          style="width: 88%;margin-right: 0">
          <a-input v-model="model.title"></a-input>
        </a-form-model-item>
        <a-form-model-item
          label="视频上传"
          :label-col="{ span: 4 }"
          :wrapper-col="{span: 10}"
          prop="title"
          style="width: 88%;margin-right: 0">
          <a-upload
            list-type="picture"
            :multiple="false"
            :before-upload="beforeUpload"
            :custom-request="uploadImage"
            accept="video/mp4"
          >
            <a-button>
              <a-icon type="upload" />
              Upload
            </a-button>
          </a-upload>
        </a-form-model-item>
        <a-form-model-item
          label="描述"
          :label-col="{ span: 4 }"
          :wrapper-col="{span: 19}"
          prop="description"
          style="width: 88%;margin-right: 0">
          <a-textarea v-model="model.description"></a-textarea>
        </a-form-model-item>
        <a-form-model-item
          label="状态"
          :label-col="{ span: 4 }"
          :wrapper-col="{span: 5}"
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
      form: {
        id: 0,
        title: '',
        duration: '',
        status: 0,
        create_at: '',
        description: ''
      },
      list: [],
      rules: {
        id: [
          { required: true, message: '不能空', trigger: 'blur' },
          { type: 'integer', min: 0, message: '不正确', trigger: 'blur' }
        ],
        title: [
          { required: true, message: '不能为空', trigger: 'blur' },
          { min: 1, max: 60, message: '范围 1 - 60 字符', trigger: 'blur' }
        ],
        description: [
          { required: true, message: '不能为空', trigger: 'blur' },
          { min: 1, max: 255, message: '范围 1 - 255 字符', trigger: 'blur' }
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
  },
  methods: {
    handleOk () {
      return this.$emit('ok', this.$refs.ruleForm, this.model)
    },
    handleCancel () {
      return this.$emit('cancel', this.$refs.ruleForm)
    },
    beforeUpload (file) {
      const isJpgOrPng = file.type === 'video/mp4'
      if (!isJpgOrPng) {
        this.$message.error('请上传视频格式!')
      }
      const isLt2M = file.size / 1024 / 1024 < 100
      if (!isLt2M) {
        this.$message.error('视频大小限制 10m 以内!')
      }
      return isJpgOrPng && isLt2M
    },
    // 上传头像
    uploadImage (file) {
      this.confirmLoading = true
      const formData = new FormData()
      formData.append('file', file.file)
      // 新增
      store.dispatch('uploadVideo', formData).then(response => {
        if (response.code !== 0) {
          this.$notification.error({
            message: '错误',
            description: response.message
          })
          return false
        }
        this.confirmLoading = false
        this.model.path = response.data.path
        this.$message.info('新增成功')
        return true
      }).catch((e) => {
        this.$notification.error({
          message: '错误',
          description: '请求信息失败，请重试'
        })
        return false
      })
    }
  }
}
</script>
