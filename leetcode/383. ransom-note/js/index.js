const assert = require("node:assert");

/**
 * @param {string} ransomNote
 * @param {string} magazine
 * @return {boolean}
 */
function canConstruct(ransomNote, magazine) {
  const letters = new Map();

  for (let i = 0; i < magazine.length; i++) {
    const letter = magazine[i];
    const count = letters.get(letter) ?? 0;
    letters.set(letter, count + 1);
  }

  for (let i = 0; i < ransomNote.length; i++) {
    const letter = ransomNote[i];
    const count = letters.get(letter) ?? 0;

    if (count <= 0) {
      return false;
    }

    letters.set(letter, count - 1);
  }

  return true;
}

assert.strictEqual(canConstruct("a", "b"), false);
assert.strictEqual(canConstruct("aa", "ab"), false);
assert.strictEqual(canConstruct("aa", "aab"), true);
