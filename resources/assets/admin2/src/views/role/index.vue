<template>
  <page-header-wrapper>
    <a-card :bordered="false">
      <div class="table-operator">
        <a-button @click="$router.push('/admin/roles/form/')" type="primary" icon="plus">添加</a-button>
      </div>

      <a-tabs :default-active-key="0" :animated="false" @change="tabChange">
        <a-tab-pane :key="0" tab="总后台">
          <a-table
            :columns="columns"
            :data-source="list"
            :pagination="false"
            :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange }"
            row-key="id">
            <span slot="action" slot-scope="text, record">
              <a-button icon="edit" @click="$router.push('/admin/roles/form/'+record.id)">编辑</a-button>
              <a-popconfirm title="是否要删除角色？" @confirm="handleDelete(record.id)" v-if="record.id !== 1">
                <a-button icon="user-delete" >删除角色</a-button>
              </a-popconfirm>
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
            <span slot="action" slot-scope="text, record">
              <a-button icon="edit" @click="$router.push('/admin/roles/form/'+record.id)">编辑</a-button>
              <a-popconfirm title="是否要删除角色？" @confirm="handleDelete(record.id)" v-if="record.id !== 1">
                <a-button icon="user-delete" >删除角色</a-button>
              </a-popconfirm>
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
    </a-card>
  </page-header-wrapper>
</template>

<script>
import store from '@/store'
import Vue from 'vue'
import { Pagination } from 'ant-design-vue'
Vue.use(Pagination)

export default {
  components: {},
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
        { title: '角色名称', dataIndex: 'name', key: 'name' },
        { title: '操作', key: 'id', fixed: 'right', scopedSlots: { customRender: 'action' } }
      ],
      list: []
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
        .dispatch('getRoleList', this.params)
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
          this.list = response.data.roles
        }).catch(() => {
        this.$notification.error({
          message: '错误',
          description: '请求信息失败，请重试'
        })
      })
    },
    handleDelete (id) {
      store.dispatch('delRole', id).then(response => {
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
<style lang="less" scoped>

</style>
