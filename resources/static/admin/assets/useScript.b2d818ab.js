import{r as t,_ as i,aj as f}from"./vendor.3850bdb6.js";function m(a){const s=t(!1),o=t(!1),r=t(!1);let e;const n=new Promise((u,c)=>{i(()=>{e=document.createElement("script"),e.type="text/javascript",e.onload=function(){s.value=!1,r.value=!0,o.value=!1,u("")},e.onerror=function(l){s.value=!1,r.value=!1,o.value=!0,c(l)},e.src=a.src,document.head.appendChild(e)})});return f(()=>{e&&e.remove()}),{isLoading:s,error:o,success:r,toPromise:()=>n}}export{m as u};
