import{D as e,bH as t,j as n,a0 as i,F as s,G as a,K as d,L as o,w as r,a4 as l,N as p}from"./vendor.5da97255.js";/* empty css              */import{a as f,e as c}from"./index.7037d70a.js";import{b as m}from"./index.92931596.js";import"./index.18992b80.js";/* empty css              *//* empty css              */import"./index.f3922401.js";import"./index.88a0ffff.js";import"./useWindowSizeFn.df7b0c93.js";import"./useContentViewHeight.e053a8fa.js";/* empty css              */import"./useSortable.6c5043ea.js";import"./lock.6275374b.js";var h=e({name:"SwitchItem",components:{Switch:t},props:{event:{type:Number},disabled:{type:Boolean},title:{type:String},def:{type:Boolean}},setup(e){const{prefixCls:t}=f("setting-switch-item"),{t:i}=c();return{prefixCls:t,t:i,handleChange:function(t){e.event&&m(e.event,t)},getBindValue:n((()=>e.def?{checked:e.def}:{}))}}});h.render=function(e,t,n,f,c,m){const h=i("Switch");return s(),a("div",{class:p(e.prefixCls)},[d("span",null,o(e.title),1),r(h,l(e.getBindValue,{onChange:e.handleChange,disabled:e.disabled,checkedChildren:e.t("layout.setting.on"),unCheckedChildren:e.t("layout.setting.off")}),null,16,["onChange","disabled","checkedChildren","unCheckedChildren"])],2)},h.__scopeId="data-v-440e72fd";export{h as default};
