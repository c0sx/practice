// https://www.codewars.com/kata/554ca54ffa7d91b236000023/rust

use std::collections::HashMap;

pub fn delete_nth(list: &[u8], n: usize) -> Vec<u8> {
    let mut enough: HashMap<u8, usize> = HashMap::new();

    list
        .into_iter()
        .filter(|&item| {
            let counter = enough.entry(*item).or_insert(0);
            *counter += 1;

            return *counter <= n;
        })
        .cloned()
        .collect()
}


#[cfg(test)]
mod tests {
    use super::delete_nth;

    #[test]
    fn test_basic() {
        assert_eq!(delete_nth(&[20,37,20,21], 1), vec![20,37,21]);
        assert_eq!(delete_nth(&[1,1,3,3,7,2,2,2,2], 3), vec![1, 1, 3, 3, 7, 2, 2, 2]);
    }
}
