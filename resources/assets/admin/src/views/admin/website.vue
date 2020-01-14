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
              <el-form-item label="网站名称:" prop="web_name">
                <el-input v-model="postForm.web_name" placeholder="请输入网站名称" />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="12">
              <el-form-item label="英文名称:" prop="web_en_name">
                <el-input v-model="postForm.web_en_name" placeholder="请输入网站的英文名称" />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="12">
              <el-form-item label="网站标题:" prop="web_title">
                <el-input v-model="postForm.web_title" placeholder="请输入网站的网站标题" />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="12">
              <el-form-item label="关键词:" prop="keywords">
                <el-input
                  v-model="postForm.keywords"
                  placeholder="请输入网站的关键词，多个关键词请以英文状态下逗号隔开，5-8个关键词最佳"
                />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="12">
              <el-form-item label="描述:" prop="description">
                <el-input v-model="postForm.description" placeholder="请输入网站的描述" />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="12">
              <el-form-item label="公司名称:" prop="company_name">
                <el-input v-model="postForm.company_name" placeholder="请输入公司名称" />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="12">
              <el-form-item label="联系电话:" prop="contact_number">
                <el-input v-model="postForm.contact_number" placeholder="请输入联系电话" />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="12">
              <el-form-item label="公司地址:" prop="company_address">
                <el-input v-model="postForm.company_address" placeholder="请输入公司地址" />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="12">
              <el-form-item label="邮箱地址:" prop="email">
                <el-input v-model="postForm.email" placeholder="请输入邮箱地址" />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="12">
              <el-form-item label="ICP备案:" prop="icp">
                <el-input v-model="postForm.icp" placeholder="请输入ICP备案" />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="12">
              <el-form-item label="公安部备案:" prop="public_security_record">
                <el-input v-model="postForm.public_security_record" placeholder="请输入公安部备案" />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="18">
              <el-form-item label="网站logo" prop="web_logo">
                <Upload
                  v-model="postForm.web_logo"
                  tip="只能上传image/jpeg文件，且不超过20mb"
                  :upload-action="baseApi+'/v1/website/upload'"
                />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="18">
              <el-form-item label="二维码:" prop="qr_code">
                <Upload
                  v-model="postForm.qr_code"
                  tip="只能上传image/jpeg文件，且不超过20mb"
                  :upload-action="baseApi+'/v1/website/upload'"
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
import Upload from '@/components/Upload/singleImage4'
import { getInfo, updateInfo } from '@/api/admin'

export default {
  name: 'WebsiteSetup',
  components: {
    Upload
  },
  // eslint-disable-next-line vue/order-in-components
  data() {
    return {
      baseApi: process.env.VUE_APP_BASE_API,
      postForm: {
        web_name: '',
        web_en_name: '',
        web_title: '',
        keywords: '',
        description: '',
        web_logo: '',
        company_name: '',
        contact_number: '',
        company_address: '',
        email: '',
        icp: '',
        public_security_record: '',
        qr_code: ''
      },
      rules: {
        web_name: [
          { type: 'string', required: true, message: '请输入网站名称', trigger: 'blur' },
          { min: 1, max: 32, message: '长度在 1 到 32 个字符', trigger: 'blur' }
        ],
        web_en_name: [
          { type: 'string', required: true, message: '请输入网站的英文名称', trigger: 'blur' },
          { min: 1, max: 32, message: '长度在 1 到 32 个字符', trigger: 'blur' }
        ],
        web_title: [
          { type: 'string', required: true, message: '请输入网站的网站标题', trigger: 'blur' },
          { min: 1, max: 32, message: '长度在 1 到 32 个字符', trigger: 'blur' }
        ],
        keywords: [
          { type: 'string', required: true, message: '请输入网站的关键词', trigger: 'blur' },
          { min: 1, max: 127, message: '长度在 1 到 127 个字符', trigger: 'blur' }
        ],
        description: [
          { type: 'string', required: true, message: '请输入网站的描述', trigger: 'blur' },
          { min: 1, max: 255, message: '长度在 1 到 255 个字符', trigger: 'blur' }
        ],
        company_name: [
          { type: 'string', required: true, message: '请输入公司名称', trigger: 'blur' },
          { min: 1, max: 32, message: '长度在 1 到 32 个字符', trigger: 'blur' }
        ],
        contact_number: [
          { type: 'string', required: true, message: '请输入联系电话', trigger: 'blur' },
          { min: 1, max: 16, message: '长度在 1 到 16 个字符', trigger: 'blur' }
        ],
        company_address: [
          { type: 'string', required: true, message: '请输入公司地址', trigger: 'blur' },
          { min: 1, max: 32, message: '长度在 1 到 32 个字符', trigger: 'blur' }
        ],
        email: [
          { type: 'email', required: true, message: '请输入正确的邮箱地址', trigger: 'blur' },
          { min: 1, max: 32, message: '长度在 1 到 32 个字符', trigger: 'blur' }
        ],
        icp: [
          { type: 'string', required: true, message: '请输入ICP备案', trigger: 'blur' },
          { min: 1, max: 32, message: '长度在 1 到 32 个字符', trigger: 'blur' }
        ],
        public_security_record: [
          { type: 'string', required: true, message: '请输入公安部备案', trigger: 'blur' },
          { min: 1, max: 32, message: '长度在 1 到 32 个字符', trigger: 'blur' }
        ],
        web_logo: [
          { type: 'string', required: true, message: '上传网站logo', trigger: 'blur' },
          { min: 1, max: 255, message: '长度在 1 到 255 个字符', trigger: 'blur' }
        ],
        qr_code: [
          { type: 'string', required: true, message: '二维码', trigger: 'blur' },
          { min: 1, max: 255, message: '长度在 1 到 255 个字符', trigger: 'blur' }
        ],

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
      getInfo().then(response => {
        const data = response.data
        this.postForm.web_name = data.web_name
        this.postForm.web_en_name = data.web_en_name
        this.postForm.web_title = data.web_title
        this.postForm.keywords = data.keywords
        this.postForm.description = data.description
        this.postForm.web_logo = data.web_logo
        this.postForm.company_name = data.company_name
        this.postForm.contact_number = data.contact_number
        this.postForm.company_address = data.company_address
        this.postForm.email = data.email
        this.postForm.icp = data.icp
        this.postForm.public_security_record = data.public_security_record
        this.postForm.qr_code = data.qr_code
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
            web_name: this.postForm.web_name,
            web_en_name: this.postForm.web_en_name,
            web_title: this.postForm.web_title,
            keywords: this.postForm.keywords,
            description: this.postForm.description,
            web_logo: this.postForm.web_logo,
            company_name: this.postForm.company_name,
            contact_number: this.postForm.contact_number,
            company_address: this.postForm.company_address,
            email: this.postForm.email,
            icp: this.postForm.icp,
            public_security_record: this.postForm.public_security_record,
            qr_code: this.postForm.qr_code
          }
          updateInfo(data).then(res => {
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