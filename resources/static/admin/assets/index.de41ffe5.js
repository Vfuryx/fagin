import{w as m,_ as f}from"./index.6d3ed4e5.js";import{A as a,r,cr as k,_ as C,B as c,D as v,a7 as x,a0 as d,a1 as O,a6 as u,w as j,H as w,J as B}from"./vendor.3850bdb6.js";import{P as S}from"./index.a09675a6.js";/* empty css               *//* empty css                */import"./useWindowSizeFn.4c2ec928.js";import"./useContentViewHeight.34544381.js";const g=a({emits:["mounted","clickOutside"],setup(e,{emit:t}){const n=r(null);return k(n,()=>{t("clickOutside")}),C(()=>{t("mounted")}),(o,l)=>(c(),v("div",{ref:(s,i)=>{i.wrap=s,n.value=s}},[x(o.$slots,"default")],512))}}),$=m(g);const h=a({components:{ClickOutSide:$,PageWrapper:S},setup(){const e=r("Click");function t(){e.value="Click Out Side"}function n(){e.value="Click Inner"}return{innerClick:n,handleClickOutside:t,text:e}}});function E(e,t,n,o,l,s){const i=d("ClickOutSide"),p=d("PageWrapper");return c(),O(p,{title:"\u70B9\u5185\u5916\u90E8\u89E6\u53D1\u4E8B\u4EF6"},{default:u(()=>[j(i,{onClickOutside:e.handleClickOutside,class:"flex justify-center"},{default:u(()=>[w("div",{onClick:t[0]||(t[0]=(..._)=>e.innerClick&&e.innerClick(..._)),class:"demo-box"},B(e.text),1)]),_:1},8,["onClickOutside"])]),_:1})}var F=f(h,[["render",E],["__scopeId","data-v-4b8dd7fd"]]);export{F as default};