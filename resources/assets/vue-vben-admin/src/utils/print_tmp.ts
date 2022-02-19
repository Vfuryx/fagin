import XEUtils from 'xe-utils';

/**
 * 打印参数
 */
export interface PrintConfig {
  /**
   * 表名
   */
  sheetName?: string;
  /**
   * 打印样式
   */
  style?: string;
  /**
   * 自定义打印内容
   */
  content: string | Blob;
  /**
   * Content-type
   */
  type?: string;
  /**
   * 打印之前的方法，可以通过返回自定义打印的内容
   */
  beforePrintMethod?(params: { content: string; options: PrintConfig }): string;
}

// 打印
let printFrame: any;

// 默认导出或打印的 HTML 样式
const defaultHtmlStyle =
  'body{margin:0;color:#333333;font-size:14px;font-family:"Microsoft YaHei",微软雅黑,"MicrosoftJhengHei",华文细黑,STHeiti,MingLiu}body *{-webkit-box-sizing:border-box;box-sizing:border-box}.vxe-table{border-collapse:collapse;text-align:left;border-spacing:0}.vxe-table:not(.is--print){table-layout:fixed}.vxe-table,.vxe-table th,.vxe-table td,.vxe-table td{border-color:#D0D0D0;border-style:solid;border-width:0}.vxe-table.is--print{width:100%}.border--default,.border--full,.border--outer{border-top-width:1px}.border--default,.border--full,.border--outer{border-left-width:1px}.border--outer,.border--default th,.border--default td,.border--full th,.border--full td,.border--outer th,.border--inner th,.border--inner td{border-bottom-width:1px}.border--default,.border--outer,.border--full th,.border--full td{border-right-width:1px}.border--default th,.border--full th,.border--outer th{background-color:#f8f8f9}.vxe-table td>div,.vxe-table th>div{padding:.5em .4em}.col--center{text-align:center}.col--right{text-align:right}.vxe-table:not(.is--print) .col--ellipsis>div{overflow:hidden;text-overflow:ellipsis;white-space:nowrap;word-break:break-all}.vxe-table--tree-node{text-align:left}.vxe-table--tree-node-wrapper{position:relative}.vxe-table--tree-icon-wrapper{position:absolute;top:50%;width:1em;height:1em;text-align:center;-webkit-transform:translateY(-50%);transform:translateY(-50%);-webkit-user-select:none;-moz-user-select:none;-ms-user-select:none;user-select:none;cursor:pointer}.vxe-table--tree-unfold-icon,.vxe-table--tree-fold-icon{position:absolute;width:0;height:0;border-style:solid;border-width:.5em;border-right-color:transparent;border-bottom-color:transparent}.vxe-table--tree-unfold-icon{left:.3em;top:0;border-left-color:#939599;border-top-color:transparent}.vxe-table--tree-fold-icon{left:0;top:.3em;border-left-color:transparent;border-top-color:#939599}.vxe-table--tree-cell{display:block;padding-left:1.5em}.vxe-table input[type="checkbox"]{margin:0}.vxe-table input[type="checkbox"],.vxe-table input[type="radio"],.vxe-table input[type="checkbox"]+span,.vxe-table input[type="radio"]+span{vertical-align:middle;padding-left:0.4em}';

export const browse = XEUtils.browse();

export function createHtmlPage(opts: PrintConfig, content: string): string {
  const { style } = opts;
  return [
    '<!DOCTYPE html>',
    `<head><title>${opts.sheetName}</title>`,
    '<meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1,minimum-scale=1,maximum-scale=1,user-scalable=no,minimal-ui">',
    `<style>${defaultHtmlStyle}</style>`,
    style ? `<style>${style}</style>` : '',
    '</head>',
    `<body>
  <button>打印</button>
${content}
</body>`,
    '</html>',
  ].join('');
}

export function getExportBlobByContent(content: string, options: { type: string }): Blob | null {
  if (window.Blob) {
    return new Blob([content], { type: `${options.type};charset=utf-8;` });
  }
  return null;
}

export function createFrame(): HTMLIFrameElement {
  const frame = document.createElement('iframe');
  frame.className = 'vxe-table--print-frame';
  frame.style.position = 'fixed';
  frame.style.zIndex = '99';
  frame.style.top = '50%';
  frame.style.left = '50%';
  frame.style.width = '50%';
  frame.style.height = '500px';
  frame.style.transform = 'translate(-50%,-50%)';
  frame.style.backgroundColor = '#fff';
  return frame;
}

