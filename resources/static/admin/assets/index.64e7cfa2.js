import{C as g,M as i,J as D}from"./index.be3f98b2.js";import{P as F}from"./index.a09675a6.js";import{A as j,bF as p,c5 as b,r as h,a0 as n,B as R,a1 as B,a6 as o,w as a,u as S,v as _,z as P,ae as s}from"./vendor.3850bdb6.js";/* empty css                *//* empty css                */import{_ as k}from"./index.6d3ed4e5.js";import"./useWindowSizeFn.4c2ec928.js";/* empty css               *//* empty css                */import"./useContentViewHeight.34544381.js";const f='{"name":"BeJson","url":"http://www.xxx.com","page":88,"isNonProfit":true,"address":{"street":"\u79D1\u6280\u56ED\u8DEF.","city":"\u6C5F\u82CF\u82CF\u5DDE","country":"\u4E2D\u56FD"},"links":[{"name":"Google","url":"http://www.xxx.com"},{"name":"Baidu","url":"http://www.xxx.com"},{"name":"SoSo","url":"http://www.xxx.com"}]}',y=`
      (() => {
        var htmlRoot = document.getElementById('htmlRoot');
        var theme = window.localStorage.getItem('__APP__DARK__MODE__');
        if (htmlRoot && theme) {
          htmlRoot.setAttribute('data-theme', theme);
          theme = htmlRoot = null;
        }
      })();
  `,A=`
     <!DOCTYPE html>
<html lang="en" id="htmlRoot">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1" />
    <meta name="renderer" content="webkit" />
    <meta
      name="viewport"
      content="width=device-width,initial-scale=1.0,minimum-scale=1.0,maximum-scale=1.0,user-scalable=0"
    />
    <title><%= title %></title>
    <link rel="icon" href="/favicon.ico" />
  </head>
  <body>
    <div id="app">
    </div>
  </body>
</html>
  `,J=j({components:{CodeEditor:g,PageWrapper:F,RadioButton:p.Button,RadioGroup:p.Group,ASpace:b},setup(){const t=h(i.JSON),e=h(f);function r(c){const u=c.target.value;if(u===i.JSON){e.value=f;return}if(u===i.HTML){e.value=A;return}if(u===i.JS){e.value=y;return}}function m(){S(t)==="application/json"?_.info({title:"\u7F16\u8F91\u5668\u5F53\u524D\u503C",content:P(D,{data:JSON.parse(e.value)})}):_.info({title:"\u7F16\u8F91\u5668\u5F53\u524D\u503C",content:e.value})}return{value:e,modeValue:t,handleModeChange:r,showData:m}}}),M=s("\u83B7\u53D6\u6570\u636E"),V=s(" json\u6570\u636E "),N=s(" html\u4EE3\u7801 "),O=s(" javascript\u4EE3\u7801 ");function G(t,e,r,m,c,u){const v=n("a-button"),l=n("RadioButton"),w=n("RadioGroup"),x=n("a-space"),C=n("CodeEditor"),E=n("PageWrapper");return R(),B(E,{title:"\u4EE3\u7801\u7F16\u8F91\u5668\u7EC4\u4EF6\u793A\u4F8B",contentFullHeight:"",fixedHeight:"",contentBackground:""},{extra:o(()=>[a(x,{size:"middle"},{default:o(()=>[a(v,{onClick:t.showData,type:"primary"},{default:o(()=>[M]),_:1},8,["onClick"]),a(w,{"button-style":"solid",value:t.modeValue,"onUpdate:value":e[0]||(e[0]=d=>t.modeValue=d),onChange:t.handleModeChange},{default:o(()=>[a(l,{value:"application/json"},{default:o(()=>[V]),_:1}),a(l,{value:"htmlmixed"},{default:o(()=>[N]),_:1}),a(l,{value:"javascript"},{default:o(()=>[O]),_:1})]),_:1},8,["value","onChange"])]),_:1})]),default:o(()=>[a(C,{value:t.value,"onUpdate:value":e[1]||(e[1]=d=>t.value=d),mode:t.modeValue},null,8,["value","mode"])]),_:1})}var X=k(J,[["render",G]]);export{X as default};
