pub fn solve(a: &str, b: &str) -> String {
    let first = a.chars().into_iter()
        .filter(|&c| b.contains(c) == false)
        .collect::<String>();

    let second = b.chars().into_iter()
        .filter(|&c| a.contains(c) == false)
        .collect::<String>();

    format!("{}{}", first, second)
}

#[cfg(test)]
mod tests {
    use super::solve;

    #[test]
    fn test_basics() {
        assert_eq!(solve("xyab","xzca"),"ybzc".to_string());
        assert_eq!(solve("xyabb","xzca"),"ybbzc".to_string());
        assert_eq!(solve("abcd","xyz"),"abcdxyz".to_string());
        assert_eq!(solve("xxx","xzca"),"zca".to_string());
    }
}
