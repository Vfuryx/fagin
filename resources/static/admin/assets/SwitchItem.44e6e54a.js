import{A as d,bQ as r,j as c,a0 as l,B as p,D as m,H as u,J as h,w as f,a4 as C,K as g}from"./vendor.3850bdb6.js";/* empty css                */import{_ as j,a as _,b}from"./index.6d3ed4e5.js";import{b as v}from"./index.c0f37e81.js";import"./index.71a54c7b.js";/* empty css               *//* empty css               */import"./index.2d2673c5.js";import"./index.da3ea590.js";import"./useWindowSizeFn.4c2ec928.js";import"./useContentViewHeight.34544381.js";/* empty css               */import"./useSortable.7a35deac.js";import"./lock.1a54c299.js";const w=d({name:"SwitchItem",components:{Switch:r},props:{event:{type:Number},disabled:{type:Boolean},title:{type:String},def:{type:Boolean}},setup(e){const{prefixCls:t}=_("setting-switch-item"),{t:n}=b(),s=c(()=>e.def?{checked:e.def}:{});function a(i){e.event&&v(e.event,i)}return{prefixCls:t,t:n,handleChange:a,getBindValue:s}}});function S(e,t,n,s,a,i){const o=l("Switch");return p(),m("div",{class:g(e.prefixCls)},[u("span",null,h(e.title),1),f(o,C(e.getBindValue,{onChange:e.handleChange,disabled:e.disabled,checkedChildren:e.t("layout.setting.on"),unCheckedChildren:e.t("layout.setting.off")}),null,16,["onChange","disabled","checkedChildren","unCheckedChildren"])],2)}var J=j(w,[["render",S],["__scopeId","data-v-440e72fd"]]);export{J as default};