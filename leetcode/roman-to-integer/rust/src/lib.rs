pub fn roman_to_int(s: String) -> i32 {
    let chars = s.chars().collect::<Vec<char>>();

    let values = chars.into_iter()
        .map(|c| {
            match c {
                'I' => 1,
                'V' => 5,
                'X' => 10,
                'L' => 50,
                'C' => 100,
                'D' => 500,
                'M' => 1000,
                _ => panic!("invalid roman digit")
            }
        })
        .collect::<Vec<i32>>();

    let mut num = 0;
    let mut accumulator = 0;

    let mut i = 0;
    while i < values.len() {
        let current = values[i];
        let next = *values.get(i + 1).or(Some(&0)).unwrap();

        if current == next {
            accumulator += current;

            i += 1;
        } else if current < next {
            num += next - (current - accumulator);
            accumulator = 0;

            i += 2;
        } else {
            num += current + accumulator;
            accumulator = 0;

            i += 1;
        }
    }

    num
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let roman = String::from("III");
        let result = roman_to_int(roman);

        assert_eq!(result, 3);
    }

    #[test]
    fn it_works2() {
        let roman = String::from("LVIII");
        let result = roman_to_int(roman);

        assert_eq!(result, 58);
    }

    #[test]
    fn it_works3() {
        let roman = String::from("MCM");
        let result = roman_to_int(roman);

        assert_eq!(result, 1900)
    }
}
