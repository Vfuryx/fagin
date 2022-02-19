var m=(t,e,n)=>new Promise((s,r)=>{var l=a=>{try{u(n.next(a))}catch(c){r(c)}},o=a=>{try{u(n.throw(a))}catch(c){r(c)}},u=a=>a.done?s(a.value):Promise.resolve(a.value).then(l,o);u((n=n.apply(t,e)).next())});import{_ as C,j as I,b as h,a as z,w as S}from"./index.6d3ed4e5.js";import{ai as F,P as A,Q as T,am as g,j as w,ba as j,$ as k,t as N,u as d,r as p,ag as D,A as y,aT as P,a0 as f,B as _,a1 as $,a6 as v,ae as O,J as R,a4 as b,aD as E,ap as V,w as q,a7 as J,aC as L,aS as Q}from"./vendor.3850bdb6.js";function U(t,e="value",n="change",s){const r=F(),l=r==null?void 0:r.emit,o=A({value:t[e]}),u=T(o),a=i=>{o.value=i};return g(()=>{o.value=t[e]}),[w({get(){return o.value},set(i){j(i,u.value)||(o.value=i,k(()=>{l==null||l(n,i,...N(d(s))||[])}))}}),a,u]}function x(t){const e=p(t),n=p(!1);let s;function r(){s&&window.clearInterval(s)}function l(){n.value=!1,r(),s=null}function o(){d(n)||!!s||(n.value=!0,s=setInterval(()=>{d(e)===1?(l(),e.value=t):e.value-=1},1e3))}function u(){e.value=t,l()}function a(){u(),o()}return D(()=>{u()}),{start:o,reset:u,restart:a,clear:r,stop:l,currentCount:e,isStart:n}}const G={value:{type:[Object,Number,String,Array]},count:{type:Number,default:60},beforeStartFunc:{type:Function,default:null}},H=y({name:"CountButton",components:{Button:P},props:G,setup(t){const e=p(!1),{currentCount:n,isStart:s,start:r,reset:l}=x(t.count),{t:o}=h(),u=w(()=>d(s)?o("component.countdown.sendText",[d(n)]):o("component.countdown.normalText"));g(()=>{t.value===void 0&&l()});function a(){return m(this,null,function*(){const{beforeStartFunc:c}=t;if(c&&I(c)){e.value=!0;try{(yield c())&&r()}finally{e.value=!1}}else r()})}return{handleStart:a,currentCount:n,loading:e,getButtonText:u,isStart:s}}});function K(t,e,n,s,r,l){const o=f("Button");return _(),$(o,b(t.$attrs,{disabled:t.isStart,onClick:t.handleStart,loading:t.loading}),{default:v(()=>[O(R(t.getButtonText),1)]),_:1},16,["disabled","onClick","loading"])}var B=C(H,[["render",K]]);const M={value:{type:String},size:{type:String,validator:t=>["default","large","small"].includes(t)},count:{type:Number,default:60},sendCodeApi:{type:Function,default:null}},W=y({name:"CountDownInput",components:{CountButton:B},inheritAttrs:!1,props:M,setup(t){const{prefixCls:e}=z("countdown-input"),[n]=U(t);return{prefixCls:e,state:n}}});function X(t,e,n,s,r,l){const o=f("CountButton"),u=f("a-input");return _(),$(u,b(t.$attrs,{class:t.prefixCls,size:t.size,value:t.state}),E({addonAfter:v(()=>[q(o,{size:t.size,count:t.count,value:t.state,beforeStartFunc:t.sendCodeApi},null,8,["size","count","value","beforeStartFunc"])]),_:2},[V(Object.keys(t.$slots).filter(a=>a!=="addonAfter"),a=>({name:a,fn:v(c=>[J(t.$slots,a,L(Q(c||{})))])}))]),1040,["class","size","value"])}var Y=C(W,[["render",X]]);const nt=S(Y);S(B);export{nt as C,U as u};
