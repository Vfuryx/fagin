import{e as t,k as e,a,w as n}from"./index.7037d70a.js";import{P as s,Q as u,al as o,j as r,b8 as l,$ as i,t as c,u as d,ah as v,r as f,af as p,D as m,aS as C,a0 as y,F as S,a1 as b,a5 as g,ad as x,L as h,a4 as w,aD as z,ar as B,a6 as F,aC as A,aR as j,w as $}from"./vendor.5da97255.js";function k(t,e="value",a="change",n){const f=v(),p=null==f?void 0:f.emit,m=s({value:t[e]}),C=u(m);o((()=>{m.value=t[e]}));return[r({get:()=>m.value,set(t){l(t,C.value)||(m.value=t,i((()=>{null==p||p(a,t,...c(d(n))||[])})))}}),t=>{m.value=t},C]}var T=m({name:"CountButton",components:{Button:C},props:{value:{type:[Object,Number,String,Array]},count:{type:Number,default:60},beforeStartFunc:{type:Function,default:null}},setup(a){const n=f(!1),{currentCount:s,isStart:u,start:l,reset:i}=function(t){const e=f(t),a=f(!1);let n;function s(){n&&window.clearInterval(n)}function u(){a.value=!1,s(),n=null}function o(){d(a)||n||(a.value=!0,n=setInterval((()=>{1===d(e)?(u(),e.value=t):e.value-=1}),1e3))}function r(){e.value=t,u()}return p((()=>{r()})),{start:o,reset:r,restart:function(){r(),o()},clear:s,stop:u,currentCount:e,isStart:a}}(a.count),{t:c}=t(),v=r((()=>d(u)?c("component.countdown.sendText",[d(s)]):c("component.countdown.normalText")));return o((()=>{void 0===a.value&&i()})),{handleStart:function(){return t=this,s=null,u=function*(){const{beforeStartFunc:t}=a;if(t&&e(t)){n.value=!0;try{(yield t())&&l()}finally{n.value=!1}}else l()},new Promise(((e,a)=>{var n=t=>{try{r(u.next(t))}catch(e){a(e)}},o=t=>{try{r(u.throw(t))}catch(e){a(e)}},r=t=>t.done?e(t.value):Promise.resolve(t.value).then(n,o);r((u=u.apply(t,s)).next())}));var t,s,u},currentCount:s,loading:n,getButtonText:v,isStart:u}}});T.render=function(t,e,a,n,s,u){const o=y("Button");return S(),b(o,w(t.$attrs,{disabled:t.isStart,onClick:t.handleStart,loading:t.loading}),{default:g((()=>[x(h(t.getButtonText),1)])),_:1},16,["disabled","onClick","loading"])};var D=m({name:"CountDownInput",components:{CountButton:T},inheritAttrs:!1,props:{value:{type:String},size:{type:String,validator:t=>["default","large","small"].includes(t)},count:{type:Number,default:60},sendCodeApi:{type:Function,default:null}},setup(t){const{prefixCls:e}=a("countdown-input"),[n]=k(t);return{prefixCls:e,state:n}}});D.render=function(t,e,a,n,s,u){const o=y("CountButton"),r=y("a-input");return S(),b(r,w(t.$attrs,{class:t.prefixCls,size:t.size,value:t.state}),z({addonAfter:g((()=>[$(o,{size:t.size,count:t.count,value:t.state,beforeStartFunc:t.sendCodeApi},null,8,["size","count","value","beforeStartFunc"])])),_:2},[B(Object.keys(t.$slots).filter((t=>"addonAfter"!==t)),(e=>({name:e,fn:g((a=>[F(t.$slots,e,A(j(a||{})))]))})))]),1040,["class","size","value"])};const I=n(D);n(T);export{I as C,k as u};
