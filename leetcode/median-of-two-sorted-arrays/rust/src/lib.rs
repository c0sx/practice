pub fn find_median_sorted_arrays(nums1: Vec<i32>, nums2: Vec<i32>) -> f64 {
    let mut input = Vec::new();
    input.extend(nums1);
    input.extend(nums2);
    input.sort();
    let length = input.len() as f64;
    let right_index = (length / 2.0).floor() as usize;
    let right_value = input[right_index];

    if length % 2.0 == 0.0 {
        let left_index = right_index - 1.0 as usize;
        let left_value = input[left_index];

        (left_value + right_value) as f64 / 2.0
    } else {
        right_value as f64 / 1.0
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let nums1 = [1, 3];
        let nums2 = [2];

        let result = find_median_sorted_arrays(Vec::from(nums1), Vec::from(nums2));
        assert_eq!(result, 2.0);
    }

    #[test]
    fn it_works2() {
        let nums1 = [5, 6, 7];
        let nums2 = [3, 8];

        let result = find_median_sorted_arrays(Vec::from(nums1), Vec::from(nums2));
        assert_eq!(result, 6.0);
    }

    #[test]
    fn it_works3() {
        let nums1 = [1, 2];
        let nums2 = [3, 4];

        let result = find_median_sorted_arrays(Vec::from(nums1), Vec::from(nums2));
        assert_eq!(result, 2.5);
    }
}
