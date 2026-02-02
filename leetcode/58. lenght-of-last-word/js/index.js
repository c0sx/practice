const assert = require("node:assert");

/**
 * @param {string} s
 * @return {number}
 */
function lengthOfLastWord(s) {
  let counter = 0;
  for (let i = s.length - 1; i >= 0; i--) {
    const c = s[i];

    if (c === " " && counter === 0) {
      continue;
    } else if (c === " ") {
      return counter;
    }

    counter += 1;
  }

  return counter;
}

assert.strictEqual(lengthOfLastWord("Hello World"), 5);
assert.strictEqual(lengthOfLastWord("   fly me   to   the moon  "), 4);
assert.strictEqual(lengthOfLastWord("luffy is still joyboy"), 6);
