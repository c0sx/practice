
class ListNode {
    constructor(val, next) {
        this.val = (val === undefined ? 0 : val)
        this.next = (next === undefined ? null : next)
    }

    toString() {
        let current = this;
        let arr = [];

        while (current !== null) {
            arr.push(current.val)
            current = current.next;
        }

        return `[${arr.join(",")}]`;
    }
}

const sumTwoLinkedLists = (l1, l2, additional) => {
    if (!l1 && !l2) {
        if (additional) {
            return new ListNode(additional);
        }

        return null;
    }

    const sum = (l1?.val ?? 0) + (l2?.val ?? 0) + additional;
    return new ListNode(sum % 10, sumTwoLinkedLists(l1?.next ?? null, l2?.next ?? null, sum >= 10 ? 1 : 0))
}

/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
const addTwoNumbers = function(l1, l2) {
    return sumTwoLinkedLists(l1, l2, 0);
};

//

const arrToLinkedList = ([head, ...tail]) => {
    return tail.reduce((acc, current) => {
        return new ListNode(current, acc)
    }, new ListNode(head));
}

const reverseLinkedList = list => {
    let prev = null;
    let current = list;
    let next = null;

    while (current !== null) {
        next = current.next;
        current.next = prev;
        prev = current;
        current = next;
    }

    return prev;
}

// console.log(addTwoNumbers(
//     arrToLinkedList([1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1]),
//     arrToLinkedList([5, 6, 4])
// ).toString());
// console.log(addTwoNumbers(
//     arrToLinkedList([2, 4, 3]),
//     arrToLinkedList([5, 6, 4])
// ).toString());
// console.log(addTwoNumbers(
//     arrToLinkedList([9, 9, 9, 9, 9, 9, 9]),
//     arrToLinkedList([9, 9, 9, 9])
// ).toString())
//[8,9,9,9,0,0,0,1]
console.log(addTwoNumbers(
    reverseLinkedList(arrToLinkedList([2, 4, 9])),
    reverseLinkedList(arrToLinkedList([5, 6, 4, 9]))
).toString())
