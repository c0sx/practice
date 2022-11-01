use std::collections::HashSet;

pub fn longest_common_prefix(mut input: Vec<String>) -> String {
    if input.len() == 0 {
        return String::new();
    }

    input.sort_by(|a, b| a.len().cmp(&b.len()));

    let shortest = input.get(0)
        .ok_or("")
        .unwrap()
        .chars()
        .collect::<Vec<char>>();

    let mut longest_prefix = String::new();
    for i in 0..shortest.len() {
        let mut set = HashSet::new();
        for word in &input {
            let word = word.chars().collect::<Vec<char>>();
            set.insert(word[i]);
        }

        if set.len() == 1 {
            longest_prefix.push(shortest[i])
        } else {
            break;
        }
    }

    longest_prefix
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let input = vec![
            String::from("flower"),
            String::from("flow"),
            String::from("flight")
        ];

        let result = longest_common_prefix(input);

        assert_eq!(result, String::from("fl"));
    }

    #[test]
    fn it_works2() {
        let input = vec![
            String::from("dog"),
            String::from("racecar"),
            String::from("car")
        ];

        let result = longest_common_prefix(input);

        assert_eq!(result, String::from(""))
    }

    #[test]
    fn it_works3() {
        let input = vec![
            String::from("cir"),
            String::from("car"),
            String::from("car")
        ];

        let result = longest_common_prefix(input);

        assert_eq!(result, String::from("c"))
    }

    #[test]
    fn it_works4() {
        let input = vec![
            String::from("ab"),
            String::from("a"),
        ];

        let result = longest_common_prefix(input);

        assert_eq!(result, String::from("a"))
    }
}
