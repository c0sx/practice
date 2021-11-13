// https://www.codewars.com/kata/576bb71bbbcf0951d5000044/rust

pub fn count_positives_sum_negatives(input: Vec<i32>) -> Vec<i32> {
    if input.is_empty() {
        return vec![];
    }

    let mut positives = 0;
    let mut negatives = 0;

    for num in input {
        if num > 0 {
            positives += 1;
        }

        if num < 0 {
            negatives += num
        }
    }

    vec![positives, negatives]
}

#[cfg(test)]
mod tests {
    use super::count_positives_sum_negatives;

    #[test]
    fn returns_expected() {
        let test_data1 = vec![1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15];
        let expected1 = vec![10, -65];
        assert_eq!(count_positives_sum_negatives(test_data1), expected1);
        assert_eq!(count_positives_sum_negatives(vec![]), vec![]);
    }
}
