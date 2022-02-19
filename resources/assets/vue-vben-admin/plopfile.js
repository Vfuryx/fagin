const apiGenerator = require('./plop/api');
const viewGenerator = require('./plop/view');
const componentGenerator = require('./plop/component');

module.exports = (plop) => {
  plop.setGenerator('api', apiGenerator);
  plop.setGenerator('view', viewGenerator);
  plop.setGenerator('component', componentGenerator);
};

// // test/test_admin
// 内置替换函数
// {{name}} // test/test_admin
// {{dotCase name}} // test.test.admin
// {{pathCase name}} // test/test/admin
// {{upperCase name}} // TEST/TEST_ADMIN
// {{lowerCase name}} // test/test_admin
// {{camelCase name}} // testTestAdmin
// {{snakeCase name}} // test_test_admin
// {{titleCase name}} // Test Test Admin
// {{kebabCase name}} // test-test-admin
// {{constantCase name}} // TEST_TEST_ADMIN
// {{sentenceCase name}} // Test test admin
// {{dashCase name}} // test-test-admin
// {{kebabCase name}} // test-test-admin
// {{properCase name}} // TestTestAdmin
// {{pascalCase name}} // TestTestAdmin
