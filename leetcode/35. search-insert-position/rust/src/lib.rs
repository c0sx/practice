pub fn search_insert_position(list: Vec<i32>, target: i32) -> i32 {
    let left_index: usize = 0;
    let right_index: usize = list.len();

    let right_value = list[right_index - 1];
    if target > right_value {
        return right_index as i32;
    }

    let mut left = left_index as i32;
    let mut right = right_index as i32;

    while left <= right {
        let middle = (left + right) / 2;
        let middle_index = middle as usize;
        let middle_value = list[middle_index];

        if middle_value == target {
            return middle;
        }

        if target < middle_value {
            right = middle - 1
        } else {
            left = middle + 1
        }
    }

    left as i32
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let list = vec![1, 3, 5, 6, 10];
        let result = search_insert_position(list, 5);

        assert_eq!(result, 2);
    }

    #[test]
    fn it_works2() {
        let list = vec![1, 3, 5, 6];
        let result = search_insert_position(list, 2);

        assert_eq!(result, 1);
    }

    #[test]
    fn it_works3() {
        let list = vec![1, 3, 5, 6];
        let result = search_insert_position(list, 7);

        assert_eq!(result, 4);
    }

    #[test]
    fn it_works4() {
        let list = vec![1, 3, 5, 6];
        let result = search_insert_position(list, 5);

        assert_eq!(result, 2);
    }

    #[test]
    fn it_works5() {
        let list = vec![1, 3, 5];
        let result = search_insert_position(list, 4);

        assert_eq!(result, 2);
    }

    #[test]
    fn it_works6() {
        let list = vec![1, 4, 5];
        let result = search_insert_position(list, 4);

        assert_eq!(result, 1);
    }

    #[test]
    fn it_works7() {
        let list = vec![1, 2, 3, 5];
        let result = search_insert_position(list, 4);

        assert_eq!(result, 3);
    }

    #[test]
    fn it_works8() {
        let list = vec![2, 7, 8, 9, 10];
        let result = search_insert_position(list, 9);

        assert_eq!(result, 3);
    }

    #[test]
    fn it_works9() {
        let list = vec![1, 3, 5, 6];
        let result = search_insert_position(list, 0);

        assert_eq!(result, 0)
    }
}
