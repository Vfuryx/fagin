var e=Object.defineProperty,t=Object.defineProperties,l=Object.getOwnPropertyDescriptors,n=Object.getOwnPropertySymbols,a=Object.prototype.hasOwnProperty,s=Object.prototype.propertyIsEnumerable,c=(t,l,n)=>l in t?e(t,l,{enumerable:!0,configurable:!0,writable:!0,value:n}):t[l]=n,r=(e,t)=>{for(var l in t||(t={}))a.call(t,l)&&c(e,l,t[l]);if(n)for(var l of n(t))s.call(t,l)&&c(e,l,t[l]);return e},o=(e,n)=>t(e,l(n)),i=(e,t,l)=>new Promise(((n,a)=>{var s=e=>{try{r(l.next(e))}catch(t){a(t)}},c=e=>{try{r(l.throw(e))}catch(t){a(t)}},r=e=>e.done?n(e.value):Promise.resolve(e.value).then(s,c);r((l=l.apply(e,t)).next())}));import{z as d,D as u,az as h,aA as y,an as p,r as f,j as v,S as k,T as b,A as m,B as K,a0 as g,F as x,G as S,a6 as A,ac as L,a1 as C,a5 as E,ad as T,L as _,N,w as I,aa as w,ar as O,aC as j,aR as D,ay as B,u as M,n as P,_ as $,$ as R,ai as U,a$ as F,bu as V,aP as H,ah as X,P as Y,t as z,c8 as G,o as q,al as J,H as W,J as Q,c9 as Z,a4 as ee,aq as te,bv as le}from"./vendor.5da97255.js";import{j as ne,_ as ae,au as se,p as ce,e as re,aK as oe,aL as ie,a as de,an as ue,k as he,aM as ye,ao as pe,aH as fe,S as ve,aN as ke,D as be,ae as me}from"./index.ab59a04a.js";/* empty css              */const Ke=({icon:e})=>e?ne(e)?d(ae,{icon:e,class:"mr-1"}):ae:null;var ge,xe;(xe=ge||(ge={}))[xe.SELECT_ALL=0]="SELECT_ALL",xe[xe.UN_SELECT_ALL=1]="UN_SELECT_ALL",xe[xe.EXPAND_ALL=2]="EXPAND_ALL",xe[xe.UN_EXPAND_ALL=3]="UN_EXPAND_ALL",xe[xe.CHECK_STRICTLY=4]="CHECK_STRICTLY",xe[xe.CHECK_UN_STRICTLY=5]="CHECK_UN_STRICTLY";var Se=u({name:"BasicTreeHeader",components:{BasicTitle:se,Icon:ae,Dropdown:h,Menu:y,MenuItem:y.Item,MenuDivider:y.Divider,InputSearch:p.Search},props:{helpMessage:{type:[String,Array],default:""},title:ce.string,toolbar:ce.bool,checkable:ce.bool,search:ce.bool,checkAll:ce.func,expandAll:ce.func,searchText:ce.string},emits:["strictly-change","search"],setup(e,{emit:t,slots:l}){const{t:n}=re(),a=f(""),s=v((()=>["mr-1","w-full",{"ml-5":l.headerTitle||e.title}])),c=v((()=>{const{checkable:t}=e,l=[{label:n("component.tree.expandAll"),value:2},{label:n("component.tree.unExpandAll"),value:3,divider:t}];return t?[{label:n("component.tree.selectAll"),value:0},{label:n("component.tree.unSelectAll"),value:1,divider:t},...l,{label:n("component.tree.checkStrictly"),value:4},{label:n("component.tree.checkUnStrictly"),value:5}]:l}));const r=b((function(e){t("search",e)}),200);return k((()=>a.value),(e=>{r(e)})),k((()=>e.searchText),(e=>{e!==a.value&&(a.value=e)})),{t:n,toolbarList:c,handleMenuClick:function(l){var n,a,s,c;const{key:r}=l;switch(r){case 0:null==(n=e.checkAll)||n.call(e,!0);break;case 1:null==(a=e.checkAll)||a.call(e,!1);break;case 2:null==(s=e.expandAll)||s.call(e,!0);break;case 3:null==(c=e.expandAll)||c.call(e,!1);break;case 4:t("strictly-change",!1);break;case 5:t("strictly-change",!0)}},searchValue:a,getInputSearchCls:s}}});m("data-v-05dbfd0e");const Ae={class:"flex px-2 py-1.5 items-center basic-tree-header"},Le={key:2,class:"flex flex-1 justify-self-stretch items-center cursor-pointer"};K(),Se.render=function(e,t,l,n,a,s){const c=g("BasicTitle"),r=g("InputSearch"),o=g("Icon"),i=g("MenuItem"),d=g("MenuDivider"),u=g("Menu"),h=g("Dropdown");return x(),S("div",Ae,[e.$slots.headerTitle?A(e.$slots,"headerTitle",{key:0},void 0,!0):L("",!0),!e.$slots.headerTitle&&e.title?(x(),C(c,{key:1,helpMessage:e.helpMessage},{default:E((()=>[T(_(e.title),1)])),_:1},8,["helpMessage"])):L("",!0),e.search||e.toolbar?(x(),S("div",Le,[e.search?(x(),S("div",{key:0,class:N(e.getInputSearchCls)},[I(r,{placeholder:e.t("common.searchText"),size:"small",allowClear:"",value:e.searchValue,"onUpdate:value":t[0]||(t[0]=t=>e.searchValue=t)},null,8,["placeholder","value"])],2)):L("",!0),e.toolbar?(x(),C(h,{key:1,onClick:t[1]||(t[1]=B((()=>{}),["prevent"]))},{overlay:E((()=>[I(u,{onClick:e.handleMenuClick},{default:E((()=>[(x(!0),S(w,null,O(e.toolbarList,(e=>(x(),S(w,{key:e.value},[I(i,j(D({key:e.value})),{default:E((()=>[T(_(e.label),1)])),_:2},1040),e.divider?(x(),C(d,{key:0})):L("",!0)],64)))),128))])),_:1},8,["onClick"])])),default:E((()=>[I(o,{icon:"ion:ellipsis-vertical"})])),_:1})):L("",!0)])):L("",!0)])},Se.__scopeId="data-v-05dbfd0e";const Ce={width:{type:Number,default:156},customEvent:{type:Object,default:null},styles:{type:Object},showIcon:{type:Boolean,default:!0},axis:{type:Object,default:()=>({x:0,y:0})},items:{type:Array,default:()=>[]}},Ee=e=>{const{item:t}=e;return I("span",{style:"display: inline-block; width: 100%; ",class:"px-4",onClick:e.handler.bind(null,t)},[e.showIcon&&t.icon&&I(ae,{class:"mr-2",icon:t.icon},null),I("span",null,[t.label])])};var Te=u({name:"ContextMenu",props:Ce,setup(e){const t=f(null),l=f(!1),n=v((()=>{const{axis:t,items:l,styles:n,width:a}=e,{x:s,y:c}=t||{x:0,y:0},i=40*(l||[]).length,d=a,u=document.body,h=u.clientWidth<s+d?s-d:s,y=u.clientHeight<c+i?c-i:c;return o(r({},n),{width:`${a}px`,left:`${h+1}px`,top:`${y+1}px`})}));function a(e,t){const{handler:n,disabled:a}=e;a||(l.value=!1,null==t||t.stopPropagation(),null==t||t.preventDefault(),null==n||n())}function s(t){return t.map((t=>{const{disabled:n,label:c,children:r,divider:o=!1}=t,i={item:t,handler:a,showIcon:e.showIcon};return r&&0!==r.length?M(l)?I(y.SubMenu,{key:c,disabled:n,popupClassName:"context-menu__popup"},{title:()=>I(Ee,i,null),default:()=>s(r)}):null:I(w,null,[I(y.Item,{disabled:n,class:"context-menu__item",key:c},{default:()=>[I(Ee,i,null)]}),o?I(F,{key:`d-${c}`},null):null])}))}return $((()=>{R((()=>l.value=!0))})),U((()=>{const e=M(t);e&&document.body.removeChild(e)})),()=>{let a;if(!M(l))return null;const{items:c}=e;return I(y,{inlineIndent:12,mode:"vertical",class:"context-menu",ref:t,style:M(n)},"function"==typeof(r=a=s(c))||"[object Object]"===Object.prototype.toString.call(r)&&!V(r)?a:{default:()=>[a]});var r}}});const _e={domList:[],resolve:()=>{}},Ne=function(e){const{event:t}=e||{};if(t&&(null==t||t.preventDefault()),ie)return new Promise((l=>{const n=document.body,a=document.createElement("div"),s={};e.styles&&(s.styles=e.styles),e.items&&(s.items=e.items),e.event&&(s.customEvent=t,s.axis={x:t.clientX,y:t.clientY});const c=I(Te,s);H(c,a);const r=function(){_e.resolve("")};_e.domList.push(a);const o=function(){_e.domList.forEach((e=>{try{e&&n.removeChild(e)}catch(t){}})),n.removeEventListener("click",r),n.removeEventListener("scroll",r)};_e.resolve=function(e){o(),l(e)},o(),n.appendChild(a),n.addEventListener("click",r),n.addEventListener("scroll",r)}))},Ie=function(){_e&&(_e.resolve(""),_e.domList=[])};var we=u({name:"BasicTree",inheritAttrs:!1,props:{value:{type:[Object,Array]},renderIcon:{type:Function},helpMessage:{type:[String,Array],default:""},title:ce.string,toolbar:ce.bool,search:ce.bool,searchValue:ce.string,checkStrictly:ce.bool,clickRowToExpand:ce.bool.def(!0),checkable:ce.bool.def(!1),defaultExpandLevel:{type:[String,Number],default:""},defaultExpandAll:ce.bool.def(!1),replaceFields:{type:Object},treeData:{type:Array},actionList:{type:Array,default:()=>[]},expandedKeys:{type:Array,default:()=>[]},selectedKeys:{type:Array,default:()=>[]},checkedKeys:{type:Array,default:()=>[]},beforeRightClick:{type:Function,default:null},rightMenuList:{type:Array},filterFn:{type:Function,default:null},highlight:{type:[Boolean,String],default:!1},expandOnSearch:ce.bool.def(!1),checkOnSearch:ce.bool.def(!1),selectedOnSearch:ce.bool.def(!1)},emits:["update:expandedKeys","update:selectedKeys","update:value","change","check","update:searchValue"],setup(e,{attrs:t,slots:l,emit:n,expose:a}){const s=Y({checkStrictly:e.checkStrictly,expandedKeys:e.expandedKeys||[],selectedKeys:e.selectedKeys||[],checkedKeys:e.checkedKeys||[]}),c=Y({startSearch:!1,searchText:"",searchData:[]}),d=f([]),[u]=function(e=!0){return X()&&e&&U((()=>{Ie()})),[Ne,Ie]}(),{prefixCls:h}=de("basic-tree"),y=v((()=>{const{replaceFields:t}=e;return r({children:"children",title:"title",key:"key"},t)})),p=v((()=>{let l=o(r(r({blockNode:!0},t),e),{expandedKeys:s.expandedKeys,selectedKeys:s.selectedKeys,checkedKeys:s.checkedKeys,checkStrictly:s.checkStrictly,replaceFields:M(y),"onUpdate:expandedKeys":e=>{s.expandedKeys=e,n("update:expandedKeys",e)},"onUpdate:selectedKeys":e=>{s.selectedKeys=e,n("update:selectedKeys",e)},onCheck:(e,t)=>{let l=z(s.checkedKeys);if(ue(l)&&c.startSearch){const{key:e}=M(y);l=G(l,C(t.node.$attrs.node[e])),t.checked&&l.push(t.node.$attrs.node[e]),s.checkedKeys=l}else s.checkedKeys=e;const a=z(s.checkedKeys);n("update:value",a),n("check",a,t)},onRightClick:T});return q(l,"treeData","class")})),b=v((()=>c.startSearch?c.searchData:M(d))),m=v((()=>!b.value||0===b.value.length)),{deleteNodeByKey:K,insertNodeByKey:g,insertNodesByKey:x,filterByLevel:S,updateNodeByKey:A,getAllKeys:L,getChildrenKeys:C,getEnabledKeys:E}=function(e,t){function l(n){const a=[],s=n||M(e),{key:c,children:r}=M(t);if(!r||!c)return a;for(let e=0;e<s.length;e++){const t=s[e];a.push(t[c]);const n=t[r];n&&n.length&&a.push(...l(n))}return a}return{deleteNodeByKey:function l(n,a){if(!n)return;const s=a||M(e),{key:c,children:r}=M(t);if(r&&c)for(let e=0;e<s.length;e++){const t=s[e],a=t[r];if(t[c]===n){s.splice(e,1);break}a&&a.length&&l(n,t[r])}},insertNodeByKey:function({parentKey:l=null,node:n,push:a="push"}){const s=P(M(e));if(!l)return s[a](n),void(e.value=s);const{key:c,children:r}=M(t);r&&c&&(oe(s,(e=>{if(e[c]===l)return e[r]=e[r]||[],e[r][a](n),!0})),e.value=s)},insertNodesByKey:function({parentKey:l=null,list:n,push:a="push"}){const s=P(M(e));if(n&&!(n.length<1))if(l){const{key:c,children:r}=M(t);if(!r||!c)return;oe(s,(t=>{if(t[c]===l){t[r]=t[r]||[];for(let e=0;e<n.length;e++)t[r][a](n[e]);return e.value=s,!0}}))}else for(let e=0;e<n.length;e++)s[a](n[e])},filterByLevel:function l(n=1,a,s=1){if(!n)return[];const c=[],r=a||M(e)||[];for(let e=0;e<r.length;e++){const a=r[e],{key:o,children:i}=M(t),d=o?a[o]:"",u=i?a[i]:[];c.push(d),u&&u.length&&s<n&&(s+=1,c.push(...l(n,u,s)))}return c},updateNodeByKey:function l(n,a,s){if(!n)return;const c=s||M(e),{key:o,children:i}=M(t);if(i&&o)for(let e=0;e<c.length;e++){const t=c[e],s=t[i];if(t[o]===n){c[e]=r(r({},c[e]),a);break}s&&s.length&&l(n,a,t[i])}},getAllKeys:l,getChildrenKeys:function n(a,s){const c=[],r=s||M(e),{key:o,children:i}=M(t);if(!i||!o)return c;for(let e=0;e<r.length;e++){const t=r[e],s=t[i];a===t[o]?(c.push(t[o]),s&&s.length&&c.push(...l(s))):s&&s.length&&c.push(...n(a,s))}return c},getEnabledKeys:function l(n){const a=[],s=n||M(e),{key:c,children:r}=M(t);if(!r||!c)return a;for(let e=0;e<s.length;e++){const t=s[e];!0!==t.disabled&&!1!==t.selectable&&a.push(t[c]);const n=t[r];n&&n.length&&a.push(...l(n))}return a}}}(d,y);function T(t){return i(this,arguments,(function*({event:t,node:l}){var n;const{rightMenuList:a=[],beforeRightClick:s}=e;let c={event:t,items:[]};if(s&&he(s)){let e=yield s(l,t);Array.isArray(e)?c.items=e:Object.assign(c,e)}else c.items=a;(null==(n=c.items)?void 0:n.length)&&u(c)}))}function _(e){s.expandedKeys=e}function N(e){s.selectedKeys=e}function O(e){s.checkedKeys=e}function j(e){s.checkedKeys=e?E():[]}function D(e){s.expandedKeys=e?L():[]}function B(e){s.checkStrictly=e}function R(t){if(t!==c.searchText&&(c.searchText=t),n("update:searchValue",t),!t)return void(c.startSearch=!1);const{filterFn:l,checkable:a,expandOnSearch:s,checkOnSearch:r,selectedOnSearch:o}=M(e);c.startSearch=!0;const{title:i,key:u}=M(y),h=[];if(c.searchData=ye(M(d),(e=>{var n,a;const s=l?l(t,e,M(y)):null!=(a=null==(n=e[i])?void 0:n.includes(t))&&a;return s&&h.push(e[u]),s}),M(y)),s){const e=pe(c.searchData).map((e=>e[u]));e&&e.length&&_(e)}r&&a&&h.length&&O(h),o&&h.length&&N(h)}function F(t,l){if(e.clickRowToExpand&&l&&0!==l.length)if(s.expandedKeys.includes(t)){const e=[...s.expandedKeys],l=e.findIndex((e=>e===t));-1!==l&&e.splice(l,1),_(e)}else _([...s.expandedKeys,t])}k((()=>e.searchValue),(e=>{e!==c.searchText&&(c.searchText=e)}),{immediate:!0}),k((()=>e.treeData),(e=>{e&&R(c.searchText)})),J((()=>{d.value=e.treeData})),$((()=>{const t=parseInt(e.defaultExpandLevel);t>0?s.expandedKeys=S(t):e.defaultExpandAll&&D(!0)})),J((()=>{s.expandedKeys=e.expandedKeys})),J((()=>{s.selectedKeys=e.selectedKeys})),J((()=>{s.checkedKeys=e.checkedKeys})),k((()=>e.value),(()=>{s.checkedKeys=z(e.value||[])})),k((()=>s.checkedKeys),(()=>{const e=z(s.checkedKeys);n("update:value",e),n("change",e)})),J((()=>{s.checkStrictly=e.checkStrictly}));function H(t){const{actionList:l}=e;if(l&&0!==l.length)return l.map(((e,l)=>{var n;let a=!0;return he(e.show)?a=null==(n=e.show)?void 0:n.call(e,t):be(e.show)&&(a=e.show),a?I("span",{key:l,class:`${h}__action`},[e.render(t)]):null}))}function ne({data:t,level:n}){if(!t)return null;const a=c.searchText,{highlight:s}=M(e);return t.map((t=>{var i;const{title:d,key:u,children:f}=M(y),v=q(t,"title"),k=function(t,l){return!l&&e.renderIcon&&he(e.renderIcon)?e.renderIcon(t):l}(o(r({},t),{level:n}),t.icon),b=le(t,f)||[],m=le(t,d),K=m.indexOf(a),g=c.startSearch&&!ke(a)&&s&&-1!==K,x=`color: ${be(s)?"#f50":s}`,S=g?I("span",{class:(null==(i=M(p))?void 0:i.blockNode)?`${h}__content`:""},[I("span",null,[m.substr(0,K)]),I("span",{style:x},[a]),I("span",null,[m.substr(K+a.length)])]):m;return I(Z.TreeNode,ee(v,{node:z(t),key:le(t,u)}),{title:()=>I("span",{class:`${h}-title pl-2`,onClick:F.bind(null,t[u],t[f])},[(null==l?void 0:l.title)?me(l,"title",t):I(w,null,[k&&I(Ke,{icon:k},null),S,I("span",{class:`${h}__actions`},[H(o(r({},t),{level:n}))])])]),default:()=>ne({data:b,level:n+1})})}))}return a({setExpandedKeys:_,getExpandedKeys:function(){return s.expandedKeys},setSelectedKeys:N,getSelectedKeys:function(){return s.selectedKeys},setCheckedKeys:O,getCheckedKeys:function(){return s.checkedKeys},insertNodeByKey:g,insertNodesByKey:x,deleteNodeByKey:K,updateNodeByKey:A,checkAll:j,expandAll:D,filterByLevel:e=>{s.expandedKeys=S(e)},setSearchValue:e=>{R(e)},getSearchValue:()=>c.searchText}),()=>{let n;const{title:a,helpMessage:s,toolbar:o,search:i,checkable:d}=e,u=a||o||i||l.headerTitle;return I("div",{class:[h,"h-full",t.class]},[u&&I(Se,{checkable:d,checkAll:j,expandAll:D,title:a,search:i,toolbar:o,helpMessage:s,onStrictlyChange:B,onSearch:R,searchText:c.searchText},(y=n=fe(l),"function"==typeof y||"[object Object]"===Object.prototype.toString.call(y)&&!V(y)?n:{default:()=>[n]})),W(I(ve,{style:{height:"calc(100% - 38px)"}},{default:()=>[I(Z,ee(M(p),{showIcon:!1}),r({default:()=>ne({data:M(b),level:1})},fe(l)))]}),[[Q,!M(m)]]),W(I(te,{image:te.PRESENTED_IMAGE_SIMPLE,class:"!mt-4"},null),[[Q,M(m)]])]);var y}}});export{we as _};
