(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-2d0da705"],{"6c35":function(t,e,a){"use strict";a.r(e);var n=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("page-header-wrapper",[a("a-card",{attrs:{bordered:!1}},[a("div",{staticClass:"table-operator"},[a("a-button",{attrs:{type:"primary",icon:"plus"},on:{click:function(e){return t.$router.push("/admin/roles/form/")}}},[t._v("添加")])],1),a("a-tabs",{attrs:{"default-active-key":0,animated:!1},on:{change:t.tabChange}},[a("a-tab-pane",{key:0,attrs:{tab:"总后台"}},[a("a-table",{attrs:{columns:t.columns,"data-source":t.list,pagination:!1,"row-selection":{selectedRowKeys:t.selectedRowKeys,onChange:t.onSelectChange},"row-key":"id"},scopedSlots:t._u([{key:"action",fn:function(e,n){return a("span",{},[a("a-button",{attrs:{icon:"edit"},on:{click:function(e){return t.$router.push("/admin/roles/form/"+n.id)}}},[t._v("编辑")]),1!==n.id?a("a-popconfirm",{attrs:{title:"是否要删除角色？"},on:{confirm:function(e){return t.handleDelete(n.id)}}},[a("a-button",{attrs:{icon:"user-delete"}},[t._v("删除角色")])],1):t._e()],1)}}])})],1),a("a-tab-pane",{key:1,attrs:{tab:"商家"}},[a("a-table",{attrs:{columns:t.columns,"data-source":t.list,pagination:!1,"row-selection":{selectedRowKeys:t.selectedRowKeys,onChange:t.onSelectChange},"row-key":"id"},scopedSlots:t._u([{key:"action",fn:function(e,n){return a("span",{},[a("a-button",{attrs:{icon:"edit"},on:{click:function(e){return t.$router.push("/admin/roles/form/"+n.id)}}},[t._v("编辑")]),1!==n.id?a("a-popconfirm",{attrs:{title:"是否要删除角色？"},on:{confirm:function(e){return t.handleDelete(n.id)}}},[a("a-button",{attrs:{icon:"user-delete"}},[t._v("删除角色")])],1):t._e()],1)}}])})],1)],1),t.total_page>0?a("a-pagination",{staticClass:"ant-table-pagination",attrs:{"page-size":t.params.page_size,total:t.params.total_count,"show-less-items":""},on:{"update:pageSize":function(e){return t.$set(t.params,"page_size",e)},"update:page-size":function(e){return t.$set(t.params,"page_size",e)},change:t.onChange},model:{value:t.params.current_page,callback:function(e){t.$set(t.params,"current_page",e)},expression:"params.current_page"}}):t._e()],1)],1)},o=[],s=(a("68c7"),a("de1b")),r=a("4360"),i=a("8bbf"),c=a.n(i);c.a.use(s["a"]);var p={components:{},props:{},data:function(){return{params:{current_page:1,page_size:10,total_count:0,type:0},total_page:0,selectedRowKeys:[],columns:[{title:"角色名称",dataIndex:"name",key:"name"},{title:"操作",key:"id",fixed:"right",scopedSlots:{customRender:"action"}}],list:[]}},watch:{},computed:{},methods:{onSelectChange:function(t){this.selectedRowKeys=t},onChange:function(t){this.params.current_page=t,this.onload()},tabChange:function(t){this.list=[],this.total_page=0,this.to_nav(t)},to_nav:function(t){this.params.type=t,this.params.current_page=1,this.onload()},onload:function(){var t=this;r["a"].dispatch("getRoleList",this.params).then((function(e){0!==e.code&&t.$notification.error({message:"错误",description:e.message}),t.params.current_page=e.data.paginator.current_page,t.params.page_size=e.data.paginator.page_size,t.params.total_count=e.data.paginator.total_count,t.total_page=e.data.paginator.total_page,t.list=e.data.roles})).catch((function(){t.$notification.error({message:"错误",description:"请求信息失败，请重试"})}))},handleDelete:function(t){var e=this;r["a"].dispatch("delRole",t).then((function(t){if(0!==t.code)return e.$notification.error({message:"错误",description:t.message}),!1;e.onload(),e.$message.info("删除成功")})).catch((function(t){return e.$notification.error({message:"错误",description:"请求信息失败，请重试"}),!1}))}},created:function(){this.onload()},mounted:function(){}},u=p,l=a("2877"),d=Object(l["a"])(u,n,o,!1,null,"62896276",null);e["default"]=d.exports}}]);