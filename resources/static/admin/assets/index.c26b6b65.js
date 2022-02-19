var K=Object.defineProperty,Q=Object.defineProperties;var U=Object.getOwnPropertyDescriptors;var j=Object.getOwnPropertySymbols;var Y=Object.prototype.hasOwnProperty,ee=Object.prototype.propertyIsEnumerable;var C=(t,r,d)=>r in t?K(t,r,{enumerable:!0,configurable:!0,writable:!0,value:d}):t[r]=d,p=(t,r)=>{for(var d in r||(r={}))Y.call(r,d)&&C(t,d,r[d]);if(j)for(var d of j(r))ee.call(r,d)&&C(t,d,r[d]);return t},V=(t,r)=>Q(t,U(r));import{b as X,ai as te,af as k,s as F,a$ as W,w as A}from"./index.6d3ed4e5.js";import{A as Z,P as H,r as $,j as _,S as q,am as ne,w as h,u,b_ as se,ch as ae,a4 as ie}from"./vendor.3850bdb6.js";const{t:z}=X(),G={value:{type:Boolean,default:!1},isSlot:{type:Boolean,default:!1},text:{type:[String],default:z("component.verify.dragText")},successText:{type:[String],default:z("component.verify.successText")},height:{type:[Number,String],default:40},width:{type:[Number,String],default:220},circle:{type:Boolean,default:!1},wrapStyle:{type:Object,default:{}},contentStyle:{type:Object,default:{}},barStyle:{type:Object,default:{}},actionStyle:{type:Object,default:{}}},oe=V(p({},G),{src:{type:String},imgWidth:{type:Number,default:260},imgWrapStyle:{type:Object,default:{}},minDegree:{type:Number,default:90},maxDegree:{type:Number,default:270},diffDegree:{type:Number,default:20}});var J=Z({name:"BaseDargVerify",props:G,emits:["success","update:value","change","start","move","end"],setup(t,{emit:r,slots:d,expose:P}){const n=H({isMoving:!1,isPassing:!1,moveDistance:0,toLeft:!1,startTime:0,endTime:0}),s=$(null),w=$(null),b=$(null),T=$(null);te({el:document,name:"mouseup",listener:()=>{n.isMoving&&y()}});const O=_(()=>{const{height:e,actionStyle:i}=t,c=`${parseInt(e)}px`;return p({left:0,width:c,height:c},i)}),N=_(()=>{const{height:e,width:i,circle:c,wrapStyle:a}=t,l=parseInt(e),m=`${parseInt(i)}px`;return p({width:m,height:`${l}px`,lineHeight:`${l}px`,borderRadius:c?l/2+"px":0},a)}),E=_(()=>{const{height:e,circle:i,barStyle:c}=t,a=parseInt(e);return p({height:`${a}px`,borderRadius:i?a/2+"px 0 0 "+a/2+"px":0},c)}),I=_(()=>{const{height:e,width:i,contentStyle:c}=t,a=`${parseInt(e)}px`,l=`${parseInt(i)}px`;return p({height:a,width:l},c)});q(()=>n.isPassing,e=>{if(e){const{startTime:i,endTime:c}=n,a=(c-i)/1e3;r("success",{isPassing:e,time:a.toFixed(1)}),r("update:value",e),r("change",e)}}),ne(()=>{n.isPassing=!!t.value});function D(e){return e.pageX||e.touches[0].pageX}function R(e){if(n.isPassing)return;const i=u(T);!i||(r("start",e),n.moveDistance=D(e)-parseInt(i.style.left.replace("px",""),10),n.startTime=new Date().getTime(),n.isMoving=!0)}function o(e){const i=parseInt(e.style.width),{width:c}=t,a=parseInt(c);return{offset:a-i-6,widthNum:a,actionWidth:i}}function f(e){const{isMoving:i,moveDistance:c}=n;if(i){const a=u(T),l=u(w);if(!a||!l)return;const{offset:m,widthNum:S,actionWidth:M}=o(a),v=D(e)-c;r("move",{event:e,moveDistance:c,moveX:v}),v>0&&v<=m?(a.style.left=`${v}px`,l.style.width=`${v+M/2}px`):v>m&&(a.style.left=`${S-M}px`,l.style.width=`${S-M/2}px`,t.isSlot||x())}}function g(e){const{isMoving:i,isPassing:c,moveDistance:a}=n;if(i&&!c){r("end",e);const l=u(T),m=u(w);if(!l||!m)return;const S=D(e)-a,{offset:M,widthNum:v,actionWidth:B}=o(l);S<M?t.isSlot?setTimeout(()=>{if(!t.value)y();else{const L=u(b);L&&(L.style.width=`${parseInt(m.style.width)}px`)}},0):y():(l.style.left=`${v-B}px`,m.style.width=`${v-B/2}px`,x()),n.isMoving=!1}}function x(){if(t.isSlot){y();return}n.endTime=new Date().getTime(),n.isPassing=!0,n.isMoving=!1}function y(){n.isMoving=!1,n.isPassing=!1,n.moveDistance=0,n.toLeft=!1,n.startTime=0,n.endTime=0;const e=u(T),i=u(w),c=u(b);!e||!i||!c||(n.toLeft=!0,F(()=>{n.toLeft=!1,e.style.left="0",i.style.width="0"},300),c.style.width=u(I).width)}return P({resume:y}),()=>{const e=()=>{const a=["darg-verify-bar"];return n.toLeft&&a.push("to-left"),h("div",{class:a,ref:w,style:u(E)},null)},i=()=>{const a=["darg-verify-content"],{isPassing:l}=n,{text:m,successText:S}=t;return l&&a.push("success"),h("div",{class:a,ref:b,style:u(I)},[k(d,"text",l)||(l?S:m)])},c=()=>{const a=["darg-verify-action"],{toLeft:l,isPassing:m}=n;return l&&a.push("to-left"),h("div",{class:a,onMousedown:R,onTouchstart:R,style:u(O),ref:T},[k(d,"actionIcon",m)||(m?h(se,{class:"darg-verify-action__icon"},null):h(ae,{class:"darg-verify-action__icon"},null))])};return h("div",{class:"darg-verify",ref:s,style:u(N),onMousemove:f,onTouchmove:f,onMouseleave:g,onMouseup:g,onTouchend:g},[e(),i(),c()])}}});var re=Z({name:"ImgRotateDragVerify",inheritAttrs:!1,props:oe,emits:["success","change","update:value"],setup(t,{emit:r,attrs:d,expose:P}){const n=$(null),s=H({showTip:!1,isPassing:!1,imgStyle:{},randomRotate:0,currentRotate:0,toOrigin:!1,startTime:0,endTime:0,draged:!1}),{t:w}=X();q(()=>s.isPassing,o=>{if(o){const{startTime:f,endTime:g}=s,x=(g-f)/1e3;r("success",{isPassing:o,time:x.toFixed(1)}),r("change",o),r("update:value",o)}});const b=_(()=>{const{imgWrapStyle:o,imgWidth:f}=t;return p({width:`${f}px`,height:`${f}px`},o)}),T=_(()=>{const{minDegree:o,maxDegree:f}=t;return o===f?Math.floor(1+Math.random()*1)/10+1:1});function O(){s.startTime=new Date().getTime()}function N(o){s.draged=!0;const{imgWidth:f,height:g,maxDegree:x}=t,{moveX:y}=o,e=Math.ceil(y/(f-parseInt(g))*x*u(T));s.currentRotate=e,s.imgStyle=W("transform",`rotateZ(${s.randomRotate-e}deg)`)}function E(){const{minDegree:o,maxDegree:f}=t,g=Math.floor(o+Math.random()*(f-o));s.randomRotate=g,s.imgStyle=W("transform",`rotateZ(${g}deg)`)}function I(){const{randomRotate:o,currentRotate:f}=s,{diffDegree:g}=t;Math.abs(o-f)>=(g||20)?(s.imgStyle=W("transform",`rotateZ(${o}deg)`),s.toOrigin=!0,F(()=>{s.toOrigin=!1,s.showTip=!0},300)):D(),s.showTip=!0}function D(){s.isPassing=!0,s.endTime=new Date().getTime()}function R(){s.showTip=!1;const o=u(n);!o||(s.isPassing=!1,o.resume(),E())}return P({resume:R}),()=>{const{src:o}=t,{toOrigin:f,isPassing:g,startTime:x,endTime:y}=s,e=[];f&&e.push("to-origin");const i=(y-x)/1e3;return h("div",{class:"ir-dv"},[h("div",{class:"ir-dv-img__wrap",style:u(b)},[h("img",{src:o,onLoad:E,width:parseInt(t.width),class:e,style:s.imgStyle,onClick:()=>{R()},alt:"verify"},null),s.showTip&&h("span",{class:["ir-dv-img__tip",s.isPassing?"success":"error"]},[s.isPassing?w("component.verify.time",{time:i.toFixed(1)}):w("component.verify.error")]),!s.showTip&&!s.draged&&h("span",{class:["ir-dv-img__tip","normal"]},[w("component.verify.redoTip")])]),h(J,ie({class:"ir-dv-drag__bar",onMove:N,onEnd:I,onStart:O,ref:n},p(p({},d),t),{value:g,isSlot:!0}),null)])}}});const ue=A(J),de=A(re);export{ue as B,de as R};
