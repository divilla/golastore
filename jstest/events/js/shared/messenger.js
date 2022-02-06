function Messenger(_) {
  const subscribers = {};
  const states = {};

  this.subscribe = function subscribe(channel, handler) {
    if (subscribers[channel] !== undefined) {
      subscribers[channel].push(handler);
    } else {
      subscribers[channel] = [handler];
    }
  };

  this.subscribeToState = function subscribeToState(channel, handler) {
    this.subscribe(channel, handler);
    if (states[channel] === undefined) {
      return undefined;
    }
    return handler('state', states[channel]);
  };

  this.publish = function publish(channel, topic, data) {
    if (_.isArray(subscribers[channel])) {
      subscribers[channel].forEach((handler) => handler(topic, data));
    }
  };

  this.publishState = function publishState(channel, data) {
    states[channel] = data;
    this.publish(channel, 'state', data);
  };
}

const messenger = new Messenger(window._);

export default messenger;
