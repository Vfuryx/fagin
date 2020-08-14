<template>
  <div class="app-container">
    <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
      <el-form-item label="系统模块" prop="title">
        <el-input
          v-model="queryParams.title"
          placeholder="请输入系统模块"
          clearable
          style="width: 240px;"
          size="small"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="操作人员" prop="operName">
        <el-input
          v-model="queryParams.operName"
          placeholder="请输入操作人员"
          clearable
          style="width: 240px;"
          size="small"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="类型" prop="businessType">
        <el-select
          v-model="queryParams.businessType"
          placeholder="操作类型"
          clearable
          size="small"
          style="width: 240px"
        >
          <el-option
            v-for="dict in typeOptions"
            :key="dict.dictValue"
            :label="dict.dictLabel"
            :value="dict.dictValue"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="状态" prop="status">
        <el-select
          v-model="queryParams.status"
          placeholder="操作状态"
          clearable
          size="small"
          style="width: 240px"
        >
          <el-option
            v-for="dict in statusOptions"
            :key="dict.dictValue"
            :label="dict.dictLabel"
            :value="dict.dictValue"
          />
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
        <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
      </el-form-item>
    </el-form>

    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button
          v-permisaction="['system:sysoperlog:export']"
          type="warning"
          icon="el-icon-download"
          size="mini"
          @click="handleExport"
        >导出</el-button>
      </el-col>
    </el-row>

    <el-table v-loading="listLoading" :data="list" element-loading-text="拼命加载中">
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
          <el-button
            size="mini"
            type="text"
            icon="el-icon-view"
            @click="handleShowLog(scope.row.id)"
          >查看</el-button>
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
          <el-input v-model="temp.input" type="textarea" :rows="10" disabled />
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
      },
      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        title: undefined,
        operName: undefined,
        businessType: undefined,
        status: undefined
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
