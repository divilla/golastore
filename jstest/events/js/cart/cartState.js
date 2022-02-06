import _ from '../shared/lodash.js';
import jsonLocalStorage from '../shared/jsonLocalStorage.js';
import messenger from '../shared/messenger.js';

function CartState() {
  const channel = 'cart';
  const self = this;
  const store = jsonLocalStorage(channel, []);
  const items = store.getItem();

  const init = () => {
    self.publishState();

    messenger.subscribe({
      channel,
    }, (message) => {
      if (_.isNil(message.topic) || _.isNil(message.data)) {
        return;
      }
      switch (message.topic) {
        case 'addFirstRemoveAll':
          self.addFirstRemoveAll(message.data);
          break;
        case 'addOne':
          self.addOne(message.data);
          break;
        case 'removeOne':
          self.removeOne(message.data);
          break;
        case 'removeAll':
          self.removeAll(message.data);
          break;
        default:
      }
    });
  };

  this.publishState = () => {
    messenger.publishState({
      channel,
      data: items,
    });
  };

  this.addOne = (data) => {
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

  this.removeOne = (data) => {
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

  this.removeAll = (data) => {
    _.remove(items, (v) => v.id === data.id);
    store.setItem(items);
    this.publishState();
  };

  this.addFirstRemoveAll = (data) => {
    const item = _.find(items, data, 0);
    if (_.isObject(item)) {
      this.removeAll(data);
    } else {
      this.addOne(data);
    }
  };

  init();
}

(function init() {
  new CartState();
}());
