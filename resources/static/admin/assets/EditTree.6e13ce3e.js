import{_ as l}from"./Tree.vue_vue&type=style&index=0&lang.d5823234.js";import{t as p}from"./data.5e96733f.js";import{P as d}from"./index.a09675a6.js";import{_ as f}from"./index.6d3ed4e5.js";import{A as m,a0 as i,B as g,a1 as h,a6 as D,H as F,w as a,z as s,c9 as B,ca as C}from"./vendor.3850bdb6.js";import"./useContextMenu.f3a8e755.js";/* empty css               *//* empty css               *//* empty css                */import"./useWindowSizeFn.4c2ec928.js";import"./useContentViewHeight.34544381.js";const E=m({components:{BasicTree:l,PageWrapper:d},setup(){function t(e){console.log(e)}function n(e){return[{label:"\u65B0\u589E",handler:()=>{console.log("\u70B9\u51FB\u4E86\u65B0\u589E",e)},icon:"bi:plus"},{label:"\u5220\u9664",handler:()=>{console.log("\u70B9\u51FB\u4E86\u5220\u9664",e)},icon:"bx:bxs-folder-open"}]}const o=[{render:e=>s(B,{class:"ml-2",onClick:()=>{t(e)}})},{render:()=>s(C)}];function u({level:e}){return e===1?"ion:git-compare-outline":e===2?"ion:home":e===3?"ion:airplane":""}return{treeData:p,actionList:o,getRightMenuList:n,createIcon:u}}}),b={class:"flex"};function _(t,n,o,u,e,j){const r=i("BasicTree"),c=i("PageWrapper");return g(),h(c,{title:"Tree\u51FD\u6570\u64CD\u4F5C\u793A\u4F8B"},{default:D(()=>[F("div",b,[a(r,{class:"w-1/3",title:"\u53F3\u4FA7\u64CD\u4F5C\u6309\u94AE/\u81EA\u5B9A\u4E49\u56FE\u6807",helpMessage:"\u5E2E\u52A9\u4FE1\u606F",treeData:t.treeData,actionList:t.actionList,renderIcon:t.createIcon},null,8,["treeData","actionList","renderIcon"]),a(r,{class:"w-1/3 mx-4",title:"\u53F3\u952E\u83DC\u5355",treeData:t.treeData,beforeRightClick:t.getRightMenuList},null,8,["treeData","beforeRightClick"]),a(r,{class:"w-1/3",title:"\u5DE5\u5177\u680F\u4F7F\u7528",toolbar:"",checkable:"",search:"",treeData:t.treeData,beforeRightClick:t.getRightMenuList},null,8,["treeData","beforeRightClick"])])]),_:1})}var I=f(E,[["render",_]]);export{I as default};