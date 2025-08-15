pub fn str_str(haystack: String, needle: String) -> i32 {
    let haystack_vec: Vec<char> = haystack.chars().collect();
    let needle_vec: Vec<char> = needle.chars().collect();

    let mut i: usize = 0;
    let mut j: usize = 0;

    while i < haystack.len() {
        if haystack_vec[i] == needle_vec[j] {
            j += 1;
        } else {
            i -= j;
            j = 0;
        }

        i += 1;

        if j >= needle.len() {
            return (i - j) as i32;
        }
    }

    return -1;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works1() {
        let result = str_str(String::from("sadbutsad"), String::from("sad"));

        assert_eq!(result, 0);
    }

    #[test]
    fn it_works2() {
        let result = str_str(String::from("leetcode"), String::from("leeto"));

        assert_eq!(result, -1);
    }

    #[test]
    fn it_works3() {
        let result = str_str(String::from("hello"), String::from("ll"));

        assert_eq!(result, 2)
    }

    #[test]
    fn it_works4() {
        let result = str_str(String::from("aaab"), String::from("aab"));

        assert_eq!(result, 1)
    }
}
