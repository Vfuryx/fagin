import{y as e}from"./index.ab59a04a.js";import{u as t}from"./useWindowSizeFn.df7b0c93.js";import{r as n,j as i,u as o}from"./vendor.5da97255.js";const a=Symbol();const r=n(0),s=n(0);function u(){return{headerHeightRef:r,footerHeightRef:s,setHeaderHeight:function(e){r.value=e},setFooterHeight:function(e){s.value=e}}}function h(){const u=n(window.innerHeight),h=n(window.innerHeight),c=i((()=>o(u)-o(r)-o(s)||0));t((()=>{u.value=window.innerHeight}),100,{immediate:!0}),e({contentHeight:c,setPageHeight:function(e){return t=this,n=null,i=function*(){h.value=e},new Promise(((e,o)=>{var a=e=>{try{s(i.next(e))}catch(t){o(t)}},r=e=>{try{s(i.throw(e))}catch(t){o(t)}},s=t=>t.done?e(t.value):Promise.resolve(t.value).then(a,r);s((i=i.apply(t,n)).next())}));var t,n,i},pageHeight:h},a,{native:!0})}export{u as a,h as u};