<template>
  <div class="qingwu">
    <div class="admin_table_page_title">
      <a-button @click="$router.back()" class="float_right" icon="arrow-left">返回</a-button>
      接口分组编辑
    </div>
    <div class="unline underm"></div>
    <div class="admin_form">
      <a-form-model :label-col="{ span: 5 }" :wrapper-col="{ span: 12 }">
        <a-form-model-item label="分组类型">
          <a-select v-model="info.type">
            <a-select-option :value="0">
              总后台
            </a-select-option>
            <a-select-option :value="1">
              商家
            </a-select-option>
          </a-select>
        </a-form-model-item>
        <a-form-model-item label="接口分组">
          <a-input v-model="info.name"></a-input>
        </a-form-model-item>
        <a-form-model-item label="分组排序">
          <a-input-number v-model="info.sort" :min="0" :formatter="limitNumber" :parser="limitNumber" @blur="blurNumber" />
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
                name: '',
                sort: 0
            },
            id: 0
        }
    },
    watch: {},
    computed: {},
    methods: {
        limitNumber (value) {
            if (typeof value === 'string') {
                return !isNaN(Number(value)) ? value.replace(/\./g, '') : 0
            } else if (typeof value === 'number') {
                return !isNaN(value) ? String(value).replace(/\./g, '') : 0
            } else {
                return 0
            }
        },
        blurNumber (e) {
            if (this.info.sort !== 0 && !this.info.sort) {
                this.info.sort = 0
            }
        },
        handleSubmit () {
            // 验证代码处
            if (this.$isEmpty(this.info.name)) {
                return this.$message.error('分组名不能为空')
            }

            const api = this.$apiHandle(this.$api.adminPermissionGroups, this.id)
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
            this.$get(this.$api.adminPermissionGroups + '/' + this.id).then(res => {
                this.info = res.data
            })
        },
        // 获取列表
        onload () {
            // 判断你是否是编辑
            if (!this.$isEmpty(this.$route.params.id)) {
                this.id = this.$route.params.id
                this.get_info()
            }
        }

    },
    created () {
        this.onload()
    },
    mounted () {
    }
}
</script>
<style lang="scss" scoped>

</style>
