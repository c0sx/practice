pub fn square_sum(vec: Vec<i32>) -> i32 {
    let mut sum = 0;
    for num in vec {
        sum += num.pow(2);
    }

    sum
}

#[cfg(test)]
mod tests {
    use super::square_sum;

    #[test]
    fn returns_expected() {
        assert_eq!(square_sum(vec![1, 2]), 5);
        assert_eq!(square_sum(vec![-1, -2]), 5);
        assert_eq!(square_sum(vec![5, 3, 4]), 50);
    }
}
