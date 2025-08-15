pub fn is_palindrome(s: String) -> bool {
    let s = s
        .chars()
        .filter(|c| c.is_alphanumeric())
        .collect::<String>()
        .to_ascii_lowercase();

    let size = s.len();
    let mut l = 0;

    while l < size / 2 {
        let c1 = s.chars().nth(l).unwrap();
        let c2 = s.chars().nth(size - l - 1).unwrap();

        if c1 != c2 {
            return false;
        }

        l += 1;
    }

    return true;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works1() {
        let result = is_palindrome(String::from("A man, a plan, a canal: Panama"));
        assert_eq!(result, true);
    }

    #[test]
    fn it_works2() {
        let result = is_palindrome(String::from("race a car"));
        assert_eq!(result, false);
    }

    #[test]
    fn it_works3() {
        let result = is_palindrome(String::from(" "));
        assert_eq!(result, true);
    }
}
