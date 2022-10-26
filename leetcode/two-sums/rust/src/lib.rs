// https://leetcode.com/problems/two-sum/

extern crate core;

use std::collections::HashMap;

pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
    let mut left_index = 0;
    let mut right_index = 1;

    loop {
        let left = nums[left_index];
        let right = nums[right_index];
        let value = left + right;

        if value == target {
            return vec![left_index as i32, right_index as i32];
        }

        if right_index == nums.len() - 1 {
            left_index += 1;
            right_index = left_index + 1;
        } else {
            right_index += 1;
        }
    }
}

pub fn optimal_solution(nums: Vec<i32>, target: i32) -> Vec<i32> {
    let mut hash = HashMap::<i32, usize>::new();

    for (index, one) in nums.iter().enumerate() {
        let needle = target - one;
        if hash.contains_key(&needle) {
            let prev = *hash.get(&needle).unwrap();
            return vec![
                prev as i32,
                index as i32
            ];
        }

        hash.insert(*one, index);
    }

    panic!("solution not found")
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let a = vec![2, 1];
        let r = optimal_solution(a, 3);

        assert_eq!(r, vec![0, 1]);
    }

    #[test]
    fn it_works2() {
        let a = vec![2, 1, 3];
        let r = optimal_solution(a, 5);

        assert_eq!(r, vec![0, 2]);
    }

    #[test]
    fn it_works3() {
        let a = vec![1, 3, 5];
        let r = optimal_solution(a, 6);

        assert_eq!(r, vec![0, 2]);
    }

    #[test]
    fn it_works4() {
        let a = vec![3, 2, 4];
        let target = 6;

        let r = optimal_solution(a, target);

        assert_eq!(r, vec![1, 2]);
    }

    #[test]
    fn it_works5() {
        let a = vec![0, 4, 3, 0];
        let target = 0;

        let r = optimal_solution(a, target);

        assert_eq!(r, vec![0, 3]);
    }

    #[test]
    fn it_works6() {
        let a = vec![1, 1, 1, 1, 1, 4, 1, 1, 1, 1, 1, 7, 1, 1, 1, 1, 1];
        let target = 11;

        let r = optimal_solution(a, target);

        assert_eq!(r, vec![5, 11]);
    }
}
