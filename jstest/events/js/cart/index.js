import _ from './js/lodash.js'
import messenger from "./js/messenger.js";

$(document).find('[data-publisher]').each(function () {
  const pub = $(this);
  pub.on(pub.data("publisher"), function () {
    messenger.publish(pub.data('channel'), pub.data('topic'), pub.data());
  });
});

$(document).find('button[data-subscriber]').each(function () {
  const sub = $(this);
  messenger.subscribeToState(sub.data('channel'), function (topic, data) {
    if (topic !== 'state') {
      return undefined;
    }

    let quantity = 0
    const item = _.find(data, {'id': sub.data('id')}, 0);
    if (!_.isNil(item)) {
      quantity = item.quantity
    }
    sub.html(`<span>${quantity}</span>`);

    if (quantity < 1) {
      // sub.closest('.field.has-addons').hide();
    }
    return sub;
  }) ;
});
