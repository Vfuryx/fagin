import{D as e,a as r,u as t,F as a,G as s}from"./vendor.5da97255.js";var p=e({setup(e){const{currentRoute:p,replace:o}=r(),{params:u,query:y}=t(p),{path:c,_redirect_type:d="path"}=u;Reflect.deleteProperty(u,"_redirect_type"),Reflect.deleteProperty(u,"path");const n=Array.isArray(c)?c.join("/"):c;return o("name"===d?{name:n,query:y,params:u}:{path:n.startsWith("/")?n:"/"+n,query:y}),(e,r)=>(a(),s("div"))}});export{p as default};