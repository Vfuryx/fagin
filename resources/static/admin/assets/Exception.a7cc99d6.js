import{A as N,r as O,bb as R,j as _,u as s,w as r,aT as A,bc as S}from"./vendor.3850bdb6.js";/* empty css               */import{E as e,l as v,m as G,a as P,P as m,b as k}from"./index.6d3ed4e5.js";var C="/static/admin/assets/no-data.f7e550cc.svg",D="/static/admin/assets/net-error.61b7e6df.svg",w=N({name:"ErrorPage",props:{status:{type:Number,default:e.PAGE_NOT_FOUND},title:{type:String,default:""},subTitle:{type:String,default:""},full:{type:Boolean,default:!1}},setup(n){const a=O(new Map),{query:x}=R(),o=v(),c=G(),{t}=k(),{prefixCls:f}=P("app-exception-page"),p=_(()=>{const{status:i}=x,{status:l}=n;return Number(i)||l}),E=_(()=>s(a).get(s(p))),b=t("sys.exception.backLogin"),u=t("sys.exception.backHome");return s(a).set(e.PAGE_NOT_ACCESS,{title:"403",status:`${e.PAGE_NOT_ACCESS}`,subTitle:t("sys.exception.subTitle403"),btnText:n.full?b:u,handler:()=>n.full?o(m.BASE_LOGIN):o()}),s(a).set(e.PAGE_NOT_FOUND,{title:"404",status:`${e.PAGE_NOT_FOUND}`,subTitle:t("sys.exception.subTitle404"),btnText:n.full?b:u,handler:()=>n.full?o(m.BASE_LOGIN):o()}),s(a).set(e.ERROR,{title:"500",status:`${e.ERROR}`,subTitle:t("sys.exception.subTitle500"),btnText:u,handler:()=>o()}),s(a).set(e.PAGE_NOT_DATA,{title:t("sys.exception.noDataTitle"),subTitle:"",btnText:t("common.redo"),handler:()=>c(),icon:C}),s(a).set(e.NET_WORK_ERROR,{title:t("sys.exception.networkErrorTitle"),subTitle:t("sys.exception.networkErrorSubTitle"),btnText:t("common.redo"),handler:()=>c(),icon:D}),()=>{const{title:i,subTitle:l,btnText:d,icon:T,handler:g,status:y}=s(E)||{};return r(S,{class:f,status:y,title:n.title||i,"sub-title":n.subTitle||l},{extra:()=>d&&r(A,{type:"primary",onClick:g},{default:()=>d}),icon:()=>T?r("img",{src:T},null):null})}}});export{w as default};