import $ from '../shared/umbrella.js';
import messenger from '../shared/messenger.js';

$(document).find('[data-publisher]').each((elm) => {
  const $elm = $(elm);
  const message = {
    channel: $elm.data('channel'),
    topic: $elm.data('topic'),
    data: JSON.parse($elm.data('data')),
  };

  $elm.on($elm.data('on'),  () => {
    messenger.publish(message);
  });
});
