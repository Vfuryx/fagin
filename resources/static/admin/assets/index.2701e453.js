import{_ as x,G as j,S as b,bd as T,a as L,Z as l}from"./index.6d3ed4e5.js";import{A as k,cx as F,j as f,u as o,a0 as s,B as n,D as w,w as B,a1 as m,ad as c,K as v,ab as y}from"./vendor.3850bdb6.js";import{c as S,u as C}from"./index.2d2673c5.js";import D from"./SessionTimeoutLogin.1cc2f0ce.js";import"./index.da3ea590.js";import"./useWindowSizeFn.4c2ec928.js";import"./useContentViewHeight.34544381.js";/* empty css               *//* empty css               */import"./useSortable.7a35deac.js";import"./lock.1a54c299.js";import"./Login.d017ea99.js";import"./LoginForm.7e908ddd.js";/* empty css              *//* empty css               *//* empty css               *//* empty css               */import"./LoginFormTitle.b4a19668.js";import"./ForgetPasswordForm.29a9bb5b.js";import"./index.f4e2dde3.js";import"./RegisterForm.a5053328.js";import"./index.c99e5782.js";import"./MobileForm.faf471fb.js";import"./QrCodeForm.49a7f610.js";import"./index.7f3b171b.js";import"./download.62bc70aa.js";import"./base64Conver.08b9f4ec.js";const P=k({name:"LayoutFeatures",components:{BackTop:F,LayoutLockPage:S(()=>j(()=>import("./index.ca1fbbac.js"),["assets/index.ca1fbbac.js","assets/vendor.3850bdb6.js","assets/vendor.ced4861a.css","assets/LockPage.683cf0df.js","assets/LockPage.592307e9.css","assets/index.6d3ed4e5.js","assets/index.460d50e9.css","assets/lock.1a54c299.js","assets/header.d801b988.js"])),SettingDrawer:S(()=>j(()=>import("./index.c0f37e81.js").then(function(e){return e.i}),["assets/index.c0f37e81.js","assets/index.35b5cf30.css","assets/index.71a54c7b.js","assets/index.5c7227e9.css","assets/index.1bb9cac4.css","assets/index.6d3ed4e5.js","assets/index.460d50e9.css","assets/vendor.3850bdb6.js","assets/vendor.ced4861a.css","assets/index.2d2673c5.js","assets/index.585f44c9.css","assets/index.65674215.css","assets/index.da3ea590.js","assets/index.55076fdd.css","assets/useWindowSizeFn.4c2ec928.js","assets/useContentViewHeight.34544381.js","assets/useSortable.7a35deac.js","assets/lock.1a54c299.js"])),SessionTimeoutLogin:D},setup(){const{getUseOpenBackTop:e,getShowSettingButton:d,getSettingButtonPosition:p,getFullContent:u}=b(),g=T(),{prefixCls:_}=L("setting-drawer-fearure"),{getShowHeader:i}=C(),r=f(()=>g.getSessionTimeout),a=f(()=>{if(!o(d))return!1;const t=o(p);return t===l.AUTO?!o(i)||o(u):t===l.FIXED});return{getTarget:()=>document.body,getUseOpenBackTop:e,getIsFixedSettingDrawer:a,prefixCls:_,getIsSessionTimeout:r}}});function E(e,d,p,u,g,_){const i=s("LayoutLockPage"),r=s("BackTop"),a=s("SettingDrawer"),t=s("SessionTimeoutLogin");return n(),w(y,null,[B(i),e.getUseOpenBackTop?(n(),m(r,{key:0,target:e.getTarget},null,8,["target"])):c("",!0),e.getIsFixedSettingDrawer?(n(),m(a,{key:1,class:v(e.prefixCls)},null,8,["class"])):c("",!0),e.getIsSessionTimeout?(n(),m(t,{key:2})):c("",!0)],64)}var re=x(P,[["render",E]]);export{re as default};
