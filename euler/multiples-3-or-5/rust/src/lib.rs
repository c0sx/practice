
pub fn multiples_3_or_5(length: i32) -> i32 {
    let mut vec = Vec::<i32>::new();
    let mut i = 1;

    while i < length {
        if i % 3 == 0 || i % 5 == 0 {
            vec.push(i);
        }

        i += 1;
    }

    let r = vec.iter().sum();
    println!("{}", r);
    return r;
}

#[cfg(test)]
mod tests {
    use super::multiples_3_or_5;

    #[test]
    fn it_works_for_10() {
        let input = 10;
        let expected = 23;

        let result = multiples_3_or_5(input);
        assert_eq!(result, expected);
    }
}
