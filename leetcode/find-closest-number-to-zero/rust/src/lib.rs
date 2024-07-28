pub fn find_closest_number(nums: Vec<i32>) -> i32 {
    let mut max_value: i32 = nums[0];
    let mut max_value_distance: i32 = i32::abs(max_value);

    let mut index = 1;
    while index < nums.len() {
        let current = nums[index];
        index += 1;

        let current_distance = i32::abs(current);
        if current_distance > max_value_distance {
            continue;
        }

        if current_distance < max_value_distance || current > max_value {
            max_value = current;
            max_value_distance = current_distance;
        }
    }

    return max_value;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = find_closest_number(vec![-4, -2, 1, 4, 8]);
        assert_eq!(result, 1);
    }

    #[test]
    fn it_works2() {
        let result = find_closest_number(vec![2, -1, 1]);
        assert_eq!(result, 1)
    }

    #[test]
    fn it_works3() {
        let result = find_closest_number(vec![-2, -3, -4]);
        assert_eq!(result, -2)
    }
}
