pub fn plus_one(mut digits: Vec<i32>) -> Vec<i32> {
    let mut index = digits.len() - 1;

    loop {
        let num = digits[index] + 1;
        if num < 10 {
            digits[index] = num;

            return digits;
        }

        let remainder = num % 10;
        digits[index] = remainder;

        if index == 0 {
            break;
        }

        index -= 1;
    }

    digits.insert(0, 1);

    return digits;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = plus_one(vec![1, 2, 3]);
        assert_eq!(result, vec![1, 2, 4]);
    }

    #[test]
    fn it_works2() {
        let result = plus_one(vec![9]);
        assert_eq!(result, vec![1, 0]);
    }

    #[test]
    fn it_works3() {
        let result = plus_one(vec![9, 9, 9]);
        assert_eq!(result, vec![1, 0, 0, 0]);
    }
}
