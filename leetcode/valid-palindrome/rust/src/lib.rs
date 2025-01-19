pub fn is_palindrome(s: String) -> bool {
    let s = s
        .chars()
        .filter(|c| c.is_alphanumeric())
        .collect::<String>()
        .to_ascii_lowercase();

    let size = s.len();

    for (i, l) in s.char_indices() {
        let r = s.chars().nth(size - i - 1).unwrap();

        if l != r {
            return false;
        }
    }

    return true;

    // let rev = s.chars().rev().collect::<String>();
    // return s == rev
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
