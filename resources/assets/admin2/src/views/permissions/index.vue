<template>
  <page-header-wrapper>
    <a-card :bordered="false">
      <!--    <admin-search :searchConfig="searchConfig" @searchParams="search"></admin-search>-->

      <div class="table-operator">
        <a-button type="primary" icon="plus" @click="handleAdd">添加</a-button>
      </div>

      <a-tabs :default-active-key="0" :animated="false" @change="tabChange">
        <a-tab-pane :key="0" tab="总后台">
          <a-table
            :columns="columns"
            :data-source="list"
            :pagination="false"
            :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange }"
            row-key="id"
          >
            <span slot="group" slot-scope="text">
              <a-tag color="blue">{{ text || '暂无' }}</a-tag>
            </span>showPagination="auto"
            <span slot="method" slot-scope="rows">
              <a-tag color="blue">{{ rows || '暂无' }}</a-tag>
            </span>
            <span slot="status" slot-scope="text, record">
              <a-badge
                :color="record.status ? 'green' : 'red'"
                :status="record.status ? 'processing' : 'default'"
                :text="record.status ? '开启' : '关闭'" />
            </span>
            <span slot="action" slot-scope="text, record">
              <a-button icon="edit" @click="handleEdit(record)">编辑</a-button>
              <a-popconfirm title="是否要删除接口？" @confirm="handleDelete(record.id)">
                <a-button icon="user-delete">删除接口</a-button>
              </a-popconfirm>
            </span>
          </a-table>
        </a-tab-pane>
        <a-tab-pane :key="1" tab="商家">
          <a-table
            :columns="columns"
            :data-source="list"
            :pagination="false"
            row-key="id"
          >
            <span slot="group" slot-scope="text">
              <a-tag color="blue">{{ text || '暂无' }}</a-tag>
            </span>
            <span slot="method" slot-scope="rows">
              <a-tag color="blue">{{ rows || '暂无' }}</a-tag>
            </span>
            <span slot="status" slot-scope="text, record">
              <a-badge
                :color="record.status ? 'green' : 'red'"
                :status="record.status ? 'processing' : 'default'"
                :text="record.status ? '开启' : '关闭'" />
            </span>
            <span slot="action" slot-scope="text, record">
              <a-button icon="edit" @click="handleEdit(record)">编辑1</a-button>
              <a-popconfirm title="是否要删除接口？" @confirm="handleDelete(record.id)">
                <a-button icon="user-delete">删除接口</a-button>
              </a-popconfirm>
            </span>
          </a-table>
        </a-tab-pane>
      </a-tabs>

      <a-pagination
        v-if="total_page > 0"
        class="ant-table-pagination"
        v-model="params.current_page"
        :page-size.sync="params.page_size"
        :total="params.total_count"
        @change="onChange"
        show-less-items />

      <create-form
        ref="createModal"
        :visible="visible"
        :loading="confirmLoading"
        :model="mdl"
        @cancel="handleCancel"
        @ok="handleOk"
      />

    </a-card>
  </page-header-wrapper>
</template>

<script>
import CreateForm from './modules/CreateForm'
import store from '@/store'
// import adminSearch from '@/components/admin/search'

import Vue from 'vue'
import { Pagination } from 'ant-design-vue'

Vue.use(Pagination)

