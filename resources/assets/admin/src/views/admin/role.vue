<template>
  <div class="app-container">
    <el-form ref="queryForm" :model="queryParams" :inline="true">
      <el-form-item label="角色名称" prop="name">
        <el-input
          v-model="queryParams.name"
          placeholder="请输入角色名称"
          clearable
          size="small"
          style="width: 240px"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="权限字符" prop="key">
        <el-input
          v-model="queryParams.key"
          placeholder="请输入权限字符"
          clearable
          size="small"
          style="width: 240px"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="状态" prop="status">
        <el-select
          v-model="queryParams.status"
          placeholder="角色状态"
          clearable
          size="small"
          style="width: 240px"
        >
          <el-option
            v-for="dict in statusOptions"
            :key="dict.value"
            :label="dict.label"
            :value="dict.value"
          />
        </el-select>
      </el-form-item>
      <!-- <el-form-item label="创建时间">
        <el-date-picker
          v-model="dateRange"
          size="small"
          style="width: 240px"
          value-format="yyyy-MM-dd"
          type="daterange"
          range-separator="-"
          start-placeholder="开始日期"
          end-placeholder="结束日期"
        />
      </el-form-item>-->
      <el-form-item>
        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
        <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
      </el-form-item>
    </el-form>

    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button
          type="primary"
          icon="el-icon-plus"
          size="mini"
          @click="handleAdd"
        >新增
        </el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button
          type="success"
          icon="el-icon-edit"
          size="mini"
          :disabled="single"
          @click="handleUpdate"
        >修改
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

    <el-table v-loading="loading" :data="roleList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center"/>
      <el-table-column label="角色编号" prop="id" width="120"/>
      <el-table-column label="角色名称" prop="name" :show-overflow-tooltip="true" width="150"/>
      <el-table-column label="权限字符" prop="key" :show-overflow-tooltip="true" width="150"/>
      <el-table-column label="显示顺序" prop="sort" width="100"/>
      <el-table-column label="状态" align="center" width="100">
        <template slot-scope="scope">
          <el-switch
            v-model="scope.row.status"
            :active-value="1"
            :inactive-value="0"
            @change="handleStatusChange(scope.row)"
          />
        </template>
      </el-table-column>
      <el-table-column label="创建时间" align="center" prop="createdAt" width="180">
        <template slot-scope="scope">
          <span>{{ parseTime(scope.row.created_at) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
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
      :v-show="total > 0"
      :total="total"
      :page.sync="queryParams.page"
      :limit.sync="queryParams.limit"
      @pagination="getList"
    />

    <!-- 添加或修改角色配置对话框 -->
    <el-dialog :title="title" :visible.sync="open" width="800px">
      <el-form ref="form" :model="form" :rules="rules" label-width="80px">
        <el-form-item label="角色名称" prop="name">
          <el-input v-model="form.name" placeholder="请输入角色名称" :disabled="isEdit"/>
        </el-form-item>
        <el-form-item label="权限字符" prop="key">
          <el-input v-model="form.key" placeholder="请输入权限字符" :disabled="isEdit"/>
        </el-form-item>
        <el-form-item label="角色顺序" prop="sort">
          <el-input-number v-model="form.sort" controls-position="right" :min="0"/>
        </el-form-item>
        <el-form-item label="状态">
          <el-radio-group v-model="form.status">
            <el-radio :label="1">正常</el-radio>
            <el-radio :label="0">停用</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="菜单权限">
          <el-tree
            ref="menu"
            :data="menuOptions"
            show-checkbox
            node-key="id"
            empty-text="加载中，请稍后"
            :props="defaultProps"
          />
        </el-form-item>
        <el-form-item label="接口权限">
          <el-row class="check_list">
            <el-col v-for="(v,k) in groupOptions" :key="k" class="item">
              <el-row class="permission_title">{{ v.name }}
                <el-checkbox
                  :checked="allchecked[k]"
                  :indeterminate="allindeterminate[k]"
                  style="float:right;"
                  :disabled="v.permissions.length === 0"
                  @change="oncheck_all($event,k)"
                >
                  全选
                </el-checkbox>
              </el-row>
              <el-row class="c-block">
                <div v-for="(vo,key) in v.permissions" :key="key" class="c-box">
                  <el-checkbox v-model="vo.checked" @change="itemCheck(k)">{{ vo.name }}</el-checkbox>
                </div>
                <div class="clear"></div>
              </el-row>
            </el-col>
          </el-row>
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="form.remark" type="textarea" placeholder="请输入内容"/>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="submitForm">确 定</el-button>
        <el-button @click="cancel">取 消</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {addRole, changeRoleStatus, dataScope, delRole, delRoles, getRole, listRole, updateRole} from '@/api/role'
import {listMenu} from '@/api/menu'
import {getGroupPermission} from '@/api/permission'
import {formatJson} from '@/utils'

export default {
  name: 'Role',
  data() {
    return {
      // 遮罩层
      loading: true,
      // 选中数组
      ids: [],
      // 非单个禁用
      single: true,
      // 非多个禁用
      multiple: true,
      // 总条数
      total: 0,
      // 角色表格数据
      roleList: [],
      // 弹出层标题
      title: '',
      // 是否显示弹出层
      open: false,
      // 是否显示弹出层（数据权限）
      openDataScope: false,
      isEdit: false,
      // 日期范围
      dateRange: [],
      // 状态数据字典
      statusOptions: [
        {label: '开启', value: 1},
        {label: '关闭', value: 0}
      ],
      // 数据范围选项
      dataScopeOptions: [
        {
          value: '1',
          label: '全部数据权限'
        },
        {
          value: '2',
          label: '自定数据权限'
        },
        {
          value: '3',
          label: '本部门数据权限'
        },
        {
          value: '4',
          label: '本部门及以下数据权限'
        },
        {
          value: '5',
          label: '仅本人数据权限'
        }
      ],
      // 菜单列表
      menuOptions: [],
      // 分组权限列表
      groupOptions: [],
      // 部门列表
      deptOptions: [],
      allchecked: [],
      allindeterminate: [],
      // 查询参数
      queryParams: {
        page: 1,
        limit: 15,
        name: undefined,
        key: undefined,
        status: undefined
      },
      // 表单参数
      form: {},
      defaultProps: {
        children: 'children',
        label: 'title'
      },
      // 表单校验
      rules: {
        name: [
          {required: true, message: '角色名称不能为空', trigger: 'blur'}
        ],
        key: [
          {required: true, message: '权限字符不能为空', trigger: 'blur'}
        ],
        sort: [
          {required: true, message: '角色顺序不能为空', trigger: 'blur'}
        ],
        remark: [
          {required: true, message: '备注不能为空', trigger: 'blur'}
        ]
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    /** 查询角色列表 */
    getList() {
      this.loading = true
      listRole(this.queryParams).then(
        response => {
          this.roleList = response.data.roles

          const {total_count, current_page, page_size} = response.data.paginator
          this.total = total_count
          this.queryParams.page = current_page
          this.queryParams.limit = page_size
          this.loading = false
        }
      )
    },
    /** 查询菜单树结构 */
    getMenuTreeselect() {
      listMenu().then(response => {
        this.menuOptions = response.data
      })
    },
    /** 查询权限结构 */
    getGroupPermissions() {
      getGroupPermission().then(response => {
        this.groupOptions = response.data
        this.groupOptions.forEach((items, ind) => {
          let checked_num = 0
          items.permissions.forEach(item => {
            if (this.form.permission_id.indexOf(item.id) > -1) {
              checked_num++
              this.$set(item, 'checked', true)
            } else {
              this.$set(item, 'checked', false)
            }
          })
          if (items.permissions.length > 0) {
            if (checked_num === 0) {
              this.allchecked[ind] = false
              this.allindeterminate[ind] = false
            } else if (checked_num === items.permissions.length) {
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
      })
    },
    // 所有菜单节点数据
    getMenuAllCheckedKeys() {
      // 目前被选中的菜单节点
      const checkedKeys = this.$refs.menu.getHalfCheckedKeys()
      // 半选中的菜单节点
      const halfCheckedKeys = this.$refs.menu.getCheckedKeys()
      checkedKeys.unshift.apply(checkedKeys, halfCheckedKeys)
      return checkedKeys
    },
    // 所有部门节点数据
    getDeptAllCheckedKeys() {
      // 目前被选中的部门节点
      const checkedKeys = this.$refs.dept.getCheckedKeys()
      // 半选中的部门节点
      // const halfCheckedKeys = this.$refs.dept.getCheckedKeys()
      // checkedKeys.unshift.apply(checkedKeys, halfCheckedKeys)
      return checkedKeys
    },
    /** 根据角色ID查询部门树结构 */
    getRoleDeptTreeselect(roleId) {
      // roleDeptTreeselect(roleId).then(response => {
      //   this.deptOptions = response.depts
      //   this.$nextTick(() => {
      //     this.$refs.dept.setCheckedKeys(response.checkedKeys)
      //   })
      // })
    },
    // 角色状态修改
    handleStatusChange(row) {
      const text = row.status === 1 ? '启用' : '停用'
      this.$confirm('确认要"' + text + '""' + row.name + '"角色吗?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function () {
        return changeRoleStatus(row.id, row.status)
      }).then(() => {
        // 提示信息
        this.$notify({
          title: '成功',
          message: text + '成功',
          type: 'success',
          duration: 2000
        })
      }).catch(function () {
        row.status = row.status === 0 ? 1 : 0
      })
    },
    // 取消按钮
    cancel() {
      this.open = false
      this.reset()
    },
    // 取消按钮（数据权限）
    cancelDataScope() {
      this.openDataScope = false
      this.reset()
    },
    // 表单重置
    reset() {
      if (this.$refs.menu !== undefined) {
        this.$refs.menu.setCheckedKeys([])
      }
      this.form = {
        id: undefined,
        name: undefined,
        key: undefined,
        sort: 0,
        status: 1,
        menu_ids: [],
        dept_ids: [],
        remark: undefined
      }
      this.groupOptions = [];
      this.resetForm('form')
    },
    /** 搜索按钮操作 */
    handleQuery() {
      this.getList()
    },
    /** 重置按钮操作 */
    resetQuery() {
      this.dateRange = []
      this.resetForm('queryForm')
      this.handleQuery()
    },
    // 多选框选中数据
    handleSelectionChange(selection) {
      this.ids = selection.map(item => item.id)
      this.single = selection.length !== 1
      this.multiple = !selection.length
    },
    /** 新增按钮操作 */
    handleAdd() {
      this.reset()
      this.getMenuTreeselect()
      this.getGroupPermissions()
      this.open = true
      this.title = '添加角色'
      this.isEdit = false
    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      this.reset()
      const roleId = row.id || this.ids
      getRole(roleId).then(response => {
        response.data.permission_id = []
        response.data.permissions.forEach(item => {
          response.data.permission_id.push(item.id)
        })
        this.form = response.data
        this.open = true
        this.title = '修改角色'
        this.isEdit = true

        this.getMenuTreeselect()
        this.getGroupPermissions()
        this.$nextTick(() => {
          // this.$refs.menu.setCheckedKeys(this.form.menu_ids)
          this.$refs.menu.setCheckedKeys(this.form.menu_ids)
        })
      })
    },
    /** 分配数据权限操作 */
    handleDataScope(row) {
      // this.reset()
      // getRole(row.roleId).then(response => {
      //   this.form = response.data
      //   this.openDataScope = true
      //   this.title = '分配数据权限'
      //   this.getRoleDeptTreeselect(row.roleId)
      // })
    },
    /** 提交按钮 */
    submitForm() {
      this.$refs['form'].validate(valid => {
        if (valid) {
          this.get_check_permission() // 获取选择的接口
          if (this.form.id !== undefined) {
            this.form.menu_ids = this.getMenuAllCheckedKeys()
            updateRole(this.form).then(response => {
              if (response.code === 0) {
                // 提示信息
                this.$notify({
                  title: '成功',
                  message: '修改成功',
                  type: 'success',
                  duration: 2000
                })
                this.open = false
                this.getList()
              }
            })
          } else {
            this.form.menu_ids = this.getMenuAllCheckedKeys()
            addRole(this.form).then(response => {
              if (response.code === 0) {
                // 提示信息
                this.$notify({
                  title: '成功',
                  message: '新增成功',
                  type: 'success',
                  duration: 2000
                })
                this.open = false
                this.getList()
              }
            })
          }
        }
      })
    },
    /** 提交按钮（数据权限） */
    submitDataScope: function () {
      if (this.form.roleId !== undefined) {
        this.form.deptIds = this.getDeptAllCheckedKeys()
        console.log(this.getDeptAllCheckedKeys())
        dataScope(this.form).then(response => {
          if (response.code === 200) {
            // 提示信息
            this.$notify({
              title: '成功',
              message: '修改成功',
              type: 'success',
              duration: 2000
            })
            this.openDataScope = false
            this.getList()
          } else {
            this.msgError(response.msg)
          }
        })
      }
    },
    /** 删除按钮操作 */
    handleDelete(row) {
      const roleIds = row.id || this.ids
      this.$confirm('是否确认删除角色编号为"' + roleIds + '"的数据项?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function () {
        if (isNaN(roleIds)) {
          return delRoles(roleIds)
        } else {
          return delRole(roleIds)
        }
      }).then(() => {
        this.getList()
        // 提示信息
        this.$notify({
          title: '成功',
          message: '删除成功',
          type: 'success',
          duration: 2000
        })
      }).catch(function () {
      })
    },
    /** 导出按钮操作 */
    handleExport() {
      this.$confirm('是否确认导出所有角色数据项?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.downloadLoading = true
        import('@/export/Export2Excel').then(excel => {
          const tHeader = ['角色编号', '角色名称', '权限字符', '显示顺序', '状态', '创建时间']
          const filterVal = ['id', 'name', 'key', 'sort', 'status', 'created_at']
          const list = this.roleList
          console.log(filterVal, list)
          const data = formatJson(filterVal, list)
          excel.export_json_to_excel({
            header: tHeader,
            data,
            filename: '角色管理',
            autoWidth: true, // Optional
            bookType: 'xlsx' // Optional
          })
          this.downloadLoading = false
        })
      })
    },
    itemCheck(ind) {
      let checked_num = 0
      this.groupOptions[ind].permissions.forEach(item => {
        if (item.checked) {
          checked_num++
        }
      })
      // eslint-disable-next-line eqeqeq
      if (checked_num === 0) {
        this.allchecked[ind] = false
        this.allindeterminate[ind] = false
      } else if (checked_num === this.groupOptions[ind].permissions.length) {
        this.allchecked[ind] = true
        this.allindeterminate[ind] = false
      } else {
        this.allindeterminate[ind] = true
        this.allchecked[ind] = false
      }
    },
    oncheck_all(e, i) {
      console.log(e)
      const checked = e
      this.allchecked[i] = checked
      this.allindeterminate[i] = false
      this.groupOptions[i].permissions.forEach(item => {
        if (this.$isEmpty(item.checked)) {
          this.$set(item, 'checked', checked)
        } else {
          item.checked = checked
        }
      })
    },
    get_check_permission() {
      const permission_id = [];
      this.groupOptions.forEach(items => {
        items.permissions.forEach(item => {
          if (!this.$isEmpty(item.checked) && item.checked) {
            permission_id.push(item.id)
          }
        })
      })
      if (permission_id.length > 0) {
        this.form.permissions = permission_id
      } else {
        this.form.permissions = undefined
      }
      this.form.permission_id = undefined
    }
  }
}
</script>

<style lang="scss" scoped>
.permission_title {
  background: #efefef;
  padding: 0 20px;
  box-sizing: border-box;
  border-radius: 4px;
}

.c-block {
  margin-top: 10px;
  margin-bottom: 10px;
}

.c-box {
  width: 33%;
  float: left;
}
</style>

