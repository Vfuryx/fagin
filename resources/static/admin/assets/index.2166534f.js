var e=Object.defineProperty,t=Object.defineProperties,n=Object.getOwnPropertyDescriptors,i=Object.getOwnPropertySymbols,o=Object.prototype.hasOwnProperty,r=Object.prototype.propertyIsEnumerable,a=(t,n,i)=>n in t?e(t,n,{enumerable:!0,configurable:!0,writable:!0,value:i}):t[n]=i;import{D as s,bq as d,a0 as c,F as m,G as p,w as l,a5 as u,u as f,a1 as j,ac as b,ad as x}from"./vendor.5da97255.js";import{ap as y}from"./index.ab59a04a.js";import{u as h,_ as g}from"./useTable.52343b10.js";import{_ as v}from"./useForm.62bbed2c.js";import{a as w}from"./index.42015244.js";import{P as S}from"./permCodeEnum.0c8982e3.js";import{g as C,c as O,s as P,_ as M,d as R}from"./MenuDrawer.vue_vue&type=script&setup=true&lang.a9417435.js";/* empty css              */import"./index.6948ae8c.js";/* empty css              *//* empty css              */import"./useWindowSizeFn.df7b0c93.js";import"./useContentViewHeight.ef6fc368.js";/* empty css              *//* empty css              *//* empty css              */import"./index.89a9d4e0.js";/* empty css              */import"./useSortable.a06a1df1.js";/* empty css              *//* empty css              *//* empty css              */import"./index.351a93c9.js";import"./download.3af720f3.js";import"./index.954cca7d.js";/* empty css              */const _=x(" 新增菜单 ");var A,D=s((A=((e,t)=>{for(var n in t||(t={}))o.call(t,n)&&a(e,n,t[n]);if(i)for(var n of i(t))r.call(t,n)&&a(e,n,t[n]);return e})({},{name:"MenuManagement"}),t(A,n({setup:function(e){const{hasPermission:t}=y(),[n,{openDrawer:i}]=w(),[o,{reload:r,expandAll:a}]=h({title:"菜单列表",api:C,columns:O,formConfig:{labelWidth:120,schemas:P},isTreeTable:!0,pagination:!1,striped:!1,useSearchForm:t(S.AdminSystemMenuQuery),showTableSetting:!0,bordered:!0,showIndexColumn:!1,canResize:!1,actionColumn:{width:80,title:"操作",dataIndex:"action",slots:{customRender:"action"},fixed:void 0}});function s(){i(!0,{isUpdate:!1})}function x(e){i(!0,{record:e,isUpdate:!0})}function A(e){return t=this,n=null,i=function*(){yield R(e.id),yield r()},new Promise(((e,o)=>{var r=e=>{try{s(i.next(e))}catch(t){o(t)}},a=e=>{try{s(i.throw(e))}catch(t){o(t)}},s=t=>t.done?e(t.value):Promise.resolve(t.value).then(r,a);s((i=i.apply(t,n)).next())}));var t,n,i}function D(){r()}function F(){}return d((()=>{r()})),(e,i)=>{const r=c("a-button");return m(),p("div",null,[l(f(g),{onRegister:f(o),onFetchSuccess:F},{toolbar:u((()=>[f(t)(f(S).AdminSystemMenuCreate)?(m(),j(r,{key:0,type:"primary",onClick:s},{default:u((()=>[_])),_:1})):b("",!0)])),action:u((({record:e})=>[l(f(v),{actions:[{icon:"clarity:note-edit-line",onClick:x.bind(null,e),auth:f(S).AdminSystemMenuUpdate},{icon:"ant-design:delete-outlined",color:"error",auth:f(S).AdminSystemMenuDelete,popConfirm:{title:"是否确认删除",confirm:A.bind(null,e)}}]},null,8,["actions"])])),_:1},8,["onRegister"]),l(M,{onRegister:f(n),onSuccess:D},null,8,["onRegister"])])}}}))));export{D as default};
