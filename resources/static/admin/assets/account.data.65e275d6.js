import{s as e,i as a,b as t}from"./account.b3167b22.js";import{z as n,bp as l,bH as o}from"./vendor.5da97255.js";/* empty css              *//* empty css              */import{P as d}from"./permCodeEnum.0c8982e3.js";import{i}from"./index.7037d70a.js";const r=[{title:"用户名",dataIndex:"username",width:120},{title:"昵称",dataIndex:"nick_name",width:120},{title:"邮箱",dataIndex:"email",width:180},{title:"角色",dataIndex:"roles",width:180,customRender:({record:e})=>{const a=[];return e.roles.forEach((e=>{a.push(n(l,{color:"blue",style:"margin-right: 8px;"},(()=>e.name)))})),n("div",{},a)}},{title:"状态",dataIndex:"status",width:120,auth:d.AdminSystemAccountUpdateStatus,customRender:({record:a})=>"admin"==a.username?n("div",null,null):(Reflect.has(a,"pendingStatus")||(a.pendingStatus=!1),n(o,{checked:1===a.status,checkedChildren:"已启用",unCheckedChildren:"已禁用",loading:a.pendingStatus,onChange(t){a.pendingStatus=!0;const n=t?1:0,{createMessage:l}=i();e(a.id,n).then((()=>{a.status=n,l.success("已成功修改账号状态")})).catch((()=>{l.error("修改账号状态失败")})).finally((()=>{a.pendingStatus=!1}))}}))},{title:"备注",dataIndex:"remark"},{title:"创建时间",dataIndex:"create_at",width:180}],s=[{field:"username",label:"用户名",component:"Input",colProps:{span:8}},{field:"nick_name",label:"昵称",component:"Input",colProps:{span:8}}],u=[{field:"password",label:"密码",component:"InputPassword",defaultValue:"",required:!0,ifShow:!0}],p=e=>[{field:"username",label:"用户名",component:"Input",helpMessage:["不能输入带有admin的用户名"],defaultValue:"",rules:[{required:!0,message:"请输入用户名"},{validator:(t,n)=>new Promise(((t,l)=>{a(n,e.value).then((()=>t())).catch((e=>{l(e.message||"验证失败")}))}))}]},{field:"password",label:"密码",component:"InputPassword",defaultValue:"",required:({})=>!e.value,ifShow:!0},{field:"nick_name",label:"昵称",component:"Input",defaultValue:"",required:!0},{field:"email",label:"邮箱",component:"Input",required:!0},{field:"phone",label:"手机",component:"Input",required:!0},{field:"roles",label:"角色",component:"ApiSelect",componentProps:{api:t,labelField:"name",valueField:"id",mode:"multiple"},required:!0},{field:"department_id",label:"所属部门",component:"TreeSelect",componentProps:{replaceFields:{title:"name",key:"id",value:"id"},getPopupContainer:()=>document.body},required:!0},{field:"remark",label:"备注",component:"InputTextArea"},{field:"status",label:"状态",component:"RadioButtonGroup",defaultValue:1,componentProps:{options:[{label:"启用",value:1},{label:"停用",value:0}]},required:!0},{field:"sex",label:"性别",component:"RadioButtonGroup",defaultValue:0,componentProps:{options:[{label:"未知",value:0},{label:"男",value:1},{label:"女",value:2}]},required:!0}];export{p as a,u as b,r as c,s};