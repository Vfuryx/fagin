<template>
  <div class="app-container">
    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button type="primary" icon="el-icon-plus" size="mini" @click="handleCreate">新增</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button
          type="danger"
          icon="el-icon-delete"
          size="mini"
          :disabled="multiple"
          @click="handleDelete"
        >删除</el-button>
      </el-col>
      <!-- <el-col :span="1.5">
        <el-button
          type="warning"
          icon="el-icon-download"
          size="mini"
          @click="handleExport"
        >导出</el-button>
      </el-col>-->
    </el-row>

    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="拼命加载中"
      fit
      highlight-current-row
      @selection-change="handleSelectionChange"
    >
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column align="center" label="ID" width="60">
        <template slot-scope="scope">{{ scope.row.id }}</template>
      </el-table-column>
      <el-table-column align="center" label="标题">
        <template slot-scope="scope">{{ scope.row.title }}</template>
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
            @click="handleShowVideo(scope.row)"
          >查看</el-button>
          <el-button
            size="mini"
            type="text"
            icon="el-icon-edit"
            @click="handleUpdateVideo(scope.row)"
          >修改</el-button>
          <el-button
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

    <el-dialog title="创建视频" :visible.sync="dialogFormVisible">
      <el-form
        ref="dataForm"
        :rules="rules"
        :model="temp"
        label-position="left"
        label-width="70px"
        style="width: 400px; margin-left:50px;"
      >
        <el-form-item label="标题" prop="title">
          <el-input v-model="temp.title" />
        </el-form-item>
        <el-form-item label="视频上传" prop="path">
          <!-- <Upload v-model="temp.path" /> -->
          <el-upload
            class="upload-demo"
            :on-success="handleAvatarSuccess"
            drag
            :action="uploadAction"
            :multiple="false"
            :headers="headers"
          >
            <i class="el-icon-upload" />
            <div class="el-upload__text">
              将文件拖到此处，或
              <em>点击上传</em>
            </div>
            <div slot="tip" class="el-upload__tip">只能上传video/mp4文件，且不超过32mb</div>
          </el-upload>
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input v-model="temp.description" type="textarea" />
        </el-form-item>
        <el-form-item label="状态">
          <el-switch v-model="temp.status" :active-value="1" :inactive-value="0" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取消</el-button>
        <el-button type="primary" @click="submitCreateForm()">提交</el-button>
      </div>
    </el-dialog>

    <el-dialog
      title="播放视频"
      :visible.sync="dialogShowVisible"
      destroy-on-close
      :before-close="showDialogClose"
    >
      <h1>{{ temp.title }}</h1>
      <video-player
        ref="videoPlayer"
        class="video-player-box"
        :options="playerOptions"
        :playsinline="true"
      />
    </el-dialog>

    <el-dialog title="更新视频" :visible.sync="dialogUpdateVisible">
      <el-form
        ref="updateForm"
        :rules="rules"
        :model="temp"
        label-position="left"
        label-width="70px"
        style="width: 400px; margin-left:50px;"
      >
        <el-form-item label="标题" prop="title">
          <el-input v-model="temp.title" />
        </el-form-item>
        <el-form-item label="视频上传" prop="path">
          <el-upload
            class="upload-demo"
            :on-success="handleAvatarSuccess"
            drag
            :action="uploadAction"
            :multiple="false"
            :headers="headers"
          >
            <i class="el-icon-upload" />
            <div class="el-upload__text">
              将文件拖到此处，或
              <em>点击上传</em>
            </div>
            <div slot="tip" class="el-upload__tip">只能上传video/mp4文件，且不超过32mb</div>
          </el-upload>
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input v-model="temp.description" type="textarea" />
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
import { getList, deleteVideo, createVideo, updateVideo, deleteVideos } from '@/api/video'
import Pagination from '@/components/Pagination' // Secondary package based on el-pagination
import 'video.js/dist/video-js.css'
import { videoPlayer } from 'vue-video-player'
import { getToken } from '@/utils/auth'

