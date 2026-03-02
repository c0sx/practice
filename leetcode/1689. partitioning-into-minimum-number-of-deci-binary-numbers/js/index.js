const assert = require("node:assert");

/**
 * @param {string} n
 * @return {number}
 */
var minPartitions = (n) => {
  let maxDigit = 0;

  for (let i = 0; i < n.length; i++) {
    const digit = n[i] - "0";

    if (digit > maxDigit) {
      maxDigit = digit;
    }

    if (maxDigit === 9) {
      break;
    }
  }

  return maxDigit;
};

assert.equal(minPartitions("32"), 3);
assert.equal(minPartitions("82734"), 8);
assert.equal(minPartitions("27346209830709182346"), 9);
