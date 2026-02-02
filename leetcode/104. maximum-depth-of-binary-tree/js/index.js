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
 * @return {number}
 */
function maxDepth(root) {
  if (!root) {
    return 0;
  }

  const right = maxDepth(root.right) + 1;
  const left = maxDepth(root.left) + 1;

  return right > left ? right : left;
}

assert.deepStrictEqual(
  maxDepth(
    new TreeNode(
      3,
      new TreeNode(9),
      new TreeNode(20, new TreeNode(15), new TreeNode(7)),
    ),
  ),
  3,
);

assert.deepStrictEqual(maxDepth(new TreeNode(1, null, new TreeNode(2))), 2);
