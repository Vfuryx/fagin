<template>
  <page-header-wrapper>
    <a-card :body-style="{padding: '24px 32px'}" :bordered="false">
      <div class="admin_table_page_title">
        <a-button @click="$router.back()" class=".ant-calendar-range-right" icon="arrow-left">返回</a-button>
      </div>
      <div class="unline underm"></div>
      <div class="admin_form">
        <a-form-model
          :model="info"
          :rules="rules"
          :confirmLoading="loading"
          ref="ruleForm"
        >
          <a-spin :spinning="loading">
            <a-form-model-item
              label="分组类型"
              prop="type"
              :labelCol="{lg: {span: 4}, sm: {span: 4}}"
              :wrapperCol="{lg: {span: 16}, sm: {span: 20} }">
              <a-select v-model="info.type" @change="typeChange">
                <a-select-option :value="0">
                  总后台
                </a-select-option>
                <a-select-option :value="1">
                  商家
                </a-select-option>
              </a-select>
            </a-form-model-item>
            <a-form-model-item
              label="角色名称"
              prop="name"
              :labelCol="{lg: {span: 4}, sm: {span: 4}}"
              :wrapperCol="{lg: {span: 16}, sm: {span: 20} }">
              <a-input v-model="info.name"></a-input>
            </a-form-model-item>
            <a-form-model-item
              label="关键词"
              prop="key"
              :labelCol="{lg: {span: 4}, sm: {span: 4}}"
              :wrapperCol="{lg: {span: 16}, sm: {span: 20} }">
              <a-input v-model="info.key" :disabled="info.id > 0"></a-input>
            </a-form-model-item>
            <a-form-model-item
              label="显示排序"
              :label-col="{lg: {span: 4}, sm: {span: 4}}"
              :wrapper-col="{lg: {span: 16}, sm: {span: 20}}"
              prop="sort"
              style="width: 100%;margin-right: 0">
              <a-input-number v-model="info.sort" controls-position="right" :min="0" />
            </a-form-model-item>
            <a-form-model-item
              label="菜单状态"
              :label-col="{lg: {span: 4}, sm: {span: 4}}"
              :wrapper-col="{lg: {span: 16}, sm: {span: 20}}"
              prop="status"
              style="width: 100%;margin-right: 0">
              <a-radio-group name="status" v-model="info.status">
                <a-radio :value="1">开启</a-radio>
                <a-radio :value="0">关闭</a-radio>
              </a-radio-group>
            </a-form-model-item>
            <a-form-model-item
              label="菜单权限"
              :labelCol="{lg: {span: 4}, sm: {span: 4}}"
              :wrapperCol="{lg: {span: 16}, sm: {span: 20} }">
              <a-tree-select
                :replace-fields="{children:'children', title:'title', key:'id', value:'id' }"
                :value="info.menu_ids"
                :tree-data="list"
                tree-checkable
                tree-check-strictly
                show-checked-strategy="SHOW_ALL"
                @change="treeChange"></a-tree-select>
            </a-form-model-item>
            <a-form-model-item
              label="接口权限"
              :labelCol="{lg: {span: 4}, sm: {span: 4}}"
              :wrapperCol="{lg: {span: 16}, sm: {span: 20} }">
              <div class="check_list">
                <div class="item" v-for="(v,k) in permissions" :key="k">
                  <div class="permission_title">{{ v.name }}
                    <a-checkbox
                      :checked="allchecked[k]"
                      :indeterminate="allindeterminate[k]"
                      style="float:right;"
                      :disabled="v.permissions.length === 0"
                      @change="oncheck_all($event,k)">
                      全选
                    </a-checkbox>
                    <!-- <a-tag @click="check_all(k)" style="float:right;margin-top:10px;">点击全选</a-tag> -->
                  </div>
                  <div class="cblock clearfix">
                    <div class="cbox" v-for="(vo,key) in v.permissions" :key="key">
                      <a-checkbox v-model="vo.checked" @change="itemCheck(k)">{{ vo.name }}</a-checkbox>
                      <a-tag color="red">{{ vo.method }}</a-tag>
                    </div>
                    <div class="clear"></div>
                  </div>
                </div>
              </div>
            </a-form-model-item>
            <a-form-model-item
              label="备注"
              prop="remark"
              :labelCol="{lg: {span: 4}, sm: {span: 4}}"
              :wrapperCol="{lg: {span: 16}, sm: {span: 20} }">
              <a-input v-model="info.remark"></a-input>
            </a-form-model-item>
            <a-form-model-item :wrapper-col="{ span: 12, offset: 5 }">
              <a-button type="primary" @click="handleSubmit">提交</a-button>
            </a-form-model-item>
          </a-spin>
        </a-form-model>
      </div>
    </a-card>
  </page-header-wrapper>
</template>

