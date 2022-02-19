import{_ as h}from"./Tree.vue_vue&type=style&index=0&lang.d5823234.js";import{t as D}from"./data.5e96733f.js";import{_ as v,i as A}from"./index.6d3ed4e5.js";import{P as F}from"./index.a09675a6.js";import{A as K,r as N,a0 as c,B as S,a1 as g,a6 as u,H as p,w as o,ae as a,u as $}from"./vendor.3850bdb6.js";import"./useContextMenu.f3a8e755.js";/* empty css               *//* empty css               *//* empty css                */import"./useWindowSizeFn.4c2ec928.js";import"./useContentViewHeight.34544381.js";const j=K({components:{BasicTree:h,PageWrapper:F},setup(){const t=N(null),{createMessage:e}=A();function l(){const n=$(t);if(!n)throw new Error("tree is null!");return n}function f(n){l().filterByLevel(n)}function m(){l().setCheckedKeys(["0-0"])}function C(){const n=l().getCheckedKeys();e.success(JSON.stringify(n))}function s(){l().setSelectedKeys(["0-0"])}function i(){const n=l().getSelectedKeys();e.success(JSON.stringify(n))}function r(){l().setExpandedKeys(["0-0"])}function d(){const n=l().getExpandedKeys();e.success(JSON.stringify(n))}function B(n){l().checkAll(n)}function k(n){l().expandAll(n)}function _(n=null){l().insertNodeByKey({parentKey:n,node:{title:"\u65B0\u589E\u8282\u70B9",key:"2-2-2"},push:"push"})}function y(n){l().deleteNodeByKey(n)}function E(n){l().updateNodeByKey(n,{title:"parent2-new"})}return{treeData:D,treeRef:t,handleLevel:f,handleSetCheckData:m,handleGetCheckData:C,handleSetSelectData:s,handleGetSelectData:i,handleSetExpandData:r,handleGetExpandData:d,appendNodeByKey:_,deleteNodeByKey:y,updateNodeByKey:E,checkAll:B,expandAll:k}}}),b={class:"mb-4"},T=a(" \u5C55\u5F00\u5168\u90E8 "),w=a(" \u6298\u53E0\u5168\u90E8 "),G=a(" \u5168\u9009 "),x=a(" \u5168\u4E0D\u9009 "),L=a(" \u663E\u793A\u5230\u7B2C2\u7EA7 "),P=a(" \u663E\u793A\u5230\u7B2C1\u7EA7 "),V={class:"mb-4"},W=a(" \u8BBE\u7F6E\u52FE\u9009\u6570\u636E "),J=a(" \u83B7\u53D6\u52FE\u9009\u6570\u636E "),M=a(" \u8BBE\u7F6E\u9009\u4E2D\u6570\u636E "),O=a(" \u83B7\u53D6\u9009\u4E2D\u6570\u636E "),H=a(" \u8BBE\u7F6E\u5C55\u5F00\u6570\u636E "),R=a(" \u83B7\u53D6\u5C55\u5F00\u6570\u636E "),z={class:"mb-4"},q=a(" \u6DFB\u52A0\u6839\u8282\u70B9 "),I=a(" \u6DFB\u52A0\u5728parent3\u5185\u6DFB\u52A0\u8282\u70B9 "),Q=a(" \u5220\u9664parent3\u8282\u70B9 "),U=a(" \u66F4\u65B0parent2\u8282\u70B9 ");function X(t,e,l,f,m,C){const s=c("a-button"),i=c("BasicTree"),r=c("PageWrapper");return S(),g(r,{title:"Tree\u51FD\u6570\u64CD\u4F5C\u793A\u4F8B",contentBackground:"",contentClass:"p-4"},{default:u(()=>[p("div",b,[o(s,{onClick:e[0]||(e[0]=d=>t.expandAll(!0)),class:"mr-2"},{default:u(()=>[T]),_:1}),o(s,{onClick:e[1]||(e[1]=d=>t.expandAll(!1)),class:"mr-2"},{default:u(()=>[w]),_:1}),o(s,{onClick:e[2]||(e[2]=d=>t.checkAll(!0)),class:"mr-2"},{default:u(()=>[G]),_:1}),o(s,{onClick:e[3]||(e[3]=d=>t.checkAll(!1)),class:"mr-2"},{default:u(()=>[x]),_:1}),o(s,{onClick:e[4]||(e[4]=d=>t.handleLevel(2)),class:"mr-2"},{default:u(()=>[L]),_:1}),o(s,{onClick:e[5]||(e[5]=d=>t.handleLevel(1)),class:"mr-2"},{default:u(()=>[P]),_:1})]),p("div",V,[o(s,{onClick:t.handleSetCheckData,class:"mr-2"},{default:u(()=>[W]),_:1},8,["onClick"]),o(s,{onClick:t.handleGetCheckData,class:"mr-2"},{default:u(()=>[J]),_:1},8,["onClick"]),o(s,{onClick:t.handleSetSelectData,class:"mr-2"},{default:u(()=>[M]),_:1},8,["onClick"]),o(s,{onClick:t.handleGetSelectData,class:"mr-2"},{default:u(()=>[O]),_:1},8,["onClick"]),o(s,{onClick:t.handleSetExpandData,class:"mr-2"},{default:u(()=>[H]),_:1},8,["onClick"]),o(s,{onClick:t.handleGetExpandData,class:"mr-2"},{default:u(()=>[R]),_:1},8,["onClick"])]),p("div",z,[o(s,{onClick:e[6]||(e[6]=d=>t.appendNodeByKey(null)),class:"mr-2"},{default:u(()=>[q]),_:1}),o(s,{onClick:e[7]||(e[7]=d=>t.appendNodeByKey("2-2")),class:"mr-2"},{default:u(()=>[I]),_:1}),o(s,{onClick:e[8]||(e[8]=d=>t.deleteNodeByKey("2-2")),class:"mr-2"},{default:u(()=>[Q]),_:1}),o(s,{onClick:e[9]||(e[9]=d=>t.updateNodeByKey("1-1")),class:"mr-2"},{default:u(()=>[U]),_:1})]),o(i,{treeData:t.treeData,title:"\u51FD\u6570\u64CD\u4F5C",ref:"treeRef",checkable:!0},null,8,["treeData"])]),_:1})}var ie=v(j,[["render",X]]);export{ie as default};
