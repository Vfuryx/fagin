var p=(i,u,n)=>new Promise((o,s)=>{var m=e=>{try{a(n.next(e))}catch(t){s(t)}},d=e=>{try{a(n.throw(e))}catch(t){s(t)}},a=e=>e.done?o(e.value):Promise.resolve(e.value).then(m,d);a((n=n.apply(i,u)).next())});import{B as C}from"./TableImg.1321f9ef.js";import{T as b}from"./BasicForm.f9b68d8c.js";import{u as x}from"./useTable.5a3ce825.js";import{o as w,t as j}from"./tree.1644d3df.js";import{d as D}from"./table.23a47f9e.js";import{_ as E,i as B}from"./index.6d3ed4e5.js";import{A as h,r as g,a0 as F,B as v,D as R,w as f,a6 as A,n as I}from"./vendor.3850bdb6.js";/* empty css                *//* empty css              */import"./useForm.ff740053.js";import"./index.a09675a6.js";/* empty css               *//* empty css                */import"./useWindowSizeFn.4c2ec928.js";import"./useContentViewHeight.34544381.js";/* empty css                *//* empty css                *//* empty css               */import"./uuid.2b29000c.js";import"./index.ab15df17.js";/* empty css               */import"./useSortable.7a35deac.js";/* empty css                *//* empty css                *//* empty css               *//* empty css               */import"./index.f4e2dde3.js";/* empty css                *//* empty css                *//* empty css                */import"./download.62bc70aa.js";import"./base64Conver.08b9f4ec.js";import"./index.c99e5782.js";const T=[{title:"\u8F93\u5165\u6846",dataIndex:"name",editRow:!0,editComponentProps:{prefix:"$"},width:150},{title:"\u9ED8\u8BA4\u8F93\u5165\u72B6\u6001",dataIndex:"name7",editRow:!0,width:150},{title:"\u8F93\u5165\u6846\u6821\u9A8C",dataIndex:"name1",editRow:!0,align:"left",editRule:!0,width:150},{title:"\u8F93\u5165\u6846\u51FD\u6570\u6821\u9A8C",dataIndex:"name2",editRow:!0,align:"right",editRule:i=>p(void 0,null,function*(){return i==="2"?"\u4E0D\u80FD\u8F93\u5165\u8BE5\u503C":""})},{title:"\u6570\u5B57\u8F93\u5165\u6846",dataIndex:"id",editRow:!0,editRule:!0,editComponent:"InputNumber",width:150},{title:"\u4E0B\u62C9\u6846",dataIndex:"name3",editRow:!0,editComponent:"Select",editComponentProps:{options:[{label:"Option1",value:"1"},{label:"Option2",value:"2"},{label:"Option3",value:"3"}]},width:200},{title:"\u8FDC\u7A0B\u4E0B\u62C9",dataIndex:"name4",editRow:!0,editComponent:"ApiSelect",editComponentProps:{api:w,resultField:"list",labelField:"name",valueField:"id"},width:200},{title:"\u8FDC\u7A0B\u4E0B\u62C9\u6811",dataIndex:"name8",editRow:!0,editComponent:"ApiTreeSelect",editRule:!1,editComponentProps:{api:j,resultField:"list"},width:200},{title:"\u65E5\u671F\u9009\u62E9",dataIndex:"date",editRow:!0,editComponent:"DatePicker",editComponentProps:{valueFormat:"YYYY-MM-DD",format:"YYYY-MM-DD"},width:150},{title:"\u65F6\u95F4\u9009\u62E9",dataIndex:"time",editRow:!0,editComponent:"TimePicker",editComponentProps:{valueFormat:"HH:mm",format:"HH:mm"},width:100},{title:"\u52FE\u9009\u6846",dataIndex:"name5",editRow:!0,editComponent:"Checkbox",editValueMap:i=>i?"\u662F":"\u5426",width:100},{title:"\u5F00\u5173",dataIndex:"name6",editRow:!0,editComponent:"Switch",editValueMap:i=>i?"\u5F00":"\u5173",width:100}],k=h({components:{BasicTable:C,TableAction:b},setup(){const{createMessage:i}=B(),u=g(""),[n]=x({title:"\u53EF\u7F16\u8F91\u884C\u793A\u4F8B",titleHelpMessage:["\u672C\u4F8B\u4E2D\u4FEE\u6539[\u6570\u5B57\u8F93\u5165\u6846]\u8FD9\u4E00\u5217\u65F6\uFF0C\u540C\u4E00\u884C\u7684[\u8FDC\u7A0B\u4E0B\u62C9]\u5217\u7684\u5F53\u524D\u7F16\u8F91\u6570\u636E\u4E5F\u4F1A\u540C\u6B65\u53D1\u751F\u6539\u53D8"],api:D,columns:T,showIndexColumn:!1,showTableSetting:!0,tableSetting:{fullScreen:!0},actionColumn:{width:160,title:"Action",dataIndex:"action",slots:{customRender:"action"}}});function o(e){var t;u.value=e.key,(t=e.onEdit)==null||t.call(e,!0)}function s(e){var t;u.value="",(t=e.onEdit)==null||t.call(e,!1,!1)}function m(e){return p(this,null,function*(){var l,r;if(i.loading({content:"\u6B63\u5728\u4FDD\u5B58...",duration:0,key:"saving"}),yield(l=e.onValid)==null?void 0:l.call(e))try{const c=I(e.editValueRefs);console.log(c),(yield(r=e.onEdit)==null?void 0:r.call(e,!1,!0))&&(u.value=""),i.success({content:"\u6570\u636E\u5DF2\u4FDD\u5B58",key:"saving"})}catch(c){i.error({content:"\u4FDD\u5B58\u5931\u8D25",key:"saving"})}else i.error({content:"\u8BF7\u586B\u5199\u6B63\u786E\u7684\u6570\u636E",key:"saving"})})}function d(e,t){return e.editable?[{label:"\u4FDD\u5B58",onClick:m.bind(null,e,t)},{label:"\u53D6\u6D88",popConfirm:{title:"\u662F\u5426\u53D6\u6D88\u7F16\u8F91",confirm:s.bind(null,e,t)}}]:[{label:"\u7F16\u8F91",disabled:u.value?u.value!==e.key:!1,onClick:o.bind(null,e)}]}function a({column:e,value:t,record:l}){e.dataIndex==="id"&&(l.editValueRefs.name4.value=`${t}`),console.log(e,t,l)}return{registerTable:n,handleEdit:o,createActions:d,onEditChange:a}}}),_={class:"p-4"};function y(i,u,n,o,s,m){const d=F("TableAction"),a=F("BasicTable");return v(),R("div",_,[f(a,{onRegister:i.registerTable,onEditChange:i.onEditChange},{action:A(({record:e,column:t})=>[f(d,{actions:i.createActions(e,t)},null,8,["actions"])]),_:1},8,["onRegister","onEditChange"])])}var fe=E(k,[["render",y]]);export{fe as default};
