import $ from '../shared/umbrella.js';
import _ from '../shared/lodash.js';
import messenger from '../shared/messenger.js';

$(document).find('span[data-subscriber="inWishlist"]').each((elm) => {
  const $elm = $(elm);
  messenger.subscribeToState({
    channel: $elm.data('channel'),
  }, (message) => {
    const item = _.find(message.data, JSON.parse($elm.data('filter')), 0)
    if (_.isObject(item)) {
      $elm.html(`Remove from wishlist`);
    } else {
      $elm.html(`Add to wishlist`);
    }
  });
});