<script>
import Vue from 'vue'
import { FormModel, TreeSelect, InputNumber } from 'ant-design-vue'
import store from '@/store'

Vue.use(FormModel)
Vue.use(TreeSelect)
Vue.use(InputNumber)
export default {
  components: {},
  props: {},
  data () {
    return {
      info: {
        type: 0,
        name: '',
        status: 1,
        sort: 0,
        key: '',
        remark: '',
        menu_ids: [],
        permission_ids: []
      },
      menu_id_copy: [],
      child_list: [],
      parent_list: [],
      id: 0,
      list: [],
      permissions: [],
      allchecked: [],
      allindeterminate: [],
      old_menu_ids: [[], []],
      loading: true,
      rules: {
        id: [
          { required: true, message: '不能空', trigger: 'blur' },
          { type: 'integer', min: 0, message: '不正确', trigger: 'blur' }
        ],
        type: [
          { required: true, message: '不能空', trigger: 'blur' },
          { type: 'enum', enum: [0, 1], message: '类型不正确', trigger: 'blur' }
        ],
        name: [
          { required: true, message: '不能为空', trigger: 'blur' },
          { min: 1, max: 32, message: '范围 1 - 32 字符', trigger: 'blur' }
        ],
        key: [
          { required: true, message: '不能为空', trigger: 'blur' },
          { min: 1, max: 32, message: '范围 1 - 32 字符', trigger: 'blur' },
          { validator: this.validateKey, trigger: 'blur' }
        ],
        remark: [
          { min: 1, max: 255, message: '范围 1 - 255 字符', trigger: 'blur' }
        ],
        sort: [
          { required: true, message: '不能空', trigger: 'blur' },
          { type: 'integer', min: 0, message: '上级不正确', trigger: 'blur' }
        ],
        status: [
          { required: true, message: '不能空', trigger: 'blur' },
          { type: 'enum', enum: [0, 1], message: '类型不正确', trigger: 'blur' }
        ]
      }
    }
  },
  watch: {},
  computed: {},
  methods: {
    setChecked (item) {
      if (item.permissions.length > 0) {
        for (let i = 0; i < item.permissions.length; i++) {
          if (!item.permissions[i].checked) {
            return false
          }
        }
        return true
      } else {
        return false
      }
    },
    itemCheck (ind) {
      let checkedNum = 0
      this.permissions[ind].permissions.forEach(item => {
        if (item.checked) {
          checkedNum++
        }
      })
      if (checkedNum === 0) {
        this.allchecked[ind] = false
        this.allindeterminate[ind] = false
      } else if (checkedNum === this.permissions[ind].permissions.length) {
        this.allchecked[ind] = true
        this.allindeterminate[ind] = false
      } else {
        this.allindeterminate[ind] = true
        this.allchecked[ind] = false
      }
    },
    oncheck_all (e, i) {
      const checked = e.target.checked
      this.allchecked[i] = checked
      this.allindeterminate[i] = false
      this.permissions[i].permissions.forEach(item => {
        if (this.$isEmpty(item.checked)) {
          this.$set(item, 'checked', checked)
        } else {
          item.checked = checked
        }
      })
    },
    handleSubmit () {
      this.$refs.ruleForm.validate(valid => {
        if (valid) {
          // 验证代码处
          if (this.$isEmpty(this.info.name)) {
            return this.$message.error('角色名称不能为空')
          }
          this.get_check_permission() // 获取选择的接口

          const info = JSON.parse(JSON.stringify(this.info))
          const menuIDs = []
          info.menu_ids.forEach(val => {
            menuIDs.push(val.value)
          })
          info.menu_ids = menuIDs

          if (info.id > 0) {
            info.menus = undefined
            info.permissions = undefined
            store
              .dispatch('updateRole', info)
              .then(response => {
                if (response.code !== 0) {
                  this.$notification.error({
                    message: '错误',
                    description: response.message
                  })
                  return
                }
                this.$message.success('保存成功')
                return this.$router.back()
              }).catch(() => {
              this.$notification.error({
                message: '错误',
                description: '请求信息失败，请重试'
              })
            })
          } else {
            store
              .dispatch('addRole', info)
              .then(response => {
                if (response.code !== 0) {
                  this.$notification.error({
                    message: '错误',
                    description: response.message
                  })
                  return
                }
                this.$message.success('保存成功')
                return this.$router.back()
              }).catch(() => {
              this.$notification.error({
                message: '错误',
                description: '请求信息失败，请重试'
              })
            })
          }
        } else {
          this.$message.warning('信息输入错误')
        }
      })
    },
    get_info () {
      store
        .dispatch('showRole', this.id)
        .then(response => {
          if (response.code !== 0) {
            this.$notification.error({
              message: '错误',
              description: response.message
            })
          }

          response.data.permission_ids = []
          response.data.permissions.forEach(item => {
            response.data.permission_ids.push(item.id)
          })
          this.menu_id_copy = response.data.menu_ids
          response.data.menu_ids = []
          this.info = response.data
          this.old_menu_ids[this.info.type] = this.info.menu_ids
          store
            .dispatch('getMenuAll', { type: this.info.type })
            .then(response => {
              if (response.code !== 0) {
                this.$notification.error({
                  message: '错误',
                  description: response.message
                })
              }
              response.data.forEach(val => {
                val.parents = []
                const obj = {
                  value: val.id,
                  lable: val.title
                }
                val = this.setParents(val, [obj])
              })
              this.list = response.data
              this.setMenuID(this.list)
            }).catch(() => {
            this.$notification.error({
              message: '错误',
              description: '请求信息失败，请重试'
            })
          })

          store
            .dispatch('getGroupPermission', { type: this.info.type })
            .then(response => {
              if (response.code !== 0) {
                this.$notification.error({
                  message: '错误',
                  description: response.message
                })
              }
              this.permissions = response.data
              this.permissions.forEach((items, ind) => {
                let checkedNum = 0
                items.permissions.forEach(item => {
                  if (this.info.permission_ids.indexOf(item.id) > -1) {
                    checkedNum++
                    this.$set(item, 'checked', true)
                  } else {
                    this.$set(item, 'checked', false)
                  }
                })
                if (items.permissions.length > 0) {
                  if (checkedNum === 0) {
                    this.allchecked[ind] = false
                    this.allindeterminate[ind] = false
                  } else if (checkedNum === items.permissions.length) {
                    this.allchecked[ind] = true
                    this.allindeterminate[ind] = false
                  } else {
                    this.allindeterminate[ind] = true
                    this.allchecked[ind] = false
                  }
                } else {
                  this.allchecked[ind] = false
                  this.allindeterminate[ind] = false
                }
              })
              console.log(this.allchecked)
              console.log(this.allindeterminate)
              this.$forceUpdate()
            }).catch(() => {
            this.$notification.error({
              message: '错误',
              description: '请求信息失败，请重试'
            })
          })
        }).catch(() => {
        this.$notification.error({
          message: '错误',
          description: '请求信息失败，请重试'
        })
      }).finally(() => {
        this.loading = false
      })
    },
    treeChange (value, label, extra) {
      const { checked, triggerNode } = extra
      let { dataRef } = triggerNode
      if (checked === undefined) {
        dataRef = triggerNode.data.props.dataRef
      }
      if (checked) {
        this.child_list = []
        this.parent_list = []
        this.getChildren(dataRef.children)
        this.getParents(this.list, dataRef.id)
        const parents = this.parent_list
        const values = [...parents, { value: dataRef.id, label: dataRef.title }, ...this.child_list]
        if (values) {
          for (let i = 0; i < values.length; i++) {
            const findInd = this.info.menu_ids.findIndex(item => {
              return item.value === values[i].value
            })
            if (findInd === -1) {
              this.info.menu_ids.push(values[i])
            }
          }
        }
      } else {
        const findInd = this.info.menu_ids.findIndex(item => {
          return item.value === dataRef.id
        })
        if (findInd > -1) {
          this.info.menu_ids.splice(findInd, 1)
          this.removeChildren(dataRef.children)
        }
      }
      this.old_menu_ids[this.info.type] = this.info.menu_ids
    },
    typeChange () {
      this.info.menu_ids = this.old_menu_ids[this.info.type]
      store
        .dispatch('getMenuAll', { type: this.info.type })
        .then(response => {
          if (response.code !== 0) {
            this.$notification.error({
              message: '错误',
              description: response.message
            })
          }
          response.data.forEach(val => {
            val.parents = []
            const obj = {
              value: val.id,
              lable: val.title
            }
            val = this.setParents(val, [obj])
          })
          this.list = response.data
        }).catch(() => {
        this.$notification.error({
          message: '错误',
          description: '请求信息失败，请重试'
        })
      })

      store
        .dispatch('getGroupPermission', { type: this.info.type })
        .then(response => {
          if (response.code !== 0) {
            this.$notification.error({
              message: '错误',
              description: response.message
            })
          }

          this.permissions = response.data
          this.permissions.forEach((items, ind) => {
            let checkedNum = 0
            items.permissions.forEach(item => {
              if (this.info.permission_ids.indexOf(item.id) > -1) {
                checkedNum++
                this.$set(item, 'checked', true)
              } else {
                this.$set(item, 'checked', false)
              }
            })
            if (items.permissions.length > 0) {
              if (checkedNum === 0) {
                this.allchecked[ind] = false
                this.allindeterminate[ind] = false
              } else if (checkedNum === items.permissions.length) {
                this.allchecked[ind] = true
                this.allindeterminate[ind] = false
              } else {
                console.log('haha')
                this.allindeterminate[ind] = true
                this.allchecked[ind] = false
              }
            } else {
              this.allchecked[ind] = false
              this.allindeterminate[ind] = false
            }
          })
          this.$forceUpdate()
        }).catch(() => {
        this.$notification.error({
          message: '错误',
          description: '请求信息失败，请重试'
        })
      })
    },
    // 获取菜单列表
    onload () {
      // 判断你是否是编辑
      if (!this.$isEmpty(this.$route.params.id)) {
        this.id = this.$route.params.id
        this.get_info()
      } else {
        this.info.type = 0
        store
          .dispatch('getMenuAll', { type: this.info.type })
          .then(response => {
            if (response.code !== 0) {
              this.$notification.error({
                message: '错误',
                description: response.message
              })
            }
            response.data.forEach(val => {
              val.parents = []
              const obj = {
                value: val.id,
                lable: val.title
              }
              val = this.setParents(val, [obj])
            })
            this.list = response.data
          }).catch(() => {
          this.$notification.error({
            message: '错误',
            description: '请求信息失败，请重试'
          })
        }).finally(() => {
          this.loading = false
        })
        store
          .dispatch('getGroupPermission', { type: this.info.type })
          .then(response => {
            if (response.code !== 0) {
              this.$notification.error({
                message: '错误',
                description: response.message
              })
            }
            this.permissions = response.data
          }).catch(() => {
          this.$notification.error({
            message: '错误',
            description: '请求信息失败，请重试'
          })
        }).finally(() => {
          this.loading = false
        })
      }
    },
    check_all (e) {
      this.permissions[e].permissions.forEach(item => {
        if (this.$isEmpty(item.checked)) {
          this.$set(item, 'checked', true)
        } else {
          item.checked = !item.checked
        }
      })
    },
    get_check_permission () {
      const permissionID = []
      this.permissions.forEach(items => {
        items.permissions.forEach(item => {
          if (!this.$isEmpty(item.checked) && item.checked) {
            permissionID.push(item.id)
          }
        })
      })
      if (permissionID.length > 0) {
        this.info.permission_ids = permissionID
      } else {
        this.info.permission_ids = []
      }
    },
    setMenuID (data) {
      data = data || []
      for (let i = 0; i < data.length; i++) {
        const findInd = this.menu_id_copy.findIndex(item => {
          return item === data[i].id
        })
        if (findInd > -1) {
          this.info.menu_ids.push({ value: data[i].id, label: data[i].title })
        }
        if (data[i].children) {
          this.setMenuID(data[i].children)
        }
      }
    },
    setParents (data, ids) {
      if (data.children && data.children.length > 0) {
        data.children.forEach(val => {
          val.parents = []
          val.parents.push(...ids)
          const obj = {
            value: val.id,
            lable: val.title
          }
          val = this.setParents(val, [obj, ...ids])
        })
      }
      return data
    },
    getParents (data, id) {
      data = data || []
      var res = data.filter(item => {
        return item.id === id
      })
      if (res.length === 0) {
        for (let i = 0; i < data.length; i++) {
          this.getParents(data[i].children, id)
        }
      } else {
        this.parent_list = res[0].parents
      }
    },
    getChildren (data) {
      if (!data) {
        return
      }
      for (let i = 0; i < data.length; i++) {
        this.child_list.push({ value: data[i].id, label: data[i].title })
        if (data[i].children) {
          this.getChildren(data[i].children)
        }
      }
    },
    removeChildren (data) {
      if (!data) {
        return
      }
      for (let i = 0; i < data.length; i++) {
        const findInd = this.info.menu_ids.findIndex(item => {
          return item.value === data[i].id
        })
        if (findInd > -1) {
          this.info.menu_ids.splice(findInd, 1)
        }
        if (data[i].children) {
          this.removeChildren(data[i].children)
        }
      }
    },
    validateKey (rule, value, callback, source, options) {
      if (this.info.id > 0) { // 梗系则不需要验证
        return callback()
      }
      store
        .dispatch('roleKeyExist', { 'key': value })
        .then(response => {
          if (response.code !== 0) {
            return callback(new Error('关键词已存在'))
          }
          return callback()
        }).catch(() => {
        return callback(new Error('关键词已存在'))
      })
    }
  },
  created () {
    this.onload()
  },
  mounted () {
  }
}
</script>
<style lang="less" scoped>
.permission_title {
  background: #efefef;
  padding: 0 20px;
  box-sizing: border-box;
  border-radius: 4px;
}

.cblock {
  margin-top: 10px;
  margin-bottom: 10px;
}

.cbox {
  width: 25%;
  float: left;
}

/* 线条 */
.unline {
  border-bottom: 1px solid #EBEEF5;
}

.unline.underm {
  margin-bottom: 20px;
  margin-top: 20px;
}
</style>
