var e,t,a=Object.defineProperty,l=Object.getOwnPropertySymbols,s=Object.prototype.hasOwnProperty,n=Object.prototype.propertyIsEnumerable,r=(e,t,l)=>t in e?a(e,t,{enumerable:!0,configurable:!0,writable:!0,value:l}):e[t]=l,o=(e,t,a)=>new Promise(((l,s)=>{var n=e=>{try{o(a.next(e))}catch(t){s(t)}},r=e=>{try{o(a.throw(e))}catch(t){s(t)}},o=e=>e.done?l(e.value):Promise.resolve(e.value).then(n,r);o((a=a.apply(e,t)).next())}));import{h as i,g as u}from"./useForm.62bbed2c.js";import{z as d,bH as p,D as m,r as c,j as f,u as h,F as v,a1 as b,a5 as g,w as y,a4 as R,t as w}from"./vendor.5da97255.js";/* empty css              */import{aj as S,i as k}from"./index.ab59a04a.js";import{P as x}from"./permCodeEnum.0c8982e3.js";import{u as I,B as j}from"./index.42015244.js";import{_ as P}from"./Tree.vue_vue&type=style&index=0&lang.521efaeb.js";(t=e||(e={})).RoleList="/roles",t.GetRoleDetail="/roles/",t.CreateRole="/roles",t.UpdateRole="/roles/",t.DeleteRole="/roles/",t.SetRoleStatus="/roles/",t.GetMenuAll="/roles/menus";const _=t=>S.get({url:e.RoleList,params:t}),D=t=>S.delete({url:e.DeleteRole+t.toString()}),C=[{title:"编号",dataIndex:"id",width:100},{title:"角色名称",dataIndex:"name",width:200},{title:"角色值",dataIndex:"key",width:180},{title:"排序",dataIndex:"sort",width:50},{title:"状态",dataIndex:"status",width:120,auth:x.AdminSystemRoleUpdateStatus,customRender:({record:t})=>"admin"==t.key?d("div",null,null):(Reflect.has(t,"pendingStatus")||(t.pendingStatus=!1),d(p,{checked:1===t.status,checkedChildren:"已启用",unCheckedChildren:"已禁用",loading:t.pendingStatus,onChange(a){t.pendingStatus=!0;const l=a?1:0,{createMessage:s}=k();var n,r;(n=t.id,r=l,S.put({url:e.SetRoleStatus+n.toString()+"/status",params:{status:r}})).then((()=>{t.status=l,s.success("已成功修改角色状态")})).catch((()=>{s.error("修改角色状态失败")})).finally((()=>{t.pendingStatus=!1}))}}))},{title:"创建时间",dataIndex:"created_at",width:180},{title:"备注",dataIndex:"remark"}],O=[{field:"name",label:"角色名称",component:"Input",colProps:{span:8}},{field:"status",label:"状态",component:"Select",componentProps:{options:[{label:"启用",value:1},{label:"停用",value:0}]},colProps:{span:8}}],F=[{field:"id",label:"编号",component:"Input",show:!1},{field:"name",label:"角色名称",required:!0,component:"Input"},{field:"key",label:"角色值",component:"Input",required:({values:e})=>!e.id,ifShow:({values:e})=>!e.id},{field:"status",label:"状态",component:"RadioButtonGroup",defaultValue:1,componentProps:{options:[{label:"启用",value:1},{label:"停用",value:0}]}},{label:"备注",field:"remark",component:"InputTextArea"},{label:" ",field:"menu_ids",slot:"menu",component:"Input"}];var G=m({emits:["success","register"],setup(t,{emit:a}){const d=c(!0),p=c([]),m=c(0),[k,{resetFields:x,setFieldsValue:_,validate:D}]=i({labelWidth:90,schemas:F,showActionButtonGroup:!1}),[C,{setDrawerProps:O,closeDrawer:G}]=I((t=>o(this,null,(function*(){if(x(),O({confirmLoading:!1}),0===h(p).length&&(p.value=yield S.get({url:e.GetMenuAll})),d.value=!!(null==t?void 0:t.isUpdate),h(d)){m.value=t.record.id;const a=yield(t=>S.get({url:e.GetRoleDetail+t.toString()}))(m.value);_(((e,t)=>{for(var a in t||(t={}))s.call(t,a)&&r(e,a,t[a]);if(l)for(var a of l(t))n.call(t,a)&&r(e,a,t[a]);return e})({},a))}})))),U=f((()=>h(d)?"编辑角色":"新增角色"));function A(){return o(this,null,(function*(){try{const l=yield D();if(O({confirmLoading:!0}),!l)return;h(d)?yield((t,a)=>S.put({url:e.UpdateRole+t.toString(),params:a}))(m.value,w({name:l.name,remark:l.remark,status:l.status,menu_ids:l.menu_ids})):yield(t=w({name:l.name,key:l.key,remark:l.remark,status:l.status,menu_ids:l.menu_ids}),S.post({url:e.CreateRole,params:t})),G(),a("success")}finally{O({confirmLoading:!1})}var t}))}return(e,t)=>(v(),b(h(j),R(e.$attrs,{onRegister:h(C),showFooter:"",title:h(U),width:"45%",onOk:A}),{default:g((()=>[y(h(u),{onRegister:h(k)},{menu:g((({model:e,field:t})=>[y(h(P),{value:e[t],"onUpdate:value":a=>e[t]=a,treeData:p.value,replaceFields:{title:"title",key:"id"},checkable:"",toolbar:"",title:"菜单分配"},null,8,["value","onUpdate:value","treeData"])])),_:1},8,["onRegister"])])),_:1},16,["onRegister","title"]))}});export{G as _,C as c,D as d,_ as g,O as s};