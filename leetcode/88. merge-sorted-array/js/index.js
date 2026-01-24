function merge(nums1, m, nums2, n) {
  let i = m - 1;
  let j = n - 1;
  let k = m + n - 1;

  while (j >= 0) {
    if (i >= 0 && nums1[i] > nums2[j]) {
      nums1[k] = nums1[i];
      i--;
      k--;
    } else {
      nums1[k] = nums2[j];
      j--;
      k--;
    }
  }
}

merge([1, 2, 3, 0, 0, 0], 3, [2, 5, 6], 3);
merge([4, 5, 6, 0, 0, 0], 3, [1, 2, 3], 3);
merge([4, 0, 0, 0, 0, 0], 1, [1, 2, 3, 5, 6], 5);
