pub fn add(a: &str, b: &str) -> String {
    let (longest, shortest) = if a.len() > b.len() {
        (a, b)
    } else {
        (b, a)
    };

    let mut reminder = 0;
    let mut result = Vec::<u32>::new();
    let longest = longest.chars().rev();
    for (i, longest_char) in longest.enumerate() {
        let shortest_char = shortest.chars().nth_back(i).or(Some('0')).unwrap();
        let a_digit = longest_char.to_digit(10).or(Some(0)).unwrap();
        let b_digit = shortest_char.to_digit(10).or(Some(0)).unwrap();
        let mut sum = a_digit + b_digit + reminder;

        if sum > 9 {
            sum = sum - 10;
            reminder = 1;
        } else {
            reminder = 0;
        };

        result.push(sum);
    }

    if reminder > 0 {
        result.push(reminder);
    }

    result.into_iter().rev().map(|x| x.to_string()).collect::<String>()
}

#[cfg(test)]
mod tests {
    use super::add;

    #[test]
    fn it_works() {
        assert_eq!(add("1", "1"), "2");
        assert_eq!(add("123", "456"), "579");
        assert_eq!(add("888", "222"), "1110");
        assert_eq!(add("1372", "69"), "1441");
        assert_eq!(add("12", "456"), "468");
        assert_eq!(add("101", "100"), "201");
        assert_eq!(add("63829983432984289347293874", "90938498237058927340892374089"), "91002328220491911630239667963");
    }
}
