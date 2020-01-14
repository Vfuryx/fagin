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
              <el-form-item label="用户名:" prop="username">
                <el-input
                  :rows="1"
                  v-model="postForm.username"
                  type="text"
                  class="article-textarea"
                  autosize
                  placeholder="请输入内容"
                  disabled
                />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="12">
              <el-form-item label="邮件:" prop="email">
                <el-input
                  :rows="1"
                  v-model="postForm.email"
                  type="text"
                  class="article-textarea"
                  autosize
                  placeholder="请输入内容"
                />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="12">
              <el-form-item label="旧密码:" prop="old_password">
                <el-input
                  :rows="1"
                  v-model="postForm.old_password"
                  type="password"
                  class="article-textarea"
                  autosize
                  placeholder="请输入内容"
                />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="12">
              <el-form-item label="新密码:" prop="new_password">
                <el-input
                  :rows="1"
                  v-model="postForm.new_password"
                  type="password"
                  class="article-textarea"
                  autosize
                  placeholder="请输入内容"
                />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="12">
              <el-form-item label="重复新密码:" prop="rep_new_password">
                <el-input
                  :rows="1"
                  v-model="postForm.rep_new_password"
                  type="password"
                  class="article-textarea"
                  autosize
                  placeholder="请输入内容"
                />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="12">
              <el-button v-loading="loading" type="primary" @click="submitForm">提交</el-button>
            </el-col>
          </el-row>
        </div>
      </el-form>
    </div>
  </div>
</template>

<script>
import { show, update } from '@/api/admin'

export default {
  name: 'UpdateUser',
  computed: {
  },
  // eslint-disable-next-line vue/order-in-components
  data() {
    return {
      postForm: {
        username: '',
        email: '',
        new_password: '',
        old_password: '',
        rep_new_password: '',
        avatar: ''
      },
      rules: {
        email: [
          { type: 'email', required: true, message: '请输入正确的邮箱地址', trigger: 'blur' }
        ],
        old_password: [
          { type: 'string', required: false, message: '请输入旧密码', trigger: 'blur' },
          { min: 8, max: 20, message: '长度在 8 到 32 个字符', trigger: 'blur' },
          { validator: this.oldPassworldValidator, trigger: 'blur' }
        ],
        new_password: [
          { type: 'string', required: false, message: '请输入密码', trigger: 'blur' },
          { min: 8, max: 20, message: '长度在 8 到 32 个字符', trigger: 'blur' },
          { validator: this.newPasswordValidator, trigger: 'blur' }
        ],
        rep_new_password: [
          { type: 'string', required: false, message: '请输入密码', trigger: 'blur' },
          { min: 8, max: 20, message: '长度在 8 到 32 个字符', trigger: 'blur' },
          { validator: this.repNewPassworldValidator, trigger: 'blur' }
        ]
      },
      loading: false
    }
  },
  created() {
    this.id = 1
    this.fetchData(this.id)
  },
  methods: {
    fetchData() {
      this.loading = true
      show().then(response => {
        const data = response.data
        this.postForm.username = data.name
        this.postForm.email = data.email
        this.postForm.avatar = data.avatar
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
            'email': this.postForm.email
          }
          console.log(this.postForm)
          if (this.postForm.old_password !== '') {
            Object.assign(data, {
              old_password: this.postForm.old_password,
              new_password: this.postForm.new_password
            })
          }
          update(this.id, data).then(res => {
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
            return
          })
        }
        console.log('error submit!!')
        this.loading = false
        return false
      })
    },
    oldPassworldValidator(rule, value, callback) {
      if (value === '' && (this.postForm.new_password !== '' || this.postForm.rep_new_password !== '')) {
        return callback(new Error('请输入旧密码'))
      }
      return callback()
    },
    newPasswordValidator(rule, value, callback) {
      if (value === '' && (this.postForm.old_password !== '' || this.postForm.rep_new_password !== '')) {
        return callback(new Error('请输入新密码'))
      }
      return callback()
    },
    repNewPassworldValidator(rule, value, callback) {
      if (value === '' && (this.postForm.old_password !== '' || this.postForm.new_password !== '')) {
        return callback(new Error('请输入重复新密码'))
      }
      if (value !== this.postForm.new_password) {
        return callback(new Error('与新密码不匹配'))
      }
      return callback()
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