function removePrintFrame() {
  if (printFrame) {
    if (printFrame.parentNode) {
      try {
        printFrame.contentDocument.write('');
      } catch (e) {}
      printFrame.parentNode.removeChild(printFrame);
    }
    printFrame = null;
  }
}

function appendPrintFrame() {
  if (!printFrame.parentNode) {
    document.body.appendChild(printFrame);
  }
}

function afterPrintEvent() {
  requestAnimationFrame(removePrintFrame);
}

function iframeOnload(event: any) {
  if (event.target.src) {
    event.target.contentWindow.onafterprint = afterPrintEvent;

    // event.target.contentWindow.print();
  }
}

export function print(opts: PrintConfig): void {
  let content = opts.content;
  // 清除frame
  removePrintFrame();

  printFrame = createFrame();

  // type 默认 html
  if (!opts.type) {
    opts.type = 'text/html';
  }
  // Blob 类型直接打印
  if (content instanceof Blob) {
    // 延时清除frame 5分钟
    setTimeout(function () {
      afterPrintEvent();
    }, 300000);

    printFrame.onload = iframeOnload;
    appendPrintFrame();
    printFrame.contentDocument.write(
      `
<div style=" width: 100%; height: 5%; ">
<button>打印</button><select><option>打印机1</option></select>
</div>
<embed src="${URL.createObjectURL(content)}" style=" width: 100%; height: 95%; " />
`,
    );

    console.log(printFrame);
    // printFrame.contentWindow.contentDocument.appendChild(embed);

    // printFrame.src = ;
    return;
  }

  const { beforePrintMethod } = opts;
  if (beforePrintMethod) {
    content = beforePrintMethod({ content, options: opts }) || '';
  }
  content = createHtmlPage(opts, content);
  const blob = getExportBlobByContent(content, { type: opts.type });

  if (browse.msie || browse.safari) {
    appendPrintFrame();
    printFrame.contentDocument.write(content);
    printFrame.contentDocument.execCommand('print');
  } else {
    printFrame.onload = iframeOnload;
    appendPrintFrame();
    printFrame.src = URL.createObjectURL(blob);
  }
}

/**
 * 例子
 * 可以产考 https://vxetable.cn/v4/#/table/module/print
 */
