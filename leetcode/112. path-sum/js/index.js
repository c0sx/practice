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
 * @param {number} targetSum
 * @return {boolean}
 */
function hasPathSum(root, targetSum) {
  if (root === null) {
    return false;
  }

  if (root.left === null && root.right === null) {
    return targetSum === root.val;
  }

  if (hasPathSum(root.left, targetSum - root.val)) {
    return true;
  }

  if (hasPathSum(root.right, targetSum - root.val)) {
    return true;
  }

  return false;
}

assert.strictEqual(
  hasPathSum(
    new TreeNode(
      5,
      new TreeNode(4, new TreeNode(11, new TreeNode(7), new TreeNode(2)), null),
      new TreeNode(8, new TreeNode(13), new TreeNode(4, null, new TreeNode(1))),
    ),
    22,
  ),
  true,
);

assert.strictEqual(
  hasPathSum(new TreeNode(1, new TreeNode(2), new TreeNode(3)), 5),
  false,
);

assert.strictEqual(hasPathSum(null, 0), false);
