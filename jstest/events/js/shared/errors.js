function InvalidArgumentError(message) {
  this.name = 'InvalidArgumentError';
  this.message = message;
  this.stack = (new Error()).stack;
}
InvalidArgumentError.prototype = new Error();

export { InvalidArgumentError };
