var l=(e,r,t)=>new Promise((a,n)=>{var s=o=>{try{p(t.next(o))}catch(c){n(c)}},i=o=>{try{p(t.throw(o))}catch(c){n(c)}},p=o=>o.done?a(o.value):Promise.resolve(o.value).then(s,i);p((t=t.apply(e,r)).next())});import{u as d}from"./useScript.b2d818ab.js";import{_ as f}from"./index.6d3ed4e5.js";import{A as m,r as u,_ as g,B as w,D as h,X as _,$ as y,u as M}from"./vendor.3850bdb6.js";const S="https://maps.googleapis.com/maps/api/js?key=AIzaSyBQWrGwj4gAzKndcbwD5favT9K0wgty_0&signed_in=true",b=m({name:"GoogleMap",props:{width:{type:String,default:"100%"},height:{type:String,default:"calc(100vh - 78px)"}},setup(){const e=u(null),{toPromise:r}=d({src:S});function t(){return l(this,null,function*(){yield r(),yield y();const a=M(e);if(!a)return;const n=window.google,s={lat:116.404,lng:39.915},i=new n.maps.Map(a,{zoom:4,center:s});new n.maps.Marker({position:s,map:i,title:"Hello World!"})})}return g(()=>{t()}),{wrapRef:e}}});function j(e,r,t,a,n,s){return w(),h("div",{ref:"wrapRef",style:_({height:e.height,width:e.width})},null,4)}var z=f(b,[["render",j]]);export{z as default};