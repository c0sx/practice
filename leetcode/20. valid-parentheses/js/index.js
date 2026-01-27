const assert = require("node:assert");

/**
 * @param {string} s
 * @return {boolean}
 */
function isValid(s) {
  const pairs = {
    ")": "(",
    "}": "{",
    "]": "[",
  };

  const stack = [];

  for (const char of s) {
    if (!pairs[char]) {
      stack.push(char);
      continue;
    }

    if (pairs[char] === stack[stack.length - 1]) {
      stack.pop();
    } else {
      return false;
    }
  }

  return stack.length === 0;
}

assert.deepStrictEqual(isValid("()"), true);
assert.deepStrictEqual(isValid("()[]{}"), true);
assert.deepStrictEqual(isValid("(]"), false);
assert.deepStrictEqual(isValid("([])"), true);
assert.deepStrictEqual(isValid("([)]"), false);
