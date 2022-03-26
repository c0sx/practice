pub fn nth_prime(n: usize) -> usize {
    let mut counter = 0;
    let mut current = 1;

    while counter <= n {
        current += 1;

        if is_prime(current) {
            counter += 1;
        }

    }

    current
}

fn is_prime(n: usize) -> bool {
    if n <= 1 {
        return false;
    }

    let mut i = 2;
    while i <= n / 2 {
        if n % i == 0 {
            return false;
        }

        i += 1;
    }

    return true;
}

#[cfg(test)]
mod tests {
    use super::nth_prime;

    #[test]
    fn it_works() {
        assert_eq!(nth_prime(0), 2);
        assert_eq!(nth_prime(1), 3);
        assert_eq!(nth_prime(3), 7);
        assert_eq!(nth_prime(5), 13);
        assert_eq!(nth_prime(10000), 104743);
        assert_eq!(nth_prime(10001), 104759);
    }
}
