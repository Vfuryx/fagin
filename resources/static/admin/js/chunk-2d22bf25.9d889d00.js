(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-2d22bf25"],{f0af:function(e,t,a){"use strict";a.r(t);var l=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("a-modal",{attrs:{title:e.model.id>0?"管理员编辑":"管理员新增",width:860,visible:e.visible,confirmLoading:e.loading,"cancel-text":"取消"},on:{ok:e.handleOk,cancel:e.handleCancel}},[a("template",{slot:"footer"},[a("a-button",{on:{click:e.handleCancel}},[e._v("关闭")])],1),a("a-spin",{attrs:{spinning:e.loading}},[a("a-form-model",{ref:"ruleForm",attrs:{layout:"inline",model:e.model,rules:e.rules}},[e.model&&e.model.id>0?a("a-form-model-item",{staticStyle:{width:"88%","margin-right":"0"},attrs:{"label-col":{span:4},"wrapper-col":{span:5},label:"日志编号",prop:"id"}},[a("a-input",{attrs:{disabled:""},model:{value:e.model.id,callback:function(t){e.$set(e.model,"id",t)},expression:"model.id"}})],1):e._e(),a("a-form-model-item",{staticStyle:{width:"50%","margin-right":"0"},attrs:{label:"操作人员","label-col":{span:7},"wrapper-col":{span:17},prop:"user"}},[a("a-input",{attrs:{disabled:""},model:{value:e.model.user,callback:function(t){e.$set(e.model,"user",t)},expression:"model.user"}})],1),a("a-form-model-item",{staticStyle:{width:"50%","margin-right":"0"},attrs:{label:"操作日期","label-col":{span:7},"wrapper-col":{span:17},prop:"created_at"}},[a("a-input",{attrs:{disabled:""},model:{value:e.model.created_at,callback:function(t){e.$set(e.model,"created_at",t)},expression:"model.created_at"}})],1),a("a-form-model-item",{staticStyle:{width:"50%","margin-right":"0"},attrs:{label:"系统模块","label-col":{span:7},"wrapper-col":{span:17},prop:"module"}},[a("a-input",{attrs:{disabled:""},model:{value:e.model.module,callback:function(t){e.$set(e.model,"module",t)},expression:"model.module"}})],1),a("a-form-model-item",{staticStyle:{width:"50%","margin-right":"0"},attrs:{label:"请求方式","label-col":{span:7},"wrapper-col":{span:17},prop:"method"}},[a("a-input",{attrs:{disabled:""},model:{value:e.model.method,callback:function(t){e.$set(e.model,"method",t)},expression:"model.method"}})],1),a("a-form-model-item",{staticStyle:{width:"50%","margin-right":"0"},attrs:{label:"IP","label-col":{span:7},"wrapper-col":{span:17},prop:"ip"}},[a("a-input",{attrs:{disabled:""},model:{value:e.model.ip,callback:function(t){e.$set(e.model,"ip",t)},expression:"model.ip"}})],1),a("a-form-model-item",{staticStyle:{width:"50%","margin-right":"0"},attrs:{label:"访问位置","label-col":{span:7},"wrapper-col":{span:17},prop:"location"}},[a("a-input",{attrs:{disabled:""},model:{value:e.model.location,callback:function(t){e.$set(e.model,"location",t)},expression:"model.location"}})],1),a("a-form-model-item",{staticStyle:{width:"50%","margin-right":"0"},attrs:{label:"耗时","label-col":{span:7},"wrapper-col":{span:17},prop:"latency_time"}},[a("a-input",{attrs:{disabled:""},model:{value:e.model.latency_time,callback:function(t){e.$set(e.model,"latency_time",t)},expression:"model.latency_time"}})],1),a("a-form-model-item",{staticStyle:{width:"88%","margin-right":"0"},attrs:{label:"body","label-col":{span:4},"wrapper-col":{span:20},prop:"input"}},[a("a-textarea",{attrs:{disabled:""},model:{value:e.model.input,callback:function(t){e.$set(e.model,"input",t)},expression:"model.input"}})],1),a("a-form-model-item",{staticStyle:{width:"88%","margin-right":"0"},attrs:{label:"UserAgent","label-col":{span:4},"wrapper-col":{span:20},prop:"user_agent"}},[a("a-textarea",{attrs:{disabled:""},model:{value:e.model.user_agent,callback:function(t){e.$set(e.model,"user_agent",t)},expression:"model.user_agent"}})],1)],1)],1)],2)},o=[],i=(a("922d"),a("09d9")),n=(a("98a7"),a("7bec")),r=(a("17ac"),a("ff57")),d=a("8bbf"),s=a.n(d);s.a.use(r["a"]),s.a.use(n["a"]),s.a.use(i["a"]);var m={props:{visible:{type:Boolean,required:!0},loading:{type:Boolean,default:function(){return!1}},model:{type:Object,default:function(){return null}}},data:function(){return{form:{id:0,user:"",method:"",path:"",ip:"",location:"",module:"",operation:"",input:"",latency_time:"",user_agent:"",created_at:""},list:[],rules:{}}},created:function(){},mounted:function(){},methods:{handleOk:function(){return this.$emit("ok",this.$refs.ruleForm,this.model)},handleCancel:function(){return this.$emit("cancel",this.$refs.ruleForm)}}},c=m,p=a("2877"),u=Object(p["a"])(c,l,o,!1,null,null,null);t["default"]=u.exports}}]);