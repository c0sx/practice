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
 * @param {TreeNode} p
 * @param {TreeNode} q
 * @return {boolean}
 */
function isSameTree(p, q) {
  if (p === null && q === null) {
    return true;
  }

  if (p && q && p.val === q.val) {
    return isSameTree(p.left, q.left) && isSameTree(p.right, q.right);
  }

  return false;
}

/**
 * Input: p = [1,2,3], q = [1,2,3]
Output: true

Input: p = [1,2], q = [1,null,2]
Output: false

Input: p = [1,2,1], q = [1,1,2]
Output: false
 */

assert.strictEqual(
  isSameTree(
    new TreeNode(1, new TreeNode(2), new TreeNode(3)),
    new TreeNode(1, new TreeNode(2), new TreeNode(3)),
  ),
  true,
);

assert.strictEqual(
  isSameTree(
    new TreeNode(1, new TreeNode(2), null),
    new TreeNode(1, null, new TreeNode(2)),
  ),
  false,
);

assert.strictEqual(
  isSameTree(
    new TreeNode(1, new TreeNode(2), new TreeNode(1)),
    new TreeNode(1, new TreeNode(1), new TreeNode(2)),
  ),
  false,
);
