(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-f08aeae4"],{"0fea":function(e,t,r){"use strict";r.d(t,"c",(function(){return i})),r.d(t,"d",(function(){return o})),r.d(t,"b",(function(){return s})),r.d(t,"a",(function(){return c}));var n=r("b775"),a={user:"/user",role:"/role",service:"/service",permission:"/permission",permissionNoPager:"/permission/no-pager",orgTree:"/org/tree"};function i(e){return Object(n["b"])({url:a.role,method:"get",params:e})}function o(e){return Object(n["b"])({url:a.service,method:"get",params:e})}function s(e){return Object(n["b"])({url:a.permissionNoPager,method:"get",params:e})}function c(e){return Object(n["b"])({url:a.orgTree,method:"get",params:e})}},"432b":function(e,t,r){"use strict";r.d(t,"a",(function(){return i}));var n=r("5530"),a=r("5880"),i={computed:Object(n["a"])(Object(n["a"])({},Object(a["mapState"])({layout:function(e){return e.app.layout},navTheme:function(e){return e.app.theme},primaryColor:function(e){return e.app.color},colorWeak:function(e){return e.app.weak},fixedHeader:function(e){return e.app.fixedHeader},fixedSidebar:function(e){return e.app.fixedSidebar},contentWidth:function(e){return e.app.contentWidth},autoHideHeader:function(e){return e.app.autoHideHeader},isMobile:function(e){return e.app.isMobile},sideCollapsed:function(e){return e.app.sideCollapsed},multiTab:function(e){return e.app.multiTab}})),{},{isTopMenu:function(){return"topmenu"===this.layout}}),methods:{isSideMenu:function(){return!this.isTopMenu}}}},"88bc":function(e,t,r){(function(t){var r=1/0,n=9007199254740991,a="[object Arguments]",i="[object Function]",o="[object GeneratorFunction]",s="[object Symbol]",c="object"==typeof t&&t&&t.Object===Object&&t,l="object"==typeof self&&self&&self.Object===Object&&self,u=c||l||Function("return this")();function d(e,t,r){switch(r.length){case 0:return e.call(t);case 1:return e.call(t,r[0]);case 2:return e.call(t,r[0],r[1]);case 3:return e.call(t,r[0],r[1],r[2])}return e.apply(t,r)}function m(e,t){var r=-1,n=e?e.length:0,a=Array(n);while(++r<n)a[r]=t(e[r],r,e);return a}function f(e,t){var r=-1,n=t.length,a=e.length;while(++r<n)e[a+r]=t[r];return e}var p=Object.prototype,h=p.hasOwnProperty,b=p.toString,v=u.Symbol,g=p.propertyIsEnumerable,y=v?v.isConcatSpreadable:void 0,k=Math.max;function j(e,t,r,n,a){var i=-1,o=e.length;r||(r=_),a||(a=[]);while(++i<o){var s=e[i];t>0&&r(s)?t>1?j(s,t-1,r,n,a):f(a,s):n||(a[a.length]=s)}return a}function O(e,t){return e=Object(e),w(e,t,(function(t,r){return r in e}))}function w(e,t,r){var n=-1,a=t.length,i={};while(++n<a){var o=t[n],s=e[o];r(s,o)&&(i[o]=s)}return i}function x(e,t){return t=k(void 0===t?e.length-1:t,0),function(){var r=arguments,n=-1,a=k(r.length-t,0),i=Array(a);while(++n<a)i[n]=r[t+n];n=-1;var o=Array(t+1);while(++n<t)o[n]=r[n];return o[t]=i,d(e,this,o)}}function _(e){return A(e)||S(e)||!!(y&&e&&e[y])}function C(e){if("string"==typeof e||E(e))return e;var t=e+"";return"0"==t&&1/e==-r?"-0":t}function S(e){return q(e)&&h.call(e,"callee")&&(!g.call(e,"callee")||b.call(e)==a)}var A=Array.isArray;function P(e){return null!=e&&T(e.length)&&!M(e)}function q(e){return N(e)&&P(e)}function M(e){var t=H(e)?b.call(e):"";return t==i||t==o}function T(e){return"number"==typeof e&&e>-1&&e%1==0&&e<=n}function H(e){var t=typeof e;return!!e&&("object"==t||"function"==t)}function N(e){return!!e&&"object"==typeof e}function E(e){return"symbol"==typeof e||N(e)&&b.call(e)==s}var F=x((function(e,t){return null==e?{}:O(e,m(j(t,1),C))}));e.exports=F}).call(this,r("c8ba"))},"98bb":function(e,t,r){"use strict";r.r(t);var n=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("a-card",{style:{height:"100%"},attrs:{bordered:!1}},[r("a-row",{attrs:{gutter:24}},[r("a-col",{attrs:{md:4}},[r("a-list",{attrs:{itemLayout:"vertical",dataSource:e.roles},scopedSlots:e._u([{key:"renderItem",fn:function(t,n){return r("a-list-item",{key:n},[r("a-list-item-meta",{style:{marginBottom:"0"}},[r("span",{staticStyle:{"text-align":"center",display:"block"},attrs:{slot:"description"},slot:"description"},[e._v(e._s(t.describe))]),r("a",{staticStyle:{"text-align":"center",display:"block"},attrs:{slot:"title"},on:{click:function(r){return e.edit(t)}},slot:"title"},[e._v(e._s(t.name))])])],1)}}])})],1),r("a-col",{attrs:{md:20}},[r("div",{staticStyle:{"max-width":"800px"}},[e.isMobile()?r("a-divider"):e._e(),e.mdl.id?r("div",[r("h3",[e._v("角色："+e._s(e.mdl.name))])]):e._e(),r("a-form",{attrs:{form:e.form,layout:e.isMobile()?"vertical":"horizontal"}},[r("a-form-item",{attrs:{label:"唯一键"}},[r("a-input",{directives:[{name:"decorator",rawName:"v-decorator",value:["id",{rules:[{required:!0,message:"Please input unique key!"}]}],expression:"[ 'id', {rules: [{ required: true, message: 'Please input unique key!' }]} ]"}],attrs:{placeholder:"请填写唯一键"}})],1),r("a-form-item",{attrs:{label:"角色名称"}},[r("a-input",{directives:[{name:"decorator",rawName:"v-decorator",value:["name",{rules:[{required:!0,message:"Please input role name!"}]}],expression:"[ 'name', {rules: [{ required: true, message: 'Please input role name!' }]} ]"}],attrs:{placeholder:"请填写角色名称"}})],1),r("a-form-item",{attrs:{label:"状态"}},[r("a-select",{directives:[{name:"decorator",rawName:"v-decorator",value:["status",{rules:[]}],expression:"[ 'status', {rules: []} ]"}]},[r("a-select-option",{attrs:{value:1}},[e._v("正常")]),r("a-select-option",{attrs:{value:2}},[e._v("禁用")])],1)],1),r("a-form-item",{attrs:{label:"备注说明"}},[r("a-textarea",{directives:[{name:"decorator",rawName:"v-decorator",value:["describe",{rules:[{required:!0,message:"Please input role name!"}]}],expression:"[ 'describe', {rules: [{ required: true, message: 'Please input role name!' }]} ]"}],attrs:{row:3,placeholder:"请填写角色名称"}})],1),r("a-form-item",{attrs:{label:"拥有权限"}},e._l(e.permissions,(function(t,n){return r("a-row",{key:n,attrs:{gutter:16}},[r("a-col",{attrs:{xl:4,lg:24}},[e._v(" "+e._s(t.name)+"： ")]),r("a-col",{attrs:{xl:20,lg:24}},[t.actionsOptions.length>0?r("a-checkbox",{attrs:{indeterminate:t.indeterminate,checked:t.checkedAll},on:{change:function(r){return e.onChangeCheckAll(r,t)}}},[e._v(" 全选 ")]):e._e(),r("a-checkbox-group",{attrs:{options:t.actionsOptions},on:{change:function(r){return e.onChangeCheck(t)}},model:{value:t.selected,callback:function(r){e.$set(t,"selected",r)},expression:"permission.selected"}})],1)],1)})),1)],1)],1)])],1)],1)},a=[],i=(r("159b"),r("d81d"),r("88bc")),o=r.n(i),s=r("0fea"),c=r("432b"),l={name:"RoleList",mixins:[c["a"]],components:{},data:function(){return{form:this.$form.createForm(this),mdl:{},roles:[],permissions:[]}},created:function(){var e=this;Object(s["c"])().then((function(t){e.roles=t.result.data,e.roles.push({id:"-1",name:"新增角色",describe:"新增一个角色"})})),this.loadPermissions()},methods:{callback:function(e){},add:function(){this.edit({id:0})},edit:function(e){var t=this;if(this.mdl=Object.assign({},e),this.mdl.permissions&&this.permissions){var r={};this.mdl.permissions.forEach((function(e){r[e.permissionId]=e.actionEntitySet.map((function(e){return e.action}))})),this.permissions.forEach((function(e){var n=r[e.id];e.selected=n||[],t.onChangeCheck(e)}))}this.$nextTick((function(){t.form.setFieldsValue(o()(t.mdl,"id","name","status","describe"))}))},onChangeCheck:function(e){e.indeterminate=!!e.selected.length&&e.selected.length<e.actionsOptions.length,e.checkedAll=e.selected.length===e.actionsOptions.length},onChangeCheckAll:function(e,t){Object.assign(t,{selected:e.target.checked?t.actionsOptions.map((function(e){return e.value})):[],indeterminate:!1,checkedAll:e.target.checked})},loadPermissions:function(){var e=this;Object(s["b"])().then((function(t){var r=t.result;e.permissions=r.map((function(e){return e.checkedAll=!1,e.selected=[],e.indeterminate=!1,e}))}))}}},u=l,d=r("2877"),m=Object(d["a"])(u,n,a,!1,null,"2fd29c24",null);t["default"]=m.exports}}]);