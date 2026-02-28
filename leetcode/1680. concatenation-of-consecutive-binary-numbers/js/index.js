const assert = require("node:assert");

/**
 * @param {number} n
 * @return {number}
 */
var concatenatedBinary = (n) => {
  const MOD = 1000000007n;
  let result = 0n;
  let bits = 0n;

  for (let i = 1n; i <= BigInt(n); i++) {
    if ((i & (i - 1n)) === 0n) {
      bits++;
    }

    result = ((result << bits) + i) % MOD;
  }

  return Number(result);
};

assert.equal(concatenatedBinary(1), 1);
assert.equal(concatenatedBinary(3), 27);
assert.equal(concatenatedBinary(12), 505379714);
