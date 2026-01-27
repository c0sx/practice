const assert = require("node:assert");

/**
 * @param {number} target
 * @param {number[]} nums
 * @return {number}
 */
function minSubArrayLen(target, nums) {
  let minSize = Number.POSITIVE_INFINITY;
  let l = 0;
  let sum = 0;

  for (let r = 0; r < nums.length; r++) {
    sum += nums[r];

    while (sum >= target) {
      const size = r - l + 1;
      minSize = Math.min(minSize, size);
      sum -= nums[l];
      l++;
    }
  }

  return minSize === Number.POSITIVE_INFINITY ? 0 : minSize;
}

assert.deepStrictEqual(minSubArrayLen(7, [2, 3, 1, 2, 4, 3]), 2);
assert.deepStrictEqual(minSubArrayLen(4, [1, 4, 4]), 1);
assert.deepStrictEqual(minSubArrayLen(11, [1, 1, 1, 1, 1, 1, 1, 1]), 0);
assert.deepStrictEqual(minSubArrayLen(11, [1, 2, 3, 4, 5]), 3);
assert.deepStrictEqual(minSubArrayLen(15, [1, 2, 3, 4, 5]), 5);
assert.deepStrictEqual(
  minSubArrayLen(213, [12, 28, 83, 4, 25, 26, 25, 2, 25, 25, 25, 12]),
  8,
);
