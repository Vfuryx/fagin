import{P as h}from"./index.a09675a6.js";import{_ as A,ag as I,f as j}from"./index.6d3ed4e5.js";import{C as B,a as E}from"./index.77d56d10.js";import{u as y}from"./upload.920018c1.js";import{h as D}from"./header.d801b988.js";import{A as w,r as s,a0 as p,B as t,a1 as k,a6 as u,w as r,H as i,D as c,ad as d,J as F}from"./vendor.3850bdb6.js";/* empty css               *//* empty css                */import"./useWindowSizeFn.4c2ec928.js";import"./useContentViewHeight.34544381.js";/* empty css                *//* empty css                *//* empty css                */import"./index.ab15df17.js";import"./base64Conver.08b9f4ec.js";const b=w({components:{PageWrapper:h,CropperImage:B,CollapseContainer:I,CropperAvatar:E},setup(){var n;const e=s(""),l=s(""),m=s(""),C=s(""),v=j(),g=s(((n=v.getUserInfo)==null?void 0:n.avatar)||"");function _({imgBase64:a,imgInfo:f}){e.value=f,l.value=a}function o({imgBase64:a,imgInfo:f}){m.value=f,C.value=a}return{img:D,info:e,circleInfo:m,cropperImg:l,circleImg:C,handleCropend:_,handleCircleCropend:o,avatar:g,uploadApi:y}}}),P={class:"container p-4"},S={class:"cropper-container mr-10"},V=["src"],W={key:0},$={class:"container p-4"},N={class:"cropper-container mr-10"},x=["src"],H={key:0};function U(e,l,m,C,v,g){const _=p("CropperAvatar"),o=p("CollapseContainer"),n=p("CropperImage"),a=p("PageWrapper");return t(),k(a,{title:"\u56FE\u7247\u88C1\u526A\u793A\u4F8B",content:"\u9700\u8981\u5F00\u542F\u6D4B\u8BD5\u63A5\u53E3\u670D\u52A1\u624D\u80FD\u8FDB\u884C\u4E0A\u4F20\u6D4B\u8BD5\uFF01"},{default:u(()=>[r(o,{title:"\u5934\u50CF\u88C1\u526A"},{default:u(()=>[r(_,{uploadApi:e.uploadApi,value:e.avatar},null,8,["uploadApi","value"])]),_:1}),r(o,{title:"\u77E9\u5F62\u88C1\u526A",class:"my-4"},{default:u(()=>[i("div",P,[i("div",S,[r(n,{ref:"refCropper",src:e.img,onCropend:e.handleCropend,style:{width:"40vw"}},null,8,["src","onCropend"])]),e.cropperImg?(t(),c("img",{key:0,src:e.cropperImg,class:"croppered",alt:""},null,8,V)):d("",!0)]),e.cropperImg?(t(),c("p",W,"\u88C1\u526A\u540E\u56FE\u7247\u4FE1\u606F\uFF1A"+F(e.info),1)):d("",!0)]),_:1}),r(o,{title:"\u5706\u5F62\u88C1\u526A"},{default:u(()=>[i("div",$,[i("div",N,[r(n,{ref:"refCropper",src:e.img,onCropend:e.handleCircleCropend,style:{width:"40vw"},circled:""},null,8,["src","onCropend"])]),e.circleImg?(t(),c("img",{key:0,src:e.circleImg,class:"croppered"},null,8,x)):d("",!0)]),e.circleImg?(t(),c("p",H,"\u88C1\u526A\u540E\u56FE\u7247\u4FE1\u606F\uFF1A"+F(e.circleInfo),1)):d("",!0)]),_:1})]),_:1})}var re=A(b,[["render",U],["__scopeId","data-v-9d983438"]]);export{re as default};
