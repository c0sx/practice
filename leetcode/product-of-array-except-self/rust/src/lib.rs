pub fn product_except_self(nums: Vec<i32>) -> Vec<i32> {
    let mut prod = vec![1; nums.len()];
    let mut left_prod = 1;
    let mut index = 0;

    while index < nums.len() {
        prod[index] *= left_prod;
        left_prod *= nums[index];

        index += 1;
    }

    let mut right_prod = 1;
    let mut index = nums.len() - 1;

    loop {
        prod[index as usize] *= right_prod;
        right_prod *= nums[index];

        if index == 0 {
            break;
        }

        index -= 1;
    }

    prod
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = product_except_self(vec![1, 2, 3, 4]);
        assert_eq!(result, vec![24, 12, 8, 6]);
    }

    #[test]
    fn it_work2() {
        let result = product_except_self(vec![-1, 1, 0, -3, 3]);
        assert_eq!(result, vec![0, 0, 9, 0, 0]);
    }
}
