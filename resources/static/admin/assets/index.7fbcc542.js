var e=Object.defineProperty,o=Object.defineProperties,t=Object.getOwnPropertyDescriptors,i=Object.getOwnPropertySymbols,r=Object.prototype.hasOwnProperty,n=Object.prototype.propertyIsEnumerable,a=(o,t,i)=>t in o?e(o,t,{enumerable:!0,configurable:!0,writable:!0,value:i}):o[t]=i;import{D as s,bq as d,F as p,G as m,w as l,u as c}from"./vendor.5da97255.js";import{aj as j,ap as b}from"./index.7037d70a.js";import{P as f}from"./permCodeEnum.0c8982e3.js";import{u as g,_ as u}from"./useTable.4b946ab8.js";import"./useForm.152ab070.js";/* empty css              */import"./index.e53eff0c.js";/* empty css              *//* empty css              */import"./useWindowSizeFn.df7b0c93.js";import"./useContentViewHeight.e053a8fa.js";/* empty css              *//* empty css              *//* empty css              */import"./index.6918c5f5.js";/* empty css              */import"./useSortable.6c5043ea.js";/* empty css              *//* empty css              *//* empty css              */import"./index.6581a3c2.js";import"./download.be553a46.js";import"./index.bf50acea.js";const x=[{title:"登录名称",dataIndex:"login_name",width:100},{title:"登录地址",dataIndex:"ip",width:100},{title:"登录地点",dataIndex:"location",width:100},{title:"浏览器",dataIndex:"browser",width:100},{title:"操作系统",dataIndex:"os",width:100},{title:"登录日期",dataIndex:"created_at",width:150}],w=[{field:"login_name",label:"登录名称",component:"Input",colProps:{span:8}},{field:"time",label:"时间范围",component:"RangePicker",colProps:{span:11},componentProps:{showTime:!0}}];var h,L;(L=h||(h={})).LoginLogList="/login-logs",L.LoginLogDetail="/login-logs/";const P=e=>j.get({url:h.LoginLogList,params:e}),y=s({name:"LoginLogManagement"});var I,O=s((I=((e,o)=>{for(var t in o||(o={}))r.call(o,t)&&a(e,t,o[t]);if(i)for(var t of i(o))n.call(o,t)&&a(e,t,o[t]);return e})({},y),o(I,t({setup:function(e){const{hasPermission:o}=b(),[t,{reload:i}]=g({title:"登录日志列表",api:P,columns:x,formConfig:{labelWidth:120,schemas:w},useSearchForm:o(f.AdminMonitorLoginLogQuery),showTableSetting:!0,bordered:!0,showIndexColumn:!1});return d((()=>{i()})),(e,o)=>(p(),m("div",null,[l(c(u),{onRegister:c(t)},null,8,["onRegister"])]))}}))));export{O as default};
