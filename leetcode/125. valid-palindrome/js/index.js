const assert = require("node:assert");

/**
 * @param {string} s
 * @return {boolean}
 */
function isPalindrome(s) {
  const str = s.toLowerCase().replace(/[^a-z\d]/g, "");
  let l = 0;
  let r = str.length - 1;

  while (l < r) {
    const leftChar = str[l];
    const rightChar = str[r];

    if (leftChar !== rightChar) {
      return false;
    }

    l++;
    r--;
  }

  return true;
}

assert.deepEqual(isPalindrome("A man, a plan, a canal: Panama"), true);
assert.deepEqual(isPalindrome("race a car"), false);
assert.deepEqual(isPalindrome(" "), true);
assert.deepEqual(isPalindrome("0P"), false);
