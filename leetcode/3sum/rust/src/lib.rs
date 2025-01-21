pub fn three_sum(nums: Vec<i32>) -> Vec<Vec<i32>> {
    let mut results = Vec::<Vec<i32>>::new();
    let mut nums = nums;
    nums.sort();

    let mut i = 0;
    while i < nums.len() - 1 && nums[i] <= 0 {
        let mut l = i + 1;
        let mut r = nums.len() - 1;

        while (i == 0 || i > 0 && nums[i] != nums[i - 1]) && l < r {
            let sum = nums[i] + nums[l] + nums[r];
            if sum < 0 {
                l += 1;
            } else if sum > 0 {
                r -= 1;
            } else {
                let result = vec![nums[i], nums[l], nums[r]];
                r -= 1;
                l += 1;
                while l < r && nums[l] == nums[l - 1] {
                    l += 1;
                }

                results.push(result)
            }
        }

        i += 1;
    }

    return results;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works1() {
        let nums = vec![-1, 0, 1, 2, -1, -4];
        let result = three_sum(nums);

        assert_eq!(result, vec![vec![-1, -1, 2], vec![-1, 0, 1]]);
    }

    #[test]
    fn it_works2() {
        let nums = vec![0, 1, 1];
        let result = three_sum(nums);

        assert_eq!(result, Vec::<Vec<i32>>::new());
    }

    #[test]
    fn it_works3() {
        let nums = vec![0, 0, 0];
        let result = three_sum(nums);

        assert_eq!(result, vec![vec![0, 0, 0]]);
    }

    #[test]
    fn it_works4() {
        let nums = vec![-2, 0, 0, 2, 2];
        let result = three_sum(nums);

        assert_eq!(result, vec![vec![-2, 0, 2]]);
    }
}
