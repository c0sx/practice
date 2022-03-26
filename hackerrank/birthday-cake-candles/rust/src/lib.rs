pub fn birthday_cake_candles<const N: usize>(input: [i32; N]) -> i32 {
    let mut value = 0;
    let mut counter = 0;

    for digit in input {
        if digit > value {
            value = digit;
            counter = 1;
        } else if digit == value {
            counter += 1;
        }
    }

    counter
}

#[cfg(test)]
mod tests {
    use super::birthday_cake_candles;

    #[test]
    fn it_works() {
        let input = [3, 2, 1, 3];
        let expected = 2;

        let result = birthday_cake_candles(input);
        assert_eq!(result, expected);
    }
}
