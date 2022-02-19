const { notEmpty, notSlash } = require('./uitls/utils.js');

module.exports = {
  description: '创建view',
  prompts: [
    {
      type: 'list',
      name: 'model',
      message: '生成模型',
      choices: ['drawer', 'modal'],
    },
    {
      type: 'list',
      name: 'path',
      message: '路径',
      choices: ['base', 'system', 'sale', '其他'],
    },
    {
      type: 'input',
      name: 'other',
      message: '请输入路径(可用/分隔),勿与之前重复,然后点击回车',
      validate: notEmpty('other'),
      when: function(answers) {
        return answers.path === '其他';
      }
    },
    {
      type: 'input',
      name: 'name',
      message: '请输入view名称(不可用/分隔),勿与之前重复,然后点击回车',
      validate: notSlash('name'),
    },
  ],
  actions: (data) => {
    console.log(data);
    // 替换路径
    if (data?.other) {
      data.path = data.other
    }
    const pathCaseName = '{{ camelCase path }}/{{ camelCase name }}';

    return [
      {
        type: 'add',
        path: `src/views/${pathCaseName}/index.vue`,
        templateFile: `plop/template/view-${data.model}/index.hbs`,
      },
      {
        type: 'add',
        path: `src/views/${pathCaseName}/{{pascalCase name}}{{ pascalCase model}}.vue`,
        templateFile: `plop/template/view-${data.model}/${data.model}.hbs`,
      },
      {
        type: 'add',
        path: `src/views/${pathCaseName}/detail{{ pascalCase model}}.vue`,
        templateFile: `plop/template/view-${data.model}/detail-${data.model}.hbs`,
      },
      {
        type: 'add',
        path: `src/views/${pathCaseName}/{{camelCase name}}.data.ts`,
        templateFile: `plop/template/view-${data.model}/data.hbs`,
      },
    ];
  },
};
