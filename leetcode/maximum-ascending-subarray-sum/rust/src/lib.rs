use std::cmp::max;

pub fn max_ascending_sum(nums: Vec<i32>) -> i32 {
    let mut max_sum = 0;
    let mut current_sum = nums[0];

    for i in 1..nums.len() {
        let prev = nums[i - 1];
        let current = nums[i];

        if current <= prev {
            max_sum = max(max_sum, current_sum);
            current_sum = current
        } else {
            current_sum += current
        }
    }

    return max(max_sum, current_sum);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works1() {
        let result = max_ascending_sum(vec![10, 20, 30, 5, 10, 50]);

        assert_eq!(result, 65);
    }

    #[test]
    fn it_works2() {
        let result = max_ascending_sum(vec![10, 20, 30, 40, 50]);

        assert_eq!(result, 150);
    }

    #[test]
    fn it_works3() {
        let result = max_ascending_sum(vec![12, 17, 15, 13, 10, 11, 12]);

        assert_eq!(result, 33);
    }

    #[test]
    fn it_works4() {
        let result = max_ascending_sum(vec![3, 6, 10, 1, 8, 9, 9, 8, 9]);

        assert_eq!(result, 19);
    }
}
