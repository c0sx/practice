// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
  pub val: i32,
  pub next: Option<Box<ListNode>>
}

impl ListNode {
  #[inline]
  fn new(val: i32) -> Self {
    ListNode {
      next: None,
      val
    }
  }
}

fn reverse_list(list: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    let mut prev = None;
    let mut next = None;
    let mut current = list;

    while let Some(mut curr) = current {
        next = curr.next;
        curr.next = prev;
        
        prev = Some(curr);
        current = next;
    }

    prev
}

pub fn add_two_numbers(l1: Option<Box<ListNode>>, l2: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    let l1_reversed = reverse_list(l1);
    let l2_reversed = reverse_list(l2);

    let l1_value = if let Some(node) = l1_reversed {
        node.val
    } else {
        0
    };

    let l2_value = if let Some(node) = l2_reversed {
        node.val 
    } else {
        0
    };

    let sum = l1_value + l2_value;
    let remainder = if sum >= 10 {
        1
    } else {
        0
    };

    None
}

#[cfg(test)]
mod tests {
    use super::*;

    fn to_list(arr: &[i32]) -> Option<Box<ListNode>> {
        let mut root = Some(Box::new(ListNode::new(arr[0])));
        let mut cursor = &mut root;

        let mut index = 1;
        while index < arr.len() {
            let node = Some(Box::new(ListNode::new(arr[index])));
            if let Some(parent) = cursor {
                parent.next = node;
                cursor = &mut parent.next;
            }

            index += 1;

        }

        root
    }

    fn from_list(mut list: Option<Box<ListNode>>) -> Vec<i32> {
        let mut v = vec![];

        while let Some( item) = list {
            v.push(item.val);

            list = item.next;
        }

        v
    }

    #[test]
    fn it_works() {
        let l1 = to_list(&[2, 4, 3]);
        let l2 = to_list(&[5, 6, 4]);

        println!("l1: {:?}", l1);
        println!("l2: {:?}", l2);

        let result = add_two_numbers(l1, l2);

        let expected = [7, 0, 8]; // Explanation: 342 + 465 = 807

        assert_eq!(from_list(result), expected);
    }
}
