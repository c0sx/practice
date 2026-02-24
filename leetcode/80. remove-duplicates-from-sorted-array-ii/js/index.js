const assert = require("assert");

/**
 * @param {number[]} nums
 * @return {number}
 */
const removeDuplicates = (nums) => {
  let k = 2;

  for (let i = 2; i < nums.length; i++) {
    if (nums[i] !== nums[k - 2]) {
      nums[k] = nums[i];
      k++;
    }
  }

  return k;
};

assert.deepEqual(removeDuplicates([1, 1, 1, 2, 2, 3]), 5);
assert.deepEqual(removeDuplicates([0, 0, 1, 1, 1, 1, 2, 3, 3]), 7);
