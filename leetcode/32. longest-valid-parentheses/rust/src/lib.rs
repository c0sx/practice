use std::cmp::max;

fn find_related_close_bracket(tokens: &Vec<char>, start_index: usize) -> i32 {
    let mut opened: Vec<usize> = Vec::with_capacity(tokens.len());
    let mut index = start_index + 1;

    while index < tokens.len() {
        let token = tokens[index];

        if token == '(' {
            opened.push(index);
        } else if let None = opened.pop() {
            return index as i32;
        }

        index += 1;
    }

    return -1;
}

pub fn longest_valid_parentheses(s: String) -> i32 {
    if s.len() == 0 {
        return 0;
    }

    let tokens: Vec<char> = s.chars().collect();

    let mut counter = 0;
    let mut max_counter = 0;

    let mut index = 0;
    while index < tokens.len() {
        let token = tokens[index];
        if token == '(' {
            let found_index = find_related_close_bracket(&tokens, index);
            if found_index != -1 {
                let length = found_index - index as i32 + 1;
                counter += length;

                index = found_index as usize + 1;
                continue;
            }
        }

        if counter > max_counter {
            max_counter = counter;
        }

        counter = 0;

        index += 1;
    }

    max(counter, max_counter)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works1() {
        let result = longest_valid_parentheses("()".to_string());
        assert_eq!(result, 2);
    }

    #[test]
    fn it_works2() {
        let result = longest_valid_parentheses(")()())".to_string());
        assert_eq!(result, 4);
    }

    #[test]
    fn it_works3() {
        let result = longest_valid_parentheses("".to_string());
        assert_eq!(result, 0);
    }

    #[test]
    fn it_works4() {
        let result = longest_valid_parentheses("()(()".to_string());
        assert_eq!(result, 2);
    }

    #[test]
    fn it_works5() {
        let result = longest_valid_parentheses(")".to_string());
        assert_eq!(result, 0);
    }

    #[test]
    fn it_works6() {
        let result = longest_valid_parentheses("()(())".to_string());
        assert_eq!(result, 6);
    }

    #[test]
    fn it_works7() {
        let result = longest_valid_parentheses(")(((((()())()()))()(()))(".to_string());
        assert_eq!(result, 22);
    }

    #[test]
    fn it_works8() {
        let result = longest_valid_parentheses(")(()(".to_string());
        assert_eq!(result, 2);
    }
}
