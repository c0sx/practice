pub fn diagonal_difference<const N: usize>(matrix: [[i32; N]; N]) -> i32 {
    let mut left: i32 = 0;
    let mut right: i32 = 0;

    let mut i = 0;
    while i < N {
        let left_index = i;
        let right_index = N - i - 1;
        left += matrix[i][left_index];
        right += matrix[i][right_index];
        i += 1;
    }

    (left - right).abs()
}

#[cfg(test)]
mod tests {
    use super::diagonal_difference;

    #[test]
    fn it_works() {
        let value1 = (
            [
                [11, 2, 4],
                [4, 5, 6],
                [10, 8, -12]
            ],
            15
        );

        let value2 = (
            [
                [1, 2, 3, 4],
                [5, 6, 7, 8],
                [9, 10, 11, 12],
                [13, 14, 15, 16]
            ],
            0
        );

        assert_eq!(diagonal_difference(value1.0), value1.1);
        assert_eq!(diagonal_difference(value2.0), value2.1);
    }
}
