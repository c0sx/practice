pub fn remove_element(nums: &mut Vec<i32>, val: i32) -> i32 {
    nums.retain(|x| *x != val);

    nums.len() as i32
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let mut v = vec![3, 2, 2, 3];
        let result = remove_element(&mut v, 3);

        assert_eq!(result, 2);
    }
}
