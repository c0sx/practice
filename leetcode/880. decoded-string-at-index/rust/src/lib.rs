pub fn decode_at_index(s: String, k: i32) -> String {
    let mut size: i64 = 0;
    let mut k = k as i64;

    for c in s.chars() {
        if c.is_digit(10) {
            size *= c.to_digit(10).unwrap() as i64;
        } else {
            size += 1;
        }
    }

    for c in s.chars().rev() {
        k %= size;

        if k == 0 && c.is_alphabetic() {
            return String::from(c);
        }

        if c.is_digit(10) {
            size /= c.to_digit(10).unwrap() as i64;
        } else {
            size -= 1;
        }
    }

    return String::new();
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works1() {
        let result = decode_at_index(String::from("leet2code3"), 10);

        assert_eq!(result, "o");
    }

    // #[test]
    // fn it_works2() {
    //     let result = decode_at_index(String::from("ha22"), 5);

    //     assert_eq!(result, "h");
    // }

    // #[test]
    // fn it_works3() {
    //     let result = decode_at_index(String::from("a2345678999999999999999"), 1);

    //     assert_eq!(result, "a");
    // }
}
