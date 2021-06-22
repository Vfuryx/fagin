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
        <span slot="roles" slot-scope="row">
          <a-tag color="blue" :key="k" v-for="(v,k) in row">{{ v['name'] }}</a-tag>
        </span>
        <span slot="action" slot-scope="text, record">
          <a-button icon="edit" @click="handleEdit(record)">编辑</a-button>
          <a-button icon="edit" @click="handleResetPassword(record)">修改密码</a-button>
          <a-popconfirm title="是否要强制下线？" @confirm="handleLogout(record.id)">
            <a-button icon="logout" >强制下线</a-button>
          </a-popconfirm>
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

      <reset-password
        ref="resetPassword"
        :visible="rs_visible"
        :loading="confirmLoading"
        :id="rw_id"
        @cancel="handleRSCancel"
        @ok="handleRSOk"
      />
    </a-card>
  </page-header-wrapper>
</template>

<script>
import CreateForm from './modules/CreateForm'
import ResetPassword from './modules/ResetPassword'
import store from '@/store'
import Vue from 'vue'
import { Pagination } from 'ant-design-vue'

Vue.use(Pagination)

export default {
  components: {
    CreateForm,
    ResetPassword
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
        { title: '昵称', dataIndex: 'nick_name', key: 'nick_name' },
        { title: '用户名', dataIndex: 'username', key: 'username' },
        { title: '角色', dataIndex: 'roles', key: 'roles', scopedSlots: { customRender: 'roles' } },
        { title: '创建时间', dataIndex: 'create_at', key: 'create_at' },
        { title: '操作', key: 'id', fixed: 'right', scopedSlots: { customRender: 'action' } }
      ],
      list: [],
      mdl: {
        id: 0,
        username: '',
        phone: '',
        email: '',
        nick_name: '',
        sex: 0,
        roles: [],
        status: 1
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
        .dispatch('getAdminList', this.params)
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
          this.list = response.data.users
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
        username: '',
        phone: '',
        email: '',
        nick_name: '',
        sex: 0,
        roles: [],
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
            store.dispatch('updateAdmin', data).then(response => {
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
            store.dispatch('addAdmin', data).then(response => {
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
      store.dispatch('showAdmin', record.id).then(response => {
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
      })
    },
    // 强制下线
    handleLogout (id) {
      // 强制下线
      store.dispatch('logoutAdmin', id).then(response => {
        if (response.code !== 0) {
          this.$notification.error({
            message: '错误',
            description: response.message
          })
          return false
        }
        this.$notification.success({
          message: '成功',
          description: '强制下线成功成功'
        })
      }).catch((e) => {
        this.$notification.error({
          message: '错误',
          description: '请求信息失败，请重试'
        })
        return false
      })
    },
    // 重置密码
    handleResetPassword (record) {
      this.rw_id = record.id
      this.rs_visible = true
      this.confirmLoading = false
    },
    handleRSCancel (form) {
      this.rs_visible = false
      form.resetFields() // 清理表单数据（可不做）
    },
    handleRSOk (form, data) {
      this.confirmLoading = true
      form.validate(valid => {
        if (valid) {
          store.dispatch('adminResetPassword', data).then(response => {
            if (response.code !== 0) {
              this.$notification.error({
                message: '错误',
                description: response.message
              })
              return false
            }
            this.rs_visible = false
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
            this.confirmLoading = false
            return false
          })
          return true
        } else {
          this.confirmLoading = false
          return false
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
