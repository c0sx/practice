const assert = require("node:assert");

/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */

function ListNode(val) {
  this.val = val;
  this.next = null;
}

/**
 * @param {ListNode} head
 * @return {boolean}
 */
function hasCycle(head) {
  const visited = new Set();
  let current = head;

  while (current) {
    if (visited.has(current)) {
      return true;
    }

    visited.add(current);
    current = current.next;
  }

  return false;
}

/**
 * Input: head = [3,2,0,-4], pos = 1
Output: true

Input: head = [1,2], pos = 0
Output: true

Input: head = [1], pos = -1
Output: false
 */

const node1 = new ListNode(3);
const node2 = new ListNode(2);
const node3 = new ListNode(0);
const node4 = new ListNode(-4);
node1.next = node2;
node2.next = node3;
node3.next = node4;
node4.next = node2; // cycle here

assert.deepStrictEqual(hasCycle(node1), true);

const nodeA = new ListNode(1);
const nodeB = new ListNode(2);
nodeA.next = nodeB;
nodeB.next = nodeA; // cycle here

assert.deepStrictEqual(hasCycle(nodeA), true);

const nodeX = new ListNode(1);

assert.deepStrictEqual(hasCycle(nodeX), false);
