import $ from '../shared/umbrella.js';
import _ from '../shared/lodash.js';
import messenger from '../shared/messenger.js';

$(document).find('span[data-subscriber="inCart"]').each((elm) => {
  const $elm = $(elm);
  messenger.subscribeToState({
    channel: $elm.data('channel'),
  }, (message) => {
    const item = _.find(message.data, JSON.parse($elm.data('filter')), 0);
    if (_.isObject(item)) {
      $elm.html('Remove all');
    } else {
      $elm.html('Add to cart');
    }
  });
});

$(document).find('div[data-subscriber="addRemoveQuantity"]').each((elm) => {
  const $elm = $(elm);
  messenger.subscribeToState({
    channel: $elm.data('channel'),
  }, (message) => {
    const item = _.find(message.data, JSON.parse($elm.data('filter')), 0);
    if (_.isObject(item)) {
      $elm.attr('style', '');
    } else {
      $elm.attr('style', 'display: none');
    }
  });
});

$(document).find('span[data-subscriber="quantity"]').each((elm) => {
  const $elm = $(elm);
  messenger.subscribeToState({
    channel: $elm.data('channel'),
  }, (message) => {
    let quantity = 0;
    const item = _.find(message.data, JSON.parse($elm.data('filter')), 0);
    if (_.isObject(item)) {
      quantity = item.quantity;
    }
    $elm.html(quantity);
  });
});

$(document).find('span[data-subscriber="totalItems"]').each((elm) => {
  const $elm = $(elm);
  messenger.subscribeToState({
    channel: $elm.data('channel'),
  }, (message) => {
    let total = 0;
    if (_.isArray(message.data)) {
      total = message.data.length;
    }
    $elm.html(total);
  });
});
