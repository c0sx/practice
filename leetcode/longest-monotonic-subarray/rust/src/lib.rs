use std::cmp::max;

pub fn longest_monotonic_subarray(nums: Vec<i32>) -> i32 {
    let mut max_length = 1;

    let mut incr = 1;
    let mut decr = 1;

    for i in 1..nums.len() {
        let n = nums[i];
        let prev = nums[i - 1];

        if n == prev {
            incr = 1;
            decr = 1;
        } else if n > prev {
            incr += 1;
            decr = 1;
        } else {
            decr += 1;
            incr = 1
        }

        let max_item = max(incr, decr);
        max_length = max(max_length, max_item)
    }

    max_length
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works0() {
        let result = longest_monotonic_subarray(vec![1, 2, 3]);

        assert_eq!(result, 3)
    }

    #[test]
    fn it_works1() {
        let result = longest_monotonic_subarray(vec![1, 2, 3, 2, 4, 5, 6, 7]);

        assert_eq!(result, 5)
    }

    #[test]
    fn it_works2() {
        let result = longest_monotonic_subarray(vec![1, 4, 3, 3, 2]);

        assert_eq!(result, 2);
    }

    #[test]
    fn it_works3() {
        let result = longest_monotonic_subarray(vec![3, 3, 3, 3]);

        assert_eq!(result, 1);
    }

    #[test]
    fn it_works4() {
        let result = longest_monotonic_subarray(vec![3, 2, 1]);

        assert_eq!(result, 3);
    }

    #[test]
    fn it_works5() {
        let result = longest_monotonic_subarray(vec![1, 1, 5]);

        assert_eq!(result, 2);
    }

    #[test]
    fn it_works6() {
        let result = longest_monotonic_subarray(vec![1, 10, 10]);

        assert_eq!(result, 2);
    }
}
