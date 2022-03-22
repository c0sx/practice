pub fn staircase(n: usize) -> Vec<String> {
    let mut result = Vec::<String>::new();

    let mut i = 0;
    while i < n {
        let amount_of_spaces = n - i - 1;
        let amount_of_symbols = n - amount_of_spaces;
        let mut spaces = " ".repeat(amount_of_spaces);
        let symbols = "#".repeat(amount_of_symbols);

        spaces.push_str(symbols.as_str());
        result.push(spaces);
        i += 1;
    }

    result
}

#[cfg(test)]
mod tests {
    use super::staircase;

    #[test]
    fn it_works() {
        let input = 6;
        let expected = [
            "     #",
            "    ##",
            "   ###",
            "  ####",
            " #####",
            "######"
        ];

        let result = staircase(input);
        println!("result {:?}", result);
        assert_eq!(result, expected);
    }
}
