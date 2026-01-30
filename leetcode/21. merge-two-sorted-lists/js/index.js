const assert = require("node:assert");

/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */

function ListNode(val, next) {
  this.val = val === undefined ? 0 : val;
  this.next = next === undefined ? null : next;
}

/**
 * @param {ListNode} list1
 * @param {ListNode} list2
 * @return {ListNode}
 */
function mergeTwoLists(list1, list2) {
  let head1 = list1;
  let head2 = list2;

  let current = new ListNode();
  const head = current;

  while (head1 && head2) {
    if (head1.val <= head2.val) {
      current.next = head1;
      head1 = head1.next;
    } else {
      current.next = head2;
      head2 = head2.next;
    }

    current = current.next;
  }

  if (head1) {
    current.next = head1;
  } else if (head2) {
    current.next = head2;
  }

  return head.next;
}

/**
 * Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]

Input: list1 = [], list2 = []
Output: []

Example 3:

Input: list1 = [], list2 = [0]
Output: [0]
 */

assert.deepStrictEqual(
  mergeTwoLists(
    new ListNode(1, new ListNode(2, new ListNode(4))),
    new ListNode(1, new ListNode(3, new ListNode(4))),
  ),
  new ListNode(
    1,
    new ListNode(
      1,
      new ListNode(2, new ListNode(3, new ListNode(4, new ListNode(4)))),
    ),
  ),
);

assert.deepStrictEqual(mergeTwoLists(null, null), null);

assert.deepStrictEqual(mergeTwoLists(null, new ListNode(0)), new ListNode(0));
