pub fn cap(s: &str) -> Vec<String> {
    let mut result = vec![String::new(), String::new()];

    for (i, c) in s.chars().enumerate() {
        let capitalized = c.to_uppercase().collect::<String>();
        let cap_index = i % 2;
        let alt_index = 1 - cap_index;

        result[cap_index].push_str(capitalized.as_str());
        result[alt_index].push(c);
    }

    result
}

#[cfg(test)]
mod tests {
    use super::cap;

    #[test]
    fn it_works() {
        assert_eq!(cap("abcdef"), ["AbCdEf", "aBcDeF"]);
        assert_eq!(cap("codewars"), ["CoDeWaRs", "cOdEwArS"]);
        assert_eq!(cap("abracadabra"), ["AbRaCaDaBrA", "aBrAcAdAbRa"]);
        assert_eq!(cap("codewarriors"), ["CoDeWaRrIoRs", "cOdEwArRiOrS"]);
    }
}
