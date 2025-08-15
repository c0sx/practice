pub fn remove_occurrences2(s: String, part: String) -> String {
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

pub fn remove_occurrences(s: String, part: String) -> String {
    let ss = s.as_bytes();
    let pp = part.as_bytes();

    let mut stack: Vec<char> = vec![];

    for i in 0..ss.len() {
        stack.push(ss[i] as char);

        let mut j = pp.len();

        while let Some(&c) = stack.last() {
            if j > 0 && pp[j - 1] as char == c {
                stack.pop();
                j -= 1;
            } else {
                break;
            }
        }

        while j > 0 && j < pp.len() {
            stack.push(pp[j] as char);
            j += 1;
        }
    }

    stack.iter().collect()
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

    #[test]
    fn it_works4() {
        let result = remove_occurrences(
            String::from("gjzgbpggjzgbpgsvpwdk"),
            String::from("gjzgbpg"),
        );

        assert_eq!(result, "svpwdk");
    }

    #[test]
    fn it_works5() {
        let result = remove_occurrences(String::from("abcabcdefgh"), String::from("abc"));

        assert_eq!(result, "defgh");
    }

    #[test]
    fn it_works6() {
        let result = remove_occurrences(
            String::from("ixcupqoixcupqokevnpokevnpoknqywmlhevgc"),
            String::from("ixcupqokevnpo"),
        );

        assert_eq!(result, "knqywmlhevgc");
    }
}
