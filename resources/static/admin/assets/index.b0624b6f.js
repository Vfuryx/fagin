import{D as e,r as a,j as s,u as t,F as n,G as i,w as r,a5 as o,K as f,N as c,X as l,aO as u,A as d,B as m}from"./vendor.5da97255.js";import{u as p}from"./useWindowSizeFn.df7b0c93.js";import{p as g,a as v}from"./index.ab59a04a.js";import{a as h}from"./useContentViewHeight.ef6fc368.js";d("data-v-179381bf");const w=["src"];m();var j=e({props:{frameSrc:g.string.def("")},setup(e){const d=a(!0),m=a(50),g=a(window.innerHeight),j=a(),{headerHeightRef:x}=h(),{prefixCls:H}=v("iframe-page");p(b,150,{immediate:!0});const _=s((()=>({height:`${t(g)}px`})));function b(){const e=t(j);if(!e)return;const a=x.value;m.value=a,g.value=window.innerHeight-a;const s=document.documentElement.clientHeight-a;e.style.height=`${s}px`}function y(){d.value=!1,b()}return(a,s)=>(n(),i("div",{class:c(t(H)),style:l(t(_))},[r(t(u),{spinning:d.value,size:"large",style:l(t(_))},{default:o((()=>[f("iframe",{src:e.frameSrc,class:c(`${t(H)}__main`),ref:(e,a)=>{a.frameRef=e,j.value=e},onLoad:y},null,42,w)])),_:1},8,["spinning","style"])],6))}});j.__scopeId="data-v-179381bf";export{j as default};
