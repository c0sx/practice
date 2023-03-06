pub fn remove_duplicates(nums: &mut Vec<i32>) -> i32 {
    nums.dedup_by(|a, b| a == b);

    nums.len() as i32
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let mut v = vec![1, 1, 2];
        let result = remove_duplicates(&mut v);

        assert_eq!(result, 2);
    }

    #[test]
    fn it_works2() {
        let mut v = vec![0, 0, 1, 1, 1, 2, 2, 3, 3, 4];
        let result = remove_duplicates(&mut v);

        assert_eq!(result, 5);
    }
}
