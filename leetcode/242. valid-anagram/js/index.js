const assert = require("node:assert");

/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
function isAnagram(s, t) {
  const map = new Map();

  for (const char of s) {
    const count = map.get(char) ?? 0;
    map.set(char, count + 1);
  }

  for (const char of t) {
    const count = map.get(char) ?? 0;

    if (count === 0) {
      return false;
    }

    map.set(char, count - 1);
  }

  for (const count of map.values()) {
    if (count !== 0) {
      return false;
    }
  }

  return true;
}

assert.deepEqual(isAnagram("anagram", "nagaram"), true);
assert.deepEqual(isAnagram("rat", "car"), false);
