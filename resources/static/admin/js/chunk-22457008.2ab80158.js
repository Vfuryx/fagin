(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-22457008"],{"27b3":function(e,i,t){"use strict";t("ab44")},2909:function(e,i,t){"use strict";t.d(i,"a",(function(){return c}));var n=t("6b75");function s(e){if(Array.isArray(e))return Object(n["a"])(e)}t("a4d3"),t("e01a"),t("d3b7"),t("d28b"),t("3ca3"),t("ddb0"),t("a630");function a(e){if("undefined"!==typeof Symbol&&null!=e[Symbol.iterator]||null!=e["@@iterator"])return Array.from(e)}var r=t("06c5");function o(){throw new TypeError("Invalid attempt to spread non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")}function c(e){return s(e)||a(e)||Object(r["a"])(e)||o()}},ab44:function(e,i,t){},cc8c:function(e,i,t){"use strict";t.r(i);var n=function(){var e=this,i=e.$createElement,t=e._self._c||i;return t("page-header-wrapper",[t("a-card",{attrs:{"body-style":{padding:"24px 32px"},bordered:!1}},[t("div",{staticClass:"admin_table_page_title"},[t("a-button",{staticClass:".ant-calendar-range-right",attrs:{icon:"arrow-left"},on:{click:function(i){return e.$router.back()}}},[e._v("返回")])],1),t("div",{staticClass:"unline underm"}),t("div",{staticClass:"admin_form"},[t("a-form-model",{ref:"ruleForm",attrs:{model:e.info,rules:e.rules,confirmLoading:e.loading}},[t("a-spin",{attrs:{spinning:e.loading}},[t("a-form-model-item",{attrs:{label:"分组类型",prop:"type",labelCol:{lg:{span:4},sm:{span:4}},wrapperCol:{lg:{span:16},sm:{span:20}}}},[t("a-select",{on:{change:e.typeChange},model:{value:e.info.type,callback:function(i){e.$set(e.info,"type",i)},expression:"info.type"}},[t("a-select-option",{attrs:{value:0}},[e._v(" 总后台 ")]),t("a-select-option",{attrs:{value:1}},[e._v(" 商家 ")])],1)],1),t("a-form-model-item",{attrs:{label:"角色名称",prop:"name",labelCol:{lg:{span:4},sm:{span:4}},wrapperCol:{lg:{span:16},sm:{span:20}}}},[t("a-input",{model:{value:e.info.name,callback:function(i){e.$set(e.info,"name",i)},expression:"info.name"}})],1),t("a-form-model-item",{attrs:{label:"关键词",prop:"key",labelCol:{lg:{span:4},sm:{span:4}},wrapperCol:{lg:{span:16},sm:{span:20}}}},[t("a-input",{attrs:{disabled:e.info.id>0},model:{value:e.info.key,callback:function(i){e.$set(e.info,"key",i)},expression:"info.key"}})],1),t("a-form-model-item",{staticStyle:{width:"100%","margin-right":"0"},attrs:{label:"显示排序","label-col":{lg:{span:4},sm:{span:4}},"wrapper-col":{lg:{span:16},sm:{span:20}},prop:"sort"}},[t("a-input-number",{attrs:{"controls-position":"right",min:0},model:{value:e.info.sort,callback:function(i){e.$set(e.info,"sort",i)},expression:"info.sort"}})],1),t("a-form-model-item",{staticStyle:{width:"100%","margin-right":"0"},attrs:{label:"菜单状态","label-col":{lg:{span:4},sm:{span:4}},"wrapper-col":{lg:{span:16},sm:{span:20}},prop:"status"}},[t("a-radio-group",{attrs:{name:"status"},model:{value:e.info.status,callback:function(i){e.$set(e.info,"status",i)},expression:"info.status"}},[t("a-radio",{attrs:{value:1}},[e._v("开启")]),t("a-radio",{attrs:{value:0}},[e._v("关闭")])],1)],1),t("a-form-model-item",{attrs:{label:"菜单权限",labelCol:{lg:{span:4},sm:{span:4}},wrapperCol:{lg:{span:16},sm:{span:20}}}},[t("a-tree-select",{attrs:{"replace-fields":{children:"children",title:"title",key:"id",value:"id"},value:e.info.menu_ids,"tree-data":e.list,"tree-checkable":"","tree-check-strictly":"","show-checked-strategy":"SHOW_ALL"},on:{change:e.treeChange}})],1),t("a-form-model-item",{attrs:{label:"接口权限",labelCol:{lg:{span:4},sm:{span:4}},wrapperCol:{lg:{span:16},sm:{span:20}}}},[t("div",{staticClass:"check_list"},e._l(e.permissions,(function(i,n){return t("div",{key:n,staticClass:"item"},[t("div",{staticClass:"permission_title"},[e._v(e._s(i.name)+" "),t("a-checkbox",{staticStyle:{float:"right"},attrs:{checked:e.allchecked[n],indeterminate:e.allindeterminate[n],disabled:0===i.permissions.length},on:{change:function(i){return e.oncheck_all(i,n)}}},[e._v(" 全选 ")])],1),t("div",{staticClass:"cblock clearfix"},[e._l(i.permissions,(function(i,s){return t("div",{key:s,staticClass:"cbox"},[t("a-checkbox",{on:{change:function(i){return e.itemCheck(n)}},model:{value:i.checked,callback:function(t){e.$set(i,"checked",t)},expression:"vo.checked"}},[e._v(e._s(i.name))]),t("a-tag",{attrs:{color:"red"}},[e._v(e._s(i.method))])],1)})),t("div",{staticClass:"clear"})],2)])})),0)]),t("a-form-model-item",{attrs:{label:"备注",prop:"remark",labelCol:{lg:{span:4},sm:{span:4}},wrapperCol:{lg:{span:16},sm:{span:20}}}},[t("a-input",{model:{value:e.info.remark,callback:function(i){e.$set(e.info,"remark",i)},expression:"info.remark"}})],1),t("a-form-model-item",{attrs:{"wrapper-col":{span:12,offset:5}}},[t("a-button",{attrs:{type:"primary"},on:{click:e.handleSubmit}},[e._v("提交")])],1)],1)],1)],1)])],1)},s=[],a=t("2909"),r=(t("922d"),t("09d9")),o=(t("98a7"),t("7bec")),c=(t("17ac"),t("ff57")),l=(t("159b"),t("b0c0"),t("d3b7"),t("99af"),t("c740"),t("a434"),t("4de4"),t("8bbf")),d=t.n(l),u=t("4360");d.a.use(c["a"]),d.a.use(o["a"]),d.a.use(r["a"]);var m={components:{},props:{},data:function(){return{info:{type:0,name:"",status:1,sort:0,key:"",remark:"",menu_ids:[],permission_ids:[]},menu_id_copy:[],child_list:[],parent_list:[],id:0,list:[],permissions:[],allchecked:[],allindeterminate:[],old_menu_ids:[[],[]],loading:!0,rules:{id:[{required:!0,message:"不能空",trigger:"blur"},{type:"integer",min:0,message:"不正确",trigger:"blur"}],type:[{required:!0,message:"不能空",trigger:"blur"},{type:"enum",enum:[0,1],message:"类型不正确",trigger:"blur"}],name:[{required:!0,message:"不能为空",trigger:"blur"},{min:1,max:32,message:"范围 1 - 32 字符",trigger:"blur"}],key:[{required:!0,message:"不能为空",trigger:"blur"},{min:1,max:32,message:"范围 1 - 32 字符",trigger:"blur"},{validator:this.validateKey,trigger:"blur"}],remark:[{min:1,max:255,message:"范围 1 - 255 字符",trigger:"blur"}],sort:[{required:!0,message:"不能空",trigger:"blur"},{type:"integer",min:0,message:"上级不正确",trigger:"blur"}],status:[{required:!0,message:"不能空",trigger:"blur"},{type:"enum",enum:[0,1],message:"类型不正确",trigger:"blur"}]}}},watch:{},computed:{},methods:{setChecked:function(e){if(e.permissions.length>0){for(var i=0;i<e.permissions.length;i++)if(!e.permissions[i].checked)return!1;return!0}return!1},itemCheck:function(e){var i=0;this.permissions[e].permissions.forEach((function(e){e.checked&&i++})),0===i?(this.allchecked[e]=!1,this.allindeterminate[e]=!1):i===this.permissions[e].permissions.length?(this.allchecked[e]=!0,this.allindeterminate[e]=!1):(this.allindeterminate[e]=!0,this.allchecked[e]=!1)},oncheck_all:function(e,i){var t=this,n=e.target.checked;this.allchecked[i]=n,this.allindeterminate[i]=!1,this.permissions[i].permissions.forEach((function(e){t.$isEmpty(e.checked)?t.$set(e,"checked",n):e.checked=n}))},handleSubmit:function(){var e=this;this.$refs.ruleForm.validate((function(i){if(i){if(e.$isEmpty(e.info.name))return e.$message.error("角色名称不能为空");e.get_check_permission();var t=JSON.parse(JSON.stringify(e.info)),n=[];t.menu_ids.forEach((function(e){n.push(e.value)})),t.menu_ids=n,t.id>0?(t.menus=void 0,t.permissions=void 0,u["a"].dispatch("updateRole",t).then((function(i){if(0===i.code)return e.$message.success("保存成功"),e.$router.back();e.$notification.error({message:"错误",description:i.message})})).catch((function(){e.$notification.error({message:"错误",description:"请求信息失败，请重试"})}))):u["a"].dispatch("addRole",t).then((function(i){if(0===i.code)return e.$message.success("保存成功"),e.$router.back();e.$notification.error({message:"错误",description:i.message})})).catch((function(){e.$notification.error({message:"错误",description:"请求信息失败，请重试"})}))}else e.$message.warning("信息输入错误")}))},get_info:function(){var e=this;u["a"].dispatch("showRole",this.id).then((function(i){0!==i.code&&e.$notification.error({message:"错误",description:i.message}),i.data.permission_ids=[],i.data.permissions.forEach((function(e){i.data.permission_ids.push(e.id)})),e.menu_id_copy=i.data.menu_ids,i.data.menu_ids=[],e.info=i.data,e.old_menu_ids[e.info.type]=e.info.menu_ids,u["a"].dispatch("getMenuAll",{type:e.info.type}).then((function(i){0!==i.code&&e.$notification.error({message:"错误",description:i.message}),i.data.forEach((function(i){i.parents=[];var t={value:i.id,lable:i.title};i=e.setParents(i,[t])})),e.list=i.data,e.setMenuID(e.list)})).catch((function(){e.$notification.error({message:"错误",description:"请求信息失败，请重试"})})),u["a"].dispatch("getGroupPermission",{type:e.info.type}).then((function(i){0!==i.code&&e.$notification.error({message:"错误",description:i.message}),e.permissions=i.data,e.permissions.forEach((function(i,t){var n=0;i.permissions.forEach((function(i){e.info.permission_ids.indexOf(i.id)>-1?(n++,e.$set(i,"checked",!0)):e.$set(i,"checked",!1)})),i.permissions.length>0?0===n?(e.allchecked[t]=!1,e.allindeterminate[t]=!1):n===i.permissions.length?(e.allchecked[t]=!0,e.allindeterminate[t]=!1):(e.allindeterminate[t]=!0,e.allchecked[t]=!1):(e.allchecked[t]=!1,e.allindeterminate[t]=!1)})),e.$forceUpdate()})).catch((function(){e.$notification.error({message:"错误",description:"请求信息失败，请重试"})}))})).catch((function(){e.$notification.error({message:"错误",description:"请求信息失败，请重试"})})).finally((function(){e.loading=!1}))},treeChange:function(e,i,t){var n=this,s=t.checked,r=t.triggerNode,o=r.dataRef;if(void 0===s&&(o=r.data.props.dataRef),s)(function(){n.child_list=[],n.parent_list=[],n.getChildren(o.children),n.getParents(n.list,o.id);var e=n.parent_list,i=[].concat(Object(a["a"])(e),[{value:o.id,label:o.title}],Object(a["a"])(n.child_list));if(i)for(var t=function(e){var t=n.info.menu_ids.findIndex((function(t){return t.value===i[e].value}));-1===t&&n.info.menu_ids.push(i[e])},s=0;s<i.length;s++)t(s)})();else{var c=this.info.menu_ids.findIndex((function(e){return e.value===o.id}));c>-1&&(this.info.menu_ids.splice(c,1),this.removeChildren(o.children))}this.old_menu_ids[this.info.type]=this.info.menu_ids},typeChange:function(){var e=this;this.info.menu_ids=this.old_menu_ids[this.info.type],u["a"].dispatch("getMenuAll",{type:this.info.type}).then((function(i){0!==i.code&&e.$notification.error({message:"错误",description:i.message}),i.data.forEach((function(i){i.parents=[];var t={value:i.id,lable:i.title};i=e.setParents(i,[t])})),e.list=i.data})).catch((function(){e.$notification.error({message:"错误",description:"请求信息失败，请重试"})})),u["a"].dispatch("getGroupPermission",{type:this.info.type}).then((function(i){0!==i.code&&e.$notification.error({message:"错误",description:i.message}),e.permissions=i.data,e.permissions.forEach((function(i,t){var n=0;i.permissions.forEach((function(i){e.info.permission_ids.indexOf(i.id)>-1?(n++,e.$set(i,"checked",!0)):e.$set(i,"checked",!1)})),i.permissions.length>0?0===n?(e.allchecked[t]=!1,e.allindeterminate[t]=!1):n===i.permissions.length?(e.allchecked[t]=!0,e.allindeterminate[t]=!1):(e.allindeterminate[t]=!0,e.allchecked[t]=!1):(e.allchecked[t]=!1,e.allindeterminate[t]=!1)})),e.$forceUpdate()})).catch((function(){e.$notification.error({message:"错误",description:"请求信息失败，请重试"})}))},onload:function(){var e=this;this.$isEmpty(this.$route.params.id)?(this.info.type=0,u["a"].dispatch("getMenuAll",{type:this.info.type}).then((function(i){0!==i.code&&e.$notification.error({message:"错误",description:i.message}),i.data.forEach((function(i){i.parents=[];var t={value:i.id,lable:i.title};i=e.setParents(i,[t])})),e.list=i.data})).catch((function(){e.$notification.error({message:"错误",description:"请求信息失败，请重试"})})).finally((function(){e.loading=!1})),u["a"].dispatch("getGroupPermission",{type:this.info.type}).then((function(i){0!==i.code&&e.$notification.error({message:"错误",description:i.message}),e.permissions=i.data})).catch((function(){e.$notification.error({message:"错误",description:"请求信息失败，请重试"})})).finally((function(){e.loading=!1}))):(this.id=this.$route.params.id,this.get_info())},check_all:function(e){var i=this;this.permissions[e].permissions.forEach((function(e){i.$isEmpty(e.checked)?i.$set(e,"checked",!0):e.checked=!e.checked}))},get_check_permission:function(){var e=this,i=[];this.permissions.forEach((function(t){t.permissions.forEach((function(t){!e.$isEmpty(t.checked)&&t.checked&&i.push(t.id)}))})),i.length>0?this.info.permission_ids=i:this.info.permission_ids=[]},setMenuID:function(e){var i=this;e=e||[];for(var t=function(t){var n=i.menu_id_copy.findIndex((function(i){return i===e[t].id}));n>-1&&i.info.menu_ids.push({value:e[t].id,label:e[t].title}),e[t].children&&i.setMenuID(e[t].children)},n=0;n<e.length;n++)t(n)},setParents:function(e,i){var t=this;return e.children&&e.children.length>0&&e.children.forEach((function(e){var n;e.parents=[],(n=e.parents).push.apply(n,Object(a["a"])(i));var s={value:e.id,lable:e.title};e=t.setParents(e,[s].concat(Object(a["a"])(i)))})),e},getParents:function(e,i){e=e||[];var t=e.filter((function(e){return e.id===i}));if(0===t.length)for(var n=0;n<e.length;n++)this.getParents(e[n].children,i);else this.parent_list=t[0].parents},getChildren:function(e){if(e)for(var i=0;i<e.length;i++)this.child_list.push({value:e[i].id,label:e[i].title}),e[i].children&&this.getChildren(e[i].children)},removeChildren:function(e){var i=this;if(e)for(var t=function(t){var n=i.info.menu_ids.findIndex((function(i){return i.value===e[t].id}));n>-1&&i.info.menu_ids.splice(n,1),e[t].children&&i.removeChildren(e[t].children)},n=0;n<e.length;n++)t(n)},validateKey:function(e,i,t,n,s){if(this.info.id>0)return t();u["a"].dispatch("roleKeyExist",{key:i}).then((function(e){return 0!==e.code?t(new Error("关键词已存在")):t()})).catch((function(){return t(new Error("关键词已存在"))}))}},created:function(){this.onload()},mounted:function(){}},p=m,f=(t("27b3"),t("2877")),h=Object(f["a"])(p,n,s,!1,null,"0605701c",null);i["default"]=h.exports}}]);