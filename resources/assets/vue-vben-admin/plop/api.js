const { notEmpty, notSlash } = require('./uitls/utils.js');

module.exports = {
  description: '创建api',
  prompts: [
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
      message: '请输入api名称(可用/分隔),勿与之前重复,然后点击回车',
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
        path: `src/api/${pathCaseName}.ts`,
        templateFile: 'plop/template/api/api.hbs',
      },
      {
        type: 'add',
        path: `src/api/{{ camelCase path }}/model/{{ camelCase name }}Model.ts`,
        templateFile: 'plop/template/api/model.hbs',
      },
    ];
  },
};
