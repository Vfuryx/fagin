import{B as f,a as g}from"./index.ab15df17.js";import{_ as h}from"./index.6d3ed4e5.js";import{A as B,r as p,S as M,a0 as m,B as o,a1 as v,a6 as i,w as C,D as s,ad as _,ab as b,ap as y,J as F,a4 as D,ae as k}from"./vendor.3850bdb6.js";import"./useWindowSizeFn.4c2ec928.js";const w=B({components:{BasicModal:f},setup(){const e=p(!0),a=p(10),[u,{setModalProps:t,redoModalHeight:d}]=g();M(()=>a.value,()=>{d()});function l(r){r&&(e.value=!0,t({loading:!0,confirmLoading:!0}),setTimeout(()=>{a.value=Math.round(Math.random()*30+5),e.value=!1,t({loading:!1,confirmLoading:!1})},3e3))}function n(){a.value=Math.round(Math.random()*20+10)}return{register:u,loading:e,handleShow:l,lines:a,setLines:n}}}),A=k("\u70B9\u6211\u66F4\u65B0\u5185\u5BB9"),L={key:0,class:"empty-tips"},S={key:1};function V(e,a,u,t,d,l){const n=m("a-button"),r=m("BasicModal");return o(),v(r,D(e.$attrs,{destroyOnClose:"",onRegister:e.register,title:"Modal Title",helpMessage:["\u63D0\u793A1","\u63D0\u793A2"],onVisibleChange:e.handleShow}),{insertFooter:i(()=>[C(n,{type:"primary",danger:"",onClick:e.setLines,disabled:e.loading},{default:i(()=>[A]),_:1},8,["onClick","disabled"])]),default:i(()=>[e.loading?(o(),s("div",L,"\u52A0\u8F7D\u4E2D\uFF0C\u7A0D\u7B493\u79D2\u2026\u2026")):_("",!0),e.loading?_("",!0):(o(),s("ul",S,[(o(!0),s(b,null,y(e.lines,c=>(o(),s("li",{key:c},"\u52A0\u8F7D\u5B8C\u6210"+F(c)+"\uFF01",1))),128))]))]),_:1},16,["onRegister","onVisibleChange"])}var E=h(w,[["render",V],["__scopeId","data-v-47fa5808"]]);export{E as default};
