use std::collections::{HashMap, HashSet};

pub fn is_valid(s: String) -> bool {
    let open_brackets: HashSet<char> = HashSet::from(['(', '{', '[']);
    let pairs: HashMap<char, char> = HashMap::from([
        (')', '('),
        ('}', '{'),
        (']', '[')
    ]);

    let mut stack: Vec<char> = vec![];

    for symbol in s.chars().into_iter() {
        if open_brackets.contains(&symbol) {
            stack.push(symbol);
            continue
        }

        if stack.len() == 0 {
            return false;
        }

        let last = stack.pop().unwrap();
        let pair = pairs.get(&symbol).unwrap();
        if last != *pair {
            return false;
        }
    }

    stack.len() == 0
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = is_valid("()".to_string());
        assert_eq!(result, true);
    }

    #[test]
    fn it_works2() {
        let result = is_valid("(]".to_string());
        assert_eq!(result, false);
    }

    #[test]
    fn it_works3() {
        let result = is_valid("()[]{}".to_string());
        assert_eq!(result, true);
    }
}
