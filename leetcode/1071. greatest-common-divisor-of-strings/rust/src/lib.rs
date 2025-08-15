fn gcd(len1: usize, len2: usize) -> usize {
    return if len2 == 0 {
        len1
    } else {
        gcd(len2, len1 % len2)
    };
}
pub fn gcd_of_strings(str1: String, str2: String) -> String {
    let s1 = format!("{}{}", str1, str2);
    let s2 = format!("{}{}", str2, str1);

    if s1 != s2 {
        return String::new();
    }

    let len = gcd(str1.len(), str2.len());
    let result = str1.chars().take(len).collect();

    return result;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works1() {
        let result = gcd_of_strings("ABCABC".to_string(), "ABC".to_string());

        assert_eq!(result, "ABC".to_string());
    }

    #[test]
    fn it_works2() {
        let result = gcd_of_strings("ABABAB".to_string(), "ABAB".to_string());

        assert_eq!(result, "AB".to_string());
    }

    #[test]
    fn it_works3() {
        let result = gcd_of_strings("LEET".to_string(), "CODE".to_string());

        assert_eq!(result, "".to_string());
    }

    #[test]
    fn it_works4() {
        let result = gcd_of_strings("ABCDEF".to_string(), "ABC".to_string());

        assert_eq!(result, "".to_string());
    }

    #[test]
    fn it_works5() {
        let result = gcd_of_strings(
            "TAUXXTAUXXTAUXXTAUXXTAUXX".to_string(),
            "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX".to_string(),
        );

        assert_eq!(result, "TAUXX".to_string());
    }
}
