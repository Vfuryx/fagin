<template>
  <div class="app-container">
    <!-- $t is vue-i18n global function to translate lang -->
    <el-row>
      <!-- <el-select v-model="wechatID" placeholder="请选择" @change="getUserlist">
        <el-option
          v-for="item in wechatList"
          :key="item.app_id"
          :label="item.app_id"
          :value="item.id"
        ></el-option>
      </el-select>
      <el-button type="primary" @click="syncUserList">同步微信公众号用户列表</el-button>
      <el-button type="primary" @click="syncUserInfo">同步微信公众号用户信息</el-button>-->
    </el-row>

    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="拼命加载中"
      border
      fit
      highlight-current-row
    >
      <el-table-column align="center" label="ID" width="60">
        <template slot-scope="scope">{{ scope.row.id }}</template>
      </el-table-column>
      <el-table-column align="center" label="方法">
        <template slot-scope="scope">{{ scope.row.method }}</template>
      </el-table-column>
      <el-table-column align="center" label="操作">
        <template slot-scope="scope">{{ scope.row.operation }}</template>
      </el-table-column>
      <el-table-column align="center" label="路径">
        <template slot-scope="scope">{{ scope.row.path }}</template>
      </el-table-column>
      <el-table-column align="center" label="IP">
        <template slot-scope="scope">{{ scope.row.ip }}</template>
      </el-table-column>
      <el-table-column align="center" label="用户">
        <template slot-scope="scope">{{ scope.row.user }}</template>
      </el-table-column>
      <el-table-column align="center" label="操作时间">
        <template slot-scope="scope">{{ scope.row.created_at }}</template>
      </el-table-column>
      <el-table-column align="center" label="操作">
        <template slot-scope="scope">
          <el-button size="mini" type="success" @click="handleShowLog(scope.row.id)">查看</el-button>
        </template>
      </el-table-column>
    </el-table>

    <pagination
      v-show="total >= 1"
      :total="total"
      :page.sync="listQuery.page"
      :limit.sync="listQuery.limit"
      @pagination="fetchData"
    />

    <el-dialog title="详情" :visible.sync="dialogShowVisible">
      <el-form label-position="left" label-width="70px" style="width: 400px; margin-left:50px;">
        <el-form-item label="操作" prop="operation">
          <el-input v-model="temp.operation" disabled />
        </el-form-item>
        <el-form-item label="用户" prop="user">
          <el-input v-model="temp.user" disabled />
        </el-form-item>
        <el-form-item label="路径" prop="path">
          <el-input v-model="temp.path" disabled />
        </el-form-item>
        <el-form-item label="IP" prop="ip">
          <el-input v-model="temp.ip" disabled />
        </el-form-item>
        <el-form-item label="输入" prop="input">
          <el-input type="textarea" :rows="10" v-model="temp.input" disabled />
        </el-form-item>
      </el-form>
    </el-dialog>
  </div>
</template>

<script>
import { logList, show } from '@/api/operation-log'
import Pagination from '@/components/Pagination' // Secondary package based on el-pagination

export default {
  name: 'OperationLog',
  components: {
    Pagination
  },
  // eslint-disable-next-line vue/order-in-components
  data() {
    return {
      listLoading: true,
      dialogShowVisible: false,
      list: [],
      total: 0,
      listQuery: {
        page: 1,
        limit: 15,
        importance: undefined,
        title: undefined,
        type: undefined,
        sort: '+id'
      },
      temp: {
        user: '',
        method: '',
        path: '',
        ip: '',
        operation: '',
        input: ''
      }
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.listLoading = true
      logList(this.listQuery).then(response => {
        this.list = response.data.logs
        const { total_count, current_page, page_size } = response.data.paginator
        this.total = total_count
        this.listQuery.page = current_page
        this.listQuery.limit = page_size
        this.listLoading = false
      })
      this.listLoading = false
    },
    handleShowLog(id) {
      this.dialogShowVisible = true
      show(id).then(response => {
        this.temp = response.data
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.el-row {
  margin-bottom: 20px;
  &:last-child {
    margin-bottom: 0;
  }
}
.clearfix:after {
  /*伪元素是行内元素 正常浏览器清除浮动方法*/
  content: "";
  display: block;
  height: 0;
  clear: both;
  visibility: hidden;
}
.clearfix {
  *zoom: 1; /*ie6清除浮动的方式 *号只有IE6-IE7执行，其他浏览器不执行*/
}
</style>