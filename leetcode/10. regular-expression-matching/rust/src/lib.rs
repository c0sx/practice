enum Pattern {
    Empty,
    Single(u8),
    Repeatable(u8),
}

fn parse(p: &[u8]) -> (Pattern, &[u8]) {
    match p.split_first() {
        None => (Pattern::Empty, p),
        Some((first, rest)) => match rest.split_first() {
            Some((b'*', rest)) => (Pattern::Repeatable(*first), rest),
            _ => (Pattern::Single(*first), rest),
        },
    }
}

fn is_match_slices(s: &[u8], p: &[u8]) -> bool {
    match parse(p) {
        (Pattern::Empty, _) => s.is_empty(),
        (Pattern::Single(c), subp) => is_match_single(s, c, subp),
        (Pattern::Repeatable(c), subp) => is_match_single(s, c, p) || is_match_slices(s, subp),
    }
}

fn is_match_single(s: &[u8], to_match: u8, p: &[u8]) -> bool {
    match s.split_first() {
        Some((c, s)) if to_match == b'.' || to_match == *c => is_match_slices(s, p),
        _ => false,
    }
}

pub fn is_match(s: String, p: String) -> bool {
    is_match_slices(s.as_bytes(), p.as_bytes())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = is_match(String::from("aa"), String::from("a"));
        assert_eq!(result, false);
    }

    #[test]
    fn it_works2() {
        let result = is_match(String::from("aa"), String::from("a."));
        assert_eq!(result, true)
    }

    #[test]
    fn it_works3() {
        let result = is_match(String::from("aab"), String::from("a.b"));
        assert_eq!(result, true)
    }

    #[test]
    fn it_works4() {
        let result = is_match(String::from("aa"), String::from("a*"));
        assert_eq!(result, true);
    }

    #[test]
    fn it_works5() {
        let result = is_match(String::from("ab"), String::from(".*"));
        assert_eq!(result, true);
    }

    #[test]
    fn it_works6() {
        let result = is_match(String::from("aaabb"), String::from("a*b*"));
        assert_eq!(result, true);
    }

    #[test]
    fn it_works7() {
        let result = is_match(String::from("aab"), String::from("c*a*b"));
        assert_eq!(result, true)
    }
}
