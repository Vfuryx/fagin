var e=Object.defineProperty,a=Object.defineProperties,t=Object.getOwnPropertyDescriptors,r=Object.getOwnPropertySymbols,s=Object.prototype.hasOwnProperty,o=Object.prototype.propertyIsEnumerable,i=(a,t,r)=>t in a?e(a,t,{enumerable:!0,configurable:!0,writable:!0,value:r}):a[t]=r;import{D as n,b9 as l,r as c,a0 as u,F as d,a1 as p,a5 as b,w as v,K as y,G as f,ar as m,L as j,aa as k,ac as O,u as g,ad as w}from"./vendor.5da97255.js";import{P as x}from"./index.6948ae8c.js";import{l as P,a4 as K}from"./index.ab59a04a.js";/* empty css              *//* empty css              */import"./useWindowSizeFn.df7b0c93.js";import"./useContentViewHeight.ef6fc368.js";const _=w(" 禁用账号 "),D=w(" 修改密码 "),T={class:"pt-4 m-4 desc-wrap"},h=n({name:"AccountDetail"});var B,F=n((B=((e,a)=>{for(var t in a||(a={}))s.call(a,t)&&i(e,t,a[t]);if(r)for(var t of r(a))o.call(a,t)&&i(e,t,a[t]);return e})({},h),a(B,t({setup:function(e){var a;const t=l(),r=P(),s=c(null==(a=t.params)?void 0:a.id),o=c("detail"),{setTitle:i}=K();function n(){r("/system/account")}return i("详情：用户"+s.value),(e,a)=>{const t=u("a-button"),r=u("a-tab-pane"),i=u("a-tabs");return d(),p(g(x),{title:"用户"+s.value+"的资料",content:"这是用户资料详情页面。本页面仅用于演示相同路由在tab中打开多个页面并且显示不同的数据",contentBackground:"",onBack:n},{extra:b((()=>[v(t,{type:"primary",danger:""},{default:b((()=>[_])),_:1}),v(t,{type:"primary"},{default:b((()=>[D])),_:1})])),footer:b((()=>[v(i,{"default-active-key":"detail",activeKey:o.value,"onUpdate:activeKey":a[0]||(a[0]=e=>o.value=e)},{default:b((()=>[v(r,{key:"detail",tab:"用户资料"}),v(r,{key:"logs",tab:"操作日志"})])),_:1},8,["activeKey"])])),default:b((()=>[y("div",T,["detail"===o.value?(d(),f(k,{key:0},m(10,(e=>y("div",{key:e},"这是用户"+j(s.value)+"资料Tab",1))),64)):O("",!0),"logs"===o.value?(d(),f(k,{key:1},m(10,(e=>y("div",{key:e},"这是用户"+j(s.value)+"操作日志Tab",1))),64)):O("",!0)])])),_:1},8,["title"])}}}))));export{F as default};
