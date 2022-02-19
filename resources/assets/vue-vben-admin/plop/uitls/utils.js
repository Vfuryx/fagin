exports.notEmpty = (name) => {
  return (v) => {
    if (!v || v.trim === '') return `${name}为必填项`;
    else return true;
  };
};

exports.notSlash = (name) => {
  return (v) => {
    if (v.indexOf("/") !== -1 ) return `${name} 存在 /`;
    else return true;
  };
};
