import _ from './lodash.js';
import { InvalidArgumentError } from './errors.js';
import { hasOwnProperty } from './utils.js';

function Messenger() {
  const subscriptions = [];
  const states = {};

  this.subscribe = function subscribe(to, handler) {
    const sub = {};
    if (_.isFunction(handler)) {
      sub.handler = handler;
    } else {
      throw new InvalidArgumentError('"handler" is not valid callback function');
    }

    if (!_.isObject(to)) {
      throw new InvalidArgumentError('"to" is not valid object');
    }

    if (hasOwnProperty(to, 'channel')) {
      sub.channel = to.channel;
    }
    if (hasOwnProperty(to, 'topic')) {
      sub.topic = to.topic;
    }

    subscriptions.push(sub);
  };

  this.subscribeToState = function subscribeToState(to, handler) {
    if (!_.isObject(to)) {
      throw new InvalidArgumentError('"to" is not valid object');
    }
    if (!hasOwnProperty(to, 'channel')) {
      throw new InvalidArgumentError('"to.channel" is required');
    }
    if (hasOwnProperty(states, to.channel)) {
      handler(states[to.channel]);
    }

    this.subscribe({
      channel: to.channel,
      topic: 'state',
    }, handler);
  };

  this.publish = function publish(message) {
    for (let i = 0; i < subscriptions.length; i++) {
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
    if (!hasOwnProperty(message, 'channel')) {
      throw new InvalidArgumentError('"message.channel" is required');
    }
    if (!hasOwnProperty(message, 'data')) {
      throw new InvalidArgumentError('"message.data" is required');
    }

    const mes = {
      channel: message.channel,
      topic: 'state',
      data: message.data,
    };
    states[message.channel] = mes;
    this.publish(mes);
  };
}

const messenger = new Messenger();

export default messenger;
