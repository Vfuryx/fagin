(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-2d0a4c1c"],{"0897":function(e,t,a){"use strict";a.r(t);var r=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("a-modal",{attrs:{title:e.model.id>0?"管理员编辑":"管理员新增",width:860,visible:e.visible,confirmLoading:e.loading,"ok-text":"确认","cancel-text":"取消"},on:{ok:e.handleOk,cancel:e.handleCancel}},[a("a-spin",{attrs:{spinning:e.loading}},[a("a-form-model",{ref:"ruleForm",attrs:{layout:"inline",model:e.model,rules:e.rules}},[e.model&&e.model.id>0?a("a-form-model-item",{staticStyle:{width:"88%","margin-right":"0"},attrs:{"label-col":{span:4},"wrapper-col":{span:5},label:"主键ID",prop:"id"}},[a("a-input",{attrs:{disabled:""},model:{value:e.model.id,callback:function(t){e.$set(e.model,"id",t)},expression:"model.id"}})],1):e._e(),a("a-form-model-item",{staticStyle:{width:"88%","margin-right":"0"},attrs:{label:"标题","label-col":{span:4},"wrapper-col":{span:5},prop:"title"}},[a("a-input",{model:{value:e.model.title,callback:function(t){e.$set(e.model,"title",t)},expression:"model.title"}})],1),a("a-form-model-item",{staticStyle:{width:"88%","margin-right":"0"},attrs:{label:"视频上传","label-col":{span:4},"wrapper-col":{span:10},prop:"title"}},[a("a-upload",{attrs:{"list-type":"picture",multiple:!1,"before-upload":e.beforeUpload,"custom-request":e.uploadImage,accept:"video/mp4"}},[a("a-button",[a("a-icon",{attrs:{type:"upload"}}),e._v(" Upload ")],1)],1)],1),a("a-form-model-item",{staticStyle:{width:"88%","margin-right":"0"},attrs:{label:"描述","label-col":{span:4},"wrapper-col":{span:19},prop:"description"}},[a("a-textarea",{model:{value:e.model.description,callback:function(t){e.$set(e.model,"description",t)},expression:"model.description"}})],1),a("a-form-model-item",{staticStyle:{width:"88%","margin-right":"0"},attrs:{label:"状态","label-col":{span:4},"wrapper-col":{span:5},prop:"status"}},[a("a-radio-group",{attrs:{name:"status"},model:{value:e.model.status,callback:function(t){e.$set(e.model,"status",t)},expression:"model.status"}},[a("a-radio",{attrs:{value:1}},[e._v("开启")]),a("a-radio",{attrs:{value:0}},[e._v("关闭")])],1)],1)],1)],1)],1)},i=[],l=(a("922d"),a("09d9")),o=(a("98a7"),a("7bec")),s=(a("17ac"),a("ff57")),n=a("8bbf"),d=a.n(n),u=a("4360");d.a.use(s["a"]),d.a.use(o["a"]),d.a.use(l["a"]);var m={props:{visible:{type:Boolean,required:!0},loading:{type:Boolean,default:function(){return!1}},model:{type:Object,default:function(){return null}}},data:function(){return{form:{id:0,title:"",duration:"",status:0,create_at:"",description:""},list:[],rules:{id:[{required:!0,message:"不能空",trigger:"blur"},{type:"integer",min:0,message:"不正确",trigger:"blur"}],title:[{required:!0,message:"不能为空",trigger:"blur"},{min:1,max:60,message:"范围 1 - 60 字符",trigger:"blur"}],description:[{required:!0,message:"不能为空",trigger:"blur"},{min:1,max:255,message:"范围 1 - 255 字符",trigger:"blur"}],status:[{required:!0,message:"不能空",trigger:"blur"},{type:"enum",enum:[0,1],message:"类型不正确",trigger:"blur"}]}}},created:function(){},mounted:function(){},methods:{handleOk:function(){return this.$emit("ok",this.$refs.ruleForm,this.model)},handleCancel:function(){return this.$emit("cancel",this.$refs.ruleForm)},beforeUpload:function(e){var t="video/mp4"===e.type;t||this.$message.error("请上传视频格式!");var a=e.size/1024/1024<100;return a||this.$message.error("视频大小限制 10m 以内!"),t&&a},uploadImage:function(e){var t=this;this.confirmLoading=!0;var a=new FormData;a.append("file",e.file),u["a"].dispatch("uploadVideo",a).then((function(e){return 0!==e.code?(t.$notification.error({message:"错误",description:e.message}),!1):(t.confirmLoading=!1,t.model.path=e.data.path,t.$message.info("新增成功"),!0)})).catch((function(e){return t.$notification.error({message:"错误",description:"请求信息失败，请重试"}),!1}))}}},c=m,p=a("2877"),g=Object(p["a"])(c,r,i,!1,null,null,null);t["default"]=g.exports}}]);