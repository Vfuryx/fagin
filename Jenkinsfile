pipeline {
  agent any
  environment { // 定义环境变量 一般服务器配置的环境变量都用不了 需要重新在下方定义
    GOPROXY     = 'https://goproxy.cn,direct'
    GO111MODULE = 'on'
    GOROOT = '/usr/local/go'
    GOPATH = '/home/mygo'
    PATH = '/sbin:/bin:/usr/sbin:/usr/bin:/usr/local/go/bin'
  }
  stages {
    stage('clone') {
      steps {
        sh 'echo env.BRANCH_NAME'
        sh 'echo 显示清空的文件'
        sh 'git clean -xdne .env'
        sh 'echo 清空的文件 排除 .env 文件'
        sh 'git clean -xdfe .env'
        sh 'echo 复制 .env 文件，重命名为 .env.yml'
        sh 'sudo cp .env .env.yml'
      }
    }
    stage('install') {
      steps {
        sh 'echo build'
        sh 'go env'
        //sh 'go get -u'                         // 更新包
        sh 'go mod tidy'
        sh 'go mod vendor'
        sh 'go run console/main.go migrate'    // 数据库迁移
      }
    }
    stage('build') {
      parallel {
        stage('build-dev') {
          when {
            branch 'dev'
          }
          steps {
            sh 'echo build-dev'
            sh 'go build -o fagin-dev ./main.go'
          }
        }
        stage('build-master') {
          when {
            branch 'master'
          }
          steps {
            sh 'echo build-master'
            sh 'sudo go build -o fagin-master ./main.go'
          }
        }
      }
    }
    stage('test') {
      steps {
        sh 'echo dep'
      }
    }
    stage('Deploy for production') {
      when {
        branch 'master'
      }
      steps {
        emailext(
            subject: "${env.JOB_NAME}-${env.BUILD_NUMBER}任务-请确认是否部署到生产服务器",
            mimeType: "text/html",
            body: """
            		<p>请测试组长确认，可以进行上线</p>
            		<p>部署用户： ${env.Stage_Submitter}</p>

        		    <p>页面链接请点击： <a href='${env.RUN_DISPLAY_URL}'>${env.JOB_NAME}</a></p>
                    <p>FAILED: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]':</p>
                    <p>Check console output at &QUOT;<a href='${env.BUILD_URL}'>${env.JOB_NAME} [${env.BUILD_NUMBER}]</a>&QUOT;</p>
                 """,
            to: "935661686@qq.com",
          )
        input message: 'Finished using the web site? (Click "Proceed" to continue)'
        sh 'echo master'
        sh 'sudo sh admin.sh restart fagin-master'
        sh 'sudo sh ./admin.sh status fagin-master'
      }
    }
    stage('Deliver for development') {
       when {
            branch 'dev'
       }
       steps {
        sh 'echo dev'
        sh 'sudo sh admin.sh restart fagin-dev'
        sh 'sudo sh ./admin.sh status fagin-dev'
      }
    }
  }
  post {
    success {
      emailext(
        subject: "${env.JOB_NAME}-${env.BUILD_NUMBER}-最新的环境部署成功",
        mimeType: "text/html",
        body: """
        <table width="95%" cellpadding="0" cellspacing="0"
            style="font-size: 11pt; font-family: Tahoma, Arial, Helvetica, sans-serif">
            <tr>
                <td><br />
                <b><font color="#0B610B">构建信息</font></b>
                <hr size="2" width="100%" align="center" /></td>
            </tr>
            <tr>
                <td>
                    <ul>
                        <li>构建名称：${JOB_NAME}</li>
                        <li>构建编号：${BUILD_NUMBER}  </li>
                        <li>GIT 地址：${git_url}</li>
                        <li>GIT 分支：${git_branch}</li>
                    </ul>
                </td>
            </tr>
            <tr>
                <td><b><font color="#0B610B">构建日志 :</font></b>
        		<p>页面链接请点击： <a href='${env.RUN_DISPLAY_URL}'>${env.JOB_NAME}</a></p>
            </tr>

        </table>
        """,
        to: "935661686@qq.com",
      )
    }
    failure {
      emailext(
        subject: "${env.JOB_NAME}-${env.BUILD_NUMBER}-最新的环境部署失败",
        mimeType: "text/html",
        body: """
        <table width="95%" cellpadding="0" cellspacing="0"
            style="font-size: 11pt; font-family: Tahoma, Arial, Helvetica, sans-serif">
            <tr>
                <td><br />
                <b><font color="#0B610B">构建信息</font></b>
                <hr size="2" width="100%" align="center" /></td>
            </tr>
            <tr>
                <td>
                    <ul>
                        <li>构建名称：${JOB_NAME}</li>
                        <li>构建编号：${BUILD_NUMBER}  </li>
                        <li>GIT 地址：${git_url}</li>
                        <li>GIT 分支：${git_branch}</li>
                    </ul>
                </td>
            </tr>
            <tr>
                <td><b><font color="#0B610B">构建日志 :</font></b>
        		<p>页面链接请点击： <a href='${env.RUN_DISPLAY_URL}'>${env.JOB_NAME}</a></p>
            </tr>

        </table>
        """,
        to: "935661686@qq.com",
      )
    }
  }
}