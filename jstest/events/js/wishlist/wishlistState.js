import _ from '../shared/lodash.js';
import jsonLocalStorage from '../shared/jsonLocalStorage.js';
import messenger from '../shared/messenger.js';

function WishlistState() {
  const channel = 'wishlist';
  const self = this;
  const store = jsonLocalStorage(channel, []);
  const items = store.getItem();

  const init = () => {
    self.publishState();

    messenger.subscribe({
      channel,
    }, (message) => {
      if (_.isNil(message.topic) || _.isNil(message.data)) {
        return undefined;
      }
      switch (message.topic) {
        case 'addToRemoveFrom':
          return self.addToRemoveFrom(message.data);
        default:
          return undefined;
      }
    });
  };

  this.publishState = () => {
    messenger.publishState({
      channel,
      data: items,
    });
  };

  this.addToRemoveFrom = (data) => {
    const item = _.find(items, { id: data.id }, 0);
    if (_.isNil(item)) {
      items.push(data);
    } else {
      _.remove(items, (v) => v.id === data.id);
    }

    store.setItem(items);
    this.publishState();
    return item;
  };

  init();
}

(function init() {
  new WishlistState();
}());