export default {
  name: 'Index',
  components: {
    Pagination, videoPlayer
  },
  // eslint-disable-next-line vue/order-in-components
  data() {
    return {
      baseApi: process.env.VUE_APP_BASE_API,
      baseURL: process.env.VUE_APP_SERVER_PUBLIC_PATH,
      uploadAction: process.env.VUE_APP_BASE_API + '/v1/video/upload',
      playAction: process.env.VUE_APP_BASE_API + '/play/av',
      headers: {
        'Authorization': `Bearer ${getToken()}`
      },
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
        status: 1,
        path: '',
        description: ''
      },
      // 选中数组
      ids: [],
      // 非单个禁用
      single: true,
      // 非多个禁用
      multiple: true,
      rules: {},
      playerOptions: {}
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.listLoading = true
      getList(this.listQuery).then(response => {
        this.list = response.data.videos
        const { total_count, current_page, page_size } = response.data.paginator
        this.total = total_count
        this.listQuery.page = current_page
        this.listQuery.limit = page_size
        this.listLoading = false
      })
      this.listLoading = false
    },
    resetTemp() {
      this.temp = {
        title: '',
        status: 1,
        path: '',
        description: ''
      }
    },
    handleCreate() {
      this.resetTemp()
      this.dialogFormVisible = true
    },
    submitCreateForm() {
      this.$refs['dataForm'].validate(valid => {
        if (valid) {
          createVideo(this.temp).then(res => {
            this.dialogFormVisible = false
            this.$notify({
              title: '成功',
              message: '创建成功',
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
        status: row.status,
        path: row.path,
        description: row.description
      }
    },
    submitUpdateForm(id) {
      this.$refs['updateForm'].validate(valid => {
        if (valid) {
          updateVideo(this.temp.id, this.temp).then(res => {
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
    showDialogClose(done) {
      // 关闭视频
      this.$refs.videoPlayer.player.pause()
      done()
    },
    handleAvatarSuccess(res, file) {
      this.temp.path = res.data.path
    },
    handleShowVideo(row) {
      this.dialogShowVisible = true
      this.temp = {
        id: row.id,
        title: row.title,
        status: row.status,
        path: row.id,
        description: row.description
      }
      // 配置视频参数
      this.playerOptions = {
        playbackRates: [0.7, 1.0, 1.5, 2.0],
        autoplay: false,
        muted: false, // 默认情况下将会消除任何音频。
        loop: false, // 导致视频一结束就重新开始。
        preload: 'auto', // 建议浏览器在<video>加载元素后是否应该开始下载视频数据。auto浏览器选择最佳行为,立即开始加载视频（如果浏览器支持）
        language: 'zh-CN',
        aspectRatio: '16:9', // 将播放器置于流畅模式，并在计算播放器的动态大小时使用该值。值应该代表一个比例 - 用冒号分隔的两个数字（例如"16:9"或"4:3"）
        fluid: true, // 当true时，Video.js player将拥有流体大小。换句话说，它将按比例缩放以适应其容器。
        sources: [{
        }, {
          src: this.playAction + this.temp.id,
          type: 'video/mp4'
        }],
        poster: '',
        // width: document.documentElement.clientWidth,
        notSupportedMessage: '此视频暂无法播放，请稍后再试',
        controlBar: {
          timeDivider: true,
          durationDisplay: true,
          remainingTimeDisplay: false,
          fullscreenToggle: true
        }
      }
    },
    /** 删除按钮操作 */
    handleDelete(row) {
      const roleIds = row.id || this.ids
      this.$confirm('是否确认删除角色编号为"' + roleIds + '"的数据项?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function() {
        if (isNaN(roleIds)) {
          return deleteVideos(roleIds)
        } else {
          return deleteVideo(roleIds)
        }
      }).then(() => {
        this.fetchData()
        // 提示信息
        this.$notify({
          title: '成功',
          message: '删除成功',
          type: 'success',
          duration: 2000
        })
      }).catch(function() { })
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
