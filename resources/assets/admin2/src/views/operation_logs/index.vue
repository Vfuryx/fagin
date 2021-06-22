<template>
  <page-header-wrapper>
    <a-card :bordered="false">
      <div class="table-operator">
        <search :searchConfig="searchConfig" @searchParams="search"></search>
      </div>
      <a-table
        :columns="columns"
        :data-source="list"
        :pagination="false"
        :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange }"
        row-key="id">
        <span slot="action" slot-scope="text, record">
          <a-button
            icon="edit"
            @click="handleEdit(record)"
          >编辑</a-button>
        </span>
      </a-table>
      <a-pagination
        v-if="total_page>0"
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
        :okButtonProps="{}"
      />
    </a-card>
  </page-header-wrapper>
</template>

<script>
import Search from '@/components/Search/search'
import CreateForm from './modules/CreateForm'
import store from '@/store'
import Vue from 'vue'
import { Pagination } from 'ant-design-vue'

Vue.use(Pagination)

export default {
  components: {
    CreateForm,
    Search
  },
  props: {},
  data () {
    return {
      params: {
        current_page: 1,
        page_size: 10,
        total_count: 0
      },
      total_page: 0, // 总页数
      selectedRowKeys: [], // 被选择的行
      columns: [
        // { title: '#', dataIndex: 'id', fixed: 'left' },
        { title: '日志编号', dataIndex: 'id', key: '#' },
        { title: '系统模块', dataIndex: 'module', key: 'module' },
        // { title: '操作类型', key: 'operation', key: 'operation' },
        { title: '请求方式', dataIndex: 'method', key: 'method' },
        { title: '操作人员', dataIndex: 'user', key: 'user' },
        { title: '请求地址', dataIndex: 'path', key: 'path' },
        { title: '操作日期', dataIndex: 'created_at', key: 'created_at' },
        { title: '操作', key: 'id', fixed: 'right', scopedSlots: { customRender: 'action' } }
      ],
      list: [],
      confirmLoading: false,
      mdl: {
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
      visible: false,
      searchConfig: [
        { label: '请求地址', name: 'path', type: 'text' },
        {
          label: '请求方式',
          name: 'method',
          type: 'select',
          data: [
            { label: 'ALL', value: 'ALL' },
            { label: 'GET', value: 'GET' },
            { label: 'POST', value: 'POST' },
            { label: 'PUT', value: 'PUT' },
            { label: 'DELETE', value: 'DELETE' }
          ]
        },
        { label: '时间范围', name: 'time', type: 'date_picker' }
      ]
    }
  },
  watch: {},
  computed: {},
  methods: {
    search (params) {
      const currentPage = this.params.current_page
      const pageSize = this.params.page_size
      if (params['time'] !== undefined) {
        const startTime = params['time'][0]
        const endTime = params['time'][1]
        this.params.start_time = startTime
        this.params.end_time = endTime
      }
      if (params['path'] !== undefined) {
        this.params.path = params['path']
      }
      if (params['method'] !== undefined) {
        this.params.method = params['method']
      }
      this.params.current_page = currentPage
      this.params.page_size = pageSize
      this.onload()
    },
    // 选择框被点击
    onSelectChange (selectedRowKeys) {
      this.selectedRowKeys = selectedRowKeys
    },
    // 选择分页
    onChange (e) {
      this.params.page = e
      this.onload()
    },
    onload () {
      store
        .dispatch('getOperationLogList', this.params)
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
          this.list = response.data.logs
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
      }
      this.visible = true
    },
    handleCancel (form) {
      this.visible = false
      form.resetFields() // 清理表单数据（可不做）
    },
    handleOk (form, data) {
    },
    handleEdit (record) {
      this.visible = true
      this.confirmLoading = true
      // 编辑
      store.dispatch('showOperationLog', record.id).then(response => {
        if (response.code !== 0) {
          this.$notification.error({
            message: '错误',
            description: response.message
          })
          this.visible = true
          return false
        }
        this.mdl = { ...response.data }
        this.confirmLoading = false
      }).catch((e) => {
        console.log(e)
        this.$notification.error({
          message: '错误',
          description: '请求信息失败，请重试'
        })
        this.visible = false
        return false
      })
    },
    handleDel (id) {
      // 删除
      this.$confirm({
        title: '你确定要删除选择的数据？',
        content: '确定删除后无法恢复.',
        okText: '是',
        okType: 'danger',
        cancelText: '取消',
        onOk: () => {
          store.dispatch('delOperationLog', id).then(response => {
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
