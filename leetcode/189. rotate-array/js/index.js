const assert = require("node:assert");

/**
 * @param {number[]} nums
 * @param {number} k
 * @return {void} Do not return anything, modify nums in-place instead.
 */
var rotate = (nums, k) => {
  const n = nums.length;
  k = k % n;
  reverse(nums, 0, n - 1);
  reverse(nums, 0, k - 1);
  reverse(nums, k, n - 1);
};

function reverse(nums, start, end) {
  while (start < end) {
    const temp = nums[end];
    nums[end] = nums[start];
    nums[start] = temp;
    start++;
    end--;
  }
}

const nums1 = [1, 2, 3, 4, 5, 6, 7];
rotate(nums1, 3);
assert.deepStrictEqual(nums1, [5, 6, 7, 1, 2, 3, 4]);

const nums2 = [-1, -100, 3, 99];
rotate(nums2, 2);
assert.deepStrictEqual(nums2, [3, 99, -1, -100]);
