function ObjLocalStorage(_, key, initData) {
  this.getItem = function () {
    return JSON.parse(localStorage.getItem(key));
  };

  this.setItem = function (val) {
    localStorage.setItem(key, JSON.stringify(val));
  };

  this.removeItem = function () {
    localStorage.removeItem(key);
  };

  let data = this.getItem();
  if (_.isNil(data)) {
    data = initData;
    this.setItem(data);
  }
}

const objLocalStorage = function (key, initVal) {
  return new ObjLocalStorage(window._, key, initVal);
};

export default objLocalStorage;
