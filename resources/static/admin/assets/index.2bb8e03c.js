import{B as f}from"./TableImg.1321f9ef.js";import{T as j}from"./BasicForm.f9b68d8c.js";import{u as b}from"./useTable.5a3ce825.js";import{a as x}from"./system.f1ed0cdb.js";import{b as h}from"./index.ab15df17.js";import{D as g,c as C,s as _}from"./DeptModal.19a6ae6d.js";import{_ as T}from"./index.6d3ed4e5.js";import{A as D,a0 as t,B,D as E,w as n,a6 as r,ae as M}from"./vendor.3850bdb6.js";/* empty css                *//* empty css              */import"./useForm.ff740053.js";import"./index.a09675a6.js";/* empty css               *//* empty css                */import"./useWindowSizeFn.4c2ec928.js";import"./useContentViewHeight.34544381.js";/* empty css                *//* empty css                *//* empty css               */import"./uuid.2b29000c.js";/* empty css               */import"./useSortable.7a35deac.js";/* empty css                *//* empty css                *//* empty css               *//* empty css               */import"./index.f4e2dde3.js";/* empty css                *//* empty css                *//* empty css                */import"./download.62bc70aa.js";import"./base64Conver.08b9f4ec.js";import"./index.c99e5782.js";const S=D({name:"DeptManagement",components:{BasicTable:f,DeptModal:g,TableAction:j},setup(){const[e,{openModal:i}]=h(),[d,{reload:l}]=b({title:"\u90E8\u95E8\u5217\u8868",api:x,columns:C,formConfig:{labelWidth:120,schemas:_},pagination:!1,striped:!1,useSearchForm:!0,showTableSetting:!0,bordered:!0,showIndexColumn:!1,canResize:!1,actionColumn:{width:80,title:"\u64CD\u4F5C",dataIndex:"action",slots:{customRender:"action"},fixed:void 0}});function c(){i(!0,{isUpdate:!1})}function m(o){i(!0,{record:o,isUpdate:!0})}function s(o){console.log(o)}function a(){l()}return{registerTable:d,registerModal:e,handleCreate:c,handleEdit:m,handleDelete:s,handleSuccess:a}}}),w=M(" \u65B0\u589E\u90E8\u95E8 ");function F(e,i,d,l,c,m){const s=t("a-button"),a=t("TableAction"),o=t("BasicTable"),u=t("DeptModal");return B(),E("div",null,[n(o,{onRegister:e.registerTable},{toolbar:r(()=>[n(s,{type:"primary",onClick:e.handleCreate},{default:r(()=>[w]),_:1},8,["onClick"])]),action:r(({record:p})=>[n(a,{actions:[{icon:"clarity:note-edit-line",onClick:e.handleEdit.bind(null,p)},{icon:"ant-design:delete-outlined",color:"error",popConfirm:{title:"\u662F\u5426\u786E\u8BA4\u5220\u9664",confirm:e.handleDelete.bind(null,p)}}]},null,8,["actions"])]),_:1},8,["onRegister"]),n(u,{onRegister:e.registerModal,onSuccess:e.handleSuccess},null,8,["onRegister","onSuccess"])])}var le=T(S,[["render",F]]);export{le as default};
