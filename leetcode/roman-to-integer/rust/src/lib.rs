pub fn roman_to_int2(s: String) -> i32 {
    let mut chars = s.chars().rev();
    let mut prev = 0;

    let map = std::collections::HashMap::from([
        ('I', 1),
        ('V', 5),
        ('X', 10),
        ('L', 50),
        ('C', 100),
        ('D', 500),
        ('M', 1000),
    ]);

    let mut result = 0;

    while let Some(c) = chars.next() {
        let value = *map.get(&c).unwrap();

        if value < prev {
            result -= value;
        } else {
            result += value;
        }

        prev = value;
    }

    return result;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let roman = String::from("III");
        let result = roman_to_int2(roman);

        assert_eq!(result, 3);
    }

    #[test]
    fn it_works2() {
        let roman = String::from("LVIII");
        let result = roman_to_int2(roman);

        assert_eq!(result, 58);
    }

    #[test]
    fn it_works3() {
        let roman = String::from("MCM");
        let result = roman_to_int2(roman);

        assert_eq!(result, 1900)
    }

    #[test]
    fn it_works4() {
        let roman = String::from("MCMXCIV");
        let result = roman_to_int2(roman);

        assert_eq!(result, 1994);
    }
}
