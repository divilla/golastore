import _ from './lodash.js';

function Messenger() {
  const subscriptions = [];
  const states = {};

  this.subscribe = function subscribe(to, handler) {
    const sub = {};
    if (_.isFunction(handler)) {
      sub.handler = handler;
    } else {
      throw 'handler argument is not valid callback function';
    }

    if (!_.isObject(to)) {
      throw 'to argument is not valid object';
    }

    if (!_.isNil(to.channel) && !_.isNil(to.topic) && to.topic === 'state' && !_.isNil(states[to.channel])) {
      handler(states[to.channel]);
    }
    if (!_.isNil(to.channel)) {
      sub.channel = to.channel;
    }
    if (!_.isNil(to.topic)) {
      sub.topic = to.topic;
    }

    subscriptions.push(sub);
  };

  this.subscribeToState = function subscribeToState (to, handler) {
    if (!_.isObject(to)) {
      throw 'to argument is not valid object';
    }
    if (_.isNil(to.channel)) {
      throw 'to.channel argument is required in subscribeToState';
    }
    to.topic = 'state';

    this.subscribe(to, handler);
  }

  this.publish = function publish (message) {
    for (let i=0; i<subscriptions.length; i++) {
      const sub = subscriptions[i];
      if (sub.channel !== undefined && sub.channel !== message.channel) {
        continue;
      }
      if (sub.topic !== undefined && sub.topic !== message.topic) {
        continue;
      }

      sub.handler(message);
    }
  };

  this.publishState = (message) => {
    if (_.isNil(message.channel)) {
      throw 'message.channel argument is required in publishState';
    }
    if (_.isNil(message.data)) {
      throw 'message.data argument is required in publishState';
    }
    message.topic = 'state';
    states[message.channel] = message;

    this.publish(message);
  };
}

const messenger = new Messenger();

export default messenger;
