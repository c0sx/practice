const assert = require("node:assert");

/**
 * @param {number[][]} grid
 * @return {number}
 */
var minSwaps = (grid) => {
  const n = grid.length;
  const pos = new Array(n).fill(-1);

  for (let i = 0; i < n; i++) {
    for (let j = n - 1; j >= 0; j--) {
      if (grid[i][j] === 1) {
        pos[i] = j;
        break;
      }
    }
  }

  let result = 0;
  for (let i = 0; i < n; i++) {
    let k = -1;
    for (let j = i; j < n; j++) {
      if (pos[j] <= i) {
        result += j - i;
        k = j;
        break;
      }
    }

    if (k !== -1) {
      for (let j = k; j > i; j--) {
        [pos[j], pos[j - 1]] = [pos[j - 1], pos[j]];
      }
    } else {
      return -1;
    }
  }

  return result;
};

assert.equal(
  minSwaps([
    [0, 0, 1],
    [1, 1, 0],
    [1, 0, 0],
  ]),
  3,
);

assert.equal(
  minSwaps([
    [0, 1, 1, 0],
    [0, 1, 1, 0],
    [0, 1, 1, 0],
    [0, 1, 1, 0],
  ]),
  -1,
);

assert.equal(
  minSwaps([
    [1, 0, 0],
    [1, 1, 0],
    [1, 1, 1],
  ]),
  0,
);
