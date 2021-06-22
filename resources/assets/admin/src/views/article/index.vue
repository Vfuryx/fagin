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
          @click=""
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
      <el-table-column align="center" label="标题">
        <template slot-scope="scope">{{ scope.row.title }}</template>
      </el-table-column>
      <el-table-column align="center" label="作者">
        <template slot-scope="scope">{{ scope.row.author }}</template>
      </el-table-column>
      <el-table-column align="center" label="分类">
        <template slot-scope="scope">{{ scope.row.category }}</template>
      </el-table-column>
      <el-table-column align="center" label="标签">
        <template slot-scope="scope">
          <el-tag
            v-for="tag in scope.row.tags"
            :key="tag.id"
            type="success">
            {{ tag.name }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column align="center" label="推送时间">
        <template slot-scope="scope">{{ scope.row.post_at }}</template>
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

    <el-dialog
      title="创建"
      :visible.sync="dialogFormVisible"
      custom-class="min-width"
      width="860px"
      :destroy-on-close="true"
      :before-close="dialogBeforeClose"
    >
      <el-form
        ref="dataForm"
        :rules="rules"
        :model="temp"
        :inline="true"
        label-position="right"
        label-width="80px"
      >
        <el-form-item label="标题" prop="title">
          <el-input style="width: 330px" v-model="temp.title"/>
        </el-form-item>
        <el-form-item label="推送时间" prop="post_at">
          <el-date-picker
            v-model="temp.post_at"
            type="datetime"
            placeholder="选择日期时间">
          </el-date-picker>
        </el-form-item>
        <el-form-item label="作者" prop="author_id">
          <el-select v-model="temp.author_id" placeholder="请选择">
            <el-option
              v-for="item in authors"
              :key="item.id"
              :label="item.name"
              :value="item.id">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="分类" prop="category_id">
          <el-select v-model="temp.category_id" placeholder="请选择">
            <el-option
              v-for="item in categories"
              :key="item.id"
              :label="item.name"
              :value="item.id">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="开关">
          <el-switch v-model="temp.status" :active-value="1" :inactive-value="0"/>
        </el-form-item>
        <el-form-item label="标签" prop="tags">
          <el-select v-model="temp.tags" style="width: 330px" multiple clearable placeholder="请选择">
            <el-option
              v-for="item in tags"
              :key="item.id"
              :label="item.name"
              :value="item.id">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="摘要" prop="summary">
          <el-input
            type="textarea"
            :rows="4"
            style="width: 640px"
            placeholder="请输入内容"
            v-model="temp.summary">
          </el-input>
        </el-form-item>
        <el-form-item label="内容" prop="content">
          <Tinymce
            ref="editor"
            :upload-action="baseApi + '/v1/company/upload'"
            :height="700"
            :width="640"
            :upload-func="imagesUploadHandler"
            v-model="temp.content"
          />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取消</el-button>
        <el-button type="primary" @click="submitCreateForm()">提交</el-button>
        <!-- <el-button type="primary" @click="dialogStatus==='create'?createData():updateData()">提交</el-button> -->
      </div>
    </el-dialog>

    <el-dialog
      title="更新"
      :visible.sync="dialogUpdateVisible"
      custom-class="min-width"
      width="860px"
      :destroy-on-close="true"
    >
      <el-form
        ref="updateForm"
        :rules="rules"
        :model="temp"
        :inline="true"
        label-position="right"
        label-width="80px"
      >
        <el-form-item label="标题" prop="title">
          <el-input style="width: 330px" v-model="temp.title"/>
        </el-form-item>
        <el-form-item label="推送时间" prop="post_at">
          <el-date-picker
            v-model="temp.post_at"
            type="datetime"
            placeholder="选择日期时间">
          </el-date-picker>
        </el-form-item>
        <el-form-item label="作者" prop="author_id">
          <el-select v-model="temp.author_id" placeholder="请选择">
            <el-option
              v-for="item in authors"
              :key="item.id"
              :label="item.name"
              :value="item.id">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="分类" prop="category_id">
          <el-select v-model="temp.category_id" placeholder="请选择">
            <el-option
              v-for="item in categories"
              :key="item.id"
              :label="item.name"
              :value="item.id">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="开关" prop="status">
          <el-switch v-model="temp.status" :active-value="1" :inactive-value="0"/>
        </el-form-item>
        <el-form-item label="标签" prop="tags">
          <el-select v-model="temp.tags" multiple clearable placeholder="请选择">
            <el-option
              v-for="item in tags"
              :key="item.id"
              :label="item.name"
              :value="item.id">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="摘要" prop="summary">
          <el-input
            type="textarea"
            :rows="4"
            style="width: 640px"
            placeholder="请输入内容"
            v-model="temp.summary">
          </el-input>
        </el-form-item>
        <el-form-item label="内容" prop="content">
          <Tinymce
            ref="editor"
            :upload-action="baseApi + '/v1/company/upload'"
            :width="640"
            :height="700"
            :upload-func="imagesUploadHandler"
            v-model="temp.content"
          />
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
// import Upload from '@/components/Upload/singleImage4'
import { list, show, create, update, del, deletes } from '@/api/article'
import { all } from '@/api/author'
import { all as allCategory } from '@/api/category'
import { all as allTag } from '@/api/tag'
import Tinymce from '@/components/Tinymce'
import { uploadCompanyImage } from '@/api/admin'

export default {
  name: 'Article',
  components: {
    // Upload,
    Tinymce
  },
  // eslint-disable-next-line vue/order-in-components
  data() {
    return {
      baseApi: process.env.VUE_APP_BASE_API,
      // baseURL: process.env.VUE_APP_SERVER_PUBLIC_PATH,
      listLoading: true,
      dialogFormVisible: false,
      dialogShowVisible: false,
      dialogUpdateVisible: false,
      list: [],
      authors: [],
      categories: [],
      tags: [],
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
        title: [
          { type: 'string', required: true, message: '请输入正确的标题', trigger: 'blur' },
          { min: 1, max: 32, message: '标题太长，不能超过 32 个字符', trigger: 'blur' }
        ],
        post_at: [
          {  required: true, message: '时间不正确', trigger: 'blur' }
        ],
        author_id: [
          { type: 'number', required: true, message: '作者不正确', trigger: 'blur' }
        ],
        category_id: [
          { type: 'number', required: true, message: '分类不正确', trigger: 'blur' }
        ],
        content: [
          { type: 'string', required: true, message: '请输入正确的内容', trigger: 'blur' }
        ],
        summary: [
          { type: 'string', required: true, message: '请输入正确的内容', trigger: 'blur' }
        ],
        sort: [
          { type: 'number', required: true, message: '范围在 0 到 1000', trigger: 'blur' }
        ]
        // status: [
        //   { type: 'number', required: true, message: '请选择正确的状态', trigger: 'blur' },
        //   { min: 0, max: 2, message: '请选择正确的状态', trigger: 'blur' }
        // ]
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
        this.list = response.data.articles
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
        content: '',
        summary: '',
        post_at: '',
        author_id: '',
        category_id: '',
        tags: [],
        status: 1
      }
    },
    handleCreate() {
      this.resetTemp()
      all().then(response => {
        this.authors = response.data.authors
      })
      allCategory().then(response => {
        this.categories = response.data.categories
      })
      allTag().then(response => {
        this.tags = response.data.tags
      })
      this.dialogFormVisible = true
    },
    submitCreateForm() {
      this.$refs['dataForm'].validate(valid => {
        if (valid) {
          const data = {
            title: this.temp.title,
            content: this.temp.content,
            summary: this.temp.summary,
            post_at: this.temp.post_at,
            author_id: this.temp.author_id,
            category_id: this.temp.category_id,
            tags: this.temp.tags,
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
            this.resetTemp()

            // 关闭页面 跳转回列表页面
            // this.$store.dispatch('UpdateUser', this.$route)
            // this.$router.push({ name: 'UpdateUser' })
          })
        }
      })
    },
    handleUpdate(row) {
      this.dialogUpdateVisible = true
      show(row.id).then(response => {
        this.temp = response.data
        const t = []
        this.temp.tags.forEach(v => {
          t.push(v.id)
        })
        this.temp.tags = t
      })
      all().then(response => {
        this.authors = response.data.authors
      })
      allCategory().then(response => {
        this.categories = response.data.categories
      })
      allTag().then(response => {
        this.tags = response.data.tags
      })
    },
    submitUpdateForm(id) {
      this.$refs['updateForm'].validate(valid => {
        if (valid) {
          const data = {
            title: this.temp.title,
            content: this.temp.content,
            summary: this.temp.summary,
            post_at: this.temp.post_at,
            author_id: this.temp.author_id,
            category_id: this.temp.category_id,
            tags: this.temp.tags,
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
    },
    imagesUploadHandler(blobInfo, success, failure, progress) {

      const formData = new FormData()
      formData.append('file', blobInfo.blob())

      uploadCompanyImage(formData).then(res => {
        const url = process.env.VUE_APP_SERVER_PUBLIC_PATH + res.data.path
        success(url)
        // 提示信息
        this.$notify({
          title: '成功',
          message: '上传成功',
          type: 'success',
          duration: 2000
        })
        return
      }).catch(res => {
        failure('上传失败，请重试')
      })
    },
    dialogBeforeClose(done) {
      this.resetTemp()
      this.$refs.editor.setContent('')
      done()
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

/deep/.min-width {
  min-width: 800px;
}
</style>
