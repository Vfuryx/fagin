import{a as e,f as a,e as s}from"./index.7037d70a.js";import{B as r,b as o}from"./index.6918c5f5.js";import{g as t,h as i}from"./useForm.152ab070.js";import{u as n}from"./lock.6275374b.js";import{h as l}from"./header.d801b988.js";import{D as d,j as c,a0 as m,F as p,a1 as f,a5 as u,K as x,N as h,L as j,w as v,ad as g,a4 as k}from"./vendor.5da97255.js";import"./useWindowSizeFn.df7b0c93.js";/* empty css              *//* empty css              *//* empty css              *//* empty css              *//* empty css              *//* empty css              *//* empty css              */import"./index.6581a3c2.js";import"./download.be553a46.js";import"./index.bf50acea.js";var _=d({name:"LockModal",components:{BasicModal:r,BasicForm:t},setup(){const{t:r}=s(),{prefixCls:t}=e("header-lock-modal"),d=a(),m=n(),p=c((()=>{var e;return null==(e=d.getUserInfo)?void 0:e.realName})),[f,{closeModal:u}]=o(),[x,{validateFields:h,resetFields:j}]=i({showActionButtonGroup:!1,schemas:[{field:"password",label:r("layout.header.lockScreenPassword"),component:"InputPassword",required:!0}]});return{t:r,prefixCls:t,getRealName:p,register:f,registerForm:x,handleLock:function(){return e=this,a=null,s=function*(){const e=(yield h()).password;u(),m.setLockInfo({isLock:!0,pwd:e}),yield j()},new Promise(((r,o)=>{var t=e=>{try{n(s.next(e))}catch(a){o(a)}},i=e=>{try{n(s.throw(e))}catch(a){o(a)}},n=e=>e.done?r(e.value):Promise.resolve(e.value).then(t,i);n((s=s.apply(e,a)).next())}));var e,a,s},avatar:c((()=>{const{avatar:e}=d.getUserInfo;return e||l}))}}});const b=["src"];_.render=function(e,a,s,r,o,t){const i=m("BasicForm"),n=m("a-button"),l=m("BasicModal");return p(),f(l,k({footer:null,title:e.t("layout.header.lockScreen")},e.$attrs,{class:e.prefixCls,onRegister:e.register}),{default:u((()=>[x("div",{class:h(`${e.prefixCls}__entry`)},[x("div",{class:h(`${e.prefixCls}__header`)},[x("img",{src:e.avatar,class:h(`${e.prefixCls}__header-img`)},null,10,b),x("p",{class:h(`${e.prefixCls}__header-name`)},j(e.getRealName),3)],2),v(i,{onRegister:e.registerForm},null,8,["onRegister"]),x("div",{class:h(`${e.prefixCls}__footer`)},[v(n,{type:"primary",block:"",class:"mt-2",onClick:e.handleLock},{default:u((()=>[g(j(e.t("layout.header.lockScreenBtn")),1)])),_:1},8,["onClick"])],2)],2)])),_:1},16,["title","class","onRegister"])};export{_ as default};