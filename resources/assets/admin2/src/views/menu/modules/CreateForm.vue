<template>
  <a-modal
    :title="model.id > 0 ? '菜单编辑' : '菜单新增'"
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
          style="width: 70%;margin-right: 0">
          <a-input v-model="model.id" disabled />
        </a-form-model-item>

        <a-form-model-item
          :label-col="{ span: 7 }"
          :wrapper-col="{span: 17}"
          label="菜单类型"
          prop="type"
          style="width: 50%;margin-right: 0">
          <a-select v-model="model.type">
            <a-select-option :value="0">后台菜单</a-select-option>
            <a-select-option :value="1">商家菜单</a-select-option>
          </a-select>
        </a-form-model-item>
        <a-form-model-item
          :label-col="{ span: 7 }"
          :wrapper-col="{span: 17}"
          label="上级菜单"
          prop="parent_id"
          style="width: 50%;margin-right: 0">
          <a-tree-select v-model="model.parent_id">
            <a-tree-select-node title="顶级菜单" :value="0"></a-tree-select-node>
            <template v-for="v in list">
              <a-tree-select-node :title="v.title" :value="v.id" :key="v.id">
                <template v-for="vo in v.children">
                  <a-tree-select-node
                    :key="vo.id"
                    :title="vo.title"
                    :value="vo.id"></a-tree-select-node>
                </template>
              </a-tree-select-node>
            </template>
          </a-tree-select>
        </a-form-model-item>
        <a-form-model-item
          label="菜单名称"
          :label-col="{ span: 7 }"
          :wrapper-col="{span: 17}"
          prop="title"
          style="width: 50%;margin-right: 0">
          <a-input v-model="model.title"></a-input>
        </a-form-model-item>
        <a-form-model-item
          label="路径"
          :label-col="{ span: 7 }"
          :wrapper-col="{span: 17}"
          prop="path"
          style="width: 50%;margin-right: 0">
          <a-input v-model="model.path"></a-input>
        </a-form-model-item>
        <a-form-model-item
          label="路由名称"
          :label-col="{ span: 7 }"
          :wrapper-col="{span: 17}"
          prop="name"
          style="width: 50%;margin-right: 0">
          <a-input v-model="model.name"></a-input>
        </a-form-model-item>
        <a-form-model-item
          label="组件名称"
          :label-col="{ span: 7 }"
          :wrapper-col="{span: 17}"
          prop="component"
          style="width: 50%;margin-right: 0">
          <a-input v-model="model.component"></a-input>
        </a-form-model-item>
        <a-form-model-item
          label="重定向"
          :label-col="{ span: 7 }"
          :wrapper-col="{span: 17}"
          prop="redirect"
          style="width: 50%;margin-right: 0">
          <a-input v-model="model.redirect"></a-input>
        </a-form-model-item>
        <a-form-model-item
          label="目标"
          :label-col="{ span: 7 }"
          :wrapper-col="{span: 17}"
          prop="target"
          style="width: 50%;margin-right: 0">
          <a-select v-model="model.target">
            <a-select-option value="">无</a-select-option>
            <a-select-option value="_blank">新窗口</a-select-option>
            <a-select-option value="_parent">父窗口</a-select-option>
            <a-select-option value="_self">默认</a-select-option>
            <a-select-option value="_top">新浏览器</a-select-option>
          </a-select>
        </a-form-model-item>
        <a-form-model-item
          label="图标"
          :label-col="{ span: 4 }"
          :wrapper-col="{span: 20}"
          prop="icon"
          style="width: 88%;margin-right: 0">
          <a-auto-complete
            class="certain-category-search"
            dropdown-class-name="certain-category-search-dropdown"
            :dropdown-match-select-width="false"
            :dropdown-style="{ width: '300px' }"
            size="large"
            style="width: 100%"
            placeholder="icon"
            option-label-prop="value"
          >
            <a-input>
              <a-icon slot="suffix" :type="model.icon" class="certain-category-icon" />
            </a-input>
            <template slot="dataSource">
              <a-select-option key="all" disabled class="show-all">
                <icon-selector @change="handleIconChange" />
              </a-select-option>
            </template>
          </a-auto-complete>
        </a-form-model-item>
        <a-form-model-item
          label="显示排序"
          :label-col="{ span: 7 }"
          :wrapper-col="{span: 17}"
          prop="sort"
          style="width: 50%;margin-right: 0">
          <a-input-number v-model="model.sort" controls-position="right" :min="0" />
        </a-form-model-item>
        <a-form-model-item
          label="菜单展示"
          :label-col="{ span: 5 }"
          :wrapper-col="{span: 18}"
          prop="is_show"
          style="width: 70%;margin-right: 0">
          <a-radio-group name="is_show" v-model="model.is_show">
            <a-radio :value="1">开启</a-radio>
            <a-radio :value="0">关闭</a-radio>
          </a-radio-group>
        </a-form-model-item>
        <a-form-model-item
          label="子菜单展示"
          :label-col="{ span: 5 }"
          :wrapper-col="{span: 18}"
          prop="is_hide_child"
          style="width: 70%;margin-right: 0">
          <a-radio-group name="is_hide_child" v-model="model.is_hide_child">
            <a-radio :value="0">展示</a-radio>
            <a-radio :value="1">隐藏</a-radio>
          </a-radio-group>
        </a-form-model-item>
        <a-form-model-item
          label="菜单状态"
          :label-col="{ span: 5 }"
          :wrapper-col="{span: 18}"
          prop="status"
          style="width: 70%;margin-right: 0">
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
import IconSelector from '@/components/IconSelector'
import { FormModel, TreeSelect, InputNumber, AutoComplete } from 'ant-design-vue'

