var e=Object.defineProperty,t=Object.defineProperties,o=Object.getOwnPropertyDescriptors,n=Object.getOwnPropertySymbols,r=Object.prototype.hasOwnProperty,i=Object.prototype.propertyIsEnumerable,a=(t,o,n)=>o in t?e(t,o,{enumerable:!0,configurable:!0,writable:!0,value:n}):t[o]=n;import{ap as s}from"./index.ab59a04a.js";import{u as d,_ as c}from"./useTable.52343b10.js";import{_ as m}from"./useForm.62bbed2c.js";import{u as p}from"./index.89a9d4e0.js";import{P as l}from"./permCodeEnum.0c8982e3.js";import{g as u,c as f,s as b,_ as j,d as y}from"./DeptModal.vue_vue&type=script&setup=true&lang.9c9c8eb1.js";import{D as x,bq as h,a0 as g,F as v,G as w,w as S,a5 as C,u as D,a1 as O,ac as P,ad as R,$ as _}from"./vendor.5da97255.js";/* empty css              */import"./index.6948ae8c.js";/* empty css              *//* empty css              */import"./useWindowSizeFn.df7b0c93.js";import"./useContentViewHeight.ef6fc368.js";/* empty css              *//* empty css              *//* empty css              *//* empty css              */import"./useSortable.a06a1df1.js";/* empty css              *//* empty css              *//* empty css              */import"./index.351a93c9.js";import"./download.3af720f3.js";import"./index.954cca7d.js";const A=R(" 新增部门 "),F=x({name:"DepartmentManagement"});var T,k=x((T=((e,t)=>{for(var o in t||(t={}))r.call(t,o)&&a(e,o,t[o]);if(n)for(var o of n(t))i.call(t,o)&&a(e,o,t[o]);return e})({},F),t(T,o({setup:function(e){const{hasPermission:t}=s(),[o,{openModal:n}]=p(),[r,{reload:i,expandAll:a}]=d({title:"部门列表",api:u,columns:f,formConfig:{labelWidth:120,schemas:b},isTreeTable:!0,pagination:!1,striped:!1,useSearchForm:t(l.AdminSystemDepartmentQuery),showTableSetting:!0,bordered:!0,showIndexColumn:!1,canResize:!1,actionColumn:{width:80,title:"操作",dataIndex:"action",slots:{customRender:"action"},fixed:void 0}});function x(){n(!0,{isUpdate:!1})}function R(e){n(!0,{record:e,isUpdate:!0})}function F(e){return t=this,o=null,n=function*(){yield y(e.id),yield i()},new Promise(((e,r)=>{var i=e=>{try{s(n.next(e))}catch(t){r(t)}},a=e=>{try{s(n.throw(e))}catch(t){r(t)}},s=t=>t.done?e(t.value):Promise.resolve(t.value).then(i,a);s((n=n.apply(t,o)).next())}));var t,o,n}function T(){i()}function k(){_(a)}return h((()=>{i()})),(e,n)=>{const i=g("a-button");return v(),w("div",null,[S(D(c),{onRegister:D(r),onFetchSuccess:k},{toolbar:C((()=>[D(t)(D(l).AdminSystemDepartmentCreate)?(v(),O(i,{key:0,type:"primary",onClick:x},{default:C((()=>[A])),_:1})):P("",!0)])),action:C((({record:e})=>[S(D(m),{actions:[{icon:"clarity:note-edit-line",onClick:R.bind(null,e),auth:D(l).AdminSystemDepartmentUpdate},{icon:"ant-design:delete-outlined",color:"error",auth:D(l).AdminSystemDepartmentDelete,popConfirm:{title:"是否确认删除",confirm:F.bind(null,e)}}]},null,8,["actions"])])),_:1},8,["onRegister"]),S(j,{onRegister:D(o),onSuccess:T},null,8,["onRegister"])])}}}))));export{k as default};
