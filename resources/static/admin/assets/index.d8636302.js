var j=Object.defineProperty,N=Object.defineProperties;var A=Object.getOwnPropertyDescriptors;var D=Object.getOwnPropertySymbols;var S=Object.prototype.hasOwnProperty,M=Object.prototype.propertyIsEnumerable;var F=(e,n,a)=>n in e?j(e,n,{enumerable:!0,configurable:!0,writable:!0,value:a}):e[n]=a,b=(e,n)=>{for(var a in n||(n={}))S.call(n,a)&&F(e,a,n[a]);if(D)for(var a of D(n))M.call(n,a)&&F(e,a,n[a]);return e},C=(e,n)=>N(e,A(n));import{P as y}from"./index.a09675a6.js";import{p as u,b as H,aV as I,aE as P,au as w,_ as E,ad as W,k as Y,F as k,w as R,ag as V}from"./index.6d3ed4e5.js";import{A as O,r as U,c6 as x,S as q,B as $,D as z,J,P as G,W as K,a0 as h,a1 as L,a6 as f,w as i,H as d}from"./vendor.3850bdb6.js";/* empty css               *//* empty css                */import"./useWindowSizeFn.4c2ec928.js";import"./useContentViewHeight.34544381.js";const p=1e3,v=p*60,_=v*60,B=_*24,Q=O({name:"Time",props:{value:u.oneOfType([u.number,u.instanceOf(Date),u.string]).isRequired,step:u.number.def(60),mode:u.oneOf(["date","datetime","relative"]).def("relative")},setup(e){const n=U(""),{t:a}=H();x(m,e.step*p),q(()=>e.value,()=>{m()},{immediate:!0});function T(){const{value:t}=e;let s=0;if(W(t)){const l=t.toString().length>10?t:t*1e3;s=new Date(l).getTime()}else Y(t)?s=new Date(t).getTime():k(t)&&(s=t.getTime());return s}function m(){const{mode:t,value:s}=e,l=T();t==="relative"?n.value=g(l):t==="datetime"?n.value=I(s):t==="date"&&(n.value=P(s))}function g(t){const s=new Date().getTime(),l=w(t).isBefore(s);let o=s-t;l||(o=-o);let r="",c=a(l?"component.time.before":"component.time.after");return o<p?r=a("component.time.just"):o<v?r=parseInt(o/p)+a("component.time.seconds")+c:o>=v&&o<_?r=Math.floor(o/v)+a("component.time.minutes")+c:o>=_&&o<B?r=Math.floor(o/_)+a("component.time.hours")+c:o>=B&&o<262386e4?r=Math.floor(o/B)+a("component.time.days")+c:o>=262386e4&&o<=3156786e4&&l?r=w(t).format("MM-DD-HH-mm"):r=w(t).format("YYYY"),r}return{date:n}}});function X(e,n,a,T,m,g){return $(),z("span",null,J(e.date),1)}var Z=E(Q,[["render",X]]);const ee=R(Z),te=O({components:{PageWrapper:y,Time:ee,CollapseContainer:V},setup(){const e=new Date().getTime(),n=G({time1:e-60*3*1e3,time2:e-86400*3*1e3});return C(b({},K(n)),{now:e})}}),ne=d("br",null,null,-1),ae=d("br",null,null,-1),oe=d("br",null,null,-1),se=d("br",null,null,-1);function ie(e,n,a,T,m,g){const t=h("Time"),s=h("CollapseContainer"),l=h("PageWrapper");return $(),L(l,{title:"\u65F6\u95F4\u7EC4\u4EF6\u793A\u4F8B"},{default:f(()=>[i(s,{title:"\u57FA\u7840\u793A\u4F8B"},{default:f(()=>[i(t,{value:e.time1},null,8,["value"]),ne,i(t,{value:e.time2},null,8,["value"])]),_:1}),i(s,{title:"\u5B9A\u65F6\u66F4\u65B0",class:"my-4"},{default:f(()=>[i(t,{value:e.now,step:1},null,8,["value"]),ae,i(t,{value:e.now,step:5},null,8,["value"])]),_:1}),i(s,{title:"\u5B9A\u65F6\u66F4\u65B0"},{default:f(()=>[i(t,{value:e.now,mode:"date"},null,8,["value"]),oe,i(t,{value:e.now,mode:"datetime"},null,8,["value"]),se,i(t,{value:e.now},null,8,["value"])]),_:1})]),_:1})}var ve=E(te,[["render",ie]]);export{ve as default};
