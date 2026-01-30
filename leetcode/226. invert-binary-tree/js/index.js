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
 * @return {TreeNode}
 */
function invertTree(root) {
  if (root === null) {
    return null;
  }

  const left = invertTree(root.left);
  const right = invertTree(root.right);

  root.left = right;
  root.right = left;
  return root;
}

/**
 * Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]

Input: root = [2,1,3]
Output: [2,3,1]

Input: root = []
Output: []
 */

assert.deepStrictEqual(
  invertTree(
    new TreeNode(
      4,
      new TreeNode(2, new TreeNode(1), new TreeNode(3)),
      new TreeNode(7, new TreeNode(6), new TreeNode(9)),
    ),
  ),
  new TreeNode(
    4,
    new TreeNode(7, new TreeNode(9), new TreeNode(6)),
    new TreeNode(2, new TreeNode(3), new TreeNode(1)),
  ),
);

assert.deepStrictEqual(
  invertTree(new TreeNode(2, new TreeNode(1), new TreeNode(3))),
  new TreeNode(2, new TreeNode(3), new TreeNode(1)),
);

assert.deepStrictEqual(invertTree(null), null);
