import{x as c}from"./index.5770a65a.js";const{utils:i,writeFile:d}=c,u="excel-list.xlsx";function w({data:e,header:t,filename:a=u,json2sheetOpts:o={},write2excelOpts:s={bookType:"xlsx"}}){const n=[...e];t&&(n.unshift(t),o.skipHeader=!0);const r=i.json_to_sheet(n,o),x={SheetNames:[a],Sheets:{[a]:r}};d(x,a,s)}function F({data:e,header:t,filename:a=u,write2excelOpts:o={bookType:"xlsx"}}){const s=[...e];t&&s.unshift(t);const n=i.aoa_to_sheet(s),r={SheetNames:[a],Sheets:{[a]:n}};d(r,a,o)}const l=[{title:"ID",dataIndex:"id",width:80},{title:"\u59D3\u540D",dataIndex:"name",width:120},{title:"\u5E74\u9F84",dataIndex:"age",width:80},{title:"\u7F16\u53F7",dataIndex:"no",width:80},{title:"\u5730\u5740",dataIndex:"address"},{title:"\u5F00\u59CB\u65F6\u95F4",dataIndex:"beginTime"},{title:"\u7ED3\u675F\u65F6\u95F4",dataIndex:"endTime"}],k=(()=>{const e=[];for(let t=0;t<40;t++)e.push({id:`${t}`,name:`${t} John Brown`,age:`${t+10}`,no:`${t}98678`,address:"New York No. 1 Lake ParkNew York No. 1 Lake Park",beginTime:new Date().toLocaleString(),endTime:new Date().toLocaleString()});return e})(),D=l.map(e=>e.title),m=k.map(e=>Object.keys(e).map(t=>e[t]));export{F as a,m as b,l as c,k as d,D as e,w as j};
