<template>
  <page-header-wrapper>
    <a-card :bordered="false">
      <div class="table-operator">
        <a-button type="primary" icon="plus" @click="handleAdd">添加</a-button>
      </div>
      <a-table
        :columns="columns"
        :data-source="list"
        :pagination="false"
        :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange }"
        row-key="id">
        <span slot="action" slot-scope="text, record">
          <a-button icon="edit" @click="handleEdit(record)">编辑</a-button>
          <a-button icon="edit" @click="handleDel(record.id)">删除</a-button>
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
      />
    </a-card>
  </page-header-wrapper>
</template>

<script>
import CreateForm from './modules/CreateForm'
import store from '@/store'
import Vue from 'vue'
import { Pagination } from 'ant-design-vue'

Vue.use(Pagination)

export default {
  components: {
    CreateForm
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
        { title: '标题', dataIndex: 'title', key: 'title' },
        { title: '播放时长', dataIndex: 'duration', key: 'duration' },
        { title: '状态', dataIndex: 'status', key: 'status' },
        { title: '创建时间', dataIndex: 'created_at', key: 'created_at' },
        { title: '操作', key: 'id', fixed: 'right', scopedSlots: { customRender: 'action' } }
      ],
      list: [],
      mdl: {
        id: 0,
        title: '',
        duration: '',
        status: 0,
        create_at: ''
      },
      confirmLoading: false,
      visible: false,
      rs_visible: false,
      rw_id: 0
    }
  },
  watch: {},
  computed: {},
  methods: {
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
        .dispatch('getVideoList', this.params)
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
          this.list = response.data.videos
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
        title: '',
        duration: '',
        status: 0,
        create_at: ''
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
            store.dispatch('updateVideo', data).then(response => {
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
            store.dispatch('addVideo', data).then(response => {
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
      store.dispatch('showVideo', record.id).then(response => {
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
          store.dispatch('delVideo', id).then(response => {
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
