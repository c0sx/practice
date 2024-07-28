pub fn is_subsequence(s: String, t: String) -> bool {
    if s.is_empty() {
        return true;
    }

    let mut s_chars = s.chars();
    let mut t_chars = t.chars();
    let mut s_char = s_chars.nth(0).unwrap();

    while let Some(t_char) = t_chars.next() {
        if t_char == s_char {
            let next = match s_chars.next() {
                Some(c) => c,
                None => return true,
            };

            s_char = next;
        }
    }

    return false;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = is_subsequence(String::from("abc"), String::from("ahbgdc"));
        assert_eq!(result, true);
    }

    #[test]
    fn it_works2() {
        let result = is_subsequence(String::from("axc"), String::from("ahbgdc"));
        assert_eq!(result, false);
    }
}
