pub fn merge_alternately(word1: String, word2: String) -> String {
    let mut result = String::new();
    let mut chars1 = word1.chars();
    let mut chars2 = word2.chars();
    let mut finished1 = false;
    let mut finished2 = false;

    while (finished1 && finished2) == false {
        let a = chars1.nth(0);
        let b = chars2.nth(0);

        match a {
            Some(c) => result.push(c),
            None => finished1 = true,
        }

        match b {
            Some(c) => result.push(c),
            None => finished2 = true,
        }
    }

    return result;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = merge_alternately(String::from("abc"), String::from("pqr"));
        assert_eq!(result, String::from("apbqcr"));
    }

    #[test]
    fn it_works2() {
        let result = merge_alternately(String::from("ab"), String::from("pqrs"));
        assert_eq!(result, String::from("apbqrs"));
    }

    #[test]
    fn it_works3() {
        let result = merge_alternately(String::from("abcd"), String::from("pq"));
        assert_eq!(result, String::from("apbqcd"));
    }
}
