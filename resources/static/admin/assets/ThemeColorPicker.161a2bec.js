import{_ as p,a as l}from"./index.6d3ed4e5.js";import{b as c}from"./index.c0f37e81.js";import{A as m,b_ as u,a0 as _,B as s,D as t,ab as f,ap as C,K as a,X as k,w as j}from"./vendor.3850bdb6.js";import"./index.71a54c7b.js";/* empty css               *//* empty css               */import"./index.2d2673c5.js";import"./index.da3ea590.js";import"./useWindowSizeFn.4c2ec928.js";import"./useContentViewHeight.34544381.js";/* empty css               */import"./useSortable.7a35deac.js";import"./lock.1a54c299.js";const h=m({name:"ThemeColorPicker",components:{CheckOutlined:u},props:{colorList:{type:Array,defualt:[]},event:{type:Number},def:{type:String}},setup(e){const{prefixCls:i}=l("setting-theme-picker");function r(o){e.event&&c(e.event,o)}return{prefixCls:i,handleClick:r}}}),v=["onClick"];function b(e,i,r,o,x,y){const d=_("CheckOutlined");return s(),t("div",{class:a(e.prefixCls)},[(s(!0),t(f,null,C(e.colorList||[],n=>(s(),t("span",{key:n,onClick:g=>e.handleClick(n),class:a([`${e.prefixCls}__item`,{[`${e.prefixCls}__item--active`]:e.def===n}]),style:k({background:n})},[j(d)],14,v))),128))],2)}var N=p(h,[["render",b]]);export{N as default};
