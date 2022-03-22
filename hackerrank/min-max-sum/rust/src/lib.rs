pub fn min_max_sum<const N: usize>(arr: [i64; N]) -> (i64, i64) {
    let mut ordered = Vec::from(arr);
    ordered.sort();

    let mut min: i64 = 0;
    let mut max: i64 = 0;

    let mut i: usize = 0;
    while i < N {
        if i < N - 1 {
            min += ordered[i];
        }

        if i > 0 {
            max += ordered[i];
        }

        i += 1;
    }

    (min, max)
}

#[cfg(test)]
mod tests {
    use super::min_max_sum;

    #[test]
    fn it_works() {
        let input = [1, 2, 3, 4, 5];
        let expected = (10, 14);
        let result = min_max_sum(input);
        assert_eq!(result, expected);
    }

    #[test]
    fn it_works2() {
        let input = [7, 69, 2, 221, 8974];
        let expected = (299, 9271);
        let result = min_max_sum(input);
        assert_eq!(result, expected);
    }
}
