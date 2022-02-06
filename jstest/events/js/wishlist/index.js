import $ from '../shared/umbrella.js';
import _ from '../shared/lodash.js';
import messenger from '../shared/messenger.js';

$(document).find('[z-publisher]').each((elm) => {
  const $elm = $(elm);
  const message = {
    channel: $elm.attr('z-channel'),
    topic: $elm.attr('z-topic'),
    data: JSON.parse($elm.attr('z-data')),
  };

  console.log(message);
  console.log($elm.attr('z-on'));
  $elm.on($elm.attr('z-on'),  () => {
    messenger.publish(message);
  });
});

$(document).find('span[z-subscriber="inCart"]').each((elm) => {
  const $elm = $(elm);
  messenger.subscribeToState({
    channel: $elm.attr('z-channel'),
  }, (message) => {
    const item = _.find(message.data, JSON.parse($elm.attr('z-filter')), 0)
    if (_.isObject(item)) {
      $elm.html(`Remove all`);
    } else {
      $elm.html(`Add to cart`);
    }
  });
});

$(document).find('div[z-subscriber="addRemoveQuantity"]').each((elm) => {
  const $elm = $(elm);
  messenger.subscribeToState({
    channel: $elm.attr('z-channel'),
  }, (message) => {
    const item = _.find(message.data, JSON.parse($elm.attr('z-filter')), 0)
    console.log(item);
    if (_.isObject(item)) {
      $elm.attr('style', '');
    } else {
      $elm.attr('style', 'display: none');
    }
  });
});

$(document).find('span[z-subscriber="quantity"]').each((elm) => {
  const $elm = $(elm);
  messenger.subscribeToState({
    channel: $elm.attr('z-channel'),
  }, (message) => {
    let quantity = 0;
    const item = _.find(message.data, JSON.parse($elm.attr('z-filter')), 0)
    if (_.isObject(item)) {
      quantity = item.quantity;
    }
    $elm.html(quantity);
  });
});
