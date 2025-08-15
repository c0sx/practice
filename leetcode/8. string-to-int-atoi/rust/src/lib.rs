const MIN: i32 = -2147483648; // -2 ^ 31
const MAX: i32 = 2147483647; // 2 ^ 31 - 1;

fn trim_start(chars: &Vec<char>) -> Option<usize> {
    let mut index = 0;

    loop {
        match chars.get(index) {
            Some(c) => {
                if !c.is_whitespace() {
                    break;
                }
            },
            None => {
                return None;
            }
        }

        index += 1;
    }

    return Some(index);
}

fn safe_mul(val: i32, digit: i32) -> Result<i32, &'static str> {
    let opt = val
        .checked_mul(10)
        .and_then(|v| v.checked_add(digit));

    match opt {
        Some(v) => Ok(v),
        None => Err("overflow")
    }
}

fn to_int(chars: Vec<char>, index: usize) -> Result<i32, &'static str> {
    let mut index = index;
    let mut value: i32 = 0;

    while index < chars.len() {
        let c = chars[index];
        if !c.is_digit(10) {
            break;
        }

        let digit = c.to_digit(10).unwrap() as i32;
        let result = safe_mul(value, digit);

        match result {
            Ok(v) => value = v,
            Err(_) => return result
        }

        index += 1;
    }

    Ok(value)
}

pub fn my_atoi(s: String) -> i32 {
    let mut sign: i32 = 1;

    let chars: Vec<char> = s.chars().collect();
    let index = trim_start(&chars);
    let mut index = match index {
        Some(i) => i,
        None => return 0
    };

    if chars[index] == '+' {
        index += 1;
    } else if chars[index] == '-' {
        index += 1;
        sign = -1;
    }

    let result = to_int(chars, index);

    match result {
        Ok(value) => {
            value * sign
        },
        Err(_) => {
            if sign == 1 {
                MAX
            } else {
                MIN
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = my_atoi(String::from("42"));
        assert_eq!(result, 42);
    }

    #[test]
    fn it_works2() {
        let result = my_atoi(String::from("  -42"));
        assert_eq!(result, -42);
    }

    #[test]
    fn it_works3() {
        let result = my_atoi(String::from("4193 with words"));
        assert_eq!(result, 4193);
    }

    #[test]
    fn it_works4() {
        let result = my_atoi(String::from("-91283472332"));
        assert_eq!(result, -2147483648);
    }

    #[test]
    fn it_works5() {
        let result = my_atoi(String::from(""));
        assert_eq!(result, 0);
    }

    #[test]
    fn it_works6() {
        let result = my_atoi(String::from(" "));
        assert_eq!(result, 0);
    }

    #[test]
    fn it_works7() {
        let result = my_atoi(String::from("  "));
        assert_eq!(result, 0);
    }

    #[test]
    fn it_works8() {
        let result = my_atoi(String::from("2147483648"));
        assert_eq!(result, 2147483647);
    }
}
