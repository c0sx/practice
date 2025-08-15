pub fn single_number(nums: Vec<i32>) -> i32 {
    let mut r = 0;

    for n in nums {
        r ^= n
    }

    return r;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works1() {
        let result = single_number(vec![2, 2, 1]);

        assert_eq!(result, 1);
    }

    #[test]
    fn it_works2() {
        let result = single_number(vec![4, 1, 2, 1, 2]);

        assert_eq!(result, 4);
    }

    #[test]
    fn it_works3() {
        let result = single_number(vec![1]);

        assert_eq!(result, 1);
    }
}
