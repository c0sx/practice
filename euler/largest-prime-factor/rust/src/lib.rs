pub fn largest_prime_factor(n: i64) -> i64 {
    let mut num = n;
    let mut largest_prime = 0;

    while num % 2 == 0 {
        largest_prime = 2;
        num /= 2;
    }

    while num % 3 == 0 {
        largest_prime = 3;
        num /= 3;
    }

    let mut i = 5;
    while i <= num {
        while num % i == 0 {
            largest_prime = i;
            num /= i;
        }

        let step = i + 2;
        while num % step == 0 {
            largest_prime = step;
            num /= step;
        }

        i += 6;
    }

    return if num > 4 {
        num
    } else {
        largest_prime
    };
}

#[cfg(test)]
mod tests {
    use super::largest_prime_factor;

    #[test]
    fn it_works() {
        let input = 13_195;
        let expected = 29;

        let result = largest_prime_factor(input);
        assert_eq!(result, expected)
    }
}
