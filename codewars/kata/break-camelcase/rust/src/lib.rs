pub fn solution(s: &str) -> String {
    let mut result = String::new();

    for c in s.chars() {
        if c.is_uppercase() {
            result.push(' ');
        }

        result.push(c);
    }

    result
}

#[cfg(test)]
mod tests {
    use super::solution;

    #[test]
    fn test_solution() {
        assert_eq!(solution("camelCasing"), "camel Casing");
        assert_eq!(solution("camelCasingTest"), "camel Casing Test");
    }
}
