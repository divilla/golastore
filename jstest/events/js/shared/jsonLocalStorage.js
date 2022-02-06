import _ from './lodash.js';

function JSONLocalStorage(key, initData) {
  this.getItem = function getItem() {
    return JSON.parse(localStorage.getItem(key));
  };

  this.setItem = function setItem(val) {
    localStorage.setItem(key, JSON.stringify(val));
  };

  this.removeItem = function removeItem() {
    localStorage.removeItem(key);
  };

  let data = this.getItem();
  if (_.isNil(data)) {
    data = initData;
    this.setItem(data);
  }
}

const jsonLocalStorage = function jsonLocalStorage(key, initVal) {
  return new JSONLocalStorage(key, initVal);
};

export default jsonLocalStorage;
