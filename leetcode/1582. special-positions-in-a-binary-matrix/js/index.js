const assert = require("node:assert");

/**
 * @param {number[][]} mat
 * @return {number}
 */
var numSpecial = (mat) => {
  const rowCount = [];
  const colCount = [];

  for (let i = 0; i < mat.length; i++) {
    for (let j = 0; j < mat[i].length; j++) {
      const num = mat[i][j];
      if (num === 1) {
        const rowValue = rowCount[i] ?? 0;
        rowCount[i] = rowValue + 1;

        const colValue = colCount[j] ?? 0;
        colCount[j] = colValue + 1;
      }
    }
  }

  let result = 0;
  for (let i = 0; i < mat.length; i++) {
    for (let j = 0; j < mat[i].length; j++) {
      if (mat[i][j] === 1) {
        if (rowCount[i] === 1 && colCount[j] === 1) {
          result += 1;
        }
      }
    }
  }

  return result;
};

assert.equal(
  numSpecial([
    [1, 0, 0],
    [0, 0, 1],
    [1, 0, 0],
  ]),
  1,
);

assert.equal(
  numSpecial([
    [1, 0, 0],
    [0, 1, 0],
    [0, 0, 1],
  ]),
  3,
);
