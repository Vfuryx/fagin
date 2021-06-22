<template>
  <page-header-wrapper>
    <a-card :bordered="false">
      <div class="table-operator">
        <a-button
          type="primary"
          icon="plus"
          @click="handleAdd">添加
        </a-button>
      </div>
      <div class="admin_table_list">
        <a-tabs :default-active-key="0" :animated="false" @change="tabChange">
          <a-tab-pane :key="0" tab="总后台">
            <a-table
              :columns="columns"
              :data-source="list"
              :pagination="false"
              :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange }"
              row-key="id">
              <span slot="sort" slot-scope="rows">
                {{ rows.sort }}
              </span>
              <span slot="action" slot-scope="text, record">
                <a-button
                  icon="edit"
                  @click="handleEdit(record)"
                >编辑</a-button>
                <!--                <a-button-->
                <!--                  icon="edit"-->
                <!--                  @click="handleDel(record.id)"-->
                <!--                >删除</a-button>-->
              </span>
            </a-table>
          </a-tab-pane>
          <a-tab-pane :key="1" tab="商家">
            <a-table
              :columns="columns"
              :data-source="list"
              :pagination="false"
              :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange }"
              row-key="id">
              <span slot="sort" slot-scope="rows">
                {{ rows.sort }}
              </span>
              <span slot="action" slot-scope="text, record">
                <a-button
                  icon="edit"
                  @click="handleEdit(record)"
                >编辑</a-button>
                <!--                <a-button-->
                <!--                  icon="edit"-->
                <!--                  @click="handleDel(record.id)"-->
                <!--                >删除</a-button>-->
              </span>
            </a-table>
          </a-tab-pane>
        </a-tabs>

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
          :list="list"
          @cancel="handleCancel"
          @ok="handleOk"
        />

      </div>
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
        total_count: 0,
        type: 0
      },
      total_page: 0, // 总页数
      selectedRowKeys: [], // 被选择的行
      columns: [
        // { title: '#', dataIndex: 'id', fixed: 'left' },
        { title: '接口分组', dataIndex: 'name' },
        { title: '排序', key: 'sort', fixed: 'right', width: '100px', scopedSlots: { customRender: 'sort' } },
        { title: '操作', key: 'id', fixed: 'right', scopedSlots: { customRender: 'action' } }
      ],
      list: [],
      confirmLoading: false,
      mdl: {
        id: 0,
        name: '',
        sort: 0,
        type: 0
      },
      visible: false
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
      this.params.current_page = e
      this.onload()
    },
    // 删除
    del () {
      // if (this.selectedRowKeys.length === 0) {
      //   return this.$message.error('未选择数据.')
      // }
      // this.$confirm({
      //   title: '你确定要删除选择的数据？',
      //   content: '确定删除后无法恢复.',
      //   okText: '是',
      //   okType: 'danger',
      //   cancelText: '取消',
      //   onOk: () => {
      //     const ids = this.selectedRowKeys.join(',')
      //     this.$delete(this.$api.adminPermissionGroups + '/' + ids).then(res => {
      //       if (res.error_code === 0) {
      //         this.onload()
      //         this.$message.success('删除成功')
      //       } else {
      //         this.$message.error(res.msg)
      //       }
      //     })
      //   }
      // })
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
        .dispatch('getPermissionGroupList', this.params)
        .then(response => {
          console.log(response)
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
        name: '',
        sort: 0,
        type: 0
      }
      this.visible = true
    },
    handleCancel (form) {
      this.visible = false
      // const form = this.$refs.createModal.$refs.ruleForm
      form.resetFields() // 清理表单数据（可不做）
    },
    handleOk (form, data) {
      form.validate(valid => {
        if (valid) {
          if (data.id > 0) {
            store.dispatch('updatePermissionGroup', data).then(response => {
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
            store.dispatch('addPermissionGroup', data).then(response => {
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
      store.dispatch('showPermissionGroup', record.id).then(response => {
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
    handleDel (id) {
      // 删除
      this.$confirm({
        title: '你确定要删除选择的数据？',
        content: '确定删除后无法恢复.',
        okText: '是',
        okType: 'danger',
        cancelText: '取消',
        onOk: () => {
          store.dispatch('delPermissionGroup', id).then(response => {
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
