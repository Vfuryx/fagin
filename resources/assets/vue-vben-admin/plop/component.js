const { notEmpty } = require('./uitls/utils.js');
module.exports = {
  description: '创建组件component',
  prompts:[{
    type: 'input',
    name: 'name',
    message: '请输入组件名称,如multi_company_picker,然后点击回车',
    validate: notEmpty('name'),
  },],
  actions: (data) => {
    data.prefix ='nb';
    data.name =`${data.prefix}_${data.name}`;
    const pathCaseName = '{{ pascalCase name }}';

    return[{
      type: 'add',
      path: `src/components/${pathCaseName}/src/${pathCaseName}.vue`,
      templateFile: 'plop/template/component/component.hbs',
    },{
      type: 'add',
      path: `src/components/${pathCaseName}/index.ts`,
      templateFile: 'plop/template/component/index.hbs',
    }];
  }
}
