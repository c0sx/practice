pub fn sum(input: &[i32]) -> i32 {
    input.iter().sum()
}

#[cfg(test)]
mod tests {
    use super::sum;

    #[test]
    fn it_works() {
        let arr = [1, 2, 3, 4, 5, 6];
        let result = sum(&arr);
        assert_eq!(result, 21);
    }
}
