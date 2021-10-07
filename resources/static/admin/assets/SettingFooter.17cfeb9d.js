import{J as e,a as t,Z as a,f as s,V as o,e as n,b0 as l,a7 as r,aV as c,aW as i,i as d}from"./index.7037d70a.js";import{D as u,cg as g,bm as f,u as p,a0 as C,F as y,G as m,w as S,a5 as k,ad as h,L as R,N as b}from"./vendor.5da97255.js";var x=u({name:"SettingFooter",components:{CopyOutlined:g,RedoOutlined:f},setup(){const u=e(),{prefixCls:g}=t("setting-footer"),{t:f}=n(),{createSuccessModal:C,createMessage:y}=d(),m=a(),S=s(),k=o();return{prefixCls:g,t:f,handleCopy:function(){const{isSuccessRef:e}=l(JSON.stringify(p(k.getProjectConfig),null,2));p(e)&&C({title:f("layout.setting.operatingTitle"),content:f("layout.setting.operatingContent")})},handleResetSetting:function(){try{k.setProjectConfig(r);const{colorWeak:e,grayMode:t}=r;c(e),i(t),y.success(f("layout.setting.resetSuccess"))}catch(e){y.error(e)}},handleClearAndRedo:function(){localStorage.clear(),k.resetAllState(),u.resetState(),m.resetState(),S.resetState(),location.reload()}}}});x.render=function(e,t,a,s,o,n){const l=C("CopyOutlined"),r=C("a-button"),c=C("RedoOutlined");return y(),m("div",{class:b(e.prefixCls)},[S(r,{type:"primary",block:"",onClick:e.handleCopy},{default:k((()=>[S(l,{class:"mr-2"}),h(" "+R(e.t("layout.setting.copyBtn")),1)])),_:1},8,["onClick"]),S(r,{color:"warning",block:"",onClick:e.handleResetSetting,class:"my-3"},{default:k((()=>[S(c,{class:"mr-2"}),h(" "+R(e.t("common.resetText")),1)])),_:1},8,["onClick"]),S(r,{color:"error",block:"",onClick:e.handleClearAndRedo},{default:k((()=>[S(c,{class:"mr-2"}),h(" "+R(e.t("layout.setting.clearBtn")),1)])),_:1},8,["onClick"])],2)},x.__scopeId="data-v-755ad017";export{x as default};
