<template>
  <div class="container">
    <div class="createPost-container">
      <el-form
        label-position="right"
        :rules="rules"
        v-loading="loading"
        label-width="100px"
        ref="postForm"
        :model="postForm"
        class="form-container"
      >
        <div class="createPost-main-container">
          <el-row>
            <el-col :span="12">
              <el-form-item label="公司名称:" prop="name">
                <el-input v-model="postForm.name" placeholder="请输入公司名称" />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="18">
              <el-form-item label="图片" prop="image">
                <Upload
                  v-model="postForm.image"
                  tip="只能上传image/jpeg文件，且不超过20mb"
                  :upload-action="baseApi+'/v1/company/upload'"
                />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="18">
              <el-form-item label="状态">
                <el-switch v-model="postForm.status" :active-value="1" :inactive-value="0" />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="24">
              <el-form-item label="内容:" prop="content">
                <!-- markdown 编辑器-->
                <!-- <markdown-editor
                  ref="markdownEditor"
                  v-model="postForm.content"
                  :options="{hideModeSwitch:true,previewStyle:'tab'}"
                  height="700px"
                />-->

                <Tinymce
                  ref="editor"
                  :upload-action="baseApi + '/v1/company/upload'"
                  :height="700"
                  :upload-func="imagesUploadHandler"
                  v-model="postForm.content"
                />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="12">
              <el-button v-loading="loading" type="primary" @click="submitForm">保存</el-button>
            </el-col>
          </el-row>
        </div>
      </el-form>
    </div>
  </div>
</template>

<script>
import Tinymce from '@/components/Tinymce'
// import MarkdownEditor from '@/components/MarkdownEditor'
import Upload from '@/components/Upload/singleImage4'
import { getCompanyIntroduction, updateCompanyIntroduction, uploadCompanyImage } from '@/api/admin'

export default {
  name: 'CompanyIntroduction',
  components: {
    Upload, Tinymce
  },
  // eslint-disable-next-line vue/order-in-components
  data() {
    return {
      baseApi: process.env.VUE_APP_BASE_API,
      postForm: {
        id: 0,
        name: '',
        image: '',
        content: '',
        status: 0
      },
      rules: {
        name: [
          { type: 'string', required: true, message: '请输入公司名称', trigger: 'blur' },
          { min: 1, max: 32, message: '长度在 1 到 32 个字符', trigger: 'blur' }
        ],
        image: [
          { type: 'string', required: true, message: '请输入图片', trigger: 'blur' },
          { min: 1, max: 255, message: '长度在 1 到 255 个字符', trigger: 'blur' }
        ],
        content: [
          { type: 'string', required: true, message: '请输入内容', trigger: 'blur' }
        ],
        status: [
          { type: 'number', required: true, message: '请选择正确的状态', trigger: 'blur' },
          { min: 0, max: 2, message: '请选择正确的状态', trigger: 'blur' }
        ]
      },
      loading: false
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.loading = true
      getCompanyIntroduction().then(response => {
        const data = response.data
        this.postForm.name = data.name
        this.postForm.image = data.image
        this.postForm.content = data.content
        this.postForm.status = data.status
        this.loading = false
      }).catch(err => {
        console.log(err)
      })
    },
    submitForm() {
      this.edit()
    },
    edit() {
      this.$refs.postForm.validate(valid => {
        if (valid) {
          this.loading = true
          const data = {
            name: this.postForm.name,
            image: this.postForm.image,
            content: this.postForm.content,
            status: this.postForm.status
          }
          updateCompanyIntroduction(data).then(res => {
            // 提示信息
            this.$notify({
              title: '成功',
              message: '修改成功',
              type: 'success',
              duration: 2000
            })
            return
          })
        }
        console.log('error submit!!')
        this.loading = false
        return false
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
    }
  }
}
</script>

<style lang="scss" scoped>
.container {
  padding: 20px;
  .filter-container {
    padding-bottom: 10px;
    .filter-item {
      float: right;
    }
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
@import "~@/styles/mixin.scss";
.createPost-container {
  position: relative;
  .createPost-main-container {
    padding: 40px 45px 20px 50px;
    .postInfo-container {
      position: relative;
      @include clearfix;
      margin-bottom: 10px;
      .postInfo-container-item {
        float: left;
      }
    }
    .editor-container {
      min-height: 500px;
      margin: 0 0 30px;
      .editor-upload-btn-container {
        text-align: right;
        margin-right: 10px;
        .editor-upload-btn {
          display: inline-block;
        }
      }
    }
  }
  .word-counter {
    width: 40px;
    position: absolute;
    right: -10px;
    top: 0px;
  }
}
</style>