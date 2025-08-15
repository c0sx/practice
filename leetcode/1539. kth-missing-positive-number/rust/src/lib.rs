pub fn find_kth_positive(arr: Vec<i32>, k: i32) -> i32 {
    let mut counter = 0;
    let mut index = 0;
    let mut expected_value = 1;

    while index < arr.len() {
        if counter == k {
            return expected_value - 1;
        }

        let value = arr[index];
        if value > expected_value {
            counter += 1;
        } else {
            index += 1;
        }

        expected_value += 1;
    }

    arr.len() as i32 + k
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = find_kth_positive(vec![2, 3, 4, 7, 11], 5);
        assert_eq!(result, 9);
    }

    #[test]
    fn it_works2() {
        let result = find_kth_positive(vec![1, 2, 3, 4], 2);
        assert_eq!(result, 6);
    }
}
