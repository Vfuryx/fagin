var e=Object.defineProperty,t=Object.getOwnPropertySymbols,r=Object.prototype.hasOwnProperty,a=Object.prototype.propertyIsEnumerable,o=(t,r,a)=>r in t?e(t,r,{enumerable:!0,configurable:!0,writable:!0,value:a}):t[r]=a;import{u as s}from"./useECharts.c30f132a.js";import{b as i}from"./props.f48aca0b.js";import{D as p,r as l,_ as n,F as b,G as c,X as f}from"./vendor.5da97255.js";var u=p({props:((e,s)=>{for(var i in s||(s={}))r.call(s,i)&&o(e,i,s[i]);if(t)for(var i of t(s))a.call(s,i)&&o(e,i,s[i]);return e})({},i),setup(e){const t=l(null),{setOptions:r}=s(t);return n((()=>{r({tooltip:{trigger:"axis",axisPointer:{lineStyle:{width:1,color:"#019680"}}},grid:{left:"1%",right:"1%",top:"2  %",bottom:0,containLabel:!0},xAxis:{type:"category",data:["1月","2月","3月","4月","5月","6月","7月","8月","9月","10月","11月","12月"]},yAxis:{type:"value",max:8e3,splitNumber:4},series:[{data:[3e3,2e3,3333,5e3,3200,4200,3200,2100,3e3,5100,6e3,3200,4800],type:"bar",barMaxWidth:80}]})})),(e,r)=>(b(),c("div",{ref:(e,r)=>{r.chartRef=e,t.value=e},style:f({height:e.height,width:e.width})},null,4))}});export{u as _};