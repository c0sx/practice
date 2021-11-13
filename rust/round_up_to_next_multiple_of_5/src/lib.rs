// https://www.codewars.com/kata/55d1d6d5955ec6365400006d/rust

pub fn round_to_next_5(n: i32) -> i32 {
    if n % 5 == 0 {
        return n;
    }

    let digit = n % 5;
    return if n > 0 {
        n + 5 - digit
    } else {
        n - digit
    };
}

#[cfg(test)]
mod tests {
    use super::round_to_next_5;

    #[test]
    fn test_basic() {
        assert_eq!(round_to_next_5(1), 5);
    }

    #[test]
    fn test_basic_neg() {
        assert_eq!(round_to_next_5(-1), 0);
    }

    #[test]
    fn test_extended() {
        let tests = [
            [0, 0],
            [1, 5],
            [-1, 0],
            [-5, -5],
            [3, 5],
            [5, 5],
            [7, 10],
            [20, 20],
            [39, 40],
            [990, 990],
            [121, 125],
            [555, 555],
        ];

        for [x, out] in tests.iter() {
            assert_eq!(round_to_next_5(*x), *out);
        }
    }
}
