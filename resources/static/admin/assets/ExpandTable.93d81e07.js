import{B as d}from"./TableImg.1321f9ef.js";import{T as c}from"./BasicForm.f9b68d8c.js";import{u as m}from"./useTable.5a3ce825.js";import{P as j}from"./index.a09675a6.js";import{getBasicColumns as F}from"./tableData.3f3da3f1.js";import{d as f}from"./table.23a47f9e.js";import{_ as b}from"./index.6d3ed4e5.js";import{A as x,a0 as n,B,a1 as C,a6 as i,w as a,H as E,J as g}from"./vendor.3850bdb6.js";/* empty css                *//* empty css              */import"./useForm.ff740053.js";/* empty css                *//* empty css                *//* empty css               */import"./uuid.2b29000c.js";import"./useWindowSizeFn.4c2ec928.js";import"./index.ab15df17.js";/* empty css               */import"./useSortable.7a35deac.js";/* empty css                *//* empty css                *//* empty css               *//* empty css               */import"./index.f4e2dde3.js";/* empty css                *//* empty css                *//* empty css                */import"./download.62bc70aa.js";import"./base64Conver.08b9f4ec.js";import"./index.c99e5782.js";/* empty css               *//* empty css                */import"./useContentViewHeight.34544381.js";const T=x({components:{BasicTable:d,TableAction:c,PageWrapper:j},setup(){const[e]=m({api:f,title:"\u53EF\u5C55\u5F00\u8868\u683C\u6F14\u793A",titleHelpMessage:["\u5DF2\u542F\u7528expandRowByClick","\u5DF2\u542F\u7528stopButtonPropagation"],columns:F(),rowKey:"id",canResize:!1,expandRowByClick:!0,actionColumn:{width:160,title:"Action",slots:{customRender:"action"}}});function u(o){console.log("\u70B9\u51FB\u4E86\u5220\u9664",o)}function s(o){console.log("\u70B9\u51FB\u4E86\u542F\u7528",o)}return{registerTable:e,handleDelete:u,handleOpen:s}}});function _(e,u,s,o,w,D){const r=n("TableAction"),p=n("BasicTable"),l=n("PageWrapper");return B(),C(l,{title:"\u53EF\u5C55\u5F00\u8868\u683C",content:"\u4E0D\u53EF\u4E0Escroll\u5171\u7528\u3002TableAction\u7EC4\u4EF6\u53EF\u914D\u7F6EstopButtonPropagation\u6765\u963B\u6B62\u64CD\u4F5C\u6309\u94AE\u7684\u70B9\u51FB\u4E8B\u4EF6\u5192\u6CE1\uFF0C\u4EE5\u4FBF\u914D\u5408Table\u7EC4\u4EF6\u7684expandRowByClick"},{default:i(()=>[a(p,{onRegister:e.registerTable},{expandedRowRender:i(({record:t})=>[E("span",null,"No: "+g(t.no),1)]),action:i(({record:t})=>[a(r,{stopButtonPropagation:"",actions:[{label:"\u5220\u9664",icon:"ic:outline-delete-outline",onClick:e.handleDelete.bind(null,t)}],dropDownActions:[{label:"\u542F\u7528",popConfirm:{title:"\u662F\u5426\u542F\u7528\uFF1F",confirm:e.handleOpen.bind(null,t)}}]},null,8,["actions","dropDownActions"])]),_:1},8,["onRegister"])]),_:1})}var se=b(T,[["render",_]]);export{se as default};
