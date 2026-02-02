const assert = require("node:assert");

/**
 * @param {string[]} strs
 * @return {string}
 */
function longestCommonPrefix(strs) {
  let prefix = "";
  let size = -1;

  for (let i = 0; i < strs.length; i++) {
    if (size === -1 || strs[i].length < size) {
      size = strs[i].length;
    }
  }

  for (let i = 0; i < size; i++) {
    const needle = strs[0][i];
    let isEqual = true;

    for (let j = 1; j < strs.length; j++) {
      const target = strs[j][i];

      if (needle !== target) {
        isEqual = false;
        break;
      }
    }

    if (isEqual) {
      prefix += needle;
    } else {
      break;
    }
  }

  return prefix;
}

assert.strictEqual(longestCommonPrefix(["flower", "flow", "flight"]), "fl");
assert.strictEqual(longestCommonPrefix(["dog", "racecar", "car"]), "");
