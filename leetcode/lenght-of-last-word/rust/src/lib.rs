pub fn length_of_last_word(s: String) -> i32 {
    let mut counter = 0;

    for c in s.chars().rev() {
        if c == ' ' && counter == 0 {
            continue;
        } else if c == ' ' {
            return counter;
        }

        counter += 1;
    }

    counter
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = length_of_last_word(String::from("Hello World"));

        assert_eq!(result, 5);
    }

    #[test]
    fn it_works2() {
        let result = length_of_last_word(String::from("   fly me   to   the moon  "));

        assert_eq!(result, 4);
    }

    #[test]
    fn it_works3() {
        let result = length_of_last_word(String::from("luffy is still joyboy"));

        assert_eq!(result, 6);
    }
}
