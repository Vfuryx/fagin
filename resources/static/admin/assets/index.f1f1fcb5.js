var l=(o,i,e)=>new Promise((u,r)=>{var a=s=>{try{t(e.next(s))}catch(d){r(d)}},n=s=>{try{t(e.throw(s))}catch(d){r(d)}},t=s=>s.done?u(s.value):Promise.resolve(s.value).then(a,n);t((e=e.apply(o,i)).next())});import{P as F}from"./index.a09675a6.js";import{B as x}from"./BasicForm.f9b68d8c.js";import{u as B}from"./useForm.ff740053.js";import{_}from"./index.6d3ed4e5.js";import{A as w,a0 as c,B as h,a1 as C,a6 as m,H as f,w as p,ae as j}from"./vendor.3850bdb6.js";/* empty css               *//* empty css                */import"./useWindowSizeFn.4c2ec928.js";import"./useContentViewHeight.34544381.js";/* empty css               *//* empty css               *//* empty css               *//* empty css                *//* empty css              *//* empty css               *//* empty css                *//* empty css                */import"./index.f4e2dde3.js";/* empty css                *//* empty css                *//* empty css                */import"./index.ab15df17.js";/* empty css                */import"./uuid.2b29000c.js";import"./download.62bc70aa.js";import"./base64Conver.08b9f4ec.js";import"./index.c99e5782.js";const b=[{field:"passwordOld",label:"\u5F53\u524D\u5BC6\u7801",component:"InputPassword",required:!0},{field:"passwordNew",label:"\u65B0\u5BC6\u7801",component:"StrengthMeter",componentProps:{placeholder:"\u65B0\u5BC6\u7801"},rules:[{required:!0,message:"\u8BF7\u8F93\u5165\u65B0\u5BC6\u7801"}]},{field:"confirmPassword",label:"\u786E\u8BA4\u5BC6\u7801",component:"InputPassword",dynamicRules:({values:o})=>[{required:!0,validator:(i,e)=>e?e!==o.passwordNew?Promise.reject("\u4E24\u6B21\u8F93\u5165\u7684\u5BC6\u7801\u4E0D\u4E00\u81F4!"):Promise.resolve():Promise.reject("\u4E0D\u80FD\u4E3A\u7A7A")}]}],g=w({name:"ChangePassword",components:{BasicForm:x,PageWrapper:F},setup(){const[o,{validate:i,resetFields:e}]=B({size:"large",labelWidth:100,showActionButtonGroup:!1,schemas:b});function u(){return l(this,null,function*(){try{const r=yield i(),{passwordOld:a,passwordNew:n}=r;console.log(a,n)}catch(r){}})}return{register:o,resetFields:e,handleSubmit:u}}}),E={class:"py-8 bg-white flex flex-col justify-center items-center"},P={class:"flex justify-center"},A=j(" \u91CD\u7F6E "),v=j(" \u786E\u8BA4 ");function y(o,i,e,u,r,a){const n=c("BasicForm"),t=c("a-button"),s=c("PageWrapper");return h(),C(s,{title:"\u4FEE\u6539\u5F53\u524D\u7528\u6237\u5BC6\u7801",content:"\u4FEE\u6539\u6210\u529F\u540E\u4F1A\u81EA\u52A8\u9000\u51FA\u5F53\u524D\u767B\u5F55\uFF01"},{default:m(()=>[f("div",E,[p(n,{onRegister:o.register},null,8,["onRegister"]),f("div",P,[p(t,{onClick:o.resetFields},{default:m(()=>[A]),_:1},8,["onClick"]),p(t,{class:"!ml-4",type:"primary",onClick:o.handleSubmit},{default:m(()=>[v]),_:1},8,["onClick"])])])]),_:1})}var te=_(g,[["render",y]]);export{te as default};
