const assert = require("node:assert");

/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */

function TreeNode(val, left, right) {
  this.val = val === undefined ? 0 : val;
  this.left = left === undefined ? null : left;
  this.right = right === undefined ? null : right;
}

/**
 * @param {TreeNode} root
 * @return {number[]}
 */
function averageOfLevels(root) {
  const map = new Map();

  const solve = (root, depth) => {
    if (!root) {
      return;
    }

    const item = map.get(depth) ?? { val: 0, count: 0 };
    map.set(depth, {
      val: item.val + root.val,
      count: item.count + 1,
    });

    solve(root.right, depth + 1);
    solve(root.left, depth + 1);
  };

  solve(root, 0);

  const result = [];
  for (const item of map.values()) {
    result.push(item.val / item.count);
  }

  return result;
}

// 3, 9 20, null, null, 15, 7
assert.deepStrictEqual(
  averageOfLevels(
    new TreeNode(
      3,
      new TreeNode(9),
      new TreeNode(20, new TreeNode(15), new TreeNode(7)),
    ),
  ),
  [3.0, 14.5, 11.0],
);

assert.deepStrictEqual(
  averageOfLevels(
    new TreeNode(
      3,
      new TreeNode(9, new TreeNode(15), new TreeNode(7)),
      new TreeNode(20),
    ),
  ),
  [3.0, 14.5, 11.0],
);

// 3,1,5,0,2,4,6
assert.deepStrictEqual(
  averageOfLevels(
    new TreeNode(
      3,
      new TreeNode(1, new TreeNode(0), new TreeNode(2)),
      new TreeNode(5, new TreeNode(4), new TreeNode(6)),
    ),
  ),
  [3.0, 3.0, 3.0],
);
