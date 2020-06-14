<template>
  <div class="app-container">
    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button
          v-permisaction="['system:syspost:add']"
          type="primary"
          icon="el-icon-plus"
          size="mini"
          @click="handleCreate"
        >新增</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button
          v-permisaction="['system:syspost:remove']"
          type="danger"
          icon="el-icon-delete"
          size="mini"
          :disabled="multiple"
          @click="handleDelete"
        >删除</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button
          v-permisaction="['system:syspost:export']"
          type="warning"
          icon="el-icon-download"
          size="mini"
          @click="handleExport"
        >导出</el-button>
      </el-col>
    </el-row>

    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="拼命加载中"
      fit
      highlight-current-row
    >
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column align="center" label="ID" width="60">
        <template slot-scope="scope">{{ scope.row.id }}</template>
      </el-table-column>
      <el-table-column align="center" label="标题">
        <template slot-scope="scope">{{ scope.row.title }}</template>
      </el-table-column>
      <el-table-column align="center" label="轮播图">
        <template slot-scope="scope">
          <img style="width: 50px; height: 50px" :src="baseURL + scope.row.banner" />
        </template>
      </el-table-column>
      <el-table-column align="center" label="路径">
        <template slot-scope="scope">{{ scope.row.path }}</template>
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
            v-permisaction="['system:syspost:edit']"
            size="mini"
            type="text"
            icon="el-icon-edit"
            @click="handleUpdateVideo(scope.row)"
          >修改</el-button>
          <el-button
            v-permisaction="['system:syspost:remove']"
            size="mini"
            type="text"
            icon="el-icon-delete"
            @click="handleDelete(scope.row)"
          >删除</el-button>
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

    <el-dialog title="创建轮播图" :visible.sync="dialogFormVisible">
      <el-form
        ref="dataForm"
        :rules="rules"
        :model="temp"
        label-position="left"
        label-width="65px"
        style="width: 550px; margin-left:50px;"
      >
        <el-form-item label="标题" prop="title">
          <el-input v-model="temp.title" />
        </el-form-item>
        <el-form-item label="轮播图" prop="banner">
          <Upload
            v-model="temp.banner"
            tip="只能上传image/jpeg文件，且不超过20mb"
            :upload-action="baseApi + '/v1/banner/upload'"
          />
        </el-form-item>
        <el-form-item label="路径" prop="path">
          <el-input v-model="temp.path" />
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input-number v-model="temp.sort" :min="0" :max="1000" label="排序" />
        </el-form-item>
        <el-form-item label="状态">
          <el-switch v-model="temp.status" :active-value="1" :inactive-value="0" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取消</el-button>
        <el-button type="primary" @click="submitCreateForm()">提交</el-button>
        <!-- <el-button type="primary" @click="dialogStatus==='create'?createData():updateData()">提交</el-button> -->
      </div>
    </el-dialog>

    <el-dialog title="更新轮播图" :visible.sync="dialogUpdateVisible">
      <el-form
        ref="updateForm"
        :rules="rules"
        :model="temp"
        label-position="left"
        label-width="70px"
        style="width: 550px; margin-left:50px;"
      >
        <el-form-item label="标题" prop="title">
          <el-input v-model="temp.title" />
        </el-form-item>
        <el-form-item label="轮播图" prop="banner">
          <Upload
            v-model="temp.banner"
            tip="只能上传image/jpeg文件，且不超过20mb"
            :upload-action="baseApi + '/v1/banner/upload'"
          />
        </el-form-item>
        <el-form-item label="路径" prop="path">
          <el-input v-model="temp.path" />
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input-number v-model="temp.sort" :min="0" :max="1000" label="排序" />
        </el-form-item>
        <el-form-item label="状态">
          <el-switch v-model="temp.status" :active-value="1" :inactive-value="0" />
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
import Upload from '@/components/Upload/singleImage4'
import { list, show, create, update } from '@/api/banner'

export default {
  name: 'BannerIndex',
  components: {
    Upload
  },
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
      temp: {
        id: 0,
        title: '',
        banner: '',
        path: '',
        sort: 100,
        status: 1
      },
      rules: {
        title: [
          { type: 'string', required: true, message: '请输入正确的标题', trigger: 'blur' },
          { min: 1, max: 32, message: '长度在 8 到 32 个字符', trigger: 'blur' }
        ],
        banner: [
          { type: 'string', required: true, message: '请上传正确的轮播图', trigger: 'blur' },
          { min: 1, max: 255, message: '长度在 1 到 255 个字符', trigger: 'blur' }
        ],
        path: [
          { type: 'string', required: true, message: '请输入正确的路径', trigger: 'blur' },
          { min: 1, max: 255, message: '长度在 1 到 255 个字符', trigger: 'blur' }
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
        this.list = response.data.banners
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
            title: this.temp.title,
            banner: this.temp.banner,
            path: this.temp.path,
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
    handleUpdateVideo(row) {
      this.dialogUpdateVisible = true
      this.temp = {
        id: row.id,
        title: row.title,
        banner: row.banner,
        path: row.path,
        sort: row.sort,
        status: row.status
      }
    },
    submitUpdateForm(id) {
      this.$refs['updateForm'].validate(valid => {
        if (valid) {
          const data = {
            title: this.temp.title,
            banner: this.temp.banner,
            path: this.temp.path,
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
    beforeUpload(file) {
      const isJPG = file.type === 'image/jpeg'
      const isLt2M = file.size / 1024 / 1024 < 20

      if (!isJPG) {
        this.$message.error('上传图片只能是 JPG 格式!222')
      }
      if (!isLt2M) {
        this.$message.error('上传图片大小不能超过 20MB!222')
      }
      return isJPG && isLt2M
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