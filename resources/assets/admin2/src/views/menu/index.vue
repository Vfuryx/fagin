<template>
  <page-header-wrapper>
    <a-card :bordered="false">
      <div class="table-operator">
        <a-button
          type="primary"
          icon="plus"
          @click="handleAdd(0)">添加
        </a-button>
        <a-button @click="clear_cache">
          <!--          <a-font type="iconitemno_0"></a-font>-->
          清除缓存
        </a-button>
        <!--        <a-button class="admin_delete_btn" type="danger" icon="delete" @click="del">批量删除</a-button>-->
      </div>
      <a-table
        ref="table"
        :columns="columns"
        :data-source="list"
        :pagination="false"
        row-key="id">
        <span slot="sort" slot-scope="rows">
          <a-input type="number" @blur="sortChange(rows)" v-model="rows.sort" />
        </span>
        <span slot="status" slot-scope="text, record">
          <a-badge
            :status="record.status ? 'processing' : 'default'"
            :color="record.status ? 'green' : 'red'"
            :text="record.status ? '开启' : '关闭'" />
        </span>
        <span slot="is_show" slot-scope="text, record">
          <a-badge :status="record.is_show ? 'processing' : 'default' " :text="record.is_show ? '开启' : '关闭'" />
        </span>
        <span slot="is_hide_child" slot-scope="text, record">
          <a-badge
            :status="record.is_hide_child ? 'default' : 'processing'"
            :text="record.is_hide_child ? '关闭' : '开启' " />
        </span>
        <span slot="action" slot-scope="text, record">
          <a-button
            icon="edit"
            @click="handleEdit(record)"
          >编辑</a-button>
          <a-button
            icon="edit"
            @click="handleAdd(record.id)"
          >新增</a-button>
          <a-button
            icon="edit"
            @click="handleDel(record.id)"
          >删除</a-button>
        </span>
      </a-table>
      <create-form
        ref="createModal"
        :visible="visible"
        :loading="confirmLoading"
        :model="mdl"
        :list="list"
        @cancel="handleCancel"
        @ok="handleOk"
      />
    </a-card>
  </page-header-wrapper>
</template>

<script>
import CreateForm from './modules/CreateForm'
import store from '@/store'

export default {
  components: {
    CreateForm
  },
  props: {},
  data () {
    return {
      type: 0,
      selectedRowKeys: [], // 被选择的行
      columns: [
        //   {title:'#',dataIndex:'id',fixed:'left'},
        { title: '菜单名称', dataIndex: 'title', key: 'title' },
        { title: '路径', dataIndex: 'path', key: 'paths' },
        { title: '组件名称', dataIndex: 'component', key: 'component' },
        // { title: '图标', dataIndex: 'icon', key: 'icon' },
        { title: '菜单状态', dataIndex: 'status', key: 'status', scopedSlots: { customRender: 'status' } },
        { title: '菜单展示', dataIndex: 'is_show', key: 'is_show', scopedSlots: { customRender: 'is_show' } },
        {
          title: '子菜单展示',
          dataIndex: 'is_hide_child',
          key: 'is_hide_child',
          scopedSlots: { customRender: 'is_hide_child' }
        },
        // { title: '排序', key: 'path', fixed: 'right', width: '100px', scopedSlots: { customRender: 'sort' } },
        { title: '操作', key: 'id', fixed: 'right', scopedSlots: { customRender: 'action' } }
      ],
      list: [],
      confirmLoading: false,
      mdl: {
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
    // 删除
    del () {
      if (this.selectedRowKeys.length === 0) {
        return this.$message.error('未选择数据.')
      }
      this.$confirm({
        title: '你确定要删除选择的数据？',
        content: '确定删除后无法恢复.',
        okText: '是',
        okType: 'danger',
        cancelText: '取消',
        onOk: () => {
          const ids = this.selectedRowKeys.join(',')
          this.$delete(this.$api.adminMenus + '/' + ids).then(res => {
            if (res.error_code === 0) {
              this.onload()
              this.$message.success('删除成功')
            } else {
              this.$message.error(res.msg)
            }
          })
        }
      })
    },
    // 清空缓存
    clear_cache () {
      this.$get(this.$api.adminMenusClearCache).then(res => {
        return this.$message.success(res.msg)
      })
    },
    onload () {
      const type = this.$route.query.type
      const params = {}
      if (!this.$isEmpty(type)) {
        params.type = type
        this.type = type
      }
      store
        .dispatch('getMenuList').then(response => {
        console.log(response)
        if (response.code !== 0) {
          this.$notification.error({
            message: '错误',
            description: response.message
          })
          return
        }
        this.list = response.data
      }).catch(() => {
        this.$notification.error({
          message: '错误',
          description: '请求信息失败，请重试'
        })
      })
    },
    // 排序移动
    sortChange (rows) {
      const api = this.$apiHandle(this.$api.adminMenus, rows.id)
      this.$put(api.url, rows).then(res => {
        if (res.error_code === 0) {
          this.onload()
          return this.$message.success(res.msg)
        } else {
          return this.$message.error(res.msg)
        }
      })
    },
    handleCancel (form) {
      this.visible = false
      form.resetFields() // 清理表单数据（可不做）
    },
    handleOk (form, data) {
      form.validate(valid => {
        if (valid) {
          if (data.id > 0) {
            store.dispatch('updateMenu', data).then(response => {
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
            store.dispatch('addMenu', data).then(response => {
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
    handleAdd (pid) {
      this.mdl = {
        id: 0,
        type: 0,
        parent_id: pid,
        icon: 'smile',
        is_show: 0,
        name: '',
        path: '',
        target: '',
        redirect: '',
        component: '',
        status: 1,
        sort: 0,
        is_hide_child: 0,
        title: ''
      }
      this.visible = true
    },
    handleEdit (record) {
      this.visible = true
      this.confirmLoading = true
      // 新增
      store.dispatch('showMenu', record.id).then(response => {
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
          store.dispatch('delMenu', id).then(response => {
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

<style lang="less" scoped>
//import "@/components/index.less";

/deep/ .ant-badge-status-processing:after {
  border: 1px solid #50f522;
}

</style>
