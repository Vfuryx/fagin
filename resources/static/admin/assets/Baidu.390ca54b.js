var d=(e,s,t)=>new Promise((o,n)=>{var r=a=>{try{p(t.next(a))}catch(c){n(c)}},i=a=>{try{p(t.throw(a))}catch(c){n(c)}},p=a=>a.done?o(a.value):Promise.resolve(a.value).then(r,i);p((t=t.apply(e,s)).next())});import{u}from"./useScript.b2d818ab.js";import{_ as l}from"./index.6d3ed4e5.js";import{A as f,r as m,_ as h,B as w,D as B,X as _,$ as v,u as M}from"./vendor.3850bdb6.js";const b="https://api.map.baidu.com/getscript?v=3.0&ak=OaBvYmKX3pjF7YFUFeeBCeGdy9Zp7xB2&services=&t=20210201100830&s=1",g=f({name:"BaiduMap",props:{width:{type:String,default:"100%"},height:{type:String,default:"calc(100vh - 78px)"}},setup(){const e=m(null),{toPromise:s}=u({src:b});function t(){return d(this,null,function*(){yield s(),yield v();const o=M(e);if(!o)return;const n=window.BMap,r=new n.Map(o),i=new n.Point(116.404,39.915);r.centerAndZoom(i,15),r.enableScrollWheelZoom(!0)})}return h(()=>{t()}),{wrapRef:e}}});function y(e,s,t,o,n,r){return w(),B("div",{ref:"wrapRef",style:_({height:e.height,width:e.width})},null,4)}var k=l(g,[["render",y]]);export{k as default};