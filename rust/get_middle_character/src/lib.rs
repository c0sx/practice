// https://www.codewars.com/kata/56747fd5cb988479af000028/rust

pub fn get_middle(s: &str) -> &str {
    let index = s.len() / 2;

    return if s.len() % 2 == 0 {
        &s[index - 1..index + 1]
    } else {
        &s[index..index + 1]
    }
}

#[cfg(test)]
mod tests {
    use super::get_middle;

    #[test]
    fn example_tests() {
        assert_eq!(get_middle("test"),"es");
        assert_eq!(get_middle("testing"),"t");
        assert_eq!(get_middle("middle"),"dd");
        assert_eq!(get_middle("A"),"A");
        assert_eq!(get_middle("of"),"of");
    }
}
