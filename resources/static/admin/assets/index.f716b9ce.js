var b=Object.defineProperty,g=Object.defineProperties;var C=Object.getOwnPropertyDescriptors;var n=Object.getOwnPropertySymbols;var h=Object.prototype.hasOwnProperty,_=Object.prototype.propertyIsEnumerable;var a=(o,e,i)=>e in o?b(o,e,{enumerable:!0,configurable:!0,writable:!0,value:i}):o[e]=i,m=(o,e)=>{for(var i in e||(e={}))h.call(e,i)&&a(o,i,e[i]);if(n)for(var i of n(e))_.call(e,i)&&a(o,i,e[i]);return o},d=(o,e)=>g(o,C(e));import{B as w}from"./TableImg.1321f9ef.js";import{T as D}from"./BasicForm.f9b68d8c.js";import{u as B}from"./useTable.5a3ce825.js";import{an as F}from"./index.6d3ed4e5.js";import{u as T}from"./index.71a54c7b.js";import{g as E,c as L,s as v,_ as A}from"./OperLogDrawer.cc2b9828.js";import{P as p}from"./permCodeEnum.0c8982e3.js";import{A as c,bs as O,B as P,D as R,w as r,a6 as S,u as s}from"./vendor.3850bdb6.js";/* empty css                *//* empty css              */import"./useForm.ff740053.js";import"./index.a09675a6.js";/* empty css               *//* empty css                */import"./useWindowSizeFn.4c2ec928.js";import"./useContentViewHeight.34544381.js";/* empty css                *//* empty css                *//* empty css               */import"./uuid.2b29000c.js";import"./index.ab15df17.js";/* empty css               */import"./useSortable.7a35deac.js";/* empty css                *//* empty css                *//* empty css               *//* empty css               */import"./index.f4e2dde3.js";/* empty css                *//* empty css                *//* empty css                */import"./download.62bc70aa.js";import"./base64Conver.08b9f4ec.js";import"./index.c99e5782.js";/* empty css               */import"./index.be3f98b2.js";import"./index.a18b2319.js";/* empty css                */import"./useDescription.a3524a85.js";const k=c({name:"OperLogManagement"});function y(o){const{hasPermission:e}=F(),[i,{openDrawer:u}]=T(),[j,{reload:l}]=B({title:"\u64CD\u4F5C\u65E5\u5FD7\u5217\u8868",api:E,columns:L,formConfig:{labelWidth:120,schemas:v},useSearchForm:e(p.AdminMonitorOperLogQuery),showTableSetting:!0,bordered:!0,showIndexColumn:!1,actionColumn:{width:80,title:"\u64CD\u4F5C",dataIndex:"action",slots:{customRender:"action"},fixed:void 0}});function x(t){u(!0,{record:t})}return O(()=>{l()}),(t,I)=>(P(),R("div",null,[r(s(w),{onRegister:s(j)},{action:S(({record:f})=>[r(s(D),{actions:[{icon:"clarity:info-standard-line",tooltip:"\u67E5\u770B\u8BE6\u60C5",onClick:x.bind(null,f),auth:s(p).AdminMonitorOperLogDetail}]},null,8,["actions"])]),_:1},8,["onRegister"]),r(A,{onRegister:s(i)},null,8,["onRegister"])]))}const Be=c(d(m({},k),{setup:y}));export{Be as default};
