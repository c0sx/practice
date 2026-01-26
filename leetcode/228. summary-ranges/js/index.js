const assert = require("node:assert");

/**
 * @param {number[]} nums
 * @return {string[]}
 */
function summaryRanges(nums) {
  let l = nums[0];
  let r = nums[0];
  const result = [];

  for (let i = 0; i < nums.length; i++) {
    const num = nums[i];

    if (num + 1 === nums[i + 1]) {
      r = nums[i + 1];
    } else {
      if (l === r) {
        result.push(`${l}`);
      } else {
        result.push(`${l}->${r}`);
      }

      l = nums[i + 1];
      r = nums[i + 1];
    }
  }

  return result;
}

assert.deepEqual(summaryRanges([0, 1, 2, 4, 5, 7]), ["0->2", "4->5", "7"]);
assert.deepEqual(summaryRanges([0, 2, 3, 4, 6, 8, 9]), [
  "0",
  "2->4",
  "6",
  "8->9",
]);
