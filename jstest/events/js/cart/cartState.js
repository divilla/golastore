import objLocalStorage from './objLocalStorage.js';
import messenger from './messenger.js';

function CartState(messenger, _) {
  const cartChannel = 'cart';
  const self = this;
  const store = objLocalStorage(cartChannel, []);
  const items = store.getItem();

  const init = function () {
    self.publishState();

    messenger.subscribe(cartChannel, (topic, data) => {
      switch (topic) {
        case 'addOne':
          return self.addOne(data);
        case 'removeOne':
          return self.removeOne(data);
        case 'removeAll':
          return self.removeAll(data);
      }
    });
  };

  this.publishState = function () {
    messenger.publishState(cartChannel, items);
  };

  this.addOne = function (data) {
    const item = _.find(items, { id: data.id }, 0);
    if (_.isNil(item)) {
      items.push({
        id: data.id,
        name: data.name,
        price: data.price,
        quantity: 1,
        total: data.price,
      });
    } else {
      item.name = data.name;
      item.price = data.price;
      item.quantity += 1;
      item.total = _.round(item.price * item.quantity, 2);
    }
    store.setItem(items);
    this.publishState();
  };

  this.removeOne = function (data) {
    const item = _.find(items, { id: data.id }, 0);
    if (_.isNil(item)) {
      return;
    }

    if (item.quantity < 2) {
      this.removeAll(data);
      return;
    }

    item.quantity--;
    store.setItem(items);
    this.publishState();
  };

  this.removeAll = function (data) {
    _.remove(items, (v) => v.id === data.id);
    store.setItem(items);
    this.publishState();
  };

  init();
}

(function (messenger, _) {
  new CartState(messenger, _);
}(messenger, window._));
