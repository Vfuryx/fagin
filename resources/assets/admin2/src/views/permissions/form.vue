<template>
  <div class="qingwu">
    <div class="admin_table_page_title"><a-button @click="$router.back()" class="float_right" icon="arrow-left">返回</a-button>接口编辑</div>
    <div class="unline underm"></div>
    <div class="admin_form">
      <a-form-model :label-col="{ span: 5 }" :wrapper-col="{ span: 12 }">
        <a-form-model-item label="分组类型">
          <a-select v-model="info.type" @change="typeChange">
            <a-select-option :value="0">
              总后台
            </a-select-option>
            <a-select-option :value="1">
              商家
            </a-select-option>
          </a-select>
        </a-form-model-item>
        <a-form-model-item label="所属分组">
          <a-select v-model="info.group_id">
            <a-select-option v-for="(v,k) in list" :value="v.id||0" :key="k">{{ v.name }}</a-select-option>
          </a-select>
        </a-form-model-item>
        <a-form-model-item label="接口名称">
          <a-input v-model="info.name"></a-input>
        </a-form-model-item>
        <a-form-model-item label="接口路由">
          <a-input v-model="info.apis"></a-input>
        </a-form-model-item>
        <a-form-model-item label="接口描述">
          <a-input v-model="info.content"></a-input>
        </a-form-model-item>

        <a-form-model-item :wrapper-col="{ span: 12, offset: 5 }">
          <a-button type="primary" @click="handleSubmit">提交</a-button>
        </a-form-model-item>
      </a-form-model>

    </div>
  </div>
</template>

<script>
export default {
    components: {},
    props: {},
    data () {
      return {
          info: {
              type: 0,
              group_id: '',
              name: '',
              apis: '',
              content: ''
          },
          list: [],
          id: 0
      }
    },
    watch: {},
    computed: {},
    methods: {
        handleSubmit () {
            // 验证代码处
            if (this.$isEmpty(this.info.group_id)) {
                return this.$message.error('所属分组不能为空')
            }
            if (this.$isEmpty(this.info.name)) {
                return this.$message.error('接口名不能为空')
            }
            if (this.$isEmpty(this.info.apis)) {
                return this.$message.error('接口路由不能为空')
            }

            const api = this.$apiHandle(this.$api.adminPermissions, this.id)
            if (api.status) {
                this.$put(api.url, this.info).then(res => {
                    if (res.error_code === 0) {
                        this.$message.success(res.msg)
                        return this.$router.back()
                    } else {
                        return this.$message.error(res.msg)
                    }
                })
            } else {
                this.$post(api.url, this.info).then(res => {
                    if (res.error_code === 0) {
                        this.$message.success(res.msg)
                        return this.$router.back()
                    } else {
                        return this.$message.error(res.msg)
                    }
                })
            }
        },
        get_info () {
            this.$get(this.$api.adminPermissions + '/' + this.id).then(res => {
                this.info = res.data
                this.getList()
            })
        },
        typeChange () {
            this.info.group_id = ''
            this.getList()
        },
        // 获取菜单列表
        onload () {
            // 判断你是否是编辑
            if (!this.$isEmpty(this.$route.params.id)) {
                this.id = this.$route.params.id
                this.get_info()
            } else {
                this.info.type = 0
                this.getList()
            }
        },
        getList () {
            this.$get(this.$api.adminPermissionGroups, { type: this.info.type }).then(res => {
                this.list = res.data.data
            })
        }

    },
    created () {
        this.onload()
    },
    mounted () {}
}
</script>
<style lang="scss" scoped>

</style>
