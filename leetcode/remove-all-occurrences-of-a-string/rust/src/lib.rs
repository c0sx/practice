pub fn remove_occurrences(s: String, part: String) -> String {
    let mut chars: Vec<char> = s.chars().collect();
    let needle: Vec<char> = part.chars().collect();

    let mut i: usize = 0;
    let mut j: usize = 0;

    while i < chars.len() {
        while i + j < chars.len() && j < needle.len() && chars[i + j] == needle[j] {
            j += 1;
        }

        if j == needle.len() {
            chars.drain(i..i + j);

            i = if j > i { 0 } else { i - j }
        } else {
            i += 1;
        }

        j = 0;
    }

    chars.iter().collect()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works1() {
        let result = remove_occurrences(String::from("daabcbaabcbc"), String::from("abc"));

        assert_eq!(result, "dab");
    }

    #[test]
    fn it_works2() {
        let result = remove_occurrences(String::from("axxxxyyyyb"), String::from("xy"));

        assert_eq!(result, "ab");
    }

    #[test]
    fn it_works3() {
        let result = remove_occurrences(String::from("ababcc"), String::from("abc"));

        assert_eq!(result, "");
    }
}