export default {
  components: { CreateForm },
  props: {},
  data () {
    return {
      params: {
        current_page: 1,
        page_size: 10,
        total_count: 0,
        type: 0
      },
      total_page: 0, // 总页数
      searchConfig: [
        { label: '接口名称', name: 'name', type: 'text' },
        { label: '接口路由', name: 'apis', type: 'text' },
        { label: '所属分组', name: 'pid', type: 'select', data: [{ label: '全部', value: '' }] }
      ],
      selectedRowKeys: [], // 被选择的行
      columns: [
        // { title: '#', dataIndex: 'id', fixed: 'left' },
        { title: '接口名称', dataIndex: 'name', key: 'name' },
        { title: '所属分组', dataIndex: 'group', key: 'group', scopedSlots: { customRender: 'group' } },
        { title: '接口路由', dataIndex: 'paths', key: 'paths' },
        { title: '请求方法', dataIndex: 'method', key: 'method', scopedSlots: { customRender: 'method' } },
        { title: '菜单状态', dataIndex: 'status', key: 'status', scopedSlots: { customRender: 'status' } },
        { title: '排序', dataIndex: 'sort', key: 'sort', width: '100px' },
        { title: '操作', key: 'id', fixed: 'right', scopedSlots: { customRender: 'action' } }
      ],
      list: [],
      confirmLoading: false,
      mdl: {
        id: 0,
        gid: 0,
        type: 0,
        name: '',
        paths: '',
        method: 'GET',
        sort: 0,
        status: 1
      },
      visible: false
    }
  },
  watch: {},
  computed: {},
  methods: {
    // search(params) {
    //   this.params = params
    //   this.params.current_page = this.params.current_page
    //   this.params.page_size = this.params.page_size
    //   this.onload()
    // },
    // 选择框被点击
    onSelectChange (selectedRowKeys) {
      this.selectedRowKeys = selectedRowKeys
    },
    // 选择分页
    onChange (e) {
      this.params.current_page = e
      this.onload()
    },
    tabChange (e) {
      this.list = []
      this.total_page = 0
      this.to_nav(e)
    },
    to_nav (e) {
      this.params.type = e
      this.params.current_page = 1
      this.onload()
    },
    onload () {
      store
        .dispatch('getPermissionList', this.params)
        .then(response => {
          if (response.code !== 0) {
            this.$notification.error({
              message: '错误',
              description: response.message
            })
          }
          this.params.current_page = response.data.paginator.current_page
          this.params.page_size = response.data.paginator.page_size
          this.params.total_count = response.data.paginator.total_count
          this.total_page = response.data.paginator.total_page
          this.list = response.data.list
        }).catch(() => {
        this.$notification.error({
          message: '错误',
          description: '请求信息失败，请重试'
        })
      })
    },
    handleAdd () {
      this.mdl = {
        id: 0,
        gid: 0,
        type: 0,
        name: '',
        paths: '',
        method: 'GET',
        sort: 0,
        status: 1
      }
      this.visible = true
    },
    handleCancel (form) {
      this.visible = false
      form.resetFields() // 清理表单数据（可不做）
    },
    handleOk (form, data) {
      form.validate(valid => {
        if (valid) {
          if (data.id > 0) {
            store.dispatch('updatePermission', data).then(response => {
              if (response.code !== 0) {
                this.$notification.error({
                  message: '错误',
                  description: response.message
                })
                return false
              }
              this.visible = false
              this.confirmLoading = false
              // 重置表单数据
              form.resetFields()
              // 刷新表格
              this.onload()
              this.$message.info('修改成功')
              return true
            }).catch((e) => {
              console.log(e)
              this.$notification.error({
                message: '错误',
                description: '请求信息失败，请重试'
              })
              return false
            })
          } else {
            // 新增
            store.dispatch('addPermission', data).then(response => {
              if (response.code !== 0) {
                this.$notification.error({
                  message: '错误',
                  description: response.message
                })
                return false
              }
              this.visible = false
              this.confirmLoading = false
              // 重置表单数据
              form.resetFields()
              // 刷新表格
              this.onload()
              this.$message.info('新增成功')
              return true
            }).catch((e) => {
              console.log(e)
              this.$notification.error({
                message: '错误',
                description: '请求信息失败，请重试'
              })
              return false
            })
          }
          return true
        } else {
          // console.log('error submit!!', data)
          this.confirmLoading = false
          return false
        }
      })
    },
    handleEdit (record) {
      this.visible = true
      this.confirmLoading = true
      // 编辑
      store.dispatch('showPermission', record.id).then(response => {
        if (response.code !== 0) {
          this.$notification.error({
            message: '错误',
            description: response.message
          })
          this.visible = false
          return false
        }
        this.mdl = { ...response.data }
        this.confirmLoading = false
      }).catch((e) => {
        this.$notification.error({
          message: '错误',
          description: '请求信息失败，请重试'
        })
        this.visible = false
        return false
      })
    },
    handleDelete (id) {
      store.dispatch('delPermission', id).then(response => {
        if (response.code !== 0) {
          this.$notification.error({
            message: '错误',
            description: response.message
          })
          return false
        }
        // 刷新表格
        this.onload()
        this.$message.info('删除成功')
      }).catch((e) => {
        console.log(e)
        this.$notification.error({
          message: '错误',
          description: '请求信息失败，请重试'
        })
        return false
      })
    }
  },
  created () {
    this.onload()
  },
  mounted () {
  }
}
</script>
<style lang="scss" scoped>

</style>
