var e=Object.defineProperty,t=Object.defineProperties,r=Object.getOwnPropertyDescriptors,a=Object.getOwnPropertySymbols,n=Object.prototype.hasOwnProperty,l=Object.prototype.propertyIsEnumerable,o=(t,r,a)=>r in t?e(t,r,{enumerable:!0,configurable:!0,writable:!0,value:a}):t[r]=a,i=(e,t)=>{for(var r in t||(t={}))n.call(t,r)&&o(e,r,t[r]);if(a)for(var r of a(t))l.call(t,r)&&o(e,r,t[r]);return e},u=(e,a)=>t(e,r(a));import{D as s,cb as c,r as d,a as m,al as p,a0 as v,F as h,G as f,w as b,a5 as y,a1 as g,ac as C,L as O,ad as j,N as w}from"./vendor.5da97255.js";/* empty css              */import{_ as x,p as k,a as I,Q as P,l as B,e as M,j as R,R as S,L as $,t as A,aM as D}from"./index.ab59a04a.js";var L=s({name:"LayoutBreadcrumb",components:{Icon:x,[c.name]:c},props:{theme:k.oneOf(["dark","light"])},setup(){const e=d([]),{currentRoute:t}=m(),{prefixCls:r}=I("layout-breadcrumb"),{getShowBreadCrumbIcon:a}=P(),n=B(),{t:l}=M();function o(e,t){const r=[];return e.forEach((e=>{var a,n;t.includes(e.path)&&r.push(u(i({},e),{name:(null==(a=e.meta)?void 0:a.title)||e.name})),(null==(n=e.children)?void 0:n.length)&&r.push(...o(e.children,t))})),r}return p((()=>{return r=this,a=null,n=function*(){var r,a,n;if(t.value.name===S)return;const l=yield $(),s=t.value.matched,c=null==s?void 0:s[s.length-1];let d=t.value.path;c&&(null==(r=null==c?void 0:c.meta)?void 0:r.currentActiveMenu)&&(d=c.meta.currentActiveMenu);const m=A(l,d),p=o(l.filter((e=>e.path===m[0])),m);if(!p||0===p.length)return;const v=D(p,(e=>{const{meta:t,name:r}=e;if(!t)return!!r;const{title:a,hideBreadcrumb:n,hideMenu:l}=t;return!(!a||n||l)})).filter((e=>{var t,r;return!(null==(t=e.meta)?void 0:t.hideBreadcrumb)||!(null==(r=e.meta)?void 0:r.hideMenu)}));(null==(a=t.value.meta)?void 0:a.currentActiveMenu)&&v.push(u(i({},t.value),{name:(null==(n=t.value.meta)?void 0:n.title)||t.value.name})),e.value=v},new Promise(((e,t)=>{var l=e=>{try{i(n.next(e))}catch(r){t(r)}},o=e=>{try{i(n.throw(e))}catch(r){t(r)}},i=t=>t.done?e(t.value):Promise.resolve(t.value).then(l,o);i((n=n.apply(r,a)).next())}));var r,a,n})),{routes:e,t:l,prefixCls:r,getIcon:function(e){var t;return e.icon||(null==(t=e.meta)?void 0:t.icon)},getShowBreadCrumbIcon:a,handleClick:function(e,t,r){null==r||r.preventDefault();const{children:a,redirect:l,meta:o}=e;if(!(null==a?void 0:a.length)||l){if(!(null==o?void 0:o.carryParam))if(l&&R(l))n(l);else{let e="";if(1===t.length)e=t[0];else{e=`${t.slice(1).pop()||""}`}e=/^\//.test(e)?e:`/${e}`,n(e)}}else null==r||r.stopPropagation()},hasRedirect:function(e,t){return e.indexOf(t)!==e.length-1}}}});const _={key:1};L.render=function(e,t,r,a,n,l){const o=v("Icon"),i=v("router-link"),u=v("a-breadcrumb");return h(),f("div",{class:w([e.prefixCls,`${e.prefixCls}--${e.theme}`])},[b(u,{routes:e.routes},{itemRender:y((({route:t,routes:r,paths:a})=>[e.getShowBreadCrumbIcon&&e.getIcon(t)?(h(),g(o,{key:0,icon:e.getIcon(t)},null,8,["icon"])):C("",!0),e.hasRedirect(r,t)?(h(),g(i,{key:2,to:"",onClick:r=>e.handleClick(t,a,r)},{default:y((()=>[j(O(e.t(t.name||t.meta.title)),1)])),_:2},1032,["onClick"])):(h(),f("span",_,O(e.t(t.name||t.meta.title)),1))])),_:1},8,["routes"])],2)};export{L as default};
