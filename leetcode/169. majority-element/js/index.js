const assert = require("node:assert");

/**
 * @param {number[]} nums
 * @return {number}
 */
function majorityElement(nums) {
  const counts = new Map();
  const target = Math.floor(nums.length / 2);

  for (const num of nums) {
    const count = counts.get(num) ?? 0;
    const increased = count + 1;
    counts.set(num, increased);

    if (increased > target) {
      return num;
    }
  }
}

assert.deepEqual(majorityElement([3, 2, 3]), 3);
assert.deepEqual(majorityElement([2, 2, 1, 1, 1, 2, 2]), 2);
