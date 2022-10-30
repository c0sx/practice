pub fn is_palindrome(x: i32) -> bool {
    if x < 0 {
        return false;
    } else if x == 0 {
        return true;
    }

    let mut value = 0;
    let mut temp = x;
    while temp > 0 {
        let digit = temp % 10;
        value = value * 10 + digit;

        temp = temp / 10;
    }

    x == value
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = is_palindrome(121);
        assert!(result);
    }
}
