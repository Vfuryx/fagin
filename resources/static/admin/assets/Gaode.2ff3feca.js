var c=(e,o,t)=>new Promise((n,r)=>{var p=a=>{try{s(t.next(a))}catch(i){r(i)}},d=a=>{try{s(t.throw(a))}catch(i){r(i)}},s=a=>a.done?n(a.value):Promise.resolve(a.value).then(p,d);s((t=t.apply(e,o)).next())});import{u as f}from"./useScript.b2d818ab.js";import{_ as u}from"./index.6d3ed4e5.js";import{A as m,r as l,_ as w,B as h,D as _,X as b,$ as M,u as v}from"./vendor.3850bdb6.js";const y="https://webapi.amap.com/maps?v=2.0&key=d7bb98e7185300250dd5f918c12f484b",A=m({name:"AMap",props:{width:{type:String,default:"100%"},height:{type:String,default:"calc(100vh - 78px)"}},setup(){const e=l(null),{toPromise:o}=f({src:y});function t(){return c(this,null,function*(){yield o(),yield M();const n=v(e);if(!n)return;const r=window.AMap;new r.Map(n,{zoom:11,center:[116.397428,39.90923],viewMode:"3D"})})}return w(()=>{t()}),{wrapRef:e}}});function g(e,o,t,n,r,p){return h(),_("div",{ref:"wrapRef",style:b({height:e.height,width:e.width})},null,4)}var j=u(A,[["render",g]]);export{j as default};
