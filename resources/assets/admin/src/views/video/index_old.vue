<template>
  <div class="container">
    <div class="filter-container clearfix">
      <el-button
        class="filter-item"
        style="margin-left: 10px;"
        type="primary"
        icon="el-icon-edit"
        @click="handleCreate"
      >添加</el-button>
    </div>
    <el-table
      v-loading="listLoading"
      :data="list"
      border
      fit
      highlight-current-row
      style="width: 100%"
    >
      <el-table-column align="center" label="ID" width="80">
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>

      <el-table-column min-width="120px" label="标题">
        <template slot-scope="scope">
          <span>{{ scope.row.title }}</span>
        </template>
      </el-table-column>

      <el-table-column align="center" label="状态" width="80">
        <template slot-scope="scope">
          <el-tag>{{ scope.row.status ? '开启' : '关闭' }}</el-tag>
        </template>
      </el-table-column>

      <el-table-column align="center" label="操作" width="250">
        <template slot-scope="scope">
          <el-button size="mini" type="success" @click="handleShowVideo(scope.row)">查看</el-button>
          <el-button size="mini" type="primary" @click="handleUpdateVideo(scope.row)">编辑</el-button>
          <el-button size="mini" type="danger" @click="handleDelete(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <pagination
      v-show="total >= 1"
      :total="total"
      :page.sync="listQuery.page"
      :limit.sync="listQuery.limit"
      @pagination="getData"
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
            <i class="el-icon-upload"></i>
            <div class="el-upload__text">
              将文件拖到此处，或
              <em>点击上传</em>
            </div>
            <div class="el-upload__tip" slot="tip">只能上传video/mp4文件，且不超过32mb</div>
          </el-upload>
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input type="textarea" v-model="temp.description" />
        </el-form-item>
        <el-form-item label="状态">
          <el-switch v-model="temp.status" :active-value="1" :inactive-value="0" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取消</el-button>
        <el-button type="primary" @click="createData()">提交</el-button>
        <!-- <el-button type="primary" @click="dialogStatus==='create'?createData():updateData()">提交</el-button> -->
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
        class="video-player-box"
        ref="videoPlayer"
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
            <div class="el-upload__tip" slot="tip">只能上传video/mp4文件，且不超过32mb</div>
          </el-upload>
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input type="textarea" v-model="temp.description" />
        </el-form-item>
        <el-form-item label="状态">
          <el-switch v-model="temp.status" :active-value="1" :inactive-value="0" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogUpdateVisible = false">取消</el-button>
        <el-button type="primary" @click="UpdateData()">提交</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { getToken } from '@/utils/auth'
import { getList, deleteVideo, createVideo, updateVideo } from '@/api/video'
import Pagination from '@/components/Pagination' // Secondary package based on el-pagination
import 'video.js/dist/video-js.css'
import { videoPlayer } from 'vue-video-player'

export default {
  name: 'Index',
  components: { Pagination, videoPlayer },
  data() {
    return {
      uploadAction: process.env.VUE_APP_BASE_API + '/v1/video/upload',
      playAction: process.env.VUE_APP_BASE_API + '/play/av',
      headers: {
        'Authorization': `Bearer ${getToken()}`
      },
      list: [
        {}
      ],
      listLoading: false,
      dialogFormVisible: false,
      dialogShowVisible: false,
      dialogUpdateVisible: false,
      temp: {
        id: 0,
        title: '',
        status: 1,
        path: '',
        description: ''
      },
      total: 0,
      listQuery: {
        page: 1,
        limit: 15,
        importance: undefined,
        title: undefined,
        type: undefined,
        sort: '+id'
      },
      rules: {
        // type: [{ required: true, message: 'type is required', trigger: 'change' }],
        // timestamp: [{ type: 'date', required: true, message: 'timestamp is required', trigger: 'change' }],
        // title: [{ required: true, message: 'title is required', trigger: 'blur' }]
      },
      playerOptions: {}
    }
  },
  created() {
    this.getData()
  },
  methods: {
    handleAvatarSuccess(res, file) {
      this.temp.path = res.data.path
    },
    getData() {
      this.listLoading = true
      getList(this.listQuery).then(response => {
        this.list = response.data.videos
        const { total_count, current_page, page_size } = response.data.paginator
        this.total = total_count
        this.listQuery.page = current_page
        this.listQuery.limit = page_size
        setTimeout(() => {
          this.listLoading = false
        }, 1.5 * 1000)
      })
    },
    handleDelete(row) {
      deleteVideo(row.id).then(response => {
        this.$notify({
          title: '成功',
          message: '删除成功',
          type: 'success',
          duration: 2000
        })
        const index = this.list.indexOf(row)
        this.list.splice(index, 1)
      })
    },
    handleCreate() {
      this.resetTemp()
      this.dialogFormVisible = true
    },
    createData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          createVideo(this.temp).then(() => {
            this.dialogFormVisible = false
            this.$notify({
              title: '成功',
              message: '创建成功',
              type: 'success',
              duration: 2000
            })
            this.getData()
          })
        }
      })
    },
    UpdateData() {
      this.$refs['updateForm'].validate((valid) => {
        if (valid) {
          updateVideo(this.temp.id, this.temp).then(() => {
            this.dialogUpdateVisible = false
            this.$notify({
              title: '成功',
              message: '创建成功',
              type: 'success',
              duration: 2000
            })
            this.getData()
          })
        }
      })
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
    resetTemp() {
      this.temp = {
        title: '',
        status: 1,
        path: '',
        description: ''
      }
    },
    showDialogClose(done) {
      // 关闭视频
      this.$refs.videoPlayer.player.pause()
      done()
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
</style>