Vue.use(FormModel)
Vue.use(TreeSelect)
Vue.use(InputNumber)
Vue.use(AutoComplete)

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
  components: {
    IconSelector
  },
  data () {
    return {
      form: {
        id: 0,
        type: 0,
        parent_id: 0,
        icon: 'smile',
        is_show: '',
        name: '',
        path: '',
        target: '',
        redirect: '',
        component: '',
        status: 1,
        sort: 0,
        is_hide_child: 0,
        title: ''
      },
      rules: {
        id: [
          { required: true, message: '不能空', trigger: 'blur' },
          { type: 'integer', min: 0, message: '不正确', trigger: 'blur' }
        ],
        type: [
          { required: true, message: '不能空', trigger: 'blur' },
          { type: 'enum', enum: [0, 1], message: '类型不正确', trigger: 'blur' }
        ],
        is_hide_child: [
          { required: true, message: '不能空', trigger: 'blur' },
          { type: 'enum', enum: [0, 1], message: '类型不正确', trigger: 'blur' }
        ],
        is_show: [
          { required: true, message: '不能空', trigger: 'blur' },
          { type: 'enum', enum: [0, 1], message: '类型不正确', trigger: 'blur' }
        ],
        status: [
          { required: true, message: '不能空', trigger: 'blur' },
          { type: 'enum', enum: [0, 1], message: '类型不正确', trigger: 'blur' }
        ],
        parent_id: [
          { required: true, message: '不能空', trigger: 'blur' },
          { type: 'integer', min: 0, message: '上级不正确', trigger: 'blur' }
        ],
        sort: [
          { required: true, message: '不能空', trigger: 'blur' },
          { type: 'integer', min: 0, message: '上级不正确', trigger: 'blur' }
        ],
        title: [
          { required: true, message: '不能为空', trigger: 'blur' },
          { min: 1, max: 32, message: '范围 1 - 32 字符', trigger: 'blur' }
        ],
        path: [
          { required: true, message: '不能为空', trigger: 'blur' },
          { min: 1, max: 128, message: '范围 1 - 128 字符', trigger: 'blur' }
        ],
        component: [
          { required: true, message: '不能为空', trigger: 'blur' },
          { min: 1, max: 32, message: '范围 1 - 32 字符', trigger: 'blur' }
        ],
        redirect: [
          { min: 1, max: 128, message: '范围 1 - 128 字符', trigger: 'blur' }
        ],
        target: [
          { min: 1, max: 32, message: '范围 1 - 32 字符', trigger: 'blur' }
        ],
        icon: [
          { min: 1, max: 32, message: '范围 1 - 32 字符', trigger: 'blur' }
        ],
        name: [
          { required: true, message: '不能为空', trigger: 'blur' },
          { min: 1, max: 32, message: '范围 1 - 32 字符', trigger: 'blur' }
        ]
      },
      dataSource: []
    }
  },
  created () {
  },
  methods: {
    handleOk () {
      return this.$emit('ok', this.$refs.ruleForm, this.model)
    },
    handleIconChange (icon) {
      console.log('change Icon', icon)
      this.model.icon = icon
    },
    handleCancel () {
      return this.$emit('cancel', this.$refs.ruleForm)
    }
  }
}
</script>
