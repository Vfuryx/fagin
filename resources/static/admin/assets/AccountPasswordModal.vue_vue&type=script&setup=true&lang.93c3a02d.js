var s=(s,a,e)=>new Promise(((o,t)=>{var r=s=>{try{i(e.next(s))}catch(a){t(a)}},n=s=>{try{i(e.throw(s))}catch(a){t(a)}},i=s=>s.done?o(s.value):Promise.resolve(s.value).then(r,n);i((e=e.apply(s,a)).next())}));import{b as a,B as e}from"./index.89a9d4e0.js";import{h as o,g as t}from"./useForm.62bbed2c.js";import{b as r}from"./account.data.54047abd.js";import{d as n}from"./account.19d69919.js";import{D as i,r as c,F as l,a1 as d,a5 as u,w as m,u as p,a4 as f}from"./vendor.5da97255.js";var h=i({emits:["success","register"],setup(i,{emit:h}){const v=c(0),[g,{resetFields:y,validate:b}]=o({labelWidth:100,schemas:r,showActionButtonGroup:!1,actionColOptions:{span:23}}),[j,{setModalProps:w,closeModal:x}]=a((a=>s(this,null,(function*(){y(),w({confirmLoading:!1}),v.value=a.record.id}))));function R(){return s(this,null,(function*(){try{const s=yield b();w({confirmLoading:!0}),yield n(v.value,s.password),x(),h("success")}finally{w({confirmLoading:!1})}}))}return(s,a)=>(l(),d(p(e),f(s.$attrs,{onRegister:p(j),title:"修改密码",onOk:R}),{default:u((()=>[m(p(t),{onRegister:p(g)},null,8,["onRegister"])])),_:1},16,["onRegister"]))}});export{h as _};
