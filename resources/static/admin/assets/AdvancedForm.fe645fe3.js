import{B as d}from"./BasicForm.f9b68d8c.js";import{u as l}from"./useForm.ff740053.js";import{_ as m,ag as B}from"./index.6d3ed4e5.js";import{P as f}from"./index.a09675a6.js";import{A as j,a0 as s,B as b,a1 as x,a6 as i,w as n}from"./vendor.3850bdb6.js";/* empty css               *//* empty css               *//* empty css               *//* empty css                *//* empty css              *//* empty css               *//* empty css                *//* empty css                */import"./index.f4e2dde3.js";/* empty css                *//* empty css                *//* empty css                */import"./index.ab15df17.js";import"./useWindowSizeFn.4c2ec928.js";/* empty css                */import"./uuid.2b29000c.js";import"./download.62bc70aa.js";import"./base64Conver.08b9f4ec.js";import"./index.c99e5782.js";/* empty css               *//* empty css                */import"./useContentViewHeight.34544381.js";const u=()=>[{field:"field1",component:"Input",label:"\u5B57\u6BB51",colProps:{span:8},componentProps:{placeholder:"\u81EA\u5B9A\u4E49placeholder",onChange:e=>{console.log(e)}}},{field:"field2",component:"Input",label:"\u5B57\u6BB52",colProps:{span:8}},{field:"field3",component:"DatePicker",label:"\u5B57\u6BB53",colProps:{span:8}},{field:"field4",component:"Select",label:"\u5B57\u6BB54",colProps:{span:8},componentProps:{options:[{label:"\u9009\u98791",value:"1",key:"1"},{label:"\u9009\u98792",value:"2",key:"2"}]}},{field:"field5",component:"CheckboxGroup",label:"\u5B57\u6BB55",colProps:{span:8},componentProps:{options:[{label:"\u9009\u98791",value:"1"},{label:"\u9009\u98792",value:"2"}]}}];function C(){return[{field:"field10",component:"Input",label:"\u5B57\u6BB510",colProps:{span:8}},{field:"field11",component:"Input",label:"\u5B57\u6BB511",colProps:{span:8}},{field:"field12",component:"Input",label:"\u5B57\u6BB512",colProps:{span:8}},{field:"field13",component:"Input",label:"\u5B57\u6BB513",colProps:{span:8}}]}const F=j({components:{BasicForm:d,CollapseContainer:B,PageWrapper:f},setup(){const[e]=l({labelWidth:120,schemas:u(),actionColOptions:{span:24},compact:!0,showAdvancedButton:!0}),t=[];for(let o=14;o<30;o++)t.push({field:"field"+o,component:"Input",label:"\u5B57\u6BB5"+o,colProps:{span:8}});const[p]=l({labelWidth:120,schemas:[...u(),...C(),{field:"",component:"Divider",label:"\u66F4\u591A\u5B57\u6BB5"},...t],actionColOptions:{span:24},compact:!0,showAdvancedButton:!0,alwaysShowLines:2});return{register:e,register1:p}}});function h(e,t,p,o,P,g){const r=s("BasicForm"),a=s("CollapseContainer"),c=s("PageWrapper");return b(),x(c,{title:"\u53EF\u6298\u53E0\u8868\u5355\u793A\u4F8B"},{default:i(()=>[n(a,{title:"\u57FA\u7840\u6536\u7F29\u793A\u4F8B"},{default:i(()=>[n(r,{onRegister:e.register},null,8,["onRegister"])]),_:1}),n(a,{title:"\u8D85\u8FC73\u884C\u81EA\u52A8\u6536\u8D77\uFF0C\u6298\u53E0\u65F6\u4FDD\u75592\u884C",class:"mt-4"},{default:i(()=>[n(r,{onRegister:e.register1},null,8,["onRegister"])]),_:1})]),_:1})}var X=m(F,[["render",h]]);export{X as default};
