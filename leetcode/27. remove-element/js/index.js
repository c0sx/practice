const assert = require("node:assert");

/**
 * @param {number[]} nums
 * @param {number} val
 * @return {number}
 */
function removeElement(nums, val) {
  let i = 0;

  for (let j = 0; j < nums.length; j++) {
    if (nums[j] !== val) {
      nums[i] = nums[j];
      i++;
    }
  }

  return i;
}

assert.deepEqual(removeElement([3, 2, 2, 3], 3), 2);
assert.deepEqual(removeElement([0, 1, 2, 2, 3, 0, 4, 2], 2), 5); // 5
