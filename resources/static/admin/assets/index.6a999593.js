import{A as b,bK as O,j as L,a0 as _,B as C,D as M,w as a,a6 as x,a7 as q,ab as J,ap as ee,a1 as N,aD as te,a4 as R,aC as ne,K as ae,P as se,r as j,am as oe,u as f,bW as ie,ae as B,bY as re,bZ as le,aQ as ce}from"./vendor.3850bdb6.js";/* empty css                */import{p as ue,a as me,k as de,_ as T,aZ as ge}from"./index.6d3ed4e5.js";import{P as fe}from"./index.a09675a6.js";/* empty css               *//* empty css                */import"./useWindowSizeFn.4c2ec928.js";import"./useContentViewHeight.34544381.js";const pe=b({name:"ImagePreview",components:{Image:O,PreviewGroup:O.PreviewGroup},props:{functional:ue.bool,imageList:{type:Array}},setup(n){const{prefixCls:u}=me("image-preview"),r=L(()=>{const{imageList:i}=n;return i?i.map(t=>de(t)?{src:t,placeholder:!1}:t):[]});return{prefixCls:u,getImageList:r}}});function ve(n,u,r,i,t,h){const d=_("Image"),p=_("PreviewGroup");return C(),M("div",{class:ae(n.prefixCls)},[a(p,null,{default:x(()=>[!n.imageList||n.$slots.default?q(n.$slots,"default",{key:0}):(C(!0),M(J,{key:1},ee(n.getImageList,l=>(C(),N(d,ne(R({key:l.src},l)),te({_:2},[l.placeholder?{name:"placeholder",fn:x(()=>[a(d,R(l,{src:l.placeholder,preview:!1}),null,16,["src"])])}:void 0]),1040))),128))]),_:3})],2)}var _e=T(pe,[["render",ve]]),he="/static/admin/assets/resume.8f27866b.svg",we="/static/admin/assets/p-rotate.cb8bbfc3.svg",Ie="/static/admin/assets/scale.44abde22.svg",Le="/static/admin/assets/unscale.c552f416.svg",Ce="/static/admin/assets/unrotate.ef6a2daf.svg",g;(function(n){n[n.LOADING=0]="LOADING",n[n.DONE=1]="DONE",n[n.FAIL=2]="FAIL"})(g||(g={}));const xe={show:{type:Boolean,default:!1},imageList:{type:[Array],default:null},index:{type:Number,default:0},scaleStep:{type:Number},defaultWidth:{type:Number},maskClosable:{type:Boolean},rememberState:{type:Boolean}},o="img-preview";var $e=b({name:"ImagePreview",props:xe,emits:["img-load","img-error"],setup(n,{expose:u,emit:r}){const i=new Map,t=se({currentUrl:"",imgScale:1,imgRotate:0,imgTop:0,imgLeft:0,status:g.LOADING,currentIndex:0,moveX:0,moveY:0,show:n.show}),h=j(null),d=j(null);function p(){G();const{index:e,imageList:s}=n;if(!s||!s.length)throw new Error("imageList is undefined");t.currentIndex=e,F(s[e])}function l(){t.imgScale=1,t.imgRotate=0,t.imgTop=0,t.imgLeft=0}function G(){const e=f(h);!e||(e.onmousewheel=$,document.body.addEventListener("DOMMouseScroll",$),document.ondragstart=function(){return!1})}const w=L(()=>{var s;const e=(s=n==null?void 0:n.scaleStep)!=null?s:0;return(e!=null?e:0>0)?e/100:t.imgScale/10});function $(e){e=e||window.event,e.delta=e.wheelDelta||-e.detail,e.preventDefault(),e.delta>0&&I(w.value),e.delta<0&&I(-w.value)}function I(e){t.imgScale<=.2&&e<0||(t.imgScale+=e)}function k(e){t.imgRotate+=e}function X(){const e=f(d);!e||(e.onmousemove=null)}function F(e){t.status=g.LOADING;const s=new Image;s.src=e,s.onload=c=>{if(t.currentUrl!==e){const m=c.composedPath();if(n.rememberState){i.set(t.currentUrl,{scale:t.imgScale,top:t.imgTop,left:t.imgLeft,rotate:t.imgRotate});const v=i.get(e);v?(t.imgScale=v.scale,t.imgTop=v.top,t.imgRotate=v.rotate,t.imgLeft=v.left):(l(),n.defaultWidth&&(t.imgScale=n.defaultWidth/m[0].naturalWidth))}else n.defaultWidth&&(t.imgScale=n.defaultWidth/m[0].naturalWidth);m&&r("img-load",{index:t.currentIndex,dom:m[0],url:e})}t.currentUrl=e,t.status=g.DONE},s.onerror=c=>{const m=c.composedPath();m&&r("img-error",{index:t.currentIndex,dom:m[0],url:e}),t.status=g.FAIL}}function P(e){e&&e.stopPropagation(),D()}function D(){t.show=!1,document.body.removeEventListener("DOMMouseScroll",$),document.ondragstart=null}function A(){l()}u({resume:A,close:D,prev:S.bind(null,"left"),next:S.bind(null,"right"),setScale:e=>{e>0&&e<=10&&(t.imgScale=e)},setRotate:e=>{t.imgRotate=e}});function S(e){const{currentIndex:s}=t,{imageList:c}=n;e==="left"&&(t.currentIndex--,s<=0&&(t.currentIndex=c.length-1)),e==="right"&&(t.currentIndex++,s>=c.length-1&&(t.currentIndex=0)),F(c[t.currentIndex])}function U(e){e=e||window.event,t.moveX=e.clientX,t.moveY=e.clientY;const s=f(d);s&&(s.onmousemove=z)}function z(e){e=e||window.event,e.preventDefault();const s=e.clientX-t.moveX,c=e.clientY-t.moveY;t.imgLeft+=s,t.imgTop+=c,t.moveX=e.clientX,t.moveY=e.clientY}const V=L(()=>{const{imgScale:e,imgRotate:s,imgTop:c,imgLeft:m}=t;return{transform:`scale(${e}) rotate(${s}deg)`,marginTop:`${c}px`,marginLeft:`${m}px`,maxWidth:n.defaultWidth?"unset":"100%"}}),E=L(()=>{const{imageList:e}=n;return e.length>1});oe(()=>{n.show&&p(),n.imageList&&l()});const K=e=>{n.maskClosable&&e.target&&e.target.classList.contains(`${o}-content`)&&P(e)},Z=()=>a("div",{class:`${o}__close`,onClick:P},[a(ie,{class:`${o}__close-icon`},null)]),H=()=>{if(!f(E))return null;const{currentIndex:e}=t,{imageList:s}=n;return a("div",{class:`${o}__index`},[e+1,B(" / "),s.length])},Q=()=>a("div",{class:`${o}__controller`},[a("div",{class:`${o}__controller-item`,onClick:()=>I(-w.value)},[a("img",{src:Le},null)]),a("div",{class:`${o}__controller-item`,onClick:()=>I(w.value)},[a("img",{src:Ie},null)]),a("div",{class:`${o}__controller-item`,onClick:A},[a("img",{src:he},null)]),a("div",{class:`${o}__controller-item`,onClick:()=>k(-90)},[a("img",{src:Ce},null)]),a("div",{class:`${o}__controller-item`,onClick:()=>k(90)},[a("img",{src:we},null)])]),W=e=>f(E)?a("div",{class:[`${o}__arrow`,e],onClick:()=>S(e)},[e==="left"?a(re,null,null):a(le,null,null)]):null;return()=>t.show&&a("div",{class:o,ref:h,onMouseup:X,onClick:K},[a("div",{class:`${o}-content`},[a("img",{style:f(V),class:[`${o}-image`,t.status===g.DONE?"":"hidden"],ref:d,src:t.currentUrl,onMousedown:U},null),Z(),H(),Q(),W("left"),W("right")])])}});let y=null;function Se(n){var i;if(!ge)return;const u={},r=document.createElement("div");return Object.assign(u,{show:!0,index:0,scaleStep:100},n),y=a($e,u),ce(y,r),document.body.appendChild(r),(i=y.component)==null?void 0:i.exposed}const Y=["https://picsum.photos/id/66/346/216","https://picsum.photos/id/67/346/216","https://picsum.photos/id/68/346/216"],be=b({components:{PageWrapper:fe,ImagePreview:_e},setup(){function n(){Se({imageList:Y,defaultWidth:700,rememberState:!0,onImgLoad:({index:r,url:i,dom:t})=>{console.log(`\u7B2C${r+1}\u5F20\u56FE\u7247\u5DF2\u52A0\u8F7D\uFF0CURL\u4E3A\uFF1A${i}`,t)}})}return{imgList:Y,openImg:n}}}),ye=B("\u65E0\u9884\u89C8\u56FE");function ke(n,u,r,i,t,h){const d=_("ImagePreview"),p=_("a-button"),l=_("PageWrapper");return C(),N(l,{title:"\u56FE\u7247\u9884\u89C8\u793A\u4F8B"},{default:x(()=>[a(d,{imageList:n.imgList},null,8,["imageList"]),a(p,{onClick:n.openImg,type:"primary"},{default:x(()=>[ye]),_:1},8,["onClick"])]),_:1})}var Ne=T(be,[["render",ke]]);export{Ne as default};