export function example() {
  // 打印样式
  const printStyle = `
              .page-1 {
                height: 1000px;
              }
              .page-2 {
                padding: 15px 0;
              }
              .fill-row {
                display: block;
                font-size: 14px;
                height: 36px;
              }
              .fill-span {
                display: inline-block;
                font-size: 14px;
                height: 36px;
              }
              .fill-title {
                display: inline-block;
                vertical-align: middle;
              }
              .fill-empty,
              .fill-part {
                display: inline-block;
                vertical-align: bottom;
                border-bottom: 1px solid #000;
              }
              .number {
                width: 250px;
                margin-top: 40px;
              }
              .number .fill-empty {
                width: 160px;
              }
              .title {
                text-align: center;
                margin: 80px 0;
              }
              .info-a,
              .info-b {
                margin: 0 auto;
                width: 400px;
                text-align: right;
              }
              .info-a .fill-row,
              .info-b .fill-row {
                height: 48px;
              }
              .info-a .fill-empty,
              .info-b .fill-empty {
                width: 200px;
              }
              .info-b {
                margin-top: 80px;
              }
              .list-desc {
                padding-left: 15px;
              }
              .title {
                text-align: center;
              }
              .my-list-row {
                display: inline-block;
                width: 100%;
              }
              .my-list-col {
                float: left;
                width: 33.33%;
                height: 28px;
                line-height: 28px;
              }
              .my-top,
              .my-bottom {
                font-size: 12px;
              }
              .my-top {
                margin-bottom: 5px;
              }
              .my-bottom {
                margin-top: 30px;
                text-align: right;
              }
              `;
  // 打印模板
  const printTmpl = `
              <div class="page-1">
                <div class="fill-row number">
                  <span class="fill-title">编号：</span>
                  <span class="fill-empty"></span>
                </div>
                <h1 class="title">劳动合同书</h1>
                <div class="info-a">
                  <div class="fill-row">
                    <span class="fill-title">甲方（用人单位名称）名称：</span>
                    <span class="fill-empty"></span>
                  </div>
                  <div class="fill-row">
                    <span class="fill-title">住址：</span>
                    <span class="fill-empty"></span>
                  </div>
                  <div class="fill-row">
                    <span class="fill-title">法定代表人（委托代理人）：</span>
                    <span class="fill-empty"></span>
                  </div>
                </div>
                <div class="info-b">
                  <div class="fill-row">
                    <span class="fill-title">乙方（劳动者）姓名：</span>
                    <span class="fill-empty"></span>
                  </div>
                  <div class="fill-row">
                    <span class="fill-title">性别：</span>
                    <span class="fill-empty"></span>
                  </div>
                  <div class="fill-row">
                    <span class="fill-title">住址：</span>
                    <span class="fill-empty"></span>
                  </div>
                  <div class="fill-row">
                    <span class="fill-title">居民身份证号：</span>
                    <span class="fill-empty"></span>
                  </div>
                  <div class="fill-row">
                    <span class="fill-title">联系电话：</span>
                    <span class="fill-empty"></span>
                  </div>
                </div>
              </div>

              <div class="page-2">
                <p>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;甲乙双方根据《中国人民共和国劳动合同法》等法律、法规、规章的规定，在平等、自愿、协商一致的基础上，同意订立本劳动合同，共同遵守本合同所列条款。</p>
                <h2>一：合同类型的期限</h2>
                <div class="list-desc">
                  <p>第一条 甲、乙双方选择以下第<span class="fill-part" style="width: 100px"></span>等形式确定本合同期限：</p>
                  <div class="list-desc">
                    <p>1、固定期限：自<span class="fill-part" style="width: 60px"></span>年<span class="fill-part" style="width: 40px"></span>月<span class="fill-part" style="width: 40px"></span>日起至<span class="fill-part" style="width: 60px"></span>年<span class="fill-part" style="width: 40px"></span>月<span class="fill-part" style="width: 40px"></span>日止。</p>
                    <p>2、无固定期限：自<span class="fill-part" style="width: 60px"></span>年<span class="fill-part" style="width: 40px"></span>月<span class="fill-part" style="width: 40px"></span>日起至<span class="fill-part" style="width: 60px"></span>年<span class="fill-part" style="width: 40px"></span>月<span class="fill-part" style="width: 40px"></span>日止</p>
                    <p>3、以完成一定的工作（任务）为期限：自<span class="fill-part" style="width: 60px"></span>年<span class="fill-part" style="width: 40px"></span>月<span class="fill-part" style="width: 40px"></span>日起至工作（任务）完成时即行终止。</p>
                    <p>双方约定的试用期限<span class="fill-part" style="width: 60px"></span>年<span class="fill-part" style="width: 40px"></span>月<span class="fill-part" style="width: 40px"></span>日只，期限为<span class="fill-part" style="width: 40px"></span>月</p>
                  </div>
                </div>
                <h2>二：工作内容和工作地点</h2>
                <div class="list-desc">...省略</div>
                <h2>三：工作时间和休息休假</h2>
                <div class="list-desc">...省略</div>
                <div style="margin-top: 15px;">如果对您有帮助，点击右上角支持我们吧！</div>
              </div>
              `;
  // 打印顶部内容模板
  const topHtml = `
            <h1 class="title">出货单据</h1>
            <div class="my-top">
              <div class="my-list-row">
                <div class="my-list-col">商品名称：vxe-table</div>
                <div class="my-list-col">发货单号：X2665847132654</div>
                <div class="my-list-col">发货日期：2020-09-20</div>
              </div>
              <div class="my-list-row">
                <div class="my-list-col">收货姓名：小徐</div>
                <div class="my-list-col">收货地址：火星第七区18号001</div>
                <div class="my-list-col">联系电话：10086</div>
              </div>
            </div>
            `;

  // 打印底部内容模板
  const bottomHtml = `
            <div class="my-bottom">
              <div class="my-list-row">
                <div class="my-list-col"></div>
                <div class="my-list-col">创建人：小徐</div>
                <div class="my-list-col">创建日期：2020-09-20</div>
              </div>
            </div>
            `;
  print({
    sheetName: '打印自定义模板',
    style: printStyle,
    content: printTmpl,
    beforePrintMethod: ({ content }) => {
      // 拦截打印之前，返回自定义的 html 内容
      return topHtml + content + bottomHtml;
    },
  });
}
