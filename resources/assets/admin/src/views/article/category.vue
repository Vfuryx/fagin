<template>
  <div class="app-container">
    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button
          type="primary"
          icon="el-icon-plus"
          size="mini"
          @click="handleCreate"
        >新增
        </el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button
          type="danger"
          icon="el-icon-delete"
          size="mini"
          :disabled="multiple"
          @click="handleDelete"
        >删除
        </el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button
          type="warning"
          icon="el-icon-download"
          size="mini"
          @click="handleExport"
        >导出
        </el-button>
      </el-col>
    </el-row>

    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="拼命加载中"
      fit
      highlight-current-row
      @selection-change="handleSelectionChange"
    >
      <el-table-column type="selection" width="55" align="center"/>
      <el-table-column align="center" label="ID" width="60">
        <template slot-scope="scope">{{ scope.row.id }}</template>
      </el-table-column>
      <el-table-column align="center" label="名称">
        <template slot-scope="scope">{{ scope.row.name }}</template>
      </el-table-column>
      <el-table-column align="center" label="排序">
        <template slot-scope="scope">{{ scope.row.sort }}</template>
      </el-table-column>
      <el-table-column align="center" label="状态">
        <template slot-scope="scope">
          <el-tag>{{ scope.row.status ? '开启' : '关闭' }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column align="center" label="操作">
        <template slot-scope="scope">
          <el-button
            size="mini"
            type="text"
            icon="el-icon-edit"
            @click="handleUpdate(scope.row)"
          >修改
          </el-button>
          <el-button
            size="mini"
            type="text"
            icon="el-icon-delete"
            @click="handleDelete(scope.row)"
          >删除
          </el-button>
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

    <el-dialog title="创建" :visible.sync="dialogFormVisible">
      <el-form
        ref="dataForm"
        :rules="rules"
        :model="temp"
        label-position="left"
        label-width="65px"
        style="width: 550px; margin-left:50px;"
      >
        <el-form-item label="名称" prop="name">
          <el-input v-model="temp.name"/>
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input-number v-model="temp.sort" :min="0" :max="1000" label="排序" />
        </el-form-item>
        <el-form-item label="状态">
          <el-switch v-model="temp.status" :active-value="1" :inactive-value="0"/>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取消</el-button>
        <el-button type="primary" @click="submitCreateForm()">提交</el-button>
      </div>
    </el-dialog>

    <el-dialog title="更新" :visible.sync="dialogUpdateVisible">
      <el-form
        ref="updateForm"
        :rules="rules"
        :model="temp"
        label-position="left"
        label-width="70px"
        style="width: 550px; margin-left:50px;"
      >
        <el-form-item label="名称" prop="name">
          <el-input v-model="temp.name"/>
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input-number v-model="temp.sort" :min="0" :max="1000" label="排序" />
        </el-form-item>
        <el-form-item label="状态">
          <el-switch v-model="temp.status" :active-value="1" :inactive-value="0"/>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogUpdateVisible = false">取消</el-button>
        <el-button type="primary" @click="submitUpdateForm()">提交</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { list, show, create, update, del, deletes } from '@/api/category'

export default {
  name: 'Category',
  components: {},
  // eslint-disable-next-line vue/order-in-components
  data() {
    return {
      baseApi: process.env.VUE_APP_BASE_API,
      baseURL: process.env.VUE_APP_SERVER_PUBLIC_PATH,
      listLoading: true,
      dialogFormVisible: false,
      dialogShowVisible: false,
      dialogUpdateVisible: false,
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
      // 选中数组
      ids: [],
      // 非单个禁用
      single: true,
      // 非多个禁用
      multiple: true,
      temp: {
        id: 0,
        title: '',
        banner: '',
        path: '',
        sort: 100,
        status: 1
      },
      rules: {
        name: [
          { type: 'string', required: true, message: '请输入正确的标题', trigger: 'blur' },
          { min: 1, max: 32, message: '长度在 8 到 32 个字符', trigger: 'blur' }
        ],
        sort: [
          { type: 'number', required: true, message: '范围在 0 到 1000', trigger: 'blur' }
        ],
        status: [
          { type: 'number', required: true, message: '请选择正确的状态', trigger: 'blur' },
          { min: 0, max: 2, message: '请选择正确的状态', trigger: 'blur' }
        ]
      }
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.listLoading = true
      list(this.listQuery).then(response => {
        this.list = response.data.categories
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
    },
    resetTemp() {
      this.temp = {
        id: 0,
        title: '',
        banner: '',
        path: '',
        sort: 100,
        status: 1
      }
    },
    handleCreate() {
      this.resetTemp()
      this.dialogFormVisible = true
    },
    submitCreateForm() {
      this.$refs['dataForm'].validate(valid => {
        if (valid) {
          const data = {
            name: this.temp.name,
            sort: this.temp.sort,
            status: this.temp.status
          }
          create(data).then(res => {
            this.dialogFormVisible = false
            // 提示信息
            this.$notify({
              title: '成功',
              message: '新增成功',
              type: 'success',
              duration: 2000
            })
            this.fetchData()

            // 关闭页面 跳转回列表页面
            // this.$store.dispatch('UpdateUser', this.$route)
            // this.$router.push({ name: 'UpdateUser' })
          })
        }
      })
    },
    handleUpdate(row) {
      this.dialogUpdateVisible = true
      this.temp = {
        id: row.id,
        name: row.name,
        sort: row.sort,
        status: row.status
      }
    },
    submitUpdateForm(id) {
      this.$refs['updateForm'].validate(valid => {
        if (valid) {
          const data = {
            name: this.temp.name,
            sort: this.temp.sort,
            status: this.temp.status
          }
          update(this.temp.id, data).then(res => {
            this.dialogUpdateVisible = false
            this.fetchData()

            // 提示信息
            this.$notify({
              title: '成功',
              message: '修改成功',
              type: 'success',
              duration: 2000
            })
            // 关闭页面 跳转回列表页面
            // this.$store.dispatch('UpdateUser', this.$route)
            // this.$router.push({ name: 'UpdateUser' })
          })
        }
      })
    },
    /** 删除按钮操作 */
    handleDelete(row) {
      const ids = row.id || this.ids
      this.$confirm('是否确认删除名称为"' + ids + '"的数据项?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function() {
        if (isNaN(ids)) {
          return deletes(ids)
        } else {
          return del(ids)
        }
      }).then(() => {
        this.fetchData()
        this.$notify({
          title: '成功',
          message: '删除成功',
          type: 'success',
          duration: 2000
        })
      }).catch(function() {
      })
    },
    // 多选框选中数据
    handleSelectionChange(selection) {
      this.ids = selection.map(item => item.id)
      this.single = selection.length !== 1
      this.multiple = !selection.length
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
