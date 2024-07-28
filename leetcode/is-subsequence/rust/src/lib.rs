pub fn is_subsequence(s: String, t: String) -> bool {
    if s.is_empty() {
        return true;
    }

    let s_chars = s.chars().collect::<Vec<char>>();
    let mut index = 0;

    for t_char in t.chars() {
        if index == s_chars.len() {
            return true;
        }

        let s_char = s_chars[index];
        if t_char == s_char {
            index += 1;
        }
    }

    return index == s_chars.len();
